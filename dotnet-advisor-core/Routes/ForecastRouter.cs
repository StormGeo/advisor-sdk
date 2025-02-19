using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class ForecastRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public async Task<AdvisorResponse<string>> GetDailyAsync(WeatherPayload payload)
    {
        return await base.GetAsync("/v1/forecast/daily" + payload.GetQueryParams());
    }

    public async Task<AdvisorResponse<string>> GetHourlyAsync(WeatherPayload payload)
    {
        return await base.GetAsync("/v1/forecast/hourly" + payload.GetQueryParams());
    }

    public async Task<AdvisorResponse<string>> GetPeriodAsync(WeatherPayload payload)
    {
        return await base.GetAsync("/v1/forecast/period" + payload.GetQueryParams());
    }
}
