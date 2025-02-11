namespace StormGeo.AdvisorCore.Payloads;

public class AdvisorResponse<TData>(string? error, TData? data)
{
    public TData? Data { get; } = data;
    public string? Error { get; } = error;
}
