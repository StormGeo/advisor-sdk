namespace StormGeo.AdvisorCore.Payloads;

public class RequestDetailsPayload(int page, int pageSize)
{
    public int Page { get; set; } = page;
    public int PageSize { get; set; } = pageSize;
    public string? Path { get; set; }
    public string? Status { get; set; }
    public string? StartDate { get; set; }
    public string? EndDate { get; set; }

    public string GetQueryParams()
    {
        return new QueryParameterBuilder()
            .AddParamIfValueIsNotNull("page", Page.ToString())
            .AddParamIfValueIsNotNull("pageSize", PageSize.ToString())
            .AddParamIfValueIsNotNull("path", Path)
            .AddParamIfValueIsNotNull("status", Status)
            .AddStartDate(StartDate)
            .AddEndDate(EndDate)
            .Build();
    }
}

