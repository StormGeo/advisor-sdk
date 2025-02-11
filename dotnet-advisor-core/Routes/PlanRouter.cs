using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class PlanRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public Task<AdvisorResponse<string>> GetInfo()
    {
        return base.Get($"/v1/plan/{base._Config.Token}");
    }

    public Task<AdvisorResponse<string>> GetRequestDetails(RequestDetailsPayload payload)
    {
        return base.Get("/v1/plan/request-details" + base.FormatQueryParams(payload.GetQueryParams()));
    }
}
