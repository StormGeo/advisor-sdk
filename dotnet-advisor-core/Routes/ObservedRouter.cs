using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class ObservedRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public Task<AdvisorResponse<string>> GetDailyAsync(WeatherPayload payload)
    {
        return base.GetAsync("/v1/observed/daily" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public Task<AdvisorResponse<string>> GetHourlyAsync(WeatherPayload payload)
    {
        return base.GetAsync("/v1/observed/hourly" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public Task<AdvisorResponse<string>> GetPeriodAsync(WeatherPayload payload)
    {
        return base.GetAsync("/v1/observed/period" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public Task<AdvisorResponse<string>> GetLightningAsync(RadiusPayload payload)
    {
        return base.GetAsync("/v1/observed/lightning" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public Task<AdvisorResponse<string>> GetLightningByGeometryAsync(GeometryPayload payload)
    {
        return base.PostAsync(
            "/v1/observed/lightning" + base.FormatQueryParams(payload.GetQueryParams()),
            payload.GetBody()
        );
    }

    public Task<AdvisorResponse<string>> GetFireFocusAsync(RadiusPayload payload)
    {
        return base.GetAsync("/v1/observed/fire-focus" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public Task<AdvisorResponse<string>> GetFireFocusByGeometryAsync(GeometryPayload payload)
    {
        return base.PostAsync(
            "/v1/observed/fire-focus" + base.FormatQueryParams(payload.GetQueryParams()),
            payload.GetBody()
        );
    }

    public Task<AdvisorResponse<string>> GetStationDataAsync(StationPayload payload)
    {
        return base.GetAsync("/v1/station" + base.FormatQueryParams(payload.GetQueryParams()));
    }
}
