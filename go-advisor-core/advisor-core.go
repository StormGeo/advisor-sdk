package advisorsdk

import (
	"fmt"
	"io"
	"time"
)

const BASE_URL = "http://advisor-core.climatempo.io/api"

type AdvisorCoreConfig struct {
	Token     string
	Retries   uint8
	Delay     time.Duration
	NoRetries bool
	NoDelay   bool
}

type AdvisorCore struct {
	Chart          chart
	Climatology    climatology
	CurrentWeather currentWeather
	Forecast       forecast
	Monitoring     monitoring
	Plan           plan
	Observed       observed
	Schema         schema
	Tms            tms
}

func NewAdvisorCore(config AdvisorCoreConfig) AdvisorCore {
	if config.Retries == 0 && !config.NoRetries {
		config.Retries = 5
	}

	if config.Delay == 0 && !config.NoDelay {
		config.Delay = time.Second * 5
	}

	return AdvisorCore{
		Chart: chart{
			GetForecastDaily:  makeGetImage("/v1/forecast/daily/chart", config),
			GetForecastHourly: makeGetImage("/v1/forecast/hourly/chart", config),
			GetObservedDaily:  makeGetImage("/v1/observed/daily/chart", config),
			GetObservedHourly: makeGetImage("/v1/observed/hourly/chart", config),
		},
		Climatology: climatology{
			GetDaily:   makeGetWithClimatologyPayload("/v1/climatology/daily", config),
			GetMonthly: makeGetWithClimatologyPayload("/v1/climatology/monthly", config),
		},
		CurrentWeather: currentWeather{
			Get: makeGetWithCurrentWeatherPayload("/v1/current-weather", config),
		},
		Forecast: forecast{
			GetDaily:  makeGetWithWeatherPayload("/v1/forecast/daily", config),
			GetHourly: makeGetWithWeatherPayload("/v1/forecast/hourly", config),
			GetPeriod: makeGetWithWeatherPayload("/v1/forecast/period", config),
		},
		Monitoring: monitoring{
			GetAlerts: func() (response AdvisorResponse, err error) {
				return get(BASE_URL + "/v1/monitoring/alerts?token=" + config.Token)
			},
		},
		Observed: observed{
			GetDaily:               makeGetWithWeatherPayload("/v1/observed/daily", config),
			GetHourly:              makeGetWithWeatherPayload("/v1/observed/hourly", config),
			GetPeriod:              makeGetWithWeatherPayload("/v1/observed/period", config),
			GetLightning:           makeGetWithRadiusPayload("/v1/observed/lightning", config),
			GetLightningByGeometry: makePostWithGeometryPayload("/v1/observed/lightning", config),
			GetFireFocus:           makeGetWithRadiusPayload("/v1/observed/fire-focus", config),
			GetFireFocusByGeometry: makePostWithGeometryPayload("/v1/observed/fire-focus", config),
			GetStationData:         makeGetWithStationPayload("/v1/station", config),
		},
		Plan: plan{
			GetInfo: func() (response AdvisorResponse, err error) {
				return get(BASE_URL + "/v1/plan/" + config.Token)
			},
		},
		Schema: schema{
			GetDefinition: func() (response AdvisorResponse, err error) {
				return get(BASE_URL + "/v1/schema/definition?token=" + config.Token)
			},
			PostDefinition: makePostWithSchemaPayload("/v1/schema/definition", config),
			PostParameters: makePostWithSchemaPayload("/v1/schema/parameters", config),
		},
		Tms: tms{
			Get: makeGetTmsImageV1(config),
		},
	}
}

func formatUrl(route string, token string, params string) string {
	url := BASE_URL + route + "?token=" + token

	if params != "" {
		url += "&" + params
	}

	return url
}

func makeGetWithWeatherPayload(route string, config AdvisorCoreConfig) RequestWithWeatherPayload {
	return func(payload WeatherPayload) (response AdvisorResponse, err error) {
		return retryRequest(
			"get",
			config.Retries,
			config.Delay,
			formatUrl(route, config.Token, payload.toQueryParams()),
			nil,
		)
	}
}

func makeGetWithClimatologyPayload(route string, config AdvisorCoreConfig) RequestWithClimatologyPayload {
	return func(payload ClimatologyPayload) (response AdvisorResponse, err error) {
		return retryRequest(
			"get",
			config.Retries,
			config.Delay,
			formatUrl(route, config.Token, payload.toQueryParams()),
			nil,
		)
	}
}

func makeGetWithCurrentWeatherPayload(route string, config AdvisorCoreConfig) RequestWithCurrentWeatherPayload {
	return func(payload CurrentWeatherPayload) (response AdvisorResponse, err error) {
		return retryRequest(
			"get",
			config.Retries,
			config.Delay,
			formatUrl(route, config.Token, payload.toQueryParams()),
			nil,
		)
	}
}

func makeGetWithRadiusPayload(route string, config AdvisorCoreConfig) RequestWithRadiusPayload {
	return func(payload RadiusPayload) (response AdvisorResponse, err error) {
		return retryRequest(
			"get",
			config.Retries,
			config.Delay,
			formatUrl(route, config.Token, payload.toQueryParams()),
			nil,
		)
	}
}

func makeGetWithStationPayload(route string, config AdvisorCoreConfig) RequestWithStationPayload {
	return func(payload StationPayload) (response AdvisorResponse, err error) {
		return retryRequest(
			"get",
			config.Retries,
			config.Delay,
			formatUrl(route, config.Token, payload.toQueryParams()),
			nil,
		)
	}
}

func makeGetImage(route string, config AdvisorCoreConfig) ImageRequestWithWeatherPayload {
	return func(payload WeatherPayload) (imageBody io.ReadCloser, err error) {
		return retryGetImage(
			config.Retries,
			config.Delay,
			formatUrl(route, config.Token, payload.toQueryParams()),
		)
	}
}

func makeGetTmsImageV1(config AdvisorCoreConfig) TmsRequest {
	return func(payload TmsPayload) (imageBody io.ReadCloser, err error) {
		url := fmt.Sprintf(
			"%s/v1/tms/%s/%s/%s/%s/%d/%d/%d.png?token=%s&istep=%s&fstep=%s",
			BASE_URL,
			payload.Server,
			payload.Mode,
			payload.Variable,
			payload.Aggregation,
			payload.X,
			payload.Y,
			payload.Z,
			config.Token,
			payload.Istep,
			payload.Fstep,
		)

		return retryGetImage(config.Retries, config.Delay, url)
	}
}

func makePostWithGeometryPayload(route string, config AdvisorCoreConfig) RequestWithGeometryPayload {
	return func(payload GeometryPayload) (response AdvisorResponse, err error) {
		return retryRequest(
			"postGeometry",
			config.Retries,
			config.Delay,
			formatUrl(route, config.Token, payload.toQueryParams()),
			payload.toBodyBytes(),
		)
	}
}

func makePostWithSchemaPayload(route string, config AdvisorCoreConfig) RequestWithSchemaPayload {
	return func(payload SchemaPayload) (response AdvisorResponse, err error) {
		return retryRequest(
			"post",
			config.Retries,
			config.Delay,
			formatUrl(route, config.Token, ""),
			payload.toBodyBytes(),
		)
	}
}
