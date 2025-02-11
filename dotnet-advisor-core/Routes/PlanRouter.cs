using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class PlanRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public Task<AdvisorResponse<string>> GetInfoAsync()
    {
        return base.GetAsync($"/v1/plan/{base._config.Token}");
    }

    public Task<AdvisorResponse<string>> GetRequestDetailsAsync(RequestDetailsPayload payload)
    {
        return base.GetAsync("/v1/plan/request-details" + base.FormatQueryParams(payload.GetQueryParams()));
    }
}
