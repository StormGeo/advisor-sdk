using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class MonitoringRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public Task<AdvisorResponse<string>> GetAlerts()
    {
        return base.Get("/v1/monitoring/alerts" + base.FormatQueryParams(""));
    }
}
