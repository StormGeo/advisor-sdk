namespace StormGeo.AdvisorCore.Payloads;

public class StorageListPayload(int page, int pageSize)
{
    public int Page { get; set; } = page;
    public int PageSize { get; set; } = pageSize;
    public string? StartDate { get; set; }
    public string? EndDate { get; set; }
    public string? FileName { get; set; }
    public string? FileExtension { get; set; }

    public string GetQueryParams()
    {
        return new QueryParameterBuilder()
            .AddParamIfValueIsNotNull("page", Page.ToString())
            .AddParamIfValueIsNotNull("pageSize", PageSize.ToString())
            .AddParamIfValueIsNotNull("fileName", FileName)
            .AddParamIfValueIsNotNull("fileExtension", FileExtension)
            .AddStartDate(StartDate)
            .AddEndDate(EndDate)
            .Build();
    }
}

