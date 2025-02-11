namespace StormGeo.AdvisorCore.Payloads;

public class RadiusPayload
{
    public int? LocaleId { get; set; }
    public string? Latitude { get; set; }
    public string? Longitude { get; set; }
    public string? StartDate { get; set; }
    public string? EndDate { get; set; }
    public int? Radius { get; set; }

    public string GetQueryParams()
    {
        return new QueryParameterBuilder()
            .AddLocaleId(LocaleId)
            .AddLatLon(Latitude, Longitude)
            .AddStartDate(StartDate)
            .AddEndDate(EndDate)
            .AddRadius(Radius)
            .Build();
    }
}
