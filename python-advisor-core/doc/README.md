# Python SDK

Advisor Software Development Kit for python.

## Contents
- [Importing](#importing)
- [Routes](#routes)
    - [Chart](#chart)
    - [Climatology](#climatology)
    - [Current Weather](#current-weather)
    - [Forecast](#forecast)
    - [Monitoring](#monitoring)
    - [Observed](#observed)
    - [Plan Information](#plan-information)
    - [Schema/Parameter](#schemaparameter)
    - [Tms (Tiles Map Server)](#tms-tiles-map-server)
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
## Importing

To install this package, use the following command:`

```bash
pip install python-advisor-core
```

Make sure you're using python 3.8 or higher.


## Routes

First you need to import the SDK on your application and instancy the `AdvisorCore` class setting up your access token and needed configurations:

```python
from advisor_core_sdk import AdvisorCore

advisor = AdvisorCore("<your_token>", retries=5, delay=5)
```

### Examples
Get data from different routes with theses examples

#### Chart
```python
from payloads import WeatherPayload

payload = WeatherPayload(
  locale_id=1234,
  variables=["temperature", "precipitation"]
)

# requesting daily forecast chart image
response = advisor.chart.get_forecast_daily(payload)

# requesting hourly forecast chart image
response = advisor.chart.get_forecast_hourly(payload)

# requesting daily observed chart image
response = advisor.chart.get_observed_daily(payload)

# requesting hourly observed chart image
response = advisor.chart.get_observed_hourly(payload)

if response['error']:
  print('Error trying to get data!')
  print(response['error'])
else:
  with open("response.png", "wb") as f:
    f.write(response["data"])
```

#### Climatology
```python
from payloads import ClimatologyPayload

payload = ClimatologyPayload(
  locale_id=1234,
  variables=["temperature", "precipitation"]
)

# requesting daily climatology data
response = advisor.climatology.get_daily(payload)

# requesting monthly climatology data
response = advisor.climatology.get_monthly(payload)

if response['error']:
  print('Error trying to get data!')
  print(response['error'])
else:
  print(response['data'])
```

#### Current Weather
```python
from payloads import CurrentWeatherPayload

payload = CurrentWeatherPayload(
  locale_id=1234
)

response = advisor.current_weather.get(payload)

if response['error']:
  print('Error trying to get data!')
  print(response['error'])
else:
  print(response['data'])
```


#### Forecast
```python
from payloads import WeatherPayload

payload = WeatherPayload(
  locale_id=1234,
  variables=["temperature", "precipitation"]
)

# requesting daily forecast data
response = advisor.forecast.get_daily(payload)

# requesting hourly forecast data
response = advisor.forecast.get_hourly(payload)

# requesting period forecast data
response = advisor.forecast.get_period(payload)

if response['error']:
  print('Error trying to get data!')
  print(response['error'])
else:
  print(response['data'])
```

#### Monitoring
```python
response = advisor.monitoring.get_alerts()

if response['error']:
  print('Error trying to get data!')
  print(response['error'])
else:
  print(response['data'])
```

#### Observed
```python
from payloads import (WeatherPayload, StationPayload, RadiusPayload, GeometryPayload)

payload = WeatherPayload(
  locale_id=1234,
)

payload_for_station = StationPayload(
  station_id="ABC123abc321CBA"
)

payload_for_radius = RadiusPayload(
  locale_id=1234,
  radius=1000
)

payload_for_geometry = GeometryPayload(
  geometry="{\"type\": \"MultiPoint\", \"coordinates\": [[-41.88, -22.74]]}",
  radius=10000
)

# requesting daily observed data
response = advisor.observed.get_daily(payload)

# requesting hourly observed data
response = advisor.observed.get_hourly(payload)

# requesting period observed data
response = advisor.observed.get_period(payload)

# requesting station observed data
response = advisor.observed.get_station_data(payload_for_station)

# requesting fire-focus observed data
response = advisor.observed.get_fire_focus(payload_for_radius)

# requesting lightning observed data
response = advisor.observed.get_lightning(payload_for_radius)

# requesting fire-focus observed data by geometry
response = advisor.observed.get_fire_focus_by_geometry(payload_for_geometry)

# requesting lightning observed data by geometry
response = advisor.observed.get_lightning_by_geometry(payload_for_geometry)

if response['error']:
  print('Error trying to get data!')
  print(response['error'])
else:
  print(response['data'])
```

#### Plan Information
```python
response = advisor.plan.get_info()

if response['error']:
  print('Error trying to get data!')
  print(response['error'])
else:
  print(response['data'])
```

#### Schema/Parameter
```python
# Arbitrary example on how to define a schema
payload_schema_definition = {
  "identifier": "arbitraryIdentifier",
  "arbitraryField1": {
      "type": "boolean",
      "required": True,
      "length": 125,
  },
  "arbitraryField2": {
      "type": "number",
      "required": True,
  },
  "arbitraryField3": {
      "type": "string",
      "required": False,
  }
}

# Arbitrary example on how to upload data to parameters from schema 
payload_schema_parameters = {
  "identifier": "arbitraryIdentifier",
  "arbitraryField1": True,
  "arbitraryField2": 15
}

# requesting all schemas from token
response = advisor.schema.get_definition()

# requesting to upload a new schema
response = advisor.schema.post_definition(payload_schema_definition)

# requesting to upload data to parameters from schema
response = advisor.schema.post_parameters(payload_schema_parameters)

if response['error']:
  print('Error trying to get data!')
  print(response['error'])
else:
  print(response['data'])
```

#### Tms (Tiles Map Server)
```python
from payloads import TmsPayload

payload = TmsPayload(
  istep="2024-12-25 10:00:00",
  fstep="2024-12-25 12:00:00",
  server="a",
  mode="forecast",
  variable="precipitation",
  aggregation="sum",
  x=2,
  y=3,
  z=4
)

response = advisor.tms.get(payload)

if response['error']:
  print('Error trying to get data!')
  print(response['error'])
else:
  with open("response.png", "wb") as f:
    f.write(response["data"])
```

---
## Response Format

All the methods will return the same pattern:

```python
{
  "data": Any | None,
  "error": Any | None,
}
```

## Payload Types

### WeatherPayload

- **locale_id**: int
- **station_id**: str
- **latitude**: float
- **longitude**: float
- **timezone**: int
- **variables**: List[str]
- **start_date**: str
- **end_date**: str

### StationPayload

- **station_id**: str
- **layer**: str
- **variables**: List[str]
- **start_date**: str
- **end_date**: str

### ClimatologyPayload

- **locale_id**: int
- **station_id**: str
- **latitude**: float
- **longitude**: float
- **variables**: List[str]

### CurrentWeatherPayload

- **locale_id**: int
- **station_id**: str
- **latitude**: float
- **longitude**: float
- **timezone**: int
- **variables**: List[str]

### RadiusPayload

- **locale_id**: int
- **station_id**: str
- **latitude**: float
- **longitude**: float
- **start_date**: str
- **end_date**: str
- **radius**: int

### GeometryPayload

- **start_date**: str
- **end_date**: str
- **radius**: int
- **geometry**: str

### TmsPayload

- **server**: str
- **mode**: str
- **variable**: str
- **aggregation**: str
- **x**: int
- **y**: int
- **z**: int
- **istep**: str
- **fstep**: str
---