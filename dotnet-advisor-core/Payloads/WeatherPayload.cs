namespace StormGeo.AdvisorCore.Payloads;

public class WeatherPayload
{
    public int? LocaleId { get; set; }
    public string? Latitude { get; set; }
    public string? Longitude { get; set; }
    public string? StationId { get; set; }
    public string? StartDate { get; set; }
    public string? EndDate { get; set; }
    public string[]? Variables { get; set; }
    public int? Timezone { get; set; }

    public string GetQueryParams()
    {
        return new QueryParameterBuilder()
            .AddLocaleId(LocaleId)
            .AddLatLon(Latitude, Longitude)
            .AddStationId(StationId)
            .AddStartDate(StartDate)
            .AddEndDate(EndDate)
            .AddVariables(Variables)
            .AddTimezone(Timezone)
            .Build();
    }
}
