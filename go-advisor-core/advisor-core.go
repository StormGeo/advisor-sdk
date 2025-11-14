package advisorsdk

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const BASE_URL = "http://advisor-core.climatempo.io/api"

type AdvisorCoreConfig struct {
	Token     string
	Retries   uint8
	Delay     uint8
	NoRetries bool
	NoDelay   bool
}

type AdvisorCore struct {
	header         *http.Header
	Chart          chart
	Climatology    climatology
	CurrentWeather currentWeather
	Forecast       forecast
	Monitoring     monitoring
	Plan           plan
	Observed       observed
	Storage        storage
	Schema         schema
	StaticMap      staticMap
	Tms            tms
}

func (a *AdvisorCore) SetHeaderAccept(value string) {
	a.header.Set("Accept", value)
}

func (a *AdvisorCore) SetHeaderAcceptLanguage(value string) {
	a.header.Set("Accept-Language", value)
}

func NewAdvisorCore(config AdvisorCoreConfig) AdvisorCore {
	if config.Retries == 0 && !config.NoRetries {
		config.Retries = 5
	}

	if config.Delay == 0 && !config.NoDelay {
		config.Delay = 5
	}

	header := http.Header{}
	header.Set("Accept", "application/json")
	header.Set("Content-Type", "application/json")
	header.Set("Accept-Language", "en-US")
	header.Set("x-advisor-token", config.Token)
	header.Set("User-Agent", "Go-AdvisorCore-SDK")

	return AdvisorCore{
		header: &header,
		Chart: chart{
			GetForecastDaily:  makeGetImage("/v1/forecast/daily/chart", config, header),
			GetForecastHourly: makeGetImage("/v1/forecast/hourly/chart", config, header),
			GetObservedDaily:  makeGetImage("/v1/observed/daily/chart", config, header),
			GetObservedHourly: makeGetImage("/v1/observed/hourly/chart", config, header),
		},
		Climatology: climatology{
			GetDaily:   makeGetWithClimatologyPayload("/v1/climatology/daily", config, header),
			GetMonthly: makeGetWithClimatologyPayload("/v1/climatology/monthly", config, header),
		},
		CurrentWeather: currentWeather{
			Get: makeGetWithCurrentWeatherPayload("/v1/current-weather", config, header),
		},
		Forecast: forecast{
			GetDaily:  makeGetWithWeatherPayload("/v1/forecast/daily", config, header),
			GetHourly: makeGetWithWeatherPayload("/v1/forecast/hourly", config, header),
			GetPeriod: makeGetWithWeatherPayload("/v1/forecast/period", config, header),
		},
		Monitoring: monitoring{
			GetAlerts: func() (response AdvisorResponse, err error) {
				return formatResponse(retryReq(
					"GET",
					config.Retries,
					config.Delay,
					BASE_URL+"/v1/monitoring/alerts",
					nil,
					header,
				))
			},
		},
		Observed: observed{
			GetDaily:               makeGetWithWeatherPayload("/v1/observed/daily", config, header),
			GetHourly:              makeGetWithWeatherPayload("/v1/observed/hourly", config, header),
			GetPeriod:              makeGetWithWeatherPayload("/v1/observed/period", config, header),
			GetLightning:           makeGetWithRadiusPayload("/v1/observed/lightning", config, header),
			GetLightningByGeometry: makePostWithGeometryPayload("/v1/observed/lightning", config, header),
			GetFireFocus:           makeGetWithRadiusPayload("/v1/observed/fire-focus", config, header),
			GetFireFocusByGeometry: makePostWithGeometryPayload("/v1/observed/fire-focus", config, header),
			GetStationData:         makeGetWithStationPayload("/v1/station", config, header),
		},
		Plan: plan{
			GetInfo:           makeGetWithPlanInfoPayload("/v1/plan", config, header),
			GetLocale:         makeGetWithPlanLocalePayload("/v1/plan/locale", config, header),
			GetRequestDetails: makeGetWithRequestDetailsPayload("/v1/plan/request-details", config, header),
		},
		Storage: storage{
			DownloadFile: makeGetFile("/v1/storage/download", config, header),
			ListFiles:    makeGetWithListFilesPayload("/v1/storage/list", config, header),
		},
		StaticMap: staticMap{
			Get: makeGetStaticMapImage("/v1/map", config, header),
		},
		Schema: schema{
			GetDefinition: func() (response AdvisorResponse, err error) {
				return formatResponse(retryReq(
					"GET",
					config.Retries,
					config.Delay,
					BASE_URL+"/v1/schema/definition",
					nil,
					header,
				))
			},
			PostDefinition: makePostWithSchemaPayload("/v1/schema/definition", config, header),
			PostParameters: makePostWithSchemaPayload("/v1/schema/parameters", config, header),
		},
		Tms: tms{
			Get: makeGetTmsImageV1(config, header),
		},
	}
}

