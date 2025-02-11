using StormGeo.AdvisorCore.Routes;

namespace StormGeo.AdvisorCore;

public class AdvisorCore
{
    private readonly AdvisorCoreConfig _Config;

    public ChartRouter Chart { get; }
    public ClimatologyRouter Climatology { get; }
    public CurrentWeatherRouter CurrentWeather { get; }
    public ForecastRouter Forecast { get; }
    public MonitoringRouter Monitoring { get; }
    public ObservedRouter Observed { get; }
    public PlanRouter Plan { get; }
    public SchemaRouter Schema { get; set; }
    public TmsRouter Tms { get; set; }

    public AdvisorCore(string token, int attempts = 5, int delayInSeconds = 5)
    {
        var headers = new Dictionary<string,string>() {
            { "Accept", "application/json" },
            { "Accept-Language", "en-US" }
        };
        _Config = new AdvisorCoreConfig(token, attempts, delayInSeconds * 1000, headers);

        Chart = new ChartRouter(_Config);
        Climatology = new ClimatologyRouter(_Config);
        CurrentWeather = new CurrentWeatherRouter(_Config);
        Forecast = new ForecastRouter(_Config);
        Monitoring = new MonitoringRouter(_Config);
        Observed = new ObservedRouter(_Config);
        Plan = new PlanRouter(_Config);
        Schema = new SchemaRouter(_Config);
        Tms = new TmsRouter(_Config);
    }

    public void SetHeaderAccept(string value)
    {
        _Config.Headers["Accept"] = value;
    }

    public void SetHeaderAcceptLanguage(string value)
    {
        _Config.Headers["Accept-Language"] = value;
    }
}
