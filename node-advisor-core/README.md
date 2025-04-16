# Node SDK

Advisor Software Development Kit for nodeJS.

## Contents
- [Node SDK](#node-sdk)
  - [How to get your token](https://www.climatempoconsultoria.com.br/contato/)
  - [Contents](#contents)
  - [Installation](#installation)
  - [Routes](#routes)
    - [Examples](#examples)
      - [Chart:](#chart)
      - [Climatology:](#climatology)
      - [Current Weather:](#current-weather)
      - [Forecast:](#forecast)
      - [Monitoring:](#monitoring)
      - [Observed:](#observed)
      - [Storage:](#storage)
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
npm i @stormgeo/advisor-core
```

Make sure you're using node v18.17 or higher.

## Routes

First you need to import the SDK on your application and instancy the `AdvisorCore` class setting up your access token and needed configurations:

```javascript
import { AdvisorCore } from '@stormgeo/advisor-core'

const advisor = new AdvisorCore({
  token: '<your-token>',
  retries: 2,
  delay: 500,
})
```

### Examples

#### Chart:
```javascript
const payload = {
  variables: ['temperature', 'precipitation'],
  localeId: 1234,
}

// requesting daily forecast chart image
let response = await advisor.chart.getForecastDaily(payload)

// requesting hourly forecast chart image
let response = await advisor.chart.getForecastHourly(payload)

// requesting daily observed chart image
let response = await advisor.chart.getObservedDaily(payload)

// requesting hourly observed chart image
let response = await advisor.chart.getObservedHourly(payload)

if (response.error) {
  console.log(response.error)
  console.log('Error trying to get data!')
} else {
  writeFileSync('test.png', Buffer.from(response.data))
}
```


#### Climatology:
```javascript
const payload = {
  variables: ['temperature', 'precipitation'],
  localeId: 1234,
}

// requesting daily climatology data
let response = await advisor.climatology.getDaily(payload)

// requesting monthly climatology data
let response = await advisor.climatology.getMonthly(payload)


if (response.error) {
  console.log(response.error)
  console.log('Error trying to get data!')
} else {
  console.log(response.data)
}
```


#### Current Weather:
```javascript
const payload = {
  localeId: 1234,
}

let response = await advisor.currentWeather.get(payload)

if (response.error) {
  console.log(response.error)
  console.log('Error trying to get data!')
} else {
  console.log(response.data)
}
```

#### Forecast:
```javascript
const payload = {
  variables: ['temperature', 'precipitation'],
  localeId: 1234,
}

// requesting daily forecast data
let response = await advisor.forecast.getDaily(payload)

// requesting hourly forecast data
let response = await advisor.forecast.getHourly(payload)

// requesting period forecast data
let response = await advisor.forecast.getPeriod(payload)

if (response.error) {
  console.log(response.error)
  console.log('Error trying to get data!')
} else {
  console.log(response.data)
}
```


#### Monitoring:
```javascript
let response = await advisor.monitoring.getAlerts()

if (response.error) {
  console.log(response.error)
  console.log('Error trying to get data!')
} else {
  console.log(response.data)
}
```


#### Observed:
```javascript
const payload = {
  localeId: 1234,
}

// requesting daily observed data
let response = await advisor.observed.getDaily(payload)

// requesting hourly observed data
let response = await advisor.observed.getHourly(payload)

// requesting period observed data
let response = await advisor.observed.getPeriod(payload)

const stationPayload = {
  stationId: "ABC123abc321CBA",
}

// requesting station observed data
let response = await advisor.observed.getStationData(stationPayload)

const radiusPayload = {
  localeId: 1234,
  radius: 100,
}

// requesting fire-focus observed data
let response = await advisor.observed.getFireFocus(radiusPayload)

// requesting lightning observed data
let response = await advisor.observed.getLightning(radiusPayload)

const geometryPayload = {
  geometry: "{\"type\": \"MultiPoint\", \"coordinates\": [[-41.88, -22.74]]}",
  radius: 10000
}

// requesting fire-focus observed data by geometry
let response = await advisor.observed.getFireFocusByGeometry(geometryPayload)

// requesting lightning observed data by geometry
let response = await advisor.observed.getLightningByGeometry(geometryPayload)

if (response.error) {
  console.log(response.error)
  console.log('Error trying to get data!')
} else {
  console.log(response.data)
}
```

#### Storage:
```javascript
  const payload = {
    page: 1,
    pageSize: 2,
  }

  // Requesting the files list
  let response = await advisor.storage.listFiles(payload)

  if (response.error) {
    console.log(response.error)
    console.log('Error trying to list files!')
  } else {
    console.log(response.data)
  }
```

```javascript
  const fileName = 'Example.pdf'
  const payload = {
    fileName,
    accessKey: 'a1b2c3d4-0010',
  }

  // Download de file as a Buffer
  let response = await advisor.storage.downloadFile(payload)

  if (!response.error && response.data) {
    writeFileSync(fileName, Buffer.from(response.data))
  } else {
    console.log(response.error)
    console.log('Error trying to get data!')
  }
  
  // Downloading the file by stream
  let response = await advisor.storage.downloadFileByStream(payload)
  
  if (!response.error && response.data) {
    response.data.pipe(createWriteStream(fileName))
  } else {
    console.log(response.error)
    console.log('Error trying to get data!')
  }
```

#### Plan Information:
```javascript
//Requesting plan information
let response = await advisor.plan.getInfo()

const payload = {
  page: 1,
  pageSize: 10,
}
// Requesting access history
let response = await advisor.plan.getRequestDetails(payload)

if (response.error) {
  console.log(response.error)
  console.log('Error trying to get data!')
} else {
  console.log(response.data)
}
```

#### Schema/Parameter:
```javascript
// Arbitrary example on how to define a schema
const schemaPayload = {
  "identifier": "arbitraryIdentifier",
  "arbitraryField1": {
      "type": "boolean",
      "required": True,
      "length": 125,
  },
}

// Arbitrary example on how to upload data to parameters from schema 
parametersPayload = {
  "identifier": "arbitraryIdentifier",
  "arbitraryField1": True,
}

// requesting all schemas from token
let response = await advisor.schema.getDefinition()

// requesting to upload a new schema
let response = await advisor.schema.postDefinition(schemaPayload)

// requesting to upload data to parameters from schema
let response = await advisor.schema.postParameters(parametersPayload)

if (response.error) {
  console.log(response.error)
  console.log('Error trying to get data!')
} else {
  console.log(response.data)
}
```


#### Tms (Tiles Map Server):
```javascript
const payload = {
  istep: "2024-12-25 10:00:00",
  fstep: "2024-12-25 12:00:00",
  server: "a",
  mode: "forecast",
  variable: "precipitation",
  aggregation: "sum",
  x: 2,
  y: 3,
  z: 4
}

let response = await advisor.tms.get(payload)

if (response.error) {
  console.log(response.error)
  console.log('Error trying to get data!')
} else {
  writeFileSync('test.png', Buffer.from(response.data))
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

```javascript
const advisor = new AdvisorCore({
  token: 'invalid-token',
})

advisor.setHeaderAccept('application/xml')
advisor.setHeaderAcceptLanguage('es-ES')

let response = await advisor.plan.getInfo()

console.log(response.error)

// <response>
//   <error>
//     <type>UNAUTHORIZED_ACCESS</type>
//     <message>UNAUTHORIZED_REQUEST</message>
//     <description>La solicitud no está autorizada.</description>
//   </error>
// </response>
```


## Response Format

All the methods returns the same pattern:

```javascript
{
  "data": Any | null,
  "error": Any | null,
}
```

## Payload Types

### WeatherPayload

- **localeId**: string
- **stationId**: string
- **latitude**: number
- **longitude**: number
- **timezone**: number
- **variables**: string[]
- **startDate**: string
- **endDate**: string

### StationPayload

- **stationId**: string
- **layer**: string
- **variables**: string[]
- **startDate**: string
- **endDate**: string

### ClimatologyPayload

- **localeId**: string
- **stationId**: string
- **latitude**: number
- **longitude**: number
- **variables**: string[]

### CurrentWeatherPayload

- **localeId**: string
- **stationId**: string
- **latitude**: number
- **longitude**: number
- **timezone**: number
- **variables**: string[]

### RadiusPayload

- **localeId**: string
- **stationId**: string
- **latitude**: number
- **longitude**: number
- **startDate**: string
- **endDate**: string
- **radius**: number

### GeometryPayload

- **startDate**: string
- **endDate**: string
- **radius**: number
- **geometry**: string

### TmsPayload

- **server**: string
- **mode**: string
- **variable**: string
- **aggregation**: string
- **x**: number
- **y**: number
- **z**: number
- **istep**: string
- **fstep**: string
