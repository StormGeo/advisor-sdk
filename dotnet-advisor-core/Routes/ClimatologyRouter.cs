using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class ClimatologyRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public async Task<AdvisorResponse<string>> GetDailyAsync(ClimatologyPayload payload)
    {
        return await base.GetAsync("/v1/climatology/daily" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public async Task<AdvisorResponse<string>> GetMonthlyAsync(ClimatologyPayload payload)
    {
        return await base.GetAsync("/v1/climatology/monthly" + base.FormatQueryParams(payload.GetQueryParams()));
    }
}
