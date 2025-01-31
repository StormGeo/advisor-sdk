# Go SDK

Advisor Software Development Kit for Go.

## Contents
- [Go SDK](#go-sdk)
  - [If you do not have a token, please contact us](https://www.climatempoconsultoria.com.br/contato/)
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
      - [Schema/Parameter:](#schemaparameter)
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
resp, respErr := advisor.Plan.GetInfo()

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

- **localeId**: string
- **stationId**: string
- **latitude**: uint32
- **longitude**: uint32
- **timezone**: int8
- **variables**: []string
- **startDate**: string
- **endDate**: string

### StationPayload

- **stationId**: string
- **layer**: string
- **variables**: []string
- **startDate**: string
- **endDate**: string

### ClimatologyPayload

- **localeId**: string
- **stationId**: string
- **latitude**: uint32
- **longitude**: uint32
- **variables**: []string

### CurrentWeatherPayload

- **localeId**: string
- **stationId**: string
- **latitude**: uint32
- **longitude**: uint32
- **timezone**: int8
- **variables**: []string

### RadiusPayload

- **localeId**: string
- **stationId**: string
- **latitude**: uint32
- **longitude**: uint32
- **startDate**: string
- **endDate**: string
- **radius**: uint32

### GeometryPayload

- **startDate**: string
- **endDate**: string
- **radius**: uint32
- **geometry**: string

### TmsPayload

- **server**: string
- **mode**: string
- **variable**: string
- **aggregation**: string
- **x**: uint16
- **y**: uint16
- **z**: uint16
- **istep**: string
- **fstep**: string
