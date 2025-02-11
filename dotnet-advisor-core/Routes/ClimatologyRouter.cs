using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class ClimatologyRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public Task<AdvisorResponse<string>> GetDailyAsync(ClimatologyPayload payload)
    {
        return base.GetAsync("/v1/climatology/daily" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public Task<AdvisorResponse<string>> GetMonthlyAsync(ClimatologyPayload payload)
    {
        return base.GetAsync("/v1/climatology/monthly" + base.FormatQueryParams(payload.GetQueryParams()));
    }
}
