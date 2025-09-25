using Microsoft.Extensions.Logging;
using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public abstract class BaseRouter(AdvisorCoreConfig config)
{
    protected readonly AdvisorCoreConfig _config = config;
    private readonly string _baseUrl = "https://advisor-core.climatempo.io/api";
    private static readonly HttpClient _httpClient = new();

    protected async Task<AdvisorResponse<string>> GetAsync(string route)
    {
        return await GetDefaultAdvisorResponseAsync(
            await RetryRequestAsync(() => MakeRequestAsync(HttpMethod.Get, _baseUrl + route))
        );
    }

    protected async Task<AdvisorResponse<Stream>> GetFileAsync(string route)
    {
        var httpResponse = await RetryRequestAsync(
            () => MakeRequestAsync(HttpMethod.Get, _baseUrl + route)
        );

        return await GetAdvisorResponseStreamAsync(httpResponse);
    }

    protected async Task<AdvisorResponse<string>> PostAsync(string route, HttpContent body)
    {
        return await GetDefaultAdvisorResponseAsync(
            await RetryRequestAsync(() => MakeRequestAsync(HttpMethod.Post, _baseUrl + route, body))
        );
    }

    private async Task<HttpResponseMessage> MakeRequestAsync(HttpMethod method, string uri, HttpContent? body = null)
    {
        using var request = new HttpRequestMessage(method, new Uri(uri));

        request.Headers.Add("User-Agent", "Csharp-AdvisorCore-SDK");
        request.Headers.Add("x-advisor-token", _config.Token);
        foreach (var header in _config.Headers.Keys)
        {
            request.Headers.Add(header, _config.Headers[header]);
        }

        if (body != null)
        {
            request.Content = body;
        }

        return await _httpClient.SendAsync(request);
    }

    private async Task<HttpResponseMessage?> RetryRequestAsync(Func<Task<HttpResponseMessage>> request)
    {
        HttpResponseMessage? response = null;

        for (int retryNumber = _config.Attempts; retryNumber >= 0; retryNumber--)
        {
            try
            {
                response = await request();
                if (response != null)
                {
                    var statusCode = (int) response.StatusCode;

                    if (retryNumber == 0 || (statusCode < 500))
                    {
                        return response;
                    }

                    response.Dispose();
                }
            } catch (Exception error)
            {
                _config.Logger?.LogWarning("AdvisorCore request failed with message: {errorMessage}", error.Message);
            }

            if (retryNumber > 0)
            {
                await Task.Delay((int) _config.Delay);
            }
        }

        return response;
    }

    private static async Task<AdvisorResponse<string>> GetDefaultAdvisorResponseAsync(HttpResponseMessage? response)
    {
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

    private static async Task<AdvisorResponse<Stream>> GetAdvisorResponseStreamAsync(HttpResponseMessage? response)
    {
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
