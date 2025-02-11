using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public abstract class BaseRouter(AdvisorCoreConfig config)
{
    private readonly string BaseUrl = "https://advisor-core.climatempo.io/api";
    protected readonly AdvisorCoreConfig _Config = config;

    protected string FormatQueryParams(string parameters)
    {
        string queryParams = $"?token={_Config.Token}";
        if (parameters.Length > 0)
        {
            queryParams += "&" + parameters;
        }

        return queryParams;
    }

    protected Task<AdvisorResponse<string>> Get(string route)
    {
        return GetDefaultAdvisorResponseAsync(
            RetryRequestAsync(() => MakeRequest(HttpMethod.Get, BaseUrl + route))
        );
    }

    protected Task<AdvisorResponse<Stream>> GetImage(string route)
    {
        return GetAdvisorResponseStreamAsync(
            RetryRequestAsync(() => MakeRequest(HttpMethod.Get, BaseUrl + route))
        );
    }

    protected Task<AdvisorResponse<string>> Post(string route, HttpContent body)
    {
        return GetDefaultAdvisorResponseAsync(
            RetryRequestAsync(() => MakeRequest(HttpMethod.Post, BaseUrl + route, body))
        );
    }

    private Task<HttpResponseMessage> MakeRequest(HttpMethod method, string uri, HttpContent? body = null)
    {
        var request = new HttpRequestMessage()
        {
            Method = method,
            RequestUri = new Uri(uri)
        };

        request.Headers.Add("User-Agent", "Csharp-AdvisorCore-SDK");
        foreach (var header in _Config.Headers.Keys)
        {
            request.Headers.Add(header, _Config.Headers[header]);
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

        for (int retryNumber = _Config.Attempts; retryNumber >= 0; retryNumber--)
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

            Thread.Sleep(_Config.Delay);
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
