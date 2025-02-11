namespace StormGeo.AdvisorCore.Payloads;

public class StationPayload
{
    public string? StationId { get; set; }
    public string? StartDate { get; set; }
    public string? EndDate { get; set; }
    public string[]? Variables { get; set; }
    public string? Layer { get; set; }

    public string GetQueryParams()
    {
        return new QueryParameterBuilder()
            .AddStationId(StationId)
            .AddStartDate(StartDate)
            .AddEndDate(EndDate)
            .AddVariables(Variables)
            .AddLayer(Layer)
            .Build();
    }
}
