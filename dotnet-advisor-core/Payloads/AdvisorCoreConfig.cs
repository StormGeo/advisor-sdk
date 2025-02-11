namespace StormGeo.AdvisorCore.Payloads;

public class AdvisorCoreConfig(string token, int attempts, int delay, Dictionary<string, string> headers)
{
    public string Token { get; } = token;
    public int Attempts { get; } = attempts;
    public int Delay { get; } = delay;
    public Dictionary<string, string> Headers { get; } = headers;
}
