# Go SDK

Advisor Software Development Kit for Go.

## Contents
- [Go SDK](#go-sdk)
  - [How to get your token](https://www.climatempoconsultoria.com.br/contato/)
  - [Installation](#installation)
  - [Contents](#contents)
  - [Routes](#routes)
    - [Examples](#examples)
      - [Chart:](#chart)
      - [Climatology:](#climatology)
      - [Current Weather:](#current-weather)
      - [Forecast:](#forecast)
      - [Monitoring:](#monitoring)
      - [Observed:](#observed)
      - [Plan Information:](#plan-information)
      - [Plan Locale:](#plan-locale)
      - [Schema/Parameter:](#schemaparameter)
      - [Storage:](#storage)
      - [Static Map:](#static-map)
      - [Tms (Tiles Map Server):](#tms-tiles-map-server)
  - [Headers Configuration](#headers-configuration)
  - [Response Format](#response-format)
  - [Payload Types](#payload-types)
    - [WeatherPayload](#weatherpayload)
    - [StationPayload](#stationpayload)
    - [ClimatologyPayload](#climatologypayload)
    - [CurrentWeatherPayload](#currentweatherpayload)
    - [RadiusPayload](#radiuspayload)
    - [GeometryPayload](#geometrypayload)
    - [TmsPayload](#tmspayload)
    - [StaticMapPayload](#staticmappayload)
    - [RequestDetailsPayload](#requestdetailspayload)
    - [PlanInfoPayload](#planinfopayload)
    - [PlanLocalePayload](#planlocalepayload)
    - [StorageListPayload](#storagelistpayload)
    - [StorageDownloadPayload](#storagedownloadpayload)
---

## Installation

To install this package, use the following command:
```bash
go get github.com/StormGeo/advisor-sdk/go-advisor-core
```

## Routes

First you need to import the SDK on your application and instancy the `AdvisorCore` class setting up your access token and needed configurations:

```go
  import (
    sdk "github.com/StormGeo/advisor-sdk/go-advisor-core"
  )

  config := sdk.AdvisorCoreConfig{
    Token: "<your-token>",
  }
  advisor := sdk.NewAdvisorCore(config)
```

The recommended Go version to use this library is the 1.23.3 or a later version. Using earlier Go versions may result in unexpected behavior or incompatibilities.

### Examples

#### Chart:
```go
payload := sdk.WeatherPayload{
  LocaleId: 3477,
  Variables: []string{"temperature"},
}

// requesting daily forecast chart image
resp, respErr := advisor.Chart.GetForecastDaily(payload)

// requesting hourly forecast chart image
resp, respErr := advisor.Chart.GetForecastHourly(payload)

// requesting daily observed chart image
resp, respErr := advisor.Chart.GetObservedDaily(payload)

// requesting hourly observed chart image
resp, respErr := advisor.Chart.GetObservedHourly(payload)

if respErr != nil {
  fmt.Println(respErr)
  return
}
defer resp.Close()

file, fileErr := os.Create("./chart.png")
if fileErr != nil {
  fmt.Println(fileErr)
  return
}
defer file.Close()

_, err := io.Copy(file, resp)
if err != nil {
  fmt.Println(err)
  return
}

fmt.Println("Chart saved with success!")
```


#### Climatology:
```go
payload := sdk.ClimatologyPayload{
  LocaleId:  3477,
  Variables: []string{"precipitation"},
}

// requesting daily climatology data
resp, respErr := advisor.Climatology.GetDaily(payload)

// requesting monthly climatology data
resp, respErr := advisor.Climatology.GetMonthly(payload)

if respErr != nil {
  fmt.Println(respErr)
  fmt.Println("Error trying to get data!")
} else {
  fmt.Println(resp)
}
```


#### Current Weather:
```go
payload := sdk.CurrentWeatherPayload{
  LocaleId: 3477,
  Variables: []string{"temperature"},
}

resp, respErr := advisor.CurrentWeather.Get(payload)

if respErr != nil {
  fmt.Println(respErr)
  fmt.Println("Error trying to get data!")
} else {
  fmt.Println(resp)
}
```


#### Forecast:
```go
payload := sdk.WeatherPayload{
  LocaleId: 3477,
  Variables: []string{"temperature"},
}

// requesting daily forecast data
resp, respErr := advisor.Forecast.GetDaily(payload)

// requesting hourly forecast data
resp, respErr := advisor.Forecast.GetHourly(payload)

// requesting period forecast data
resp, respErr := advisor.Forecast.GetPeriod(payload)

if respErr != nil {
  fmt.Println(respErr)
  fmt.Println("Error trying to get data!")
} else {
  fmt.Println(resp)
}
```


#### Monitoring:
```go
resp, respErr := advisor.Monitoring.GetAlerts()

if respErr != nil {
  fmt.Println(respErr)
  fmt.Println("Error trying to get data!")
} else {
  fmt.Println(resp)
}
```


#### Observed:
```go
payload := sdk.WeatherPayload{
  LocaleId: 3477,
  Variables: []string{"temperature"},
}

// requesting daily observed data
resp, respErr := advisor.Observed.GetDaily(payload)

// requesting hourly observed data
resp, respErr := advisor.Observed.GetHourly(payload)

// requesting period observed data
resp, respErr := advisor.Observed.GetPeriod(payload)

stationPayload := sdk.StationPayload{
  StationId: "ABC123abc321CBA",
}

// requesting station observed data
resp, respErr := advisor.Observed.GetStationData(stationPayload)

radiusPayload := sdk.RadiusPayload{
  LocaleId: 3477,
  Radius: 100,
}

// requesting fire-focus observed data
resp, respErr := advisor.Observed.GetFireFocus(radiusPayload)

// requesting lightning observed data
resp, respErr := advisor.Observed.GetLightning(radiusPayload)


geometryPayload := sdk.GeometryPayload{
  StartDate: "2024-11-28 00:00:00",
  EndDate:   "2024-11-28 23:59:59",
  Geometry:  "{\"type\": \"MultiPoint\", \"coordinates\": [[-41.88, -22.74]]}",
}

// requesting fire-focus observed data by geometry
resp, respErr := advisor.Observed.GetFireFocusByGeometry(geometryPayload)

// requesting lightning observed data by geometry
resp, respErr := advisor.Observed.GetLightningByGeometry(geometryPayload)

if respErr != nil {
  fmt.Println(respErr)
  fmt.Println("Error trying to get data!")
} else {
  fmt.Println(resp)
}
```


#### Plan Information:
```go
payload := sdk.RequestDetailsPayload{
  Page:     1,
  PageSize: 10,
}

resp, respErr := advisor.Plan.GetRequestDetails(payload)

payload := sdk.PlanInfoPayload{
  Timezone: -3, // Set the timezone offset the default is 0 (UTC)
}

resp, respErr := advisor.Plan.GetInfo(payload)

if respErr != nil {
  fmt.Println(respErr)
  fmt.Println("Error trying to get data!")
} else {
  fmt.Println(resp)
}
```

#### Plan Locale:
```go
payload := sdk.PlanLocalePayload{
  LocaleId: 3477,
  // You can also set Latitude/Longitude or StationId instead of LocaleId
}

resp, respErr := advisor.Plan.GetLocale(payload)

if respErr != nil {
  fmt.Println(respErr)
  fmt.Println("Error trying to get data!")
} else {
  fmt.Println(resp)
}
```

#### Schema/Parameter:
```go
// Arbitrary example on how to define a schema
schemaPayload := sdk.SchemaPayload{
  "identifier": "arbitraryIdentifier",
  "arbitraryField1": map[string]interface{}{
    "type": "boolean",
    "required": true,
  },
}

// Arbitrary example on how to upload data to parameters from schema
parametersPayload := sdk.SchemaPayload{
  "identifier": "arbitraryIdentifier",
  "arbitraryField1": true,
}

// requesting all schemas from token
resp, respErr := advisor.Schema.GetDefinition()

// requesting to upload a new schema
resp, respErr := advisor.Schema.PostDefinition(schemaPayload)

// requesting to upload data to parameters from schema
resp, respErr := advisor.Schema.PostParameters(parametersPayload)

if respErr != nil {
  fmt.Println(respErr)
  fmt.Println("Error trying to get data!")
} else {
  fmt.Println(resp)
}
```

#### Storage:
```go
payload := sdk.StorageListPayload{
  Page:      1,
  PageSize:  10,
}

resp, respErr := advisor.Storage.ListFiles(payload)

if respErr != nil {
  fmt.Println(respErr)
  fmt.Println("Error trying to get data!")
} else {
  fmt.Println(resp)
}

payload := sdk.StorageDownloadPayload{
  FileName:  "file.pdf",
  AccessKey: "<file-access-key>",
}

resp, respErr := advisor.Storage.DownloadFile(payload)

if respErr != nil {
  fmt.Println(respErr)
  return
}
defer resp.Close()

file, fileErr := os.Create("./file.pdf")
if fileErr != nil {
  fmt.Println(fileErr)
  return
}
defer file.Close()

_, err := io.Copy(file, resp)
if err != nil {
  fmt.Println(err)
  return
}

fmt.Println("File saved with success!")
```

#### Static Map:
```go
payload := sdk.StaticMapPayload{
  Type:          "periods",
  Category:      "observed",
  Variable:      "temperature",
  Aggregation:   "max",
  StartDate:     "2025-07-28 00:00:00",
  EndDate:       "2025-07-29 23:59:59",
  Dpi:           50,
  Title:         true,
  Titlevariable: "temperature",
}

resp, respErr := advisor.StaticMap.Get(payload)

if respErr != nil {
  fmt.Println(respErr)
  return
}
defer resp.Close()

file, fileErr := os.Create("./map_image.png")
if fileErr != nil {
  fmt.Println(fileErr)
  return
}
defer file.Close()

_, err := io.Copy(file, resp)
if err != nil {
  fmt.Println(err)
  return
}

fmt.Println("File saved with success!")
```

#### Tms (Tiles Map Server):
```go
payload := sdk.TmsPayload{
  Istep: "2024-12-25 10:00:00",
  Fstep: "2024-12-25 12:00:00",
  Server: "a",
  Mode: "forecast",
  Variable: "precipitation",
  Aggregation: "sum",
  X: 2,
  Y: 3,
  Z: 4,
  Timezone: -3, // Set the timezone offset the default is 0 (UTC)
}

resp, respErr := advisor.Tms.Get(payload)

if respErr != nil {
  fmt.Println(respErr)
  return
}
defer resp.Close()

file, fileErr := os.Create("./tile.png")
if fileErr != nil {
  fmt.Println(fileErr)
  return
}
defer file.Close()

_, err := io.Copy(file, resp)
if err != nil {
  fmt.Println(err)
  return
}

fmt.Println("Tile saved with success!")
```

## Headers Configuration

You can also set headers to translate the error descriptions or to receive the response in a different format type. This functionality is only available for some routes, consult the API documentation to find out which routes have this functionality.

Available languages: 
- en-US (default)
- pt-BR
- es-ES

Available response types:
- application/json (default)
- application/xml
- text/csv

Example:

```go
config := sdk.AdvisorCoreConfig{Token: "invalid-token"}
advisor := sdk.NewAdvisorCore(config)

advisor.SetHeaderAccept("application/xml")
advisor.SetHeaderAcceptLanguage("es-ES")

_, respErr := advisor.Plan.GetInfo()

fmt.Println(respErr)

// <response>
//
//	<error>
//	  <type>UNAUTHORIZED_ACCESS</type>
//	  <message>UNAUTHORIZED_REQUEST</message>
//	  <description>La solicitud no está autorizada.</description>
//	</error>
//
// </response>
```


## Response Format

All the methods returns two parameters data and error:

- data -> interface{} | nil
- error -> error | nil

## Payload Types

### WeatherPayload

- **LocaleId**: string
- **StationId**: string
- **Latitude**: uint32
- **Longitude**: uint32
- **Timezone**: int8
- **Variables**: []string
- **StartDate**: string
- **EndDate**: string

### StationPayload

- **StationId**: string
- **Layer**: string
- **Variables**: []string
- **StartDate**: string
- **EndDate**: string

### ClimatologyPayload

- **LocaleId**: string
- **StationId**: string
- **Latitude**: uint32
- **Longitude**: uint32
- **Variables**: []string

### CurrentWeatherPayload

- **LocaleId**: string
- **StationId**: string
- **Latitude**: uint32
- **Longitude**: uint32
- **Timezone**: int8
- **Variables**: []string

### RadiusPayload

- **LocaleId**: string
- **StationId**: string
- **Latitude**: uint32
- **Longitude**: uint32
- **StartDate**: string
- **EndDate**: string
- **Radius**: uint32

### GeometryPayload

- **StartDate**: string
- **EndDate**: string
- **Radius**: uint32
- **Geometry**: string

### TmsPayload

- **Server**: string
- **Mode**: string
- **Variable**: string
- **Aggregation**: string
- **X**: uint16
- **Y**: uint16
- **Z**: uint16
- **Istep**: string
- **Fstep**: string
- **Timezone**: int32

### StaticMapPayload

- **Type**: string
- **Category**: string
- **Variable**: string
- **StartDate**: string
- **EndDate**: string
- **Aggregation**: string
- **Model**: string
- **Lonmin**: string
- **Lonmax**: string
- **Latmin**: string
- **Latmax**: string
- **Dpi**: uint32
- **Title**: bool
- **TitleVariable**: string
- **Hours**: int32

### RequestDetailsPayload
- **Page**: int
- **PageSize**: int
- **Path**: string
- **Status**: string
- **StartDate**: string
- **EndDate**: string

### PlanInfoPayload
- **Timezone**: int32

### PlanLocalePayload
- **LocaleId**: uint32
- **Latitude**: string
- **Longitude**: string
- **StationId**: string

### StorageListPayload
- **Page**: uint32
- **PageSize**: uint32
- **FileName**: string
- **FileExtension**: string
- **FileTypes**: []string
- **StartDate**: string
- **EndDate**: string

### StorageDownloadPayload
- **FileName**: string
- **AccessKey**: string
