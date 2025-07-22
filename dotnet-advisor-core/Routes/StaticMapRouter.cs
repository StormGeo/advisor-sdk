using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class StaticMapRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
  public async Task<AdvisorResponse<Stream>> GetStaticMapAsync(StaticMapPayload payload)
  {
    var path = String.Format(
      "/v1/map/{0}/{1}/{2}",
      payload.Type,
      payload.Category,
      payload.Variable
    );

    return await base.GetFileAsync(path + payload.GetQueryParams());
  }
}