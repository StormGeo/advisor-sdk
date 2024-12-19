# Python SDK

Advisor Software Development Kit for python.

## Installation
To install this package, use the following command:`

```bash
pip install python-advisor-core
```
Make sure you're using python 3.8 or higher.

## How to use
First you need to import the SDK on your application and instancy the `AdvisorCore` class setting up your access token and needed configurations:

```python
from advisor_core_sdk import AdvisorCore

advisor = AdvisorCore("<your_token>", retries=5, delay=1000)
```

### Examples for getting data
Daily Forecast:
```python
from payloads import WeatherPayload

payload = WeatherPayload(
  locale_id="1234",
)

daily_forecast = advisor.forecast.get_daily(payload)

if daily_forecast['error']:
  print('Error trying to get data!')
  print(daily_forecast['error'])
else:
  print(daily_forecast['data'])
```
- **[WeatherPayload](#weatherPayload)**: Payload type for getting weather data type.
- **advisor**: Variable with the instance of the class
- **forecast**: Class attribute responsible for getting forecast data by different methods
- **get_daily**: Method to collect daily forecast data

Hourly Observed:
```python
from payloads import WeatherPayload

payload = WeatherPayload(
  locale_id="1234",
)

hourly_observed = advisor.observed.get_hourly(payload)

if hourly_observed['error']:
  print('Error trying to get data!')
  print(hourly_observed['error'])
else:
  print(hourly_observed['data'])
```

Observed Lightning by Geometry:
```python
from payloads import GeometryPayload

payload = GeometryPayload(
  geometry="{\"type\": \"MultiPoint\", \"coordinates\": [[-41.88, -22.74]]}",
  radius=10000
)

lightning_by_geometry = advisor.observed.get_lightning_by_geometry(payload)

if lightning_by_geometry['error']:
  print('Error trying to get data!')
  print(lightning_by_geometry['error'])
else:
  print(lightning_by_geometry['data'])
```

Tms (Tiles Map Server):
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

tms_image = advisor.tms.get(payload)

if tms_image['error']:
    print('Error trying to get data!')
    print(tms_image['error'])
else:
    with open("tms.png", "wb") as f:
        f.write(tms_image["data"])
```

Plan Information:
```python
plan_info = advisor.plan.get_info(payload)

if plan_info['error']:
    print('Error trying to get data!')
    print(plan_info['error'])
else:
    print(plan_info['data'])
```

## Response Format
All the methods returns the same pattern: 

```python
{
  "data": Any | None,
  "error": Any | None,
}
```

## Payload Types
### WeatherPayload
  - **locale_id**: str
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
  - **locale_id**: str
  - **station_id**: str
  - **latitude**: float
  - **longitude**: float
  - **variables**: List[str]

### CurrentWeatherPayload
  - **locale_id**: str
  - **station_id**: str
  - **latitude**: float
  - **longitude**: float
  - **timezone**: int
  - **variables**: List[str]

### RadiusPayload
  - **locale_id**: str
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
