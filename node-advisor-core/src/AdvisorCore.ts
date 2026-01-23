import { Readable } from "stream"
import axios, {
  AxiosInstance,
  AxiosRequestConfig,
} from "axios"
import {
  ChartRoutes,
  ClimatologyRoutes,
  CurrentWeatherRoutes,
  ForecastRoutes,
  MonitoringRoutes,
  ObservedRoutes,
  PlanRoutes,
  SchemaRoutes,
  StaticMapRoutes,
  StorageRoutes,
  TmsRoutes,
} from "./interfaces"
import {
  AdvisorCoreConfig,
  ApiFileResponse,
  ApiResponse,
  ClimatologyPayload,
  CurrentWeatherPayload,
  StationPayload,
  GeometryPayload,
  RadiusPayload,
  TmsPayload,
  WeatherPayload,
  RequestDetailsPayload,
  StorageListPayload,
  StorageDownloadPayload,
  ApiStreamResponse,
  PlanInfoPayload,
  PlanLocalePayload,
  StaticMapPayload,
} from "./payloads"

function sleep(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms))
}

/**
 * @typedef {Object} WeatherPayload
 * @property {string} stationId
 * @property {number} localeId
 * @property {number} latitude
 * @property {number} longitude
 * @property {number} timezone
 * @property {Array<string>} variables
 * @property {string} startDate
 * @property {string} endDate
 */

/**
 * @typedef {Object} StationPayload
 * @property {string} stationId
 * @property {string} layer
 * @property {number} timezone
 * @property {Array<string>} variables
 * @property {string} startDate
 * @property {string} endDate
 */

/**
 * @typedef {Object} ClimatologyPayload
 * @property {string} stationId
 * @property {number} localeId
 * @property {number} latitude
 * @property {number} longitude
 * @property {Array<string>} variables
 */

/**
 * @typedef {Object} CurrentWeatherPayload
 * @property {number} localeId
 * @property {number} latitude
 * @property {number} longitude
 * @property {number} timezone
 * @property {Array<string>} variables
 */

/**
 * @typedef {Object} RadiusPayload
 * @property {number} localeId
 * @property {number} latitude
 * @property {number} longitude
 * @property {string} startDate
 * @property {string} endDate
 * @property {number} radius
 */

/**
 * @typedef {Object} GeometryPayload
 * @property {string} startDate
 * @property {string} endDate
 * @property {number} radius
 * @property {string} geometry
 */

/**
 * @typedef {Object} PlanInfoPayload
 * @property {number} timezone
 */

/**
 * @typedef {Object} RequestDetailsPayload
 * @property {number} page
 * @property {number} pageSize
 * @property {string} path
 * @property {number} status
 * @property {string} startDate
 * @property {string} endDate
*/

/**
 * @typedef {Object} StorageListPayload
 * @property {number} page
 * @property {number} pageSize
 * @property {string} startDate
 * @property {string} endDate
 * @property {string} fileName
 * @property {string} fileExtension
 * @property {Array<string>} fileTypes
*/

/**
 * @typedef {Object} StorageDownloadPayload
 * @property {string} fileName
 * @property {string} accessKey
*/

/**
 * @typedef {Object} TmsPayload
 * @property {string} server
 * @property {string} mode
 * @property {string} variable
 * @property {string} aggregation
 * @property {string} x
 * @property {string} y
 * @property {string} z
 * @property {string} istep
 * @property {string} fstep
 */

export class AdvisorCore {
  private readonly baseURL: string
  private readonly token: string
  private readonly retries: number
  private readonly delay: number
  private readonly client: AxiosInstance
  private headers: Record<string, string>

