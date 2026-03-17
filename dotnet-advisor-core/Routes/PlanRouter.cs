using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class PlanRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public async Task<AdvisorResponse<string>> GetInfoAsync(PlanInfoPayload payload)
    {
        return await base.GetAsync("/v2/plan" + payload.GetQueryParams());
    }

    public async Task<AdvisorResponse<string>> GetRequestDetailsAsync(RequestDetailsPayload payload)
    {
        return await base.GetAsync("/v1/plan/request-details" + payload.GetQueryParams());
    }
}
