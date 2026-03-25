using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class PmtilesRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public async Task<AdvisorResponse<Stream>> Get(PmtilesPayload payload)
    {
        var path = String.Format(
            "/v1/pmtiles/{0}/{1}/{2}/{3}.pmtiles",
            payload.Mode,
            payload.Model,
            payload.Aggregation,
            payload.Variable
        );

        return await base.GetFileAsync(path + payload.GetQueryParams());
    }
}
