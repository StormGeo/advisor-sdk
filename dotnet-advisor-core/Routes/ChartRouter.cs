using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class ChartRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public Task<AdvisorResponse<Stream>> GetForecastDaily(WeatherPayload payload)
    {
        return base.GetImage("/v1/forecast/daily/chart" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public Task<AdvisorResponse<Stream>> GetForecastHourly(WeatherPayload payload)
    {
        return base.GetImage("/v1/forecast/hourly/chart" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public Task<AdvisorResponse<Stream>> GetObservedDaily(WeatherPayload payload)
    {
        return base.GetImage("/v1/observed/daily/chart" + base.FormatQueryParams(payload.GetQueryParams()));
    }

    public Task<AdvisorResponse<Stream>> GetObservedHourly(WeatherPayload payload)
    {
        return base.GetImage("/v1/observed/hourly/chart" + base.FormatQueryParams(payload.GetQueryParams()));
    }
}
