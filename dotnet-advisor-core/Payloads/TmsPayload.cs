namespace StormGeo.AdvisorCore.Payloads;

public class TmsPayload
{
    public required char Server { get; set; }
    public required string Mode { get; set; }
    public required string Variable { get; set; }
    public required string Aggregation { get; set; }
    public int X { get; set; }
    public int Y { get; set; }
    public int Z { get; set; }
    public required string Istep { get; set; }
    public required string Fstep { get; set; }

    public string GetQueryParams()
    {
        return new QueryParameterBuilder()
            .AddParamIfValueIsNotNull("istep", Istep)
            .AddParamIfValueIsNotNull("fstep", Fstep)
            .Build();
    }
}
