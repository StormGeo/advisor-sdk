namespace StormGeo.AdvisorCore.Payloads;

public class StorageDownloadPayload()
{
    public string? FileName { get; set; }
    public string? AccessKey { get; set; }

    public string GetQueryParams()
    {
        return new QueryParameterBuilder()
            .AddParamIfValueIsNotNull("accessKey", AccessKey)
            .Build();
    }
}

