package advisorsdk

import "io"

type ImageRequestWithWeatherPayload func(WeatherPayload) (io.ReadCloser, error)
type RequestWithWeatherPayload func(WeatherPayload) (AdvisorResponse, error)
type RequestWithClimatologyPayload func(ClimatologyPayload) (AdvisorResponse, error)
type RequestWithCurrentWeatherPayload func(CurrentWeatherPayload) (AdvisorResponse, error)
type RequestWithGeometryPayload func(GeometryPayload) (AdvisorResponse, error)
type RequestWithSchemaPayload func(SchemaPayload) (AdvisorResponse, error)
type RequestWithStationPayload func(StationPayload) (AdvisorResponse, error)
type RequestWithRadiusPayload func(RadiusPayload) (AdvisorResponse, error)
type RequestWithPayload func(RadiusPayload) (AdvisorResponse, error)
type TmsRequest func(TmsPayload) (io.ReadCloser, error)
type RequestWithPlanInfoPayload func(PlanInfoPayload) (AdvisorResponse, error)
type RequestWithRequestDetailsPayload func(RequestDetailsPayload) (AdvisorResponse, error)
type RequestWithStorageDownloadPayload func(StorageDownloadPayload) (io.ReadCloser, error)
type RequestWithStorageListPayload func(StorageListPayload) (AdvisorResponse, error)
type ImageRequestWithStaticMapPayload func(StaticMapPayload) (io.ReadCloser, error)

type chart struct {
	GetForecastDaily  ImageRequestWithWeatherPayload
	GetForecastHourly ImageRequestWithWeatherPayload
	GetObservedDaily  ImageRequestWithWeatherPayload
	GetObservedHourly ImageRequestWithWeatherPayload
}

type climatology struct {
	GetDaily   RequestWithClimatologyPayload
	GetMonthly RequestWithClimatologyPayload
}

type currentWeather struct {
	Get RequestWithCurrentWeatherPayload
}

type forecast struct {
	GetDaily  RequestWithWeatherPayload
	GetHourly RequestWithWeatherPayload
	GetPeriod RequestWithWeatherPayload
}

type monitoring struct {
	GetAlerts func() (response AdvisorResponse, err error)
}

type plan struct {
	GetInfo RequestWithPlanInfoPayload
	GetRequestDetails RequestWithRequestDetailsPayload
}

type observed struct {
	GetDaily               RequestWithWeatherPayload
	GetHourly              RequestWithWeatherPayload
	GetPeriod              RequestWithWeatherPayload
	GetLightning           RequestWithRadiusPayload
	GetLightningByGeometry RequestWithGeometryPayload
	GetFireFocus           RequestWithRadiusPayload
	GetFireFocusByGeometry RequestWithGeometryPayload
	GetStationData         RequestWithStationPayload
}

type storage struct {
	DownloadFile RequestWithStorageDownloadPayload
	ListFiles    RequestWithStorageListPayload
}

type schema struct {
	GetDefinition  func() (response AdvisorResponse, err error)
	PostDefinition RequestWithSchemaPayload
	PostParameters RequestWithSchemaPayload
}

type staticMap struct {
	Get ImageRequestWithStaticMapPayload
}

type tms struct {
	Get TmsRequest
}
