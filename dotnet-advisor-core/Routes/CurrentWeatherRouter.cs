using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class CurrentWeatherRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public Task<AdvisorResponse<string>> GetAsync(CurrentWeatherPayload payload)
    {
        return base.GetAsync("/v1/current-weather" + base.FormatQueryParams(payload.GetQueryParams()));
    }
}
