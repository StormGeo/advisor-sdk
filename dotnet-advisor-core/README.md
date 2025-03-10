# .NET SDK

Advisor Software Development Kit for .NET. 

## Contents
- [.NET SDK](#net-sdk)
  - [How to get your token](https://www.climatempoconsultoria.com.br/contato/)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Routes](#routes)
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
dotnet add package StormGeo.AdvisorCore
```

## Usage

First you need to import the SDK on your application and instancy the `AdvisorCore` class setting up your access token:

```c#
using StormGeo.AdvisorCore;

// Default configuration
var sdk = new AdvisorCore("<your-token>");

// Set number of attempts to get data (default=5)
var sdk = new AdvisorCore("<your-token>", attempts: 3);

// Set the delay in seconds between one request and another (default=5)
var sdk = new AdvisorCore("<your-token>", delayInSeconds: 2);

// Set a logger class (default=null)
using Microsoft.Extensions.Logging;

ILogger loggerInstance = ...;
var sdk = new AdvisorCore("<your-token>", logger: loggerInstance);
```

The recommended .NET version to use this library is the net7, net8, net9 or a later version. Using earlier .NET versions may result in unexpected behavior or incompatibilities.

### Routes

#### Chart:
```c#
using StormGeo.AdvisorCore.Payloads;

var payload = new WeatherPayload()
{
  LocaleId = 3477,
  Variables = [ "temperature" ],
};

// requesting daily forecast chart image
var response = await sdk.Chart.GetForecastDailyAsync(payload);

// requesting hourly forecast chart image
var response = await sdk.Chart.GetForecastHourlyAsync(payload);

// requesting daily observed chart image
var response = await sdk.Chart.GetObservedDailyAsync(payload);

// requesting hourly observed chart image
var response = await sdk.Chart.GetObservedHourlyAsync(payload);

if (response.Error == null && response.Data != null)
{
    var filename = "chart.png";

    using (var fileStream = new FileStream(filename, FileMode.Create))
    {
        await response.Data.CopyToAsync(fileStream);
    }
}
else
{
    Console.WriteLine(response.Error); // string
    Console.WriteLine(response.JsonDeserializeError()); // ExpandoObject if 'Error' is a json string
}
```


#### Climatology:
```c#
using StormGeo.AdvisorCore.Payloads;

var payload = new ClimatologyPayload()
{
    LocaleId = 3477,
    Variables = ["temperature"],
};

// requesting daily climatology data
var response = await sdk.Climatology.GetDailyAsync(payload);

// requesting monthly climatology data
var response = await sdk.Climatology.GetMonthlyAsync(payload);

if (response.Error == null && response.Data != null)
{
    Console.WriteLine(response.Data); // string
    Console.WriteLine(response.JsonDeserializeData()); // ExpandoObject if 'Data' is a json string
}
else
{
    Console.WriteLine(response.Error); // string
    Console.WriteLine(response.JsonDeserializeError()); // ExpandoObject if 'Error' is a json string
}
```


#### Current Weather:
```c#
using StormGeo.AdvisorCore.Payloads;

var payload = new CurrentWeatherPayload()
{
    LocaleId = 3477,
    Variables = ["temperature"],
};

var response = await sdk.CurrentWeather.GetAsync(payload);

if (response.Error == null && response.Data != null)
{
    Console.WriteLine(response.Data); // string
    Console.WriteLine(response.JsonDeserializeData()); // ExpandoObject if 'Data' is a json string
}
else
{
    Console.WriteLine(response.Error); // string
    Console.WriteLine(response.JsonDeserializeError()); // ExpandoObject if 'Error' is a json string
}
```


#### Forecast:
```c#
using StormGeo.AdvisorCore.Payloads;

var payload = new WeatherPayload()
{
    LocaleId = 3477,
};

// requesting daily forecast data
var response = await sdk.Forecast.GetDailyAsync(payload);

// requesting hourly forecast data
var response = await sdk.Forecast.GetHourlyAsync(payload);

// requesting period forecast data
var response = await sdk.Forecast.GetPeriodAsync(payload);

if (response.Error == null && response.Data != null)
{
    Console.WriteLine(response.Data); // string
    Console.WriteLine(response.JsonDeserializeData()); // ExpandoObject if 'Data' is a json string
}
else
{
    Console.WriteLine(response.Error); // string
    Console.WriteLine(response.JsonDeserializeError()); // ExpandoObject if 'Error' is a json string
}
```


#### Monitoring:
```c#
var response = await sdk.Monitoring.GetAlertsAsync();

if (response.Error == null && response.Data != null)
{
    Console.WriteLine(response.Data); // string
    Console.WriteLine(response.JsonDeserializeData()); // ExpandoObject if 'Data' is a json string
}
else
{
    Console.WriteLine(response.Error); // string
    Console.WriteLine(response.JsonDeserializeError()); // ExpandoObject if 'Error' is a json string
}
```


#### Observed:
```c#
using StormGeo.AdvisorCore.Payloads;

var payload = new WeatherPayload() {
    LocaleId = 3477,
};

// requesting daily observed data
var response = await sdk.Observed.GetDailyAsync(payload);

// requesting hourly observed data
var response = await sdk.Observed.GetHourlyAsync(payload);

// requesting period observed data
var response = await sdk.Observed.GetPeriodAsync(payload);

var stationPayload = new StationPayload() {
    StationId = "<station-id>",
    Layer = "daily",
};

// requesting station observed data
var response = await sdk.Observed.GetStationDataAsync(stationPayload);

var radiusPayload = new RadiusPayload() {
    LocaleId = 3477,
    Radius = 100,
};

// requesting fire-focus observed data
var response = await sdk.Observed.GetFireFocusAsync(radiusPayload);

// requesting lightning observed data
var response = await sdk.Observed.GetLightningAsync(radiusPayload);

var geometryPayload = new GeometryPayload() {
    StartDate = "2024-11-28 00:00:00",
    EndDate = "2024-11-28 23:59:59",
    Geometry = "{\"type\": \"MultiPoint\", \"coordinates\": [[-41.88, -22.74]]}",
};

// requesting fire-focus observed data by geometry
var response = await sdk.Observed.GetFireFocusByGeometryAsync(geometryPayload);

// requesting lightning observed data by geometry
var response = await sdk.Observed.GetLightningByGeometryAsync(geometryPayload);

if (response.Error == null && response.Data != null)
{
    Console.WriteLine(response.Data); // string
    Console.WriteLine(response.JsonDeserializeData()); // ExpandoObject if 'Data' is a json string
}
else
{
    Console.WriteLine(response.Error); // string
    Console.WriteLine(response.JsonDeserializeError()); // ExpandoObject if 'Error' is a json string
}
```


#### Plan Information:
```c#
using StormGeo.AdvisorCore.Payloads;

// requesting plan details
var response = await sdk.Plan.GetInfoAsync();

// requesting api access logs
var requestDetailsPayload = new RequestDetailsPayload(1, 10);
var response = await sdk.Plan.GetRequestDetailsAsync(requestDetailsPayload);

if (response.Error == null && response.Data != null)
{
    Console.WriteLine(response.Data); // string
    Console.WriteLine(response.JsonDeserializeData()); // ExpandoObject if 'Data' is a json string
}
else
{
    Console.WriteLine(response.Error); // string
    Console.WriteLine(response.JsonDeserializeError()); // ExpandoObject if 'Error' is a json string
}
```

#### Schema/Parameter:
```c#
using StormGeo.AdvisorCore.Payloads;

// Arbitrary example on how to define a schema
var schemaDefinition = new Dictionary<string, SchameDefinitionField>()
{
    { "arbitraryField1", new() { Type = "string", Required = true } }
};

// Arbitrary example on how to upload data to parameters from schema
var schemaParameters = new Dictionary<string, object>()
{
    { "arbitraryField1", "testing" }
};

// requesting all schemas from token
var response = await sdk.Schema.GetDefinitionAsync();

// requesting to upload a new schema
var response = await sdk.Schema.PostDefinitionAsync("test", schemaDefinition);

// requesting to upload data to parameters from schema
var response = await sdk.Schema.PostParametersAsync("test", schemaParameters);

if (response.Error == null && response.Data != null)
{
    Console.WriteLine(response.Data); // string
    Console.WriteLine(response.JsonDeserializeData()); // ExpandoObject if 'Data' is a json string
}
else
{
    Console.WriteLine(response.Error); // string
    Console.WriteLine(response.JsonDeserializeError()); // ExpandoObject if 'Error' is a json string
}
```


#### Tms (Tiles Map Server):
```c#
using StormGeo.AdvisorCore.Payloads;

var tmsPayload = new TmsPayload() {
    Server = 'a',
    Mode = "forecast",
    Variable = "precipitation",
    Aggregation = "sum",
    X = 2,
    Y = 3,
    Z = 3,
    Istep = "2025-02-13 00:00:00",
    Fstep = "2025-02-13 23:59:59"
};

var response = await sdk.Tms.Get(tmsPayload);

if (response.Error == null && response.Data != null)
{
    var filename = "tile.png";

    using (var fileStream = new FileStream(filename, FileMode.Create))
    {
        await response.Data.CopyToAsync(fileStream);
    }
}
else
{
    Console.WriteLine(response.Error); // string
    Console.WriteLine(response.JsonDeserializeError()); // ExpandoObject if 'Error' is a json string
}
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

```c#
var sdk = new AdvisorCore("invalid-token");

sdk.SetHeaderAccept("application/xml");
sdk.SetHeaderAcceptLanguage("es-ES");

var response = await sdk.Plan.GetInfoAsync();
Console.WriteLine(response.Error);

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

All the methods return AdvisorResponse, this class contains the attributes and methods bellow:

- Data -> string | Stream | null
- Error -> string | null
- JsonDeserializeData() -> Return a ExpandoObject if 'Data' is a json string
- JsonDeserializeError() -> Return a ExpandoObject if 'Error' is a json string

## Payload Types

### WeatherPayload

- **LocaleId**: int
- **StationId**: string
- **Latitude**: string
- **Longitude**: string
- **Timezone**: int
- **Variables**: string[]
- **StartDate**: string
- **EndDate**: string

### StationPayload

- **StationId**: string
- **Layer**: string
- **Variables**: string[]
- **StartDate**: string
- **EndDate**: string

### ClimatologyPayload

- **LocaleId**: int
- **StationId**: string
- **Latitude**: string
- **Longitude**: string
- **Variables**: string[]

### CurrentWeatherPayload

- **LocaleId**: int
- **StationId**: string
- **Latitude**: string
- **Longitude**: string
- **Timezone**: int
- **Variables**: string[]

### RadiusPayload

- **LocaleId**: int
- **Latitude**: string
- **Longitude**: string
- **StartDate**: string
- **EndDate**: string
- **Radius**: int

### GeometryPayload

- **StartDate**: string
- **EndDate**: string
- **Radius**: int
- **Geometry**: string

### TmsPayload

- **Server**: char
- **Mode**: string
- **Variable**: string
- **Aggregation**: string
- **X**: int
- **Y**: int
- **Z**: int
- **Istep**: string
- **Fstep**: string
