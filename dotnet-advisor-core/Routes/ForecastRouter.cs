using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class ForecastRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public Task<AdvisorResponse<string>> GetDaily(WeatherPayload payload)
    {
        return base.Get("/v1/forecast/daily" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public Task<AdvisorResponse<string>> GetHourly(WeatherPayload payload)
    {
        return base.Get("/v1/forecast/hourly" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public Task<AdvisorResponse<string>> GetPeriod(WeatherPayload payload)
    {
        return base.Get("/v1/forecast/period" + base.FormatQueryParams(payload.GetQueryParams()));
    }
}
