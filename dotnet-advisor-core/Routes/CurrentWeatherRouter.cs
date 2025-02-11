using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class CurrentWeatherRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public Task<AdvisorResponse<string>> Get(CurrentWeatherPayload payload)
    {
        return base.Get("/v1/current-weather" + base.FormatQueryParams(payload.GetQueryParams()));
    }
}
