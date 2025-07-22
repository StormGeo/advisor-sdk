import { Readable } from "stream"

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
  stationId?: string
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

export interface StaticMapPayload {
  startDate?: string
  endDate?: string
  aggregation?: string
  model?: string
  lonmin?: number
  lonmax?: number
  latmin?: number
  latmax?: number
  dpi?: number
  title?: boolean
  titlevariable?: string
  hours?: number
  type?: string
  category?: string
  variable?: string
}

export interface TmsPayload {
  server?: string
  mode?: string
  variable?: string
  aggregation?: string
  x?: string
  y?: string
  z?: string
  timezone?: number
  istep?: string
  fstep?: string
}

export interface PlanInfoPayload {
  timezone?: number
}

export interface RequestDetailsPayload {
  page?: number
  pageSize?: number
  path?: string
  status?: number
  startDate?: string
  endDate?: string
}

export interface StorageListPayload {
  page?: number
  pageSize?: number
  startDate?: string
  endDate?: string
  fileName?: string
  fileExtension?: string
}

export interface StorageDownloadPayload {
  fileName?: string
  accessKey?: string
}

export interface ApiResponse {
  data?: any
  error?: any
}

export interface ApiFileResponse {
  data?: Buffer | null
  error?: any
}

export interface ApiStreamResponse {
  data?: Readable | null
  error?: any
}

export interface AdvisorCoreConfig {
  token: string
  retries?: number
  delay?: number
  headers?: Record<string, string>
}
