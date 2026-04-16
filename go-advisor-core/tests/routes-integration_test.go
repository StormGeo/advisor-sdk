//go:build integration

package test

import (
	"io"
	"testing"
)

func TestForecastRoutes(t *testing.T) {
	advisor := newIntegrationAdvisor(t)
	payloads := createIntegrationPayloads()

	routes := map[string]func(WeatherPayload) (AdvisorResponse, error){
		"GetDaily":  advisor.Forecast.GetDaily,
		"GetHourly": advisor.Forecast.GetHourly,
		"GetPeriod": advisor.Forecast.GetPeriod,
	}

	for name, route := range routes {
		t.Run(name, func(t *testing.T) {
			response, err := route(payloads.weatherPayload)
			assertJSONSuccess(t, response, err)
		})
	}
}

func TestObservedWeatherRoutes(t *testing.T) {
	advisor := newIntegrationAdvisor(t)
	payloads := createIntegrationPayloads()

	routes := map[string]func(WeatherPayload) (AdvisorResponse, error){
		"GetDaily":  advisor.Observed.GetDaily,
		"GetHourly": advisor.Observed.GetHourly,
		"GetPeriod": advisor.Observed.GetPeriod,
	}

	for name, route := range routes {
		t.Run(name, func(t *testing.T) {
			response, err := route(payloads.weatherPayload)
			assertJSONSuccess(t, response, err)
		})
	}
}

func TestObservedStationData(t *testing.T) {
	advisor := newIntegrationAdvisor(t)
	payloads := createIntegrationPayloads()

	response, err := advisor.Observed.GetStationData(payloads.stationPayload)
	assertJSONSuccess(t, response, err)
}

func TestObservedRadiusRoutes(t *testing.T) {
	advisor := newIntegrationAdvisor(t)
	payloads := createIntegrationPayloads()

	routes := map[string]func(RadiusPayload) (AdvisorResponse, error){
		"GetFireFocus":        advisor.Observed.GetFireFocus,
		"GetLightning":        advisor.Observed.GetLightning,
		"GetLightningDetails": advisor.Observed.GetLightningDetails,
	}

	for name, route := range routes {
		t.Run(name, func(t *testing.T) {
			payload := payloads.radiusPayload
			if name == "GetLightningDetails" {
				payload = payloads.lightningDetailsPayload
			}

			response, err := route(payload)
			assertJSONSuccess(t, response, err)
		})
	}
}

func TestObservedGeometryRoutes(t *testing.T) {
	advisor := newIntegrationAdvisor(t)
	payloads := createIntegrationPayloads()

	routes := map[string]func(GeometryPayload) (AdvisorResponse, error){
		"GetFireFocusByGeometry": advisor.Observed.GetFireFocusByGeometry,
		"GetLightningByGeometry": advisor.Observed.GetLightningByGeometry,
	}

	for name, route := range routes {
		t.Run(name, func(t *testing.T) {
			response, err := route(payloads.geometryPayload)
			assertJSONSuccess(t, response, err)
		})
	}
}

func TestObservedLightningLite(t *testing.T) {
	advisor := newIntegrationAdvisor(t)
	payloads := createIntegrationPayloads()

	response, err := advisor.Observed.GetLightningLite(payloads.lightningLitePayload)
	assertJSONSuccess(t, response, err)
}

func TestCurrentWeather(t *testing.T) {
	advisor := newIntegrationAdvisor(t)
	payloads := createIntegrationPayloads()

	response, err := advisor.CurrentWeather.Get(payloads.currentWeatherPayload)
	assertJSONSuccess(t, response, err)
}

func TestClimatologyRoutes(t *testing.T) {
	advisor := newIntegrationAdvisor(t)
	payloads := createIntegrationPayloads()

	routes := map[string]func(ClimatologyPayload) (AdvisorResponse, error){
		"GetDaily":   advisor.Climatology.GetDaily,
		"GetMonthly": advisor.Climatology.GetMonthly,
	}

	for name, route := range routes {
		t.Run(name, func(t *testing.T) {
			response, err := route(payloads.climatologyPayload)
			assertJSONSuccess(t, response, err)
		})
	}
}

