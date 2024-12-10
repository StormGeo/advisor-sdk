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
  ObservedByGeometryPayload,
  SpecificObservedPayload,
  TMSPayload,
  WeatherPayload,
} from "./payloads"

function sleep(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms))
}

export class AdvisorCore {
  private readonly baseURL: string
  private readonly token: string
  private readonly retries: number
  private readonly delay: number
  private readonly client: AxiosInstance

  constructor({ token, retries = 5, delay = 1000 }: AdvisorCoreConfig) {
    if (!token) {
      throw new Error("Token is required.")
    }
    if (typeof retries !== "number" || retries < 0) {
      throw new Error("Retries must be a non-negative number.")
    }
    if (typeof delay !== "number" || delay < 0) {
      throw new Error("Delay must be a non-negative number.")
    }

    this.baseURL = "https://advisor-core.climatempo.io/api/v1/"
    this.token = token
    this.retries = retries
    this.delay = delay
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
        ...(method === "GET" ? { params } : { params, data }),
      })

      return { data: response.data, error: null }
    } catch (error: any) {
      if (retries > 0 && error?.response?.status >= 500) {
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
        ...(method === "GET" ? { params } : { params, data }),
        responseType: 'arraybuffer',
      })

      const byteArray  = new Uint8Array(response.data)
      return { data: byteArray, error: null }
    } catch (error: any) {
      if (retries > 0 && error?.response?.status >= 500) {
        await sleep(this.delay)
        console.log(`Re-trying in ${this.delay}ms... attempts left: ${retries}`)

        return this.makeRequestImage(method, url, params, data, --retries)
      }

      return { data: null, error: error?.response?.data ?? error }
    }
  }

  chart: ChartRoutes = {
    getForecastDaily: async (payload: WeatherPayload): Promise<ApiImgResponse> => {
      return this.makeRequestImage("GET", "forecast/daily/chart", { ...payload, token: this.token })
    },
    getForecastHourly: async (payload: WeatherPayload): Promise<ApiImgResponse> => {
      return this.makeRequestImage("GET", "forecast/hourly/chart", { ...payload, token: this.token })
    },
    getObservedDaily: async (payload: WeatherPayload): Promise<ApiImgResponse> => {
      return this.makeRequestImage("GET", "forecast/daily/chart", { ...payload, token: this.token })
    },
    getObservedHourly: async (payload: WeatherPayload): Promise<ApiImgResponse> => {
      return this.makeRequestImage("GET", "forecast/hourly/chart", { ...payload, token: this.token })
    },
  }

  forecast: ForecastRoutes = {
    getDaily: async (payload: WeatherPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "forecast/daily", { ...payload, token: this.token })
    },
    getHourly: async (payload: WeatherPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "forecast/hourly", { ...payload, token: this.token })
    },
    getPeriod: async (payload: WeatherPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "forecast/period", { ...payload, token: this.token })
    },
  }

  observed: ObservedRoutes = {
    getDaily: async (payload: WeatherPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "forecast/daily", { ...payload, token: this.token })
    },
    getHourly: async (payload: WeatherPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "forecast/hourly", { ...payload, token: this.token })
    },
    getPeriod: async (payload: WeatherPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "forecast/period", { ...payload, token: this.token })
    },
    getStationData: async (payload: WeatherPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "station", { ...payload, token: this.token })
    },
    getLightning: async (payload: SpecificObservedPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "observed/lightning", { ...payload, token: this.token })
    },
    postLightning: async (payload: ObservedByGeometryPayload): Promise<ApiResponse> => {
      const { geometry, ...restData } = payload
      return this.makeRequest("POST", "observed/lightning", { ...restData, token: this.token }, { geometry })
    },
    getFireFocus: async (payload: SpecificObservedPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "observed/fire-focus", { ...payload, token: this.token })
    },
    postFireFocus: async (payload: ObservedByGeometryPayload): Promise<ApiResponse> => {
      const { geometry, ...restData } = payload
      return this.makeRequest("POST", "observed/fire-focus", { ...restData, token: this.token }, { geometry })
    },
  }

  climatology: ClimatologyRoutes = {
    getDaily: async (payload: ClimatologyPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "climatology/daily", { ...payload, token: this.token })
    },
    getMonthly: async (payload: ClimatologyPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "climatology/monthly", { ...payload, token: this.token })
    },
  }

  currentWeather: CurrentWeatherRoutes = {
    get: async (payload: CurrentWeatherPayload): Promise<ApiResponse> => {
      return this.makeRequest("GET", "current-weather", { ...payload, token: this.token })
    },
  }

  monitoring: MonitoringRoutes = {
    getAlerts: async (): Promise<ApiResponse> => {
      return this.makeRequest("GET", "monitoring/alerts", { token: this.token })
    },
  }

  plan: PlanRoutes = {
    getInfo: async (): Promise<ApiResponse> => {
      return this.makeRequest("GET", `plan/${this.token}`)
    },
  }

  schema: SchemaRoutes = {
    getDefinition: async (): Promise<ApiResponse> => {
      return this.makeRequest("GET", "schema/definition", { token: this.token })
    },
    postDefinition: async (payload: any): Promise<ApiResponse> => {
      return this.makeRequest("POST", "schema/definition", { ...payload, token: this.token }, payload)
    },
    postParameters: async (payload: any): Promise<ApiResponse> => {
      return this.makeRequest("POST", "schema/parameters", { ...payload, token: this.token }, payload)
    },
  }

  tms: TmsRoutes = {
    get: async (payload: TMSPayload): Promise<ApiImgResponse> => {
      const path = `/tms/${payload.server}/${payload.mode}/${payload.variable}/${payload.aggregation}/${payload.x}/${payload.y}/${payload.z}.png`
      return this.makeRequestImage("GET", path, { istep: payload.istep, fstep: payload.fstep, token: this.token })
    },
  }
}

export default AdvisorCore
