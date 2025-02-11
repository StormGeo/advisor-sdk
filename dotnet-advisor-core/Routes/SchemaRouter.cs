using System.Text;
using System.Text.Json;
using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class SchemaRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public Task<AdvisorResponse<string>> GetDefinition()
    {
        return base.Get("/v1/schema/definition" + base.FormatQueryParams(""));
    }

    public Task<AdvisorResponse<string>> PostDefinition(
        string identifier,
        Dictionary<string, SchameDefinitionField> fields
    ) {
        var payload = new Dictionary<string, object>()
        {
            { "identifier", identifier }
        };

        foreach (var field in fields.Keys)
        {
            payload.Add(field, fields[field].ToDictionary());
        }

        return base.Post(
            "/v1/schema/definition" + base.FormatQueryParams(""),
            new StringContent(JsonSerializer.Serialize(payload), Encoding.UTF8, "application/json")
        );
    }

    public Task<AdvisorResponse<string>> PostParameters(string identifier, Dictionary<string, object> parameters)
    {
        var payload = new Dictionary<string, object>(parameters)
        {
            { "identifier", identifier }
        };

        return base.Post(
            "/v1/schema/parameters" + base.FormatQueryParams(""),
            new StringContent(JsonSerializer.Serialize(payload), Encoding.UTF8, "application/json")
        );
    }
}
