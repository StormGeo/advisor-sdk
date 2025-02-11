using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class ObservedRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public async Task<AdvisorResponse<string>> GetDailyAsync(WeatherPayload payload)
    {
        return await base.GetAsync("/v1/observed/daily" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public async Task<AdvisorResponse<string>> GetHourlyAsync(WeatherPayload payload)
    {
        return await base.GetAsync("/v1/observed/hourly" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public async Task<AdvisorResponse<string>> GetPeriodAsync(WeatherPayload payload)
    {
        return await base.GetAsync("/v1/observed/period" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public async Task<AdvisorResponse<string>> GetLightningAsync(RadiusPayload payload)
    {
        return await base.GetAsync("/v1/observed/lightning" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public async Task<AdvisorResponse<string>> GetLightningByGeometryAsync(GeometryPayload payload)
    {
        return await base.PostAsync(
            "/v1/observed/lightning" + base.FormatQueryParams(payload.GetQueryParams()),
            payload.GetBody()
        );
    }

    public async Task<AdvisorResponse<string>> GetFireFocusAsync(RadiusPayload payload)
    {
        return await base.GetAsync("/v1/observed/fire-focus" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public async Task<AdvisorResponse<string>> GetFireFocusByGeometryAsync(GeometryPayload payload)
    {
        return await base.PostAsync(
            "/v1/observed/fire-focus" + base.FormatQueryParams(payload.GetQueryParams()),
            payload.GetBody()
        );
    }

    public async Task<AdvisorResponse<string>> GetStationDataAsync(StationPayload payload)
    {
        return await base.GetAsync("/v1/station" + base.FormatQueryParams(payload.GetQueryParams()));
    }
}