func formatUrl(route string, params string) string {
	url := BASE_URL + route

	if params != "" {
		separator := "?"
		if strings.Contains(url, "?") {
			separator = "&"
		}

		url += separator + params
	}

	return url
}

func makeGetWithWeatherPayload(route string, config AdvisorCoreConfig, header http.Header) RequestWithWeatherPayload {
	return func(payload WeatherPayload) (response AdvisorResponse, err error) {
		return formatResponse(retryReq(
			"GET",
			config.Retries,
			config.Delay,
			formatUrl(route, payload.toQueryParams()),
			nil,
			header,
		))
	}
}

func makeGetWithPlanInfoPayload(route string, config AdvisorCoreConfig, header http.Header) RequestWithPlanInfoPayload {
	return func(payload PlanInfoPayload) (response AdvisorResponse, err error) {
		return formatResponse(retryReq(
			"GET",
			config.Retries,
			config.Delay,
			formatUrl(route+"/"+config.Token, payload.toQueryParams()),
			nil,
			header,
		))
	}
}

func makeGetWithPlanLocalePayload(route string, config AdvisorCoreConfig, header http.Header) RequestWithPlanLocalePayload {
	return func(payload PlanLocalePayload) (response AdvisorResponse, err error) {
		return formatResponse(retryReq(
			"GET",
			config.Retries,
			config.Delay,
			formatUrl(route, payload.toQueryParams()),
			nil,
			header,
		))
	}
}

func makeGetWithRequestDetailsPayload(route string, config AdvisorCoreConfig, header http.Header) RequestWithRequestDetailsPayload {
	return func(payload RequestDetailsPayload) (response AdvisorResponse, err error) {
		return formatResponse(retryReq(
			"GET",
			config.Retries,
			config.Delay,
			formatUrl(route, payload.toQueryParams()),
			nil,
			header,
		))
	}
}

func makeGetWithClimatologyPayload(route string, config AdvisorCoreConfig, header http.Header) RequestWithClimatologyPayload {
	return func(payload ClimatologyPayload) (response AdvisorResponse, err error) {
		return formatResponse(retryReq(
			"GET",
			config.Retries,
			config.Delay,
			formatUrl(route, payload.toQueryParams()),
			nil,
			header,
		))
	}
}

func makeGetWithCurrentWeatherPayload(route string, config AdvisorCoreConfig, header http.Header) RequestWithCurrentWeatherPayload {
	return func(payload CurrentWeatherPayload) (response AdvisorResponse, err error) {
		return formatResponse(retryReq(
			"GET",
			config.Retries,
			config.Delay,
			formatUrl(route, payload.toQueryParams()),
			nil,
			header,
		))
	}
}

func makeGetWithRadiusPayload(route string, config AdvisorCoreConfig, header http.Header) RequestWithRadiusPayload {
	return func(payload RadiusPayload) (response AdvisorResponse, err error) {
		return formatResponse(retryReq(
			"GET",
			config.Retries,
			config.Delay,
			formatUrl(route, payload.toQueryParams()),
			nil,
			header,
		))
	}
}

