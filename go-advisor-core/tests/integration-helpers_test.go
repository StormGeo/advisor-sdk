//go:build integration

package test

import (
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	sdk "github.com/StormGeo/advisor-sdk/go-advisor-core"
)

const defaultStationID = "bWV0b3M6MDM0MTMyRjM6LTIyLjIzMTQ1MjQ4MDg0NDU2Oi00NC4yNTEzNTMwMzgzMTcx"
const defaultGeometry = "{\"type\":\"Polygon\",\"coordinates\":[[[-47.09861059094109,-23.280351816702165],[-47.09861059094109,-23.895097240590488],[-46.12890390857018,-23.895097240590488],[-46.12890390857018,-23.280351816702165],[-47.09861059094109,-23.280351816702165]]]}"
const defaultStorageFileName = "Boletim_Meteorologico_-_Andre_Alves-22_03_2026_07_02.pdf"
const defaultStorageAccessKey = "ad441946268c2227a96ad254fafe71565acd02e2-1774173734244"

type AdvisorCore = sdk.AdvisorCore
type AdvisorCoreConfig = sdk.AdvisorCoreConfig
type AdvisorResponse = sdk.AdvisorResponse
type ClimatologyPayload = sdk.ClimatologyPayload
type CurrentWeatherPayload = sdk.CurrentWeatherPayload
type GeometryPayload = sdk.GeometryPayload
type LightningLitePayload = sdk.LightningLitePayload
type PlanInfoPayload = sdk.PlanInfoPayload
type PlanLocalePayload = sdk.PlanLocalePayload
type PmtilesPayload = sdk.PmtilesPayload
type RadiusPayload = sdk.RadiusPayload
type RequestDetailsPayload = sdk.RequestDetailsPayload
type SchemaPayload = sdk.SchemaPayload
type StaticMapPayload = sdk.StaticMapPayload
type StationPayload = sdk.StationPayload
type StationsLastDataPayload = sdk.StationsLastDataPayload
type StorageDownloadPayload = sdk.StorageDownloadPayload
type StorageListPayload = sdk.StorageListPayload
type TmsPayload = sdk.TmsPayload
type WeatherPayload = sdk.WeatherPayload

var NewAdvisorCore = sdk.NewAdvisorCore

type integrationPayloads struct {
	weatherPayload          WeatherPayload
	weatherChartPayload     WeatherPayload
	climatologyPayload      ClimatologyPayload
	currentWeatherPayload   CurrentWeatherPayload
	stationPayload          StationPayload
	stationsLastDataPayload StationsLastDataPayload
	radiusPayload           RadiusPayload
	geometryPayload         GeometryPayload
	lightningDetailsPayload RadiusPayload
	lightningLitePayload    LightningLitePayload
	storageListPayload      StorageListPayload
	planInfoPayload         PlanInfoPayload
	planLocalePayload       PlanLocalePayload
	requestDetailsPayload   RequestDetailsPayload
	staticMapPayload        StaticMapPayload
	tmsPayload              TmsPayload
	pmtilesPayload          PmtilesPayload
	schemaDefinitionPayload SchemaPayload
	schemaParametersPayload SchemaPayload
}

func requireEnv(t *testing.T, name string) string {
	t.Helper()

	value := os.Getenv(name)
	if value == "" {
		t.Skipf("set %s before running the Go integration tests", name)
	}

	return value
}

func envOrDefault(name string, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}

	return value
}

func envUint32(name string, defaultValue uint32) uint32 {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}

	var parsed uint32
	fmt.Sscanf(value, "%d", &parsed)
	if parsed == 0 {
		return defaultValue
	}

	return parsed
}

func formatDateTime(value time.Time) string {
	return value.Format("2006-01-02 15:04:05")
}

func startOfDay(value time.Time) string {
	return formatDateTime(time.Date(
		value.Year(),
		value.Month(),
		value.Day(),
		0,
		0,
		0,
		0,
		value.Location(),
	))
}

func endOfDay(value time.Time) string {
	return formatDateTime(time.Date(
		value.Year(),
		value.Month(),
		value.Day(),
		23,
		59,
		59,
		0,
		value.Location(),
	))
}

func newIntegrationAdvisor(t *testing.T) AdvisorCore {
	t.Helper()

	advisor := NewAdvisorCore(AdvisorCoreConfig{
		Token:   requireEnv(t, "ADVISOR_TOKEN"),
		Retries: 1,
		Delay:   0,
	})
	advisor.SetHeaderAccept("application/json")
	advisor.SetHeaderAcceptLanguage(envOrDefault("ADVISOR_ACCEPT_LANGUAGE", "en-US"))
	return advisor
}

