using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class MonitoringRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public async Task<AdvisorResponse<string>> GetAlertsAsync()
    {
        return await base.GetAsync("/v1/monitoring/alerts");
    }
}
