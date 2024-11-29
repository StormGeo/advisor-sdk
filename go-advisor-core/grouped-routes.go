package advisorsdk

import "io"

type ImageRequestWithBasePayload func(BasePayload) (io.ReadCloser, error)
type RequestWithBasePayload func(BasePayload) (AdvisorResponse, error)
type RequestWithClimatologyPayload func(ClimatologyPayload) (AdvisorResponse, error)
type RequestWithCurrentWeatherPayload func(CurrentWeatherPayload) (AdvisorResponse, error)
type RequestWithGeometryPayload func(GeometryPayload) (AdvisorResponse, error)
type RequestWithSchemaPayload func(SchemaPayload) (AdvisorResponse, error)
type RequestWithStationPayload func(StationPayload) (AdvisorResponse, error)
type RequestWithRadiusPayload func(RadiusPayload) (AdvisorResponse, error)
type RequestWithPayload func(RadiusPayload) (AdvisorResponse, error)
type TmsRequest func(TmsPayload) (io.ReadCloser, error)

type chart struct {
	GetForecastDaily  ImageRequestWithBasePayload
	GetForecastHourly ImageRequestWithBasePayload
	GetObservedDaily  ImageRequestWithBasePayload
	GetObservedHourly ImageRequestWithBasePayload
}

type climatology struct {
	GetDaily   RequestWithClimatologyPayload
	GetMonthly RequestWithClimatologyPayload
}

type currentWeather struct {
	Get RequestWithCurrentWeatherPayload
}

type forecast struct {
	GetDaily  RequestWithBasePayload
	GetHourly RequestWithBasePayload
	GetPeriod RequestWithBasePayload
}

type monitoring struct {
	GetAlerts func() (response AdvisorResponse, err error)
}

type plan struct {
	GetInfo func() (response AdvisorResponse, err error)
}

type observed struct {
	GetDaily       RequestWithBasePayload
	GetHourly      RequestWithBasePayload
	GetPeriod      RequestWithBasePayload
	GetLightning   RequestWithRadiusPayload
	PostLightning  RequestWithGeometryPayload
	GetFireFocus   RequestWithRadiusPayload
	PostFireFocus  RequestWithGeometryPayload
	GetStationData RequestWithStationPayload
}

type schema struct {
	GetDefinition  func() (response AdvisorResponse, err error)
	PostDefinition RequestWithSchemaPayload
	PostParameters RequestWithSchemaPayload
}

type tms struct {
	Get TmsRequest
}
