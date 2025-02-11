using System.Web;

namespace StormGeo.AdvisorCore.Payloads;

public class QueryParameterBuilder
{
    private readonly List<string> _Params = new();

    public QueryParameterBuilder AddLocaleId(int? localeId)
    {
        AddParamIfValueIsNotNull("localeId", localeId?.ToString());
        return this;
    }

    public QueryParameterBuilder AddLatLon(string? latitude, string? longitude)
    {
        AddParamIfValueIsNotNull("latitude", latitude);
        AddParamIfValueIsNotNull("longitude", longitude);
        return this;
    }

    public QueryParameterBuilder AddStationId(string? stationId)
    {
        AddParamIfValueIsNotNull("stationId", stationId);
        return this;
    }

    public QueryParameterBuilder AddStartDate(string? startDate)
    {
        AddParamIfValueIsNotNull("startDate", startDate);
        return this;
    }

    public QueryParameterBuilder AddEndDate(string? endDate)
    {
        AddParamIfValueIsNotNull("endDate", endDate);
        return this;
    }

    public QueryParameterBuilder AddVariables(string[]? variables)
    {
        if (variables != null) {
            foreach (string variable in variables)
            {
                AddParamIfValueIsNotNull("variables[]", variable);
            }
        }
        return this;
    }

    public QueryParameterBuilder AddTimezone(int? timezone)
    {
        AddParamIfValueIsNotNull("timezone", timezone?.ToString());
        return this;
    }

    public QueryParameterBuilder AddLayer(string? layer)
    {
        AddParamIfValueIsNotNull("layer", layer);
        return this;
    }

    public QueryParameterBuilder AddRadius(int? radius)
    {
        AddParamIfValueIsNotNull("radius", radius?.ToString());
        return this;
    }

    public string Build()
    {
        return string.Join('&', _Params) ?? "";
    }

    public QueryParameterBuilder AddParamIfValueIsNotNull(string param, string? value)
    {
        if (value != null)
        {
            _Params.Add($"{param}={HttpUtility.UrlEncode(value)}");
        }
        return this;
    }
}
