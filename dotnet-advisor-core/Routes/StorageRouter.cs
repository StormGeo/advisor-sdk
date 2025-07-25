using StormGeo.AdvisorCore.Payloads;

namespace StormGeo.AdvisorCore.Routes;

public class StorageRouter(AdvisorCoreConfig config) : BaseRouter(config)
{
    public async Task<AdvisorResponse<string>> ListFilesAsync(StorageListPayload payload)
    {
        return await base.GetAsync($"/v1/storage/list" + payload.GetQueryParams());
    }

    public async Task<AdvisorResponse<Stream>> DownloadFileAsync(StorageDownloadPayload payload)
    {
        return await base.GetFileAsync("/v1/storage/download/" + payload.FileName + payload.GetQueryParams());
    }
}
