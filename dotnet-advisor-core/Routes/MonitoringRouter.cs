using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class MonitoringRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public Task<AdvisorResponse<string>> GetAlertsAsync()
    {
        return base.GetAsync("/v1/monitoring/alerts" + base.FormatQueryParams(""));
    }
}
