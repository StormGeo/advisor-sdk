using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class ChartRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public async Task<AdvisorResponse<Stream>> GetForecastDailyAsync(WeatherPayload payload)
    {
        return await base.GetFileAsync("/v1/forecast/daily/chart" + payload.GetQueryParams());
    }

    public async Task<AdvisorResponse<Stream>> GetForecastHourlyAsync(WeatherPayload payload)
    {
        return await base.GetFileAsync("/v1/forecast/hourly/chart" + payload.GetQueryParams());
    }

    public async Task<AdvisorResponse<Stream>> GetObservedDailyAsync(WeatherPayload payload)
    {
        return await base.GetFileAsync("/v1/observed/daily/chart" + payload.GetQueryParams());
    }

    public async Task<AdvisorResponse<Stream>> GetObservedHourlyAsync(WeatherPayload payload)
    {
        return await base.GetFileAsync("/v1/observed/hourly/chart" + payload.GetQueryParams());
    }
}
