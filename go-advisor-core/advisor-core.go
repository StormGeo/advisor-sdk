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
			GetDaily:  makeGetWithBasePayload("/v1/forecast/daily", config),
			GetHourly: makeGetWithBasePayload("/v1/forecast/hourly", config),
			GetPeriod: makeGetWithBasePayload("/v1/forecast/period", config),
		},
		Monitoring: monitoring{
			GetAlerts: func() (response AdvisorResponse, err error) {
				return Get(BASE_URL + "/v1/monitoring/alerts?token=" + config.Token)
			},
		},
		Observed: observed{
			GetDaily:       makeGetWithBasePayload("/v1/observed/daily", config),
			GetHourly:      makeGetWithBasePayload("/v1/observed/hourly", config),
			GetPeriod:      makeGetWithBasePayload("/v1/observed/period", config),
			GetLightning:   makeGetWithRadiusPayload("/v1/observed/lightning", config),
			PostLightning:  makePostWithGeometryPayload("/v1/observed/lightning", config),
			GetFireFocus:   makeGetWithRadiusPayload("/v1/observed/fire-focus", config),
			PostFireFocus:  makePostWithGeometryPayload("/v1/observed/fire-focus", config),
			GetStationData: makeGetWithStationPayload("/v1/station", config),
		},
		Plan: plan{
			GetInfo: func() (response AdvisorResponse, err error) {
				return Get(BASE_URL + "/v1/plan/" + config.Token)
			},
		},
		Schema: schema{
			GetDefinition: func() (response AdvisorResponse, err error) {
				return Get(BASE_URL + "/v1/schema/definition?token=" + config.Token)
			},
			PostDefinition: makePostWithSchemaPayload("/v1/schema/definition", config),
			PostParameters: makePostWithSchemaPayload("/v1/schema/parameters", config),
		},
		Tms: tms{
			Get: makeGetTmsImageV1(config),
		},
	}
}

func makeGetWithBasePayload(route string, advisorCoreConfig AdvisorCoreConfig) RequestWithBasePayload {
	url := BASE_URL + route + "?token=" + advisorCoreConfig.Token + "&"

	return func(payload BasePayload) (response AdvisorResponse, err error) {
		for attempts := advisorCoreConfig.Retries + 1; attempts > 0; attempts-- {
			response, err = Get(url + payload.toQueryParams())
			if err != nil {
				time.Sleep(advisorCoreConfig.Delay)
				fmt.Println("Waiting, attempts: ", attempts)
				continue
			}

			break
		}

		return response, err
	}
}

func makeGetWithClimatologyPayload(route string, advisorCoreConfig AdvisorCoreConfig) RequestWithClimatologyPayload {
	url := BASE_URL + route + "?token=" + advisorCoreConfig.Token + "&"

	return func(payload ClimatologyPayload) (response AdvisorResponse, err error) {
		for attempts := advisorCoreConfig.Retries + 1; attempts > 0; attempts-- {
			response, err = Get(url + payload.toQueryParams())
			if err != nil {
				time.Sleep(advisorCoreConfig.Delay)
				fmt.Println("Waiting, attempts: ", attempts)
				continue
			}

			break
		}

		return response, err
	}
}

func makeGetWithCurrentWeatherPayload(route string, advisorCoreConfig AdvisorCoreConfig) RequestWithCurrentWeatherPayload {
	url := BASE_URL + route + "?token=" + advisorCoreConfig.Token + "&"

	return func(payload CurrentWeatherPayload) (response AdvisorResponse, err error) {
		for attempts := advisorCoreConfig.Retries + 1; attempts > 0; attempts-- {
			response, err = Get(url + payload.toQueryParams())
			if err != nil {
				time.Sleep(advisorCoreConfig.Delay)
				fmt.Println("Waiting, attempts: ", attempts)
				continue
			}

			break
		}

		return response, err
	}
}

