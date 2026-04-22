using StormGeo.AdvisorCore;
using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Tests.Integration;

public class RoutesIntegrationTests
{
    private static (AdvisorCore sdk, IntegrationPayloads payloads) CreateContext()
    {
        return (IntegrationHelpers.CreateSdk(), IntegrationHelpers.CreatePayloads());
    }

    [IntegrationFact]
    public async Task ForecastRoutesAsync()
    {
        var (sdk, payloads) = CreateContext();
        var routes = new Func<WeatherPayload, Task<AdvisorResponse<string>>>[]
        {
            sdk.Forecast.GetDailyAsync,
            sdk.Forecast.GetHourlyAsync,
            sdk.Forecast.GetPeriodAsync,
        };

        foreach (var route in routes)
        {
            IntegrationHelpers.AssertJsonSuccess(await route(payloads.WeatherPayload));
        }
    }

    [IntegrationFact]
    public async Task ObservedWeatherRoutesAsync()
    {
        var (sdk, payloads) = CreateContext();
        var routes = new Func<WeatherPayload, Task<AdvisorResponse<string>>>[]
        {
            sdk.Observed.GetDailyAsync,
            sdk.Observed.GetHourlyAsync,
            sdk.Observed.GetPeriodAsync,
        };

        foreach (var route in routes)
        {
            IntegrationHelpers.AssertJsonSuccess(await route(payloads.WeatherPayload));
        }
    }

    [IntegrationFact]
    public async Task ObservedStationDataAsync()
    {
        var (sdk, payloads) = CreateContext();
        IntegrationHelpers.AssertJsonSuccess(await sdk.Observed.GetStationDataAsync(payloads.StationPayload));
    }

    [IntegrationFact]
    public async Task ObservedRadiusRoutesAsync()
    {
        var (sdk, payloads) = CreateContext();
        IntegrationHelpers.AssertJsonSuccess(await sdk.Observed.GetFireFocusAsync(payloads.RadiusPayload));
        IntegrationHelpers.AssertJsonSuccess(await sdk.Observed.GetLightningAsync(payloads.RadiusPayload));
        IntegrationHelpers.AssertJsonSuccess(await sdk.Observed.GetLightningDetailsAsync(payloads.LightningDetailsPayload));
    }

    [IntegrationFact]
    public async Task ObservedGeometryRoutesAsync()
    {
        var (sdk, payloads) = CreateContext();
        IntegrationHelpers.AssertJsonSuccess(await sdk.Observed.GetFireFocusByGeometryAsync(payloads.GeometryPayload));
        IntegrationHelpers.AssertJsonSuccess(await sdk.Observed.GetLightningByGeometryAsync(payloads.GeometryPayload));
    }

    [IntegrationFact]
    public async Task ObservedLightningLiteAsync()
    {
        var (sdk, payloads) = CreateContext();
        IntegrationHelpers.AssertJsonSuccess(await sdk.Observed.GetLightningLiteAsync(payloads.LightningLitePayload));
    }

    [IntegrationFact]
    public async Task CurrentWeatherAsync()
    {
        var (sdk, payloads) = CreateContext();
        IntegrationHelpers.AssertJsonSuccess(await sdk.CurrentWeather.GetAsync(payloads.CurrentWeatherPayload));
    }

    [IntegrationFact]
    public async Task ClimatologyRoutesAsync()
    {
        var (sdk, payloads) = CreateContext();
        IntegrationHelpers.AssertJsonSuccess(await sdk.Climatology.GetDailyAsync(payloads.ClimatologyPayload));
        IntegrationHelpers.AssertJsonSuccess(await sdk.Climatology.GetMonthlyAsync(payloads.ClimatologyPayload));
    }

    [IntegrationFact]
    public async Task MonitoringAlertsAsync()
    {
        var (sdk, _) = CreateContext();
        IntegrationHelpers.AssertJsonSuccess(await sdk.Monitoring.GetAlertsAsync());
    }

    [IntegrationFact]
    public async Task PlanRoutesAsync()
    {
        var (sdk, payloads) = CreateContext();
        IntegrationHelpers.AssertJsonSuccess(await sdk.Plan.GetInfoAsync(payloads.PlanInfoPayload));
        IntegrationHelpers.AssertJsonSuccess(await sdk.Plan.GetRequestDetailsAsync(payloads.RequestDetailsPayload));
    }

    [IntegrationFact]
    public async Task ChartRoutesAsync()
    {
        var (sdk, payloads) = CreateContext();
        var routes = new Func<WeatherPayload, Task<AdvisorResponse<Stream>>>[]
        {
            sdk.Chart.GetForecastDailyAsync,
            sdk.Chart.GetForecastHourlyAsync,
            sdk.Chart.GetObservedDailyAsync,
            sdk.Chart.GetObservedHourlyAsync,
        };

        foreach (var route in routes)
        {
            await IntegrationHelpers.AssertStreamSuccessAsync(await route(payloads.WeatherChartPayload));
        }
    }

    [IntegrationFact]
    public async Task StorageRoutesAsync()
    {
        var (sdk, payloads) = CreateContext();
        IntegrationHelpers.AssertJsonSuccess(await sdk.Storage.ListFilesAsync(payloads.StorageListPayload));
        await IntegrationHelpers.AssertStreamSuccessAsync(await sdk.Storage.DownloadFileAsync(IntegrationHelpers.ResolveStorageDownloadPayload()));
    }

    [IntegrationFact]
    public async Task StaticMapAsync()
    {
        var (sdk, payloads) = CreateContext();
        await IntegrationHelpers.AssertStreamSuccessAsync(await sdk.StaticMap.GetStaticMapAsync(payloads.StaticMapPayload));
    }

    [IntegrationFact]
    public async Task TmsAsync()
    {
        var (sdk, payloads) = CreateContext();
        await IntegrationHelpers.AssertStreamSuccessAsync(await sdk.Tms.Get(payloads.TmsPayload));
    }

    [IntegrationFact]
    public async Task PmtilesAsync()
    {
        var (sdk, payloads) = CreateContext();
        await IntegrationHelpers.AssertStreamSuccessAsync(await sdk.Pmtiles.Get(payloads.PmtilesPayload));
    }

    [IntegrationFact]
    public async Task SchemaRoutesAsync()
    {
        var (sdk, payloads) = CreateContext();
        IntegrationHelpers.AssertJsonSuccess(await sdk.Schema.GetDefinitionAsync());
        IntegrationHelpers.AssertJsonSuccess(await sdk.Schema.PostDefinitionAsync(payloads.SchemaIdentifier, payloads.SchemaDefinitionFields));
    }
}
