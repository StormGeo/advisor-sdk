using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public abstract class BaseRouter(AdvisorCoreConfig config)
{
    private readonly string _baseUrl = "https://advisor-core.climatempo.io/api";
    protected readonly AdvisorCoreConfig _config = config;

    protected string FormatQueryParams(string parameters)
    {
        string queryParams = $"?token={_config.Token}";
        if (parameters.Length > 0)
        {
            queryParams += "&" + parameters;
        }

        return queryParams;
    }

    protected Task<AdvisorResponse<string>> GetAsync(string route)
    {
        return GetDefaultAdvisorResponseAsync(
            RetryRequestAsync(() => MakeRequestAsync(HttpMethod.Get, _baseUrl + route))
        );
    }

    protected Task<AdvisorResponse<Stream>> GetImageAsync(string route)
    {
        return GetAdvisorResponseStreamAsync(
            RetryRequestAsync(() => MakeRequestAsync(HttpMethod.Get, _baseUrl + route))
        );
    }

    protected Task<AdvisorResponse<string>> PostAsync(string route, HttpContent body)
    {
        return GetDefaultAdvisorResponseAsync(
            RetryRequestAsync(() => MakeRequestAsync(HttpMethod.Post, _baseUrl + route, body))
        );
    }

    private Task<HttpResponseMessage> MakeRequestAsync(HttpMethod method, string uri, HttpContent? body = null)
    {
        var request = new HttpRequestMessage()
        {
            Method = method,
            RequestUri = new Uri(uri)
        };

        request.Headers.Add("User-Agent", "Csharp-AdvisorCore-SDK");
        foreach (var header in _config.Headers.Keys)
        {
            request.Headers.Add(header, _config.Headers[header]);
        }

        if (body != null)
        {
            request.Content = body;
        }

        using var client = new HttpClient();
        return client.SendAsync(request);
    }

    private async Task<HttpResponseMessage?> RetryRequestAsync(Func<Task<HttpResponseMessage>> request)
    {
        HttpResponseMessage? response = null;

        for (int retryNumber = _config.Attempts; retryNumber >= 0; retryNumber--)
        {
            response = await request();
            if (response != null)
            {
                var statusCode = (int) response.StatusCode;

                if (retryNumber == 0 || (statusCode < 500 && statusCode != 429))
                {
                    return response;
                }

                response.Dispose();
            }

            Thread.Sleep(_config.Delay);
        }

        return response;
    }

    private static async Task<AdvisorResponse<string>> GetDefaultAdvisorResponseAsync(Task<HttpResponseMessage?> request)
    {
        using var response = await request;
        if (response == null)
        {
            return new(null, null);
        }

        var statusCode = (int) response.StatusCode;
        var contentText = await response.Content.ReadAsStringAsync();
        if (statusCode < 400)
        {
            return new(null, contentText);
        }

        return new(contentText, null);
    }

    private static async Task<AdvisorResponse<Stream>> GetAdvisorResponseStreamAsync(Task<HttpResponseMessage?> request)
    {
        using var response = await request;
        if (response == null)
        {
            return new(null, null);
        }

        var statusCode = (int) response.StatusCode;
        if (statusCode < 400)
        {
            var contentStream = await response.Content.ReadAsStreamAsync();
            return new(null, contentStream);
        }

        var contentText = await response.Content.ReadAsStringAsync();
        return new(contentText, null);
    }
}
