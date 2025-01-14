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
  TmsRoutes,
} from "./interfaces"
import {
  AdvisorCoreConfig,
  ApiImgResponse,
  ApiResponse,
  ClimatologyPayload,
  CurrentWeatherPayload,
  StationPayload,
  GeometryPayload,
  RadiusPayload,
  TmsPayload,
  WeatherPayload,
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

  constructor({ token, retries = 5, delay = 5}: AdvisorCoreConfig) {
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
      Accept: "application/json",
      "Accept-Language": "en-US",
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
      if (retries > 0 && (error?.response?.status >= 500 || error?.response?.status == 429)) {
        await sleep(this.delay)
        console.log(`Re-trying in ${this.delay}ms... attempts left: ${retries}`)
        return this.makeRequest(method, url, params, data, --retries)
      }

      return { data: null, error: error?.response?.data ?? error }
    }
  }

  private async makeRequestImage(
    method: AxiosRequestConfig["method"],
    url: string,
    params: Record<string, any> = {},
    data: any = null,
    retries: number = this.retries
  ): Promise<ApiImgResponse> {
    try {
      const response = await this.client({
        method,
        url,
        headers: this.headers,
        ...(method === "GET" ? { params } : { params, data }),
        responseType: 'arraybuffer',
      })

      const byteArray  = new Uint8Array(response.data)
      return { data: byteArray, error: null }
    } catch (error: any) {
      if (retries > 0 && (error?.response?.status >= 500 || error?.response?.status == 429)) {
        await sleep(this.delay)
        console.log(`Re-trying in ${this.delay}ms... attempts left: ${retries}`)

        return this.makeRequestImage(method, url, params, data, --retries)
      }

      return { data: null, error: error?.response?.data ?? error }
    }
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
    getForecastDaily: async (payload: WeatherPayload): Promise<ApiImgResponse> => {
      return this.makeRequestImage("GET", "v1/forecast/daily/chart", { ...payload, token: this.token })
    },
    /**
     * Fetch hourly weather forecast chart.
     * GET /v1/forecast/hourly/chart
     * @param {WeatherPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getForecastHourly: async (payload: WeatherPayload): Promise<ApiImgResponse> => {
      return this.makeRequestImage("GET", "v1/forecast/hourly/chart", { ...payload, token: this.token })
    },
    /**
     * Fetch daily observed weather chart.
     * GET /v1/observed/daily/chart
     * @param {WeatherPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getObservedDaily: async (payload: WeatherPayload): Promise<ApiImgResponse> => {
      return this.makeRequestImage("GET", "v1/forecast/daily/chart", { ...payload, token: this.token })
    },
    /**
     * Fetch hourly observed weather chart.
     * GET /v1/observed/hourly/chart
     * @param {WeatherPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getObservedHourly: async (payload: WeatherPayload): Promise<ApiImgResponse> => {
      return this.makeRequestImage("GET", "v1/forecast/hourly/chart", { ...payload, token: this.token })
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
      return this.makeRequest("GET", "v1/forecast/daily", { ...payload, token: this.token })
    },
    /**
     * Fetch hourly weather forecast.
     * GET /v1/forecast/hourly
     * @param {WeatherPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getHourly: async (payload: WeatherPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/forecast/hourly", { ...payload, token: this.token })
    },
    /**
     * Fetch period weather forecast.
     * GET /v1/forecast/period
     * @param {WeatherPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getPeriod: async (payload: WeatherPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/forecast/period", { ...payload, token: this.token })
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
      return this.makeRequest("GET", "v1/forecast/daily", { ...payload, token: this.token })
    },
    /**
     * Fetch hourly observed weather.
     * GET /v1/observed/hourly
     * @param {WeatherPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getHourly: async (payload: WeatherPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/forecast/hourly", { ...payload, token: this.token })
    },
    /**
     * Fetch period observed weather.
     * GET /v1/observed/period
     * @param {WeatherPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getPeriod: async (payload: WeatherPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/forecast/period", { ...payload, token: this.token })
    },
    /**
     * Fetch station observed data.
     * GET /v1/station
     * @param {StationPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getStationData: async (payload: StationPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/station", { ...payload, token: this.token })
    },
    /**
     * Fetch observed lightning.
     * GET /v1/observed/lightning
     * @param {RadiusPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getLightning: async (payload: RadiusPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/observed/lightning", { ...payload, token: this.token })
    },
    /**
     * Fetch observed lightning by geometry.
     * POST /v1/observed/lightning
     * @param {GeometryPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getLightningByGeometry: async (payload: GeometryPayload): Promise<ApiResponse> => {
      const { geometry, ...restData } = payload
      return this.makeRequest("POST", "v1/observed/lightning", { ...restData, token: this.token }, { geometry })
    },
    /**
     * Fetch observed fire focus bu geometry.
     * GET /v1/observed/fire-focus
     * @param {GeometryPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getFireFocus: async (payload: RadiusPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/observed/fire-focus", { ...payload, token: this.token })
    },
    /**
     * Fetch observed fire focus bu geometry.
     * POST /v1/observed/fire-focus
     * @param {GeometryPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getFireFocusByGeometry: async (payload: GeometryPayload): Promise<ApiResponse> => {
      const { geometry, ...restData } = payload
      return this.makeRequest("POST", "v1/observed/fire-focus", { ...restData, token: this.token }, { geometry })
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
      return this.makeRequest("GET", "v1/climatology/daily", { ...payload, token: this.token })
    },
    /**
     * Fetch monthly climatology weather.
     * GET /v1/climatology/monthly
     * @param {ClimatologyPayload} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getMonthly: async (payload: ClimatologyPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "v1/climatology/monthly", { ...payload, token: this.token })
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
      return this.makeRequest("GET", "v1/current-weather", { ...payload, token: this.token })
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
      return this.makeRequest("GET", "v1/monitoring/alerts", { token: this.token })
    },
  }

  /**
   * Fetch plan information.
   */
  plan: PlanRoutes = {
    /**
     * Fetch plan information.
     * GET /v1/plan/{token}
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    getInfo: async (): Promise<ApiResponse> => {
      return this.makeRequest("GET", `v1/plan/${this.token}`)
    },
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
      return this.makeRequest("GET", "v1/schema/definition", { token: this.token })
    },
    /**
     * Set schema definition.
     * POST /v1/schema/definition
     * @param {Object} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    postDefinition: async (payload: any): Promise<ApiResponse> => {
      return this.makeRequest("POST", "v1/schema/definition", { ...payload, token: this.token }, payload)
    },
    /**
     * Post schema parameters.
     * POST /v1/schema/parameters
     * @param {Object} payload
     * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
     */
    postParameters: async (payload: any): Promise<ApiResponse> => {
      return this.makeRequest("POST", "v1/schema/parameters", { ...payload, token: this.token }, payload)
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
    get: async (payload: TmsPayload): Promise<ApiImgResponse> => {
      const path = `v1/tms/${payload.server}/${payload.mode}/${payload.variable}/${payload.aggregation}/${payload.x}/${payload.y}/${payload.z}.png`
      return this.makeRequestImage("GET", path, { istep: payload.istep, fstep: payload.fstep, token: this.token })
    },
  }
}

export default AdvisorCore
