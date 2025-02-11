using System.Text;
using System.Text.Json;

namespace StormGeo.AdvisorCore.Payloads;

public class GeometryPayload
{
    public string? StartDate { get; set; }
    public string? EndDate { get; set; }
    public int? Radius { get; set; }
    public string? Geometry { get; set; }

    public string GetQueryParams()
    {
        return new QueryParameterBuilder()
            .AddStartDate(StartDate)
            .AddEndDate(EndDate)
            .AddRadius(Radius)
            .Build();
    }

    public HttpContent GetBody()
    {
        var json = JsonSerializer.Serialize(new Dictionary<string, string>() {
            { "geometry", Geometry ?? "" },
        });

        return new StringContent(json, Encoding.UTF8, "application/json");
    }
}
