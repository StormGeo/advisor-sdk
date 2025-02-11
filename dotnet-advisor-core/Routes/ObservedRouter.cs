using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class ObservedRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public Task<AdvisorResponse<string>> GetDaily(WeatherPayload payload)
    {
        return base.Get("/v1/observed/daily" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public Task<AdvisorResponse<string>> GetHourly(WeatherPayload payload)
    {
        return base.Get("/v1/observed/hourly" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public Task<AdvisorResponse<string>> GetPeriod(WeatherPayload payload)
    {
        return base.Get("/v1/observed/period" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public Task<AdvisorResponse<string>> GetLightning(RadiusPayload payload)
    {
        return base.Get("/v1/observed/lightning" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public Task<AdvisorResponse<string>> GetLightningByGeometry(GeometryPayload payload)
    {
        return base.Post(
            "/v1/observed/lightning" + base.FormatQueryParams(payload.GetQueryParams()),
            payload.GetBody()
        );
    }

    public Task<AdvisorResponse<string>> GetFireFocus(RadiusPayload payload)
    {
        return base.Get("/v1/observed/fire-focus" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public Task<AdvisorResponse<string>> GetFireFocusByGeometry(GeometryPayload payload)
    {
        return base.Post(
            "/v1/observed/fire-focus" + base.FormatQueryParams(payload.GetQueryParams()),
            payload.GetBody()
        );
    }

    public Task<AdvisorResponse<string>> GetStationData(StationPayload payload)
    {
        return base.Get("/v1/station" + base.FormatQueryParams(payload.GetQueryParams()));
    }
}
