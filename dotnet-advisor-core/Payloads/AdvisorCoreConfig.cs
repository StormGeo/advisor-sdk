using Microsoft.Extensions.Logging;

namespace StormGeo.AdvisorCore.Payloads;

public class AdvisorCoreConfig(string token, int attempts, double delay, Dictionary<string, string> headers, ILogger? logger)
{
    public string Token { get; } = token;
    public int Attempts { get; } = attempts;
    public double Delay { get; } = delay;
    public Dictionary<string, string> Headers { get; } = headers;
    public ILogger? Logger { get; } = logger;
}
