import {
  ApiResponse,
  ClimatologyPayload,
  CurrentWeatherPayload,
  GeometryPayload,
  StationPayload,
  RadiusPayload,
  TmsPayload,
  WeatherPayload
} from "./payloads"
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

export interface ChartRoutes {
  /**
   * Fetch daily weather forecast chart.
   * GET /v1/forecast/daily/chart
   * @param {WeatherPayload} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  getForecastDaily: (payload: WeatherPayload) => Promise<ApiResponse>
  /**
   * Fetch hourly weather forecast chart.
   * GET /v1/forecast/hourly/chart
   * @param {WeatherPayload} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  getForecastHourly: (payload: WeatherPayload) => Promise<ApiResponse>
  /**
   * Fetch daily observed weather chart.
   * GET /v1/observed/daily/chart
   * @param {WeatherPayload} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  getObservedDaily: (payload: WeatherPayload) => Promise<ApiResponse>
  /**
   * Fetch hourly observed weather chart.
   * GET /v1/observed/hourly/chart
   * @param {WeatherPayload} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  getObservedHourly: (payload: WeatherPayload) => Promise<ApiResponse>
}

export interface ForecastRoutes {
  /**
   * Fetch daily weather forecast.
   * GET /v1/forecast/daily
   * @param {WeatherPayload} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  getDaily: (payload: WeatherPayload) => Promise<ApiResponse>
  /**
   * Fetch hourly weather forecast.
   * GET /v1/forecast/hourly
   * @param {WeatherPayload} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  getHourly: (payload: WeatherPayload) => Promise<ApiResponse>
  /**
   * Fetch period weather forecast.
   * GET /v1/forecast/period
   * @param {WeatherPayload} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  getPeriod: (payload: WeatherPayload) => Promise<ApiResponse>
}

export interface ObservedRoutes {
  /**
   * Fetch daily observed weather.
   * GET /v1/observed/daily
   * @param {WeatherPayload} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  getDaily: (payload: WeatherPayload) => Promise<ApiResponse>
  /**
   * Fetch hourly observed weather.
   * GET /v1/observed/hourly
   * @param {WeatherPayload} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  getHourly: (payload: WeatherPayload) => Promise<ApiResponse>
  /**
   * Fetch period observed weather.
   * GET /v1/observed/period
   * @param {WeatherPayload} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  getPeriod: (payload: WeatherPayload) => Promise<ApiResponse>
  /**
   * Fetch station observed data.
   * GET /v1/station
   * @param {StationPayload} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  getStationData: (payload: StationPayload) => Promise<ApiResponse>
  /**
   * Fetch observed lightning.
   * GET /v1/observed/lightning
   * @param {RadiusPayload} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  getLightning: (payload: RadiusPayload) => Promise<ApiResponse>
  /**
   * Fetch observed lightning by geometry.
   * POST /v1/observed/lightning
   * @param {GeometryPayload} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  getLightningByGeometry: (payload: GeometryPayload) => Promise<ApiResponse>
  /**
   * Fetch observed fire focus bu geometry.
   * GET /v1/observed/fire-focus
   * @param {GeometryPayload} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  getFireFocus: (payload: RadiusPayload) => Promise<ApiResponse>
  /**
   * Fetch observed fire focus bu geometry.
   * POST /v1/observed/fire-focus
   * @param {GeometryPayload} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  getFireFocusByGeometry: (payload: GeometryPayload) => Promise<ApiResponse>
}

export interface ClimatologyRoutes {
  /**
   * Fetch daily climatology weather.
   * GET /v1/climatology/daily
   * @param {ClimatologyPayload} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  getDaily: (payload: ClimatologyPayload) => Promise<ApiResponse>
  /**
   * Fetch monthly climatology weather.
   * GET /v1/climatology/monthly
   * @param {ClimatologyPayload} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  getMonthly: (payload: ClimatologyPayload) => Promise<ApiResponse>  
}

export interface CurrentWeatherRoutes {
  /**
   * Fetch current weather.
   * GET /v1/current-weather
   * @param {CurrentWeatherPayload} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  get: (payload: CurrentWeatherPayload) => Promise<ApiResponse>  
}

export interface MonitoringRoutes {
  /**
   * Fetch alerts.
   * GET /v1/monitoring/alerts
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  getAlerts: () => Promise<ApiResponse>
}

export interface PlanRoutes {
  /**
   * Fetch plan information.
   * GET /v1/plan/{token}
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  getInfo: () => Promise<ApiResponse>
}

export interface SchemaRoutes {
  /**
   * Fetch schema definition.
   * GET /v1/schema/definition
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  getDefinition: () => Promise<ApiResponse>
  /**
   * Set schema definition.
   * POST /v1/schema/definition
   * @param {Object} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  postDefinition: (payload: any) => Promise<ApiResponse>
  /**
   * Post schema parameters.
   * POST /v1/schema/parameters
   * @param {Object} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  postParameters: (payload: any) => Promise<ApiResponse>
}

export interface TmsRoutes {
  /**
   * Fetch daily weather forecast.
   * GET /v1/tms/{server}/{mode}/{variable}/{aggregation}/{x}/{y}/{z}.png
   * @param {TmsPayload} payload
   * @returns {Promise<{data: Object|null, error: Object|null}>} API response.
   */
  get: (payload: TmsPayload) => Promise<ApiResponse>
}