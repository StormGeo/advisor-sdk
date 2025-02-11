using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class TmsRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public async Task<AdvisorResponse<Stream>> Get(TmsPayload payload)
    {
        var path = String.Format(
            "/v1/tms/{0}/{1}/{2}/{3}/{4:D}/{5:D}/{6:D}.png",
            payload.Server,
            payload.Mode,
            payload.Variable,
            payload.Aggregation,
            payload.X,
            payload.Y,
            payload.Z
        );

        return await base.GetImageAsync(path + base.FormatQueryParams(payload.GetQueryParams()));
    }
}
