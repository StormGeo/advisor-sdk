using System.Text;
using System.Text.Json;
using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class SchemaRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public Task<AdvisorResponse<string>> GetDefinitionAsync()
    {
        return base.GetAsync("/v1/schema/definition" + base.FormatQueryParams(""));
    }

    public Task<AdvisorResponse<string>> PostDefinitionAsync(
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

        return base.PostAsync(
            "/v1/schema/definition" + base.FormatQueryParams(""),
            new StringContent(JsonSerializer.Serialize(payload), Encoding.UTF8, "application/json")
        );
    }

    public Task<AdvisorResponse<string>> PostParametersAsync(string identifier, Dictionary<string, object> parameters)
    {
        var payload = new Dictionary<string, object>(parameters)
        {
            { "identifier", identifier }
        };

        return base.PostAsync(
            "/v1/schema/parameters" + base.FormatQueryParams(""),
            new StringContent(JsonSerializer.Serialize(payload), Encoding.UTF8, "application/json")
        );
    }
}