  constructor({ token, retries = 5, delay = 5 }: AdvisorCoreConfig) {
    if (!token) {
      throw new Error("Token is required.")
    }
    if (typeof retries !== "number" || retries < 0) {
      throw new Error("Retries must be a non-negative number.")
    }
    if (typeof delay !== "number" || delay < 0) {
      throw new Error("Delay must be a non-negative number.")
    }

    this.baseURL = "https://advisor-core.climatempo.io/api/"
    this.token = token
    this.retries = retries
    this.delay = delay * 1000
    this.headers = {
      'Content-Type': 'application/json',
      Accept: "application/json",
      "Accept-Language": "en-US",
      "x-advisor-token": this.token,
      "User-Agent": "Nodejs-AdvisorCore-SDK",
    }
    this.client = axios.create({
      baseURL: this.baseURL,
    })
  }

  private async makeRequest(
    method: AxiosRequestConfig["method"],
    url: string,
    params: Record<string, any> = {},
    data: any = null,
    retries: number = this.retries
  ): Promise<ApiResponse> {
    try {
      const response = await this.client({
        method,
        url,
        headers: this.headers,
        ...(method === "GET" ? { params } : { params, data }),
      })

      return { data: response.data, error: null }
    } catch (error: any) {
      if (retries > 0 && (error?.response?.status >= 500)) {
        await sleep(this.delay)
        return this.makeRequest(method, url, params, data, --retries)
      }

      return { data: null, error: error?.response?.data ?? error }
    }
  }

  private async makeRequestFile(
    method: AxiosRequestConfig["method"],
    url: string,
    params: Record<string, any> = {},
    data: any = null,
    retries: number = this.retries
  ): Promise<ApiFileResponse> {
    try {
      const response = await this.client({
        method,
        url,
        headers: this.headers,
        ...(method === "GET" ? { params } : { params, data }),
        responseType: 'arraybuffer',
      })

      const byteArray  = Buffer.from(response.data)
      return { data: byteArray, error: null }
    } catch (error: any) {
      if (retries > 0 && (error?.response?.status >= 500)) {
        await sleep(this.delay)

        return this.makeRequestFile(method, url, params, data, --retries)
      }

      return { data: null, error: error?.response?.data ?? error }
    }
  }

  private async makeRequestFileByStream(
    method: AxiosRequestConfig["method"],
    url: string,
    params: Record<string, any> = {},
    data: any = null,
    retries: number = this.retries
  ): Promise<ApiStreamResponse> {
    try {
      const response = await this.client({
        method,
        url,
        headers: this.headers,
        ...(method === "GET" ? { params } : { params, data }),
        responseType: 'stream',
      })

      return { data: response.data, error: null }
    } catch (error: any) {
      if (retries > 0 && (error?.response?.status >= 500)) {
        await sleep(this.delay)

        return this.makeRequestFileByStream(method, url, params, data, --retries)
      }

      return {
        data: null,
        error: error?.response?.data ? await this.getAllStreamContent(error.response.data) : error,
      }
    }
  }

  private async getAllStreamContent(stream: Readable): Promise<string> {
    let content = ''

    for await (const chunk of stream) {
      content += chunk.toString('utf8')
    }

    return JSON.parse(content)
  }

  setHeaderAccept(value: string): void {
    this.headers.Accept = value
  }

  setHeaderAcceptLanguage(value: string): void {
    this.headers["Accept-Language"] = value
  }

