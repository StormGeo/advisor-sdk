using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class ClimatologyRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public Task<AdvisorResponse<string>> GetDaily(ClimatologyPayload payload)
    {
        return base.Get("/v1/climatology/daily" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public Task<AdvisorResponse<string>> GetMonthly(ClimatologyPayload payload)
    {
        return base.Get("/v1/climatology/monthly" + base.FormatQueryParams(payload.GetQueryParams()));
    }
}
