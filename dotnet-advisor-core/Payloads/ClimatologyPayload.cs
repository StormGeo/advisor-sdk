namespace StormGeo.AdvisorCore.Payloads;

public class ClimatologyPayload
{
    public int? LocaleId { get; set; }
    public string? Latitude { get; set; }
    public string? Longitude { get; set; }
    public string? StationId { get; set; }
    public string[]? Variables { get; set; }

    public string GetQueryParams()
    {
        return new QueryParameterBuilder()
            .AddLocaleId(LocaleId)
            .AddLatLon(Latitude, Longitude)
            .AddStationId(StationId)
            .AddVariables(Variables)
            .Build();
    }
}
