using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class PlanRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public async Task<AdvisorResponse<string>> GetInfoAsync()
    {
        return await base.GetAsync($"/v1/plan/{base._config.Token}");
    }

    public async Task<AdvisorResponse<string>> GetRequestDetailsAsync(RequestDetailsPayload payload)
    {
        return await base.GetAsync("/v1/plan/request-details" + base.FormatQueryParams(payload.GetQueryParams()));
    }
}
