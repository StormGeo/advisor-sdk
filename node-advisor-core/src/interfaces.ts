import {
  ApiResponse,
  ClimatologyPayload,
  CurrentWeatherPayload,
  ObservedByGeometryPayload,
  SpecificObservedPayload,
  TMSPayload,
  WeatherPayload
} from "./payloads"

export interface ChartRoutes {
  getForecastDaily: (payload: WeatherPayload) => Promise<ApiResponse>
  getForecastHourly: (payload: WeatherPayload) => Promise<ApiResponse>
  getObservedDaily: (payload: WeatherPayload) => Promise<ApiResponse>
  getObservedHourly: (payload: WeatherPayload) => Promise<ApiResponse>
}

export interface ForecastRoutes {
  getDaily: (payload: WeatherPayload) => Promise<ApiResponse>
  getHourly: (payload: WeatherPayload) => Promise<ApiResponse>
  getPeriod: (payload: WeatherPayload) => Promise<ApiResponse>
}

export interface ObservedRoutes {
  getDaily: (payload: WeatherPayload) => Promise<ApiResponse>
  getHourly: (payload: WeatherPayload) => Promise<ApiResponse>
  getPeriod: (payload: WeatherPayload) => Promise<ApiResponse>
  getStationData: (payload: WeatherPayload) => Promise<ApiResponse>
  getLightning: (payload: SpecificObservedPayload) => Promise<ApiResponse>
  postLightning: (payload: ObservedByGeometryPayload) => Promise<ApiResponse>
  getFireFocus: (payload: SpecificObservedPayload) => Promise<ApiResponse>
  postFireFocus: (payload: ObservedByGeometryPayload) => Promise<ApiResponse>
}

export interface ClimatologyRoutes {
  getDaily: (payload: ClimatologyPayload) => Promise<ApiResponse>
  getMonthly: (payload: ClimatologyPayload) => Promise<ApiResponse>  
}

export interface CurrentWeatherRoutes {
  get: (payload: CurrentWeatherPayload) => Promise<ApiResponse>  
}

export interface MonitoringRoutes {
  getAlerts: () => Promise<ApiResponse>
}

export interface PlanRoutes {
  getInfo: () => Promise<ApiResponse>
}

export interface SchemaRoutes {
  getDefinition: () => Promise<ApiResponse>
  postDefinition: (payload: any) => Promise<ApiResponse>
  postParameters: (payload: any) => Promise<ApiResponse>
}

export interface TmsRoutes {
  get: (payload: TMSPayload) => Promise<ApiResponse>
}