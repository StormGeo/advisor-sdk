using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class CurrentWeatherRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public async Task<AdvisorResponse<string>> GetAsync(CurrentWeatherPayload payload)
    {
        return await base.GetAsync("/v1/current-weather" + base.FormatQueryParams(payload.GetQueryParams()));
    }
}
