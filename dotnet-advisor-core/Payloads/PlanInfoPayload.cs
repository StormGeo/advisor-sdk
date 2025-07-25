namespace StormGeo.AdvisorCore.Payloads;

public class PlanInfoPayload
{
  public int? Timezone { get; set; }

  public string GetQueryParams()
  {
    return new QueryParameterBuilder()
      .AddParamIfValueIsNotNull("timezone", Timezone?.ToString())
      .Build();
  }
}