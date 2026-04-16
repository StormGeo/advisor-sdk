//go:build integration

package test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	sdk "github.com/StormGeo/advisor-sdk/go-advisor-core"
)

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

func init() {
	loadIntegrationEnv()
}

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

func requireConfiguredEnv(t *testing.T, name string) string {
	t.Helper()

	value := os.Getenv(name)
	if value == "" {
		t.Fatalf(
			"set %s or add it to .env.integration.local before running the Go integration tests",
			name,
		)
	}

	return value
}

func requireEnvValue(name string) string {
	value := os.Getenv(name)
	if value == "" {
		panic(fmt.Sprintf(
			"set %s or add it to .env.integration.local before running the Go integration tests",
			name,
		))
	}

	return value
}

func loadIntegrationEnv() {
	workingDir, err := os.Getwd()
	if err != nil {
		return
	}

	envFile := findIntegrationEnvFile(workingDir)
	if envFile == "" {
		return
	}

	file, err := os.Open(envFile)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		name := strings.TrimSpace(parts[0])
		if name == "" || os.Getenv(name) != "" {
			continue
		}

		value := strings.TrimSpace(parts[1])
		if len(value) >= 2 {
			if (strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"")) ||
				(strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'")) {
				value = value[1 : len(value)-1]
			}
		}

		_ = os.Setenv(name, value)
	}
}

func findIntegrationEnvFile(startDir string) string {
	currentDir := startDir
	for {
		candidate := filepath.Join(currentDir, ".env.integration.local")
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			return ""
		}

		currentDir = parentDir
	}
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
	acceptLanguage := os.Getenv("ADVISOR_ACCEPT_LANGUAGE")
	if acceptLanguage == "" {
		acceptLanguage = "en-US"
	}
	advisor.SetHeaderAcceptLanguage(acceptLanguage)
	return advisor
}

func createIntegrationPayloads() integrationPayloads {
	localeID := envUint32("ADVISOR_LOCALE_ID", 3477)
	planLocaleID := envUint32("ADVISOR_PLAN_LOCALE_ID", 5959)
	stationID := requireEnvValue("ADVISOR_STATION_ID")
	geometry := requireEnvValue("ADVISOR_GEOMETRY")
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
			Geometry:  geometry,
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
			Geometry:  geometry,
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

	return StorageDownloadPayload{
		FileName:  requireConfiguredEnv(t, "ADVISOR_STORAGE_FILE_NAME"),
		AccessKey: requireConfiguredEnv(t, "ADVISOR_STORAGE_ACCESS_KEY"),
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
