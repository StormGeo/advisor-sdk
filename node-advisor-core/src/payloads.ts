
export interface WeatherPayload {
  stationId?: string
  localeId?: number
  latitude?: number
  longitude?: number
  timezone?: number
  variables?: string[]
  startDate?: string
  endDate?: string
}

export interface StationPayload {
  stationId?: string
  layer?: string
  variables?: string[]
  startDate?: string
  endDate?: string
}

export interface ClimatologyPayload {
  stationId?: string
  localeId?: number
  latitude?: number
  longitude?: number
  variables?: string[]
}

export interface CurrentWeatherPayload {
  localeId?: number
  latitude?: number
  longitude?: number
  timezone?: number
  variables?: string[]
}

export interface RadiusPayload {
  localeId?: number
  latitude?: number
  longitude?: number
  startDate?: string
  endDate?: string
  radius?: number
}

export interface GeometryPayload {
  startDate?: string
  endDate?: string
  radius?: number
  geometry?: string
}

export interface TmsPayload {
  server?: string
  mode?: string
  variable?: string
  aggregation?: string
  x?: string
  y?: string
  z?: string
  istep?: string
  fstep?: string
}

export interface ApiResponse {
  data?: any
  error?: any
}

export interface ApiImgResponse {
  data?: ArrayBuffer | null
  error?: any
}

export interface AdvisorCoreConfig {
  token: string
  retries?: number
  delay?: number
}