  /**
   * Fetch weather data charts.
   */
  chart: ChartRoutes = {
    /**
     * Fetch daily weather forecast chart.
     * GET /v1/forecast/daily/chart
     * @param {WeatherPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getForecastDaily: async (payload: WeatherPayload): Promise<ApiFileResponse> => {
      return this.makeRequestFile("GET", "v1/forecast/daily/chart", payload)
    },
    /**
     * Fetch hourly weather forecast chart.
     * GET /v1/forecast/hourly/chart
     * @param {WeatherPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getForecastHourly: async (payload: WeatherPayload): Promise<ApiFileResponse> => {
      return this.makeRequestFile("GET", "v1/forecast/hourly/chart", payload)
    },
    /**
     * Fetch daily observed weather chart.
     * GET /v1/observed/daily/chart
     * @param {WeatherPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getObservedDaily: async (payload: WeatherPayload): Promise<ApiFileResponse> => {
      return this.makeRequestFile("GET", "v1/forecast/daily/chart", payload)
    },
    /**
     * Fetch hourly observed weather chart.
     * GET /v1/observed/hourly/chart
     * @param {WeatherPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getObservedHourly: async (payload: WeatherPayload): Promise<ApiFileResponse> => {
      return this.makeRequestFile("GET", "v1/forecast/hourly/chart", payload)
    },
  }

  /**
   * Fetch weather forecast.
   */
  forecast: ForecastRoutes = {
    /**
     * Fetch daily weather forecast.
     * GET /v1/forecast/daily
     * @param {WeatherPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getDaily: async (payload: WeatherPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/forecast/daily", payload)
    },
    /**
     * Fetch hourly weather forecast.
     * GET /v1/forecast/hourly
     * @param {WeatherPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getHourly: async (payload: WeatherPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/forecast/hourly", payload)
    },
    /**
     * Fetch period weather forecast.
     * GET /v1/forecast/period
     * @param {WeatherPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getPeriod: async (payload: WeatherPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/forecast/period", payload)
    },
  }

  /**
   * Fetch observed weather.
   */
  observed: ObservedRoutes = {
    /**
     * Fetch daily observed weather.
     * GET /v1/observed/daily
     * @param {WeatherPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getDaily: async (payload: WeatherPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/forecast/daily", payload)
    },
    /**
     * Fetch hourly observed weather.
     * GET /v1/observed/hourly
     * @param {WeatherPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getHourly: async (payload: WeatherPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/forecast/hourly", payload)
    },
    /**
     * Fetch period observed weather.
     * GET /v1/observed/period
     * @param {WeatherPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getPeriod: async (payload: WeatherPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/forecast/period", payload)
    },
    /**
     * Fetch station observed data.
     * GET /v1/station
     * @param {StationPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getStationData: async (payload: StationPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/station", payload)
    },
    /**
     * Fetch observed lightning.
     * GET /v1/observed/lightning
     * @param {RadiusPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getLightning: async (payload: RadiusPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/observed/lightning", payload)
    },
    /**
     * Fetch observed lightning by geometry.
     * POST /v1/observed/lightning
     * @param {GeometryPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getLightningByGeometry: async (payload: GeometryPayload): Promise<ApiResponse> => {
      const { geometry, ...restData } = payload
      return this.makeRequest("POST", "v1/observed/lightning", restData, { geometry })
    },
    /**
     * Fetch observed fire focus bu geometry.
     * GET /v1/observed/fire-focus
     * @param {GeometryPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getFireFocus: async (payload: RadiusPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/observed/fire-focus", payload)
    },
    /**
     * Fetch observed fire focus bu geometry.
     * POST /v1/observed/fire-focus
     * @param {GeometryPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getFireFocusByGeometry: async (payload: GeometryPayload): Promise<ApiResponse> => {
      const { geometry, ...restData } = payload
      return this.makeRequest("POST", "v1/observed/fire-focus", restData, { geometry })
    },
  }

  /**
   * Fetch climatology weather.
   */
  climatology: ClimatologyRoutes = {
    /**
     * Fetch daily climatology weather.
     * GET /v1/climatology/daily
     * @param {ClimatologyPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getDaily: async (payload: ClimatologyPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/climatology/daily", payload)
    },
    /**
     * Fetch monthly climatology weather.
     * GET /v1/climatology/monthly
     * @param {ClimatologyPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getMonthly: async (payload: ClimatologyPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/climatology/monthly", payload)
    },
  }

  /**
   * Fetch current weather.
   */
  currentWeather: CurrentWeatherRoutes = {
    /**
     * Fetch current weather.
     * GET /v1/current-weather
     * @param {CurrentWeatherPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    get: async (payload: CurrentWeatherPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/current-weather", payload)
    },
  }

  /**
   * Fetch alerts.
   */
  monitoring: MonitoringRoutes = {
    /**
     * Fetch alerts.
     * GET /v1/monitoring/alerts
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getAlerts: async (): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/monitoring/alerts")
    },
  }

  /**
   * Fetch plan information.
   */
  plan: PlanRoutes = {
    /**
     * Fetch plan information.
     * GET /v2/plan
     * @param {PlanInfoPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getInfo: async (payload: PlanInfoPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v2/plan", payload)
    },
    /**
     * Get request history of a plan
     * GET /v1/plan/request-details
     * @param {RequestDetailsPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getRequestDetails: async (payload: RequestDetailsPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", 'v1/plan/request-details', payload)
    },
    /**
     * Fetch locale information associated with a plan.
     * GET /v1/plan/locale
     * @param {PlanLocalePayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getLocale: async (payload: PlanLocalePayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", 'v1/plan/locale', payload)
    }
  }

  /**
   * Get and set schema/parameters.
   */
  schema: SchemaRoutes = {
    /**
     * Fetch schema definition.
     * GET /v1/schema/definition
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getDefinition: async (): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/schema/definition")
    },
    /**
     * Set schema definition.
     * POST /v1/schema/definition
     * @param {Object} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    postDefinition: async (payload: any): Promise<ApiResponse> => {
      return this.makeRequest("POST", "v1/schema/definition", {}, payload)
    },
    /**
     * Post schema parameters.
     * POST /v1/schema/parameters
     * @param {Object} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    postParameters: async (payload: any): Promise<ApiResponse> => {
      return this.makeRequest("POST", "v1/schema/parameters", {}, payload)
    },
  }

  /**
   * Fetch static map images.
   */
  staticMap: StaticMapRoutes = {
    /**
     * Fetch static map images.
     * GET /v1/map
     * @param {StaticMapPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    get: async (payload: StaticMapPayload): Promise<ApiResponse> => {
      const {type, category, variable, ...rest} = payload
      return this.makeRequestFile("GET", `v1/map/${type}/${category}/${variable}`, rest)
    },
  }

  /**
   * Fetch tiles map service.
   */
  tms: TmsRoutes = {
    /**
     * Fetch daily weather forecast.
     * GET /v1/tms/{server}/{mode}/{variable}/{aggregation}/{x}/{y}/{z}.png
     * @param {TmsPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    get: async (payload: TmsPayload): Promise<ApiFileResponse> => {
      const path = `v1/tms/${payload.server}/${payload.mode}/${payload.variable}/${payload.aggregation}/${payload.x}/${payload.y}/${payload.z}.png`
      return this.makeRequestFile("GET", path, { timezone: payload.timezone, istep: payload.istep, fstep: payload.fstep })
    },
  }

  /**
   * List and Download bucket files.
   */
  storage: StorageRoutes = {
    /**
     * List bucket files.
     * GET /v1/storage/list
     * @param {StorageListPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    listFiles: async (payload: StorageListPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", '/v1/storage/list', payload)
    },
    /**
     * Download a file.
     * GET /v1/storage/download/{fileName}
     * @param {StorageDownloadPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    downloadFile: async ({ fileName, ...rest }: StorageDownloadPayload): Promise<ApiFileResponse> => {
      return this.makeRequestFile("GET", `/v1/storage/download/${fileName}`, rest)
    },
    /**
     * Download a file by stream.
     * GET /v1/storage/download/{fileName}
     * @param {StorageDownloadPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    downloadFileByStream: async ({ fileName, ...rest }: StorageDownloadPayload): Promise<ApiStreamResponse> => {
      return this.makeRequestFileByStream("GET", `/v1/storage/download/${fileName}`, rest)
    },
  }
}

export default AdvisorCore
