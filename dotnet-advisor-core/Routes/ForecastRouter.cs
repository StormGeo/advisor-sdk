using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class ForecastRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public Task<AdvisorResponse<string>> GetDailyAsync(WeatherPayload payload)
    {
        return base.GetAsync("/v1/forecast/daily" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public Task<AdvisorResponse<string>> GetHourlyAsync(WeatherPayload payload)
    {
        return base.GetAsync("/v1/forecast/hourly" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public Task<AdvisorResponse<string>> GetPeriodAsync(WeatherPayload payload)
    {
        return base.GetAsync("/v1/forecast/period" + base.FormatQueryParams(payload.GetQueryParams()));
    }
}
