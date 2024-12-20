# Node SDK

Advisor Software Development Kit for nodeJS.

## Installation

To install this package, use the following command:

```bash
npm install node-advisor-core
```

Make sure you're using node v18.17 or higher.

## How to use

First you need to import the SDK on your application and instancy the `AdvisorCore` class setting up your access token and needed configurations:

```javascript
import AdvisorCore from './AdvisorCore.js'

const advisor = new AdvisorCore({
  token: '<your-token>',
  retries: 2,
  delay: 500,
})
```

### Examples for getting data

<details>
  <summary>Chart</summary>

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
</details>

<details>
  <summary>Climatology</summary>

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
</details>

<details>
  <summary>Current Weather</summary>

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
</details>

<details>
  <summary>Forecast</summary>

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
</details>

<details>
  <summary>Monitoring</summary>

  ```javascript
  let response = await advisor.monitoring.getAlerts()

  if (response.error) {
    console.log(response.error)
    console.log('Error trying to get data!')
  } else {
    console.log(response.data)
  }
  ```
</details>

<details>
  <summary>Observed</summary>

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
    geometry: "{\"type\": \"MultiPonumber\", \"coordinates\": [[-41.88, -22.74]]}",
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
</details>

<details>
  <summary>Plan Information</summary>

  ```javascript
  let response = await advisor.plan.getInfo()

  if (response.error) {
    console.log(response.error)
    console.log('Error trying to get data!')
  } else {
    console.log(response.data)
  }
  ```
</details>

<details>
  <summary>Schema/Parameter</summary>

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
</details>

<details>
  <summary>Tms (Tiles Map Server)</summary>

  ```javascript
  const payload = {
    istep="2024-12-25 10:00:00",
    fstep="2024-12-25 12:00:00",
    server="a",
    mode="forecast",
    variable="precipitation",
    aggregation="sum",
    x=2,
    y=3,
    z=4
  }

  let response = await advisor.tms.get(payload)

  if (response.error) {
    console.log(response.error)
    console.log('Error trying to get data!')
  } else {
    writeFileSync('test.png', Buffer.from(response.data))
  }
  ```
</details>

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
