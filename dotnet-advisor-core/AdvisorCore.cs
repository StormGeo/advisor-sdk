using StormGeo.AdvisorCore.Payloads;
using StormGeo.AdvisorCore.Routes;

namespace StormGeo.AdvisorCore;

public class AdvisorCore
{
    private readonly AdvisorCoreConfig _config;

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
        _config = new AdvisorCoreConfig(token, attempts, delayInSeconds * 1000, headers);

        Chart = new ChartRouter(_config);
        Climatology = new ClimatologyRouter(_config);
        CurrentWeather = new CurrentWeatherRouter(_config);
        Forecast = new ForecastRouter(_config);
        Monitoring = new MonitoringRouter(_config);
        Observed = new ObservedRouter(_config);
        Plan = new PlanRouter(_config);
        Schema = new SchemaRouter(_config);
        Tms = new TmsRouter(_config);
    }

    public void SetHeaderAccept(string value)
    {
        _config.Headers["Accept"] = value;
    }

    public void SetHeaderAcceptLanguage(string value)
    {
        _config.Headers["Accept-Language"] = value;
    }
}