func TestMonitoringAlerts(t *testing.T) {
	advisor := newIntegrationAdvisor(t)

	response, err := advisor.Monitoring.GetAlerts()
	assertJSONSuccess(t, response, err)
}

func TestStationsLastData(t *testing.T) {
	advisor := newIntegrationAdvisor(t)
	payloads := createIntegrationPayloads()

	response, err := advisor.Stations.GetLastData(payloads.stationsLastDataPayload)
	assertJSONSuccess(t, response, err)
}

func TestPlanRoutes(t *testing.T) {
	advisor := newIntegrationAdvisor(t)
	payloads := createIntegrationPayloads()

	t.Run("GetInfo", func(t *testing.T) {
		response, err := advisor.Plan.GetInfo(payloads.planInfoPayload)
		assertJSONSuccess(t, response, err)
	})

	t.Run("GetRequestDetails", func(t *testing.T) {
		response, err := advisor.Plan.GetRequestDetails(payloads.requestDetailsPayload)
		assertJSONSuccess(t, response, err)
	})

	t.Run("GetLocale", func(t *testing.T) {
		response, err := advisor.Plan.GetLocale(payloads.planLocalePayload)
		assertJSONSuccess(t, response, err)
	})
}

func TestChartRoutes(t *testing.T) {
	advisor := newIntegrationAdvisor(t)
	payloads := createIntegrationPayloads()

	routes := map[string]func(WeatherPayload) (io.ReadCloser, error){
		"GetForecastDaily":  advisor.Chart.GetForecastDaily,
		"GetForecastHourly": advisor.Chart.GetForecastHourly,
		"GetObservedDaily":  advisor.Chart.GetObservedDaily,
		"GetObservedHourly": advisor.Chart.GetObservedHourly,
	}

	for name, route := range routes {
		t.Run(name, func(t *testing.T) {
			reader, err := route(payloads.weatherChartPayload)
			assertBinarySuccess(t, reader, err)
		})
	}
}

func TestStorageRoutes(t *testing.T) {
	advisor := newIntegrationAdvisor(t)
	payloads := createIntegrationPayloads()
	downloadPayload := resolveStorageDownloadPayload(t)

	t.Run("ListFiles", func(t *testing.T) {
		response, err := advisor.Storage.ListFiles(payloads.storageListPayload)
		assertJSONSuccess(t, response, err)
	})

	t.Run("DownloadFile", func(t *testing.T) {
		reader, err := advisor.Storage.DownloadFile(downloadPayload)
		assertBinarySuccess(t, reader, err)
	})
}

func TestStaticMap(t *testing.T) {
	advisor := newIntegrationAdvisor(t)
	payloads := createIntegrationPayloads()

	reader, err := advisor.StaticMap.Get(payloads.staticMapPayload)
	assertBinarySuccess(t, reader, err)
}

func TestTms(t *testing.T) {
	advisor := newIntegrationAdvisor(t)
	payloads := createIntegrationPayloads()

	reader, err := advisor.Tms.Get(payloads.tmsPayload)
	assertBinarySuccess(t, reader, err)
}

func TestPmtiles(t *testing.T) {
	advisor := newIntegrationAdvisor(t)
	payloads := createIntegrationPayloads()

	reader, err := advisor.Pmtiles.Get(payloads.pmtilesPayload)
	assertBinarySuccess(t, reader, err)
}

func TestSchemaRoutes(t *testing.T) {
	advisor := newIntegrationAdvisor(t)
	payloads := createIntegrationPayloads()

	t.Run("GetDefinition", func(t *testing.T) {
		response, err := advisor.Schema.GetDefinition()
		assertJSONSuccess(t, response, err)
	})

	t.Run("PostDefinition", func(t *testing.T) {
		response, err := advisor.Schema.PostDefinition(payloads.schemaDefinitionPayload)
		assertJSONSuccess(t, response, err)
	})
}