func makeGetWithRadiusPayload(route string, advisorCoreConfig AdvisorCoreConfig) RequestWithRadiusPayload {
	url := BASE_URL + route + "?token=" + advisorCoreConfig.Token + "&"

	return func(payload RadiusPayload) (response AdvisorResponse, err error) {
		for attempts := advisorCoreConfig.Retries + 1; attempts > 0; attempts-- {
			response, err = Get(url + payload.toQueryParams())
			if err != nil {
				time.Sleep(advisorCoreConfig.Delay)
				fmt.Println("Waiting, attempts: ", attempts)
				continue
			}

			break
		}

		return response, err
	}
}

func makeGetWithStationPayload(route string, advisorCoreConfig AdvisorCoreConfig) RequestWithStationPayload {
	url := BASE_URL + route + "?token=" + advisorCoreConfig.Token + "&"

	return func(payload StationPayload) (response AdvisorResponse, err error) {
		for attempts := advisorCoreConfig.Retries + 1; attempts > 0; attempts-- {
			response, err = Get(url + payload.toQueryParams())
			if err != nil {
				time.Sleep(advisorCoreConfig.Delay)
				fmt.Println("Waiting, attempts: ", attempts)
				continue
			}

			break
		}

		return response, err
	}
}

func makeGetImage(route string, advisorCoreConfig AdvisorCoreConfig) ImageRequestWithBasePayload {
	url := BASE_URL + route + "?token=" + advisorCoreConfig.Token + "&"

	return func(payload BasePayload) (imageBody io.ReadCloser, err error) {
		for attempts := advisorCoreConfig.Retries + 1; attempts > 0; attempts-- {
			imageBody, err = GetImage(url + payload.toQueryParams())
			if err != nil {
				time.Sleep(advisorCoreConfig.Delay)
				fmt.Println("Waiting, attempts: ", attempts)
				continue
			}

			break
		}

		return imageBody, err
	}
}

func makeGetTmsImageV1(advisorCoreConfig AdvisorCoreConfig) TmsRequest {
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
			advisorCoreConfig.Token,
			payload.Istep,
			payload.Fstep,
		)

		for attempts := advisorCoreConfig.Retries + 1; attempts > 0; attempts-- {
			imageBody, err = GetImage(url)
			if err != nil {
				time.Sleep(advisorCoreConfig.Delay)
				fmt.Println("Waiting, attempts: ", attempts)
				continue
			}

			break
		}

		return imageBody, err
	}
}

func makePostWithGeometryPayload(route string, advisorCoreConfig AdvisorCoreConfig) RequestWithGeometryPayload {
	return func(payload GeometryPayload) (response AdvisorResponse, err error) {
		url := BASE_URL + route + "?token=" + advisorCoreConfig.Token
		body := payload.toBodyBytes()
		params := payload.toQueryParams()
		if params != "" {
			url += "&" + params
		}

		for attempts := advisorCoreConfig.Retries + 1; attempts > 0; attempts-- {
			response, err = PostGeometry(url, body)
			if err != nil && attempts > 1 {
				time.Sleep(advisorCoreConfig.Delay)
				fmt.Println("Waiting, attempts: ", attempts)
				continue
			}

			break
		}

		return response, err
	}
}

func makePostWithSchemaPayload(route string, advisorCoreConfig AdvisorCoreConfig) RequestWithSchemaPayload {
	return func(payload SchemaPayload) (response AdvisorResponse, err error) {
		url := BASE_URL + route + "?token=" + advisorCoreConfig.Token
		body := payload.toBodyBytes()

		for attempts := advisorCoreConfig.Retries + 1; attempts > 0; attempts-- {
			response, err = Post(url, body)
			if err != nil && attempts > 1 {
				time.Sleep(advisorCoreConfig.Delay)
				fmt.Println("Waiting, attempts: ", attempts)
				continue
			}

			break
		}

		return response, err
	}
}
