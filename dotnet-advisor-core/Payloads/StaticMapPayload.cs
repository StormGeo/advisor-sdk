using System.Globalization;

namespace StormGeo.AdvisorCore.Payloads;

public class StaticMapPayload
{
  public required string Type { get; set; }
  public required string Category { get; set; }
  public required string Variable { get; set; }
  public string? StartDate { get; set; }
  public string? EndDate { get; set; }
  public string? Aggregation { get; set; }
  public string? Model { get; set; }
  public double? Lonmin { get; set; }
  public double? Lonmax { get; set; }
  public double? Latmin { get; set; }
  public double? Latmax { get; set; }
  public int? Dpi { get; set; }
  public bool? Title { get; set; }
  public string? TitleVariable { get; set; }
  public int? Hours { get; set; }

  public string GetQueryParams()
  {
    return new QueryParameterBuilder()
      .AddStartDate(StartDate)
      .AddEndDate(EndDate)
      .AddParamIfValueIsNotNull("aggregation", Aggregation)
      .AddParamIfValueIsNotNull("model", Model)
      .AddParamIfValueIsNotNull("lonmin", Lonmin?.ToString(CultureInfo.InvariantCulture))
      .AddParamIfValueIsNotNull("lonmax", Lonmax?.ToString(CultureInfo.InvariantCulture))
      .AddParamIfValueIsNotNull("latmin", Latmin?.ToString(CultureInfo.InvariantCulture))
      .AddParamIfValueIsNotNull("latmax", Latmax?.ToString(CultureInfo.InvariantCulture))
      .AddParamIfValueIsNotNull("dpi", Dpi?.ToString())
      .AddParamIfValueIsNotNull("title", Title.HasValue ? Title.Value.ToString().ToLowerInvariant() : null)
      .AddParamIfValueIsNotNull("titleVariable", TitleVariable)
      .AddParamIfValueIsNotNull("hours", Hours?.ToString())
      .Build();
  }
}