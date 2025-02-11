using System.Text;
using System.Text.Json;
using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class SchemaRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public async Task<AdvisorResponse<string>> GetDefinitionAsync()
    {
        return await base.GetAsync("/v1/schema/definition" + base.FormatQueryParams(""));
    }

    public async Task<AdvisorResponse<string>> PostDefinitionAsync(
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

        return await base.PostAsync(
            "/v1/schema/definition" + base.FormatQueryParams(""),
            new StringContent(JsonSerializer.Serialize(payload), Encoding.UTF8, "application/json")
        );
    }

    public async Task<AdvisorResponse<string>> PostParametersAsync(string identifier, Dictionary<string, object> parameters)
    {
        var payload = new Dictionary<string, object>(parameters)
        {
            { "identifier", identifier }
        };

        return await base.PostAsync(
            "/v1/schema/parameters" + base.FormatQueryParams(""),
            new StringContent(JsonSerializer.Serialize(payload), Encoding.UTF8, "application/json")
        );
    }
}
