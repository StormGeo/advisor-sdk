namespace StormGeo.AdvisorCore.Payloads;

public class PmtilesPayload
{
    public required string Mode { get; set; }
    public required string Model { get; set; }
    public required string Variable { get; set; }
    public required string Aggregation { get; set; }
    public string? Istep { get; set; }
    public string? Fstep { get; set; }
    public int? Timezone { get; set; }
    public required int MaxZoom { get; set; }
    public string? Cmap { get; set; }
    public string? DynamicElevation { get; set; }
    public string? DynamicType { get; set; }
    public string? DynamicVariable { get; set; }

    public string GetQueryParams()
    {
        return new QueryParameterBuilder()
            .AddParamIfValueIsNotNull("istep", Istep)
            .AddParamIfValueIsNotNull("fstep", Fstep)
            .AddTimezone(Timezone)
            .AddParamIfValueIsNotNull("maxZoom", MaxZoom.ToString())
            .AddParamIfValueIsNotNull("cmap", Cmap)
            .AddParamIfValueIsNotNull("dynamicElevation", DynamicElevation)
            .AddParamIfValueIsNotNull("dynamicType", DynamicType)
            .AddParamIfValueIsNotNull("dynamicVariable", DynamicVariable)
            .Build();
    }
}
