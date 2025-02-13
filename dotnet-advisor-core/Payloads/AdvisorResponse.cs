using Newtonsoft.Json;

namespace StormGeo.AdvisorCore.Payloads;

public class AdvisorResponse<TData>(string? error, TData? data)
{
    public TData? Data { get; } = data;
    public string? Error { get; } = error;

    public dynamic? JsonDeserializeData()
    {
        if (Data != null && Data.GetType().Equals(typeof(string))) {
            return JsonConvert.DeserializeObject<dynamic>(Data as string ?? "");
        }

        return null;
    }

    public dynamic? JsonDeserializeError()
    {
        if (Error != null) {
            return JsonConvert.DeserializeObject<dynamic>(Error);
        }

        return null;
    }
}