func createIntegrationPayloads() integrationPayloads {
	localeID := envUint32("ADVISOR_LOCALE_ID", 3477)
	planLocaleID := envUint32("ADVISOR_PLAN_LOCALE_ID", 5959)
	stationID := envOrDefault("ADVISOR_STATION_ID", defaultStationID)
	now := time.Now()
	observedDay := now.AddDate(0, 0, -1)
	observedPeriodEnd := observedDay
	observedPeriodStart := observedPeriodEnd.AddDate(0, 0, -4)
	forecastDay := now.AddDate(0, 0, 1)
	pmtilesPeriodEnd := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		1,
		0,
		0,
		0,
		now.Location(),
	)
	schemaIdentifier := "schemaIdentifier"

	return integrationPayloads{
		weatherPayload: WeatherPayload{
			LocaleId:  localeID,
			Variables: []string{"temperature"},
		},
		weatherChartPayload: WeatherPayload{
			LocaleId:  localeID,
			Variables: []string{"temperature", "precipitation"},
		},
		climatologyPayload: ClimatologyPayload{
			LocaleId:  localeID,
			Variables: []string{"temperature"},
		},
		currentWeatherPayload: CurrentWeatherPayload{
			LocaleId: localeID,
		},
		stationPayload: StationPayload{
			StationId: stationID,
		},
		stationsLastDataPayload: StationsLastDataPayload{
			StationIds: []string{stationID},
			Variables:  []string{"temperature"},
		},
		radiusPayload: RadiusPayload{
			LocaleId: localeID,
			Radius:   10000,
		},
		geometryPayload: GeometryPayload{
			Geometry:  defaultGeometry,
			StartDate: startOfDay(observedDay),
			EndDate:   endOfDay(observedDay),
			Radius:    10000,
		},
		lightningDetailsPayload: RadiusPayload{
			Latitude:  "-22.9",
			Longitude: "-43.2",
			StartDate: startOfDay(observedDay),
			EndDate:   endOfDay(observedDay),
			Radius:    10000,
		},
		lightningLitePayload: LightningLitePayload{
			Geometry:  defaultGeometry,
			StartDate: startOfDay(observedPeriodStart),
			EndDate:   endOfDay(observedPeriodEnd),
			Radius:    10000,
			Page:      1,
			PageSize:  50,
		},
		storageListPayload: StorageListPayload{
			Page:     1,
			PageSize: 10,
		},
		planInfoPayload: PlanInfoPayload{
			Timezone: -3,
		},
		planLocalePayload: PlanLocalePayload{
			LocaleId: planLocaleID,
		},
		requestDetailsPayload: RequestDetailsPayload{
			Page:     1,
			PageSize: 3,
		},
		staticMapPayload: StaticMapPayload{
			Type:          "periods",
			Category:      "observed",
			Variable:      "temperature",
			Aggregation:   "max",
			StartDate:     startOfDay(observedPeriodStart),
			EndDate:       endOfDay(observedPeriodEnd),
			Dpi:           50,
			Title:         true,
			Titlevariable: "Static Map",
		},
		tmsPayload: TmsPayload{
			Istep:       startOfDay(forecastDay),
			Fstep:       endOfDay(forecastDay),
			Server:      "a",
			Mode:        "forecast",
			Variable:    "precipitation",
			Aggregation: "sum",
			X:           5,
			Y:           8,
			Z:           4,
		},
		pmtilesPayload: PmtilesPayload{
			Mode:        "forecast",
			Model:       "ct2w15_as",
			Variable:    "precipitation",
			Aggregation: "sum",
			Istep:       startOfDay(now),
			Fstep:       formatDateTime(pmtilesPeriodEnd),
			MaxZoom:     4,
			Timezone:    -3,
		},
		schemaDefinitionPayload: SchemaPayload{
			"identifier": schemaIdentifier,
			"arbitraryField1": map[string]any{
				"type":     "boolean",
				"required": true,
				"length":   125,
			},
			"arbitraryField2": map[string]any{
				"type":     "number",
				"required": true,
			},
			"arbitraryField3": map[string]any{
				"type":     "string",
				"required": false,
			},
		},
		schemaParametersPayload: SchemaPayload{
			"identifier":      schemaIdentifier,
			"arbitraryField1": true,
			"arbitraryField2": 15,
		},
	}
}

func resolveStorageDownloadPayload(t *testing.T) StorageDownloadPayload {
	t.Helper()

	fileName := os.Getenv("ADVISOR_STORAGE_FILE_NAME")
	accessKey := os.Getenv("ADVISOR_STORAGE_ACCESS_KEY")

	if fileName != "" || accessKey != "" {
		if fileName == "" || accessKey == "" {
			t.Fatal("set both ADVISOR_STORAGE_FILE_NAME and ADVISOR_STORAGE_ACCESS_KEY, or neither")
		}

		return StorageDownloadPayload{
			FileName:  fileName,
			AccessKey: accessKey,
		}
	}

	return StorageDownloadPayload{
		FileName:  defaultStorageFileName,
		AccessKey: defaultStorageAccessKey,
	}
}

func assertJSONSuccess(t *testing.T, response AdvisorResponse, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if response == nil {
		t.Fatal("expected non-nil response")
	}
}

func assertBinarySuccess(t *testing.T, reader io.ReadCloser, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if reader == nil {
		t.Fatal("expected non-nil reader")
	}

	defer reader.Close()

	content, readErr := io.ReadAll(reader)
	if readErr != nil {
		t.Fatalf("failed reading stream: %v", readErr)
	}

	if len(content) == 0 {
		t.Fatal("expected non-empty binary content")
	}
}