func makeGetWithStationPayload(route string, config AdvisorCoreConfig, header http.Header) RequestWithStationPayload {
	return func(payload StationPayload) (response AdvisorResponse, err error) {
		return formatResponse(retryReq(
			"GET",
			config.Retries,
			config.Delay,
			formatUrl(route, payload.toQueryParams()),
			nil,
			header,
		))
	}
}

func makeGetImage(route string, config AdvisorCoreConfig, header http.Header) ImageRequestWithWeatherPayload {
	return func(payload WeatherPayload) (imageBody io.ReadCloser, err error) {
		resp, respErr := retryReq(
			"GET",
			config.Retries,
			config.Delay,
			formatUrl(route, payload.toQueryParams()),
			nil,
			header,
		)
		if respErr != nil {
			return nil, respErr
		}

		return resp.Body, nil
	}
}

func makeGetStaticMapImage(route string, config AdvisorCoreConfig, header http.Header) ImageRequestWithStaticMapPayload {
	return func(payload StaticMapPayload) (imageBody io.ReadCloser, err error) {
		route = fmt.Sprintf(
			"%s/%s/%s/%s",
			route,
			url.QueryEscape(payload.Type),
			url.QueryEscape(payload.Category),
			url.QueryEscape(payload.Variable),
		)

		resp, respErr := retryReq(
			"GET",
			config.Retries,
			config.Delay,
			formatUrl(route, payload.toQueryParams()),
			nil,
			header,
		)

		if respErr != nil {
			return nil, respErr
		}

		return resp.Body, nil
	}
}

func makeGetFile(route string, config AdvisorCoreConfig, header http.Header) RequestWithStorageDownloadPayload {
	return func(payload StorageDownloadPayload) (fileBody io.ReadCloser, err error) {
		route = route + "/" + url.QueryEscape(payload.FileName)
		resp, respErr := retryReq(
			"GET",
			config.Retries,
			config.Delay,
			formatUrl(route, payload.toQueryParams()),
			nil,
			header,
		)
		if respErr != nil {
			return nil, respErr
		}

		return resp.Body, nil
	}
}

func makeGetWithListFilesPayload(route string, config AdvisorCoreConfig, header http.Header) RequestWithStorageListPayload {
	return func(payload StorageListPayload) (response AdvisorResponse, err error) {
		return formatResponse(retryReq(
			"GET",
			config.Retries,
			config.Delay,
			formatUrl(route, payload.toQueryParams()),
			nil,
			header,
		))
	}
}

func makeGetTmsImageV1(config AdvisorCoreConfig, header http.Header) TmsRequest {
	return func(payload TmsPayload) (imageBody io.ReadCloser, err error) {
		url := fmt.Sprintf(
			"%s/v1/tms/%s/%s/%s/%s/%d/%d/%d.png?istep=%s&fstep=%s&timezone=%d",
			BASE_URL,
			payload.Server,
			payload.Mode,
			payload.Variable,
			payload.Aggregation,
			payload.X,
			payload.Y,
			payload.Z,
			url.QueryEscape(payload.Istep),
			url.QueryEscape(payload.Fstep),
			payload.Timezone,
		)

		resp, respErr := retryReq(
			"GET",
			config.Retries,
			config.Delay,
			url,
			nil,
			header,
		)
		if respErr != nil {
			return nil, respErr
		}

		return resp.Body, nil
	}
}

func makePostWithGeometryPayload(route string, config AdvisorCoreConfig, header http.Header) RequestWithGeometryPayload {
	return func(payload GeometryPayload) (response AdvisorResponse, err error) {
		return formatResponse(retryReq(
			"POST",
			config.Retries,
			config.Delay,
			formatUrl(route, payload.toQueryParams()),
			payload.toBodyBytes(),
			header,
		))
	}
}

func makePostWithSchemaPayload(route string, config AdvisorCoreConfig, header http.Header) RequestWithSchemaPayload {
	return func(payload SchemaPayload) (response AdvisorResponse, err error) {
		return formatResponse(retryReq(
			"POST",
			config.Retries,
			config.Delay,
			formatUrl(route, ""),
			payload.toBodyBytes(),
			header,
		))
	}
}
