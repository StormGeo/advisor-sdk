import os
from datetime import date, datetime, time, timedelta
import pytest

from advisor_core import (
    ClimatologyPayload,
    CurrentWeatherPayload,
    GeometryPayload,
    LightningLitePayload,
    PlanInfoPayload,
    PlanLocalePayload,
    PmtilesPayload,
    RadiusPayload,
    RequestDetailsPayload,
    StationPayload,
    StationsLastDataPayload,
    StaticMapPayload,
    StorageDownloadPayload,
    StorageListPayload,
    TmsPayload,
    WeatherPayload,
)

pytestmark = pytest.mark.integration

def _assert_json_success(response):
    assert response["error"] is None, response["error"]
    assert response["data"] is not None

def _assert_binary_success(response):
    assert response["error"] is None, response["error"]
    assert isinstance(response["data"], bytes)
    assert response["data"]

def _start_of_day(value):
    return datetime.combine(value, time.min).strftime("%Y-%m-%d %H:%M:%S")

def _end_of_day(value):
    return datetime.combine(value, time.max).replace(microsecond=0).strftime("%Y-%m-%d %H:%M:%S")

def _require_env(name):
    value = os.getenv(name)
    if not value:
        raise pytest.UsageError(
            f"Set {name} or add it to .env.integration.local before running the Python integration tests."
        )
    return value

@pytest.fixture(scope="session")
def station_id():
    return _require_env("ADVISOR_STATION_ID")

@pytest.fixture(scope="session")
def weather_payload(locale_id):
    return WeatherPayload(locale_id=locale_id, variables=["temperature"])

@pytest.fixture(scope="session")
def weather_chart_payload(locale_id):
    return WeatherPayload(locale_id=locale_id, variables=["temperature", "precipitation"])

@pytest.fixture(scope="session")
def climatology_payload(locale_id):
    return ClimatologyPayload(locale_id=locale_id, variables=["temperature"])

@pytest.fixture(scope="session")
def current_weather_payload(locale_id):
    return CurrentWeatherPayload(locale_id=locale_id)

@pytest.fixture(scope="session")
def station_payload(station_id):
    return StationPayload(station_id=station_id)

@pytest.fixture(scope="session")
def stations_last_data_payload(station_id):
    return StationsLastDataPayload(station_ids=[station_id], variables=["temperature"])

@pytest.fixture(scope="session")
def radius_payload(locale_id):
    return RadiusPayload(locale_id=locale_id, radius=10000)

@pytest.fixture(scope="session")
def observed_day():
    return date.today() - timedelta(days=1)

@pytest.fixture(scope="session")
def observed_period():
    end_day = date.today() - timedelta(days=1)
    start_day = end_day - timedelta(days=4)
    return {
        "start_date": _start_of_day(start_day),
        "end_date": _end_of_day(end_day),
    }

@pytest.fixture(scope="session")
def forecast_day():
    return date.today() + timedelta(days=1)

@pytest.fixture(scope="session")
def geometry_payload(observed_day):
    return GeometryPayload(
        geometry=_require_env("ADVISOR_GEOMETRY"),
        start_date=_start_of_day(observed_day),
        end_date=_end_of_day(observed_day),
        radius=10000,
    )

@pytest.fixture(scope="session")
def lightning_lite_payload(observed_day):
    return LightningLitePayload(
        geometry=_require_env("ADVISOR_GEOMETRY"),
        start_date=_start_of_day(observed_day),
        end_date=_end_of_day(observed_day),
        radius=10000,
        page=1,
        page_size=10
    )

@pytest.fixture(scope="session")
def storage_list_payload():
    return StorageListPayload(page=1, page_size=10)

@pytest.fixture(scope="session")
def storage_download_payload():
    return StorageDownloadPayload(
        file_name=_require_env("ADVISOR_STORAGE_FILE_NAME"),
        access_key=_require_env("ADVISOR_STORAGE_ACCESS_KEY"),
    )

@pytest.fixture(scope="session")
def plan_info_payload():
    return PlanInfoPayload(timezone=-3)

@pytest.fixture(scope="session")
def plan_locale_payload(plan_locale_id):
    return PlanLocalePayload(locale_id=plan_locale_id)

@pytest.fixture(scope="session")
def request_details_payload():
    return RequestDetailsPayload(page=1, page_size=3)

@pytest.fixture(scope="session")
def static_map_payload(observed_period):
    return StaticMapPayload(
        type="periods",
        category="observed",
        variable="temperature",
        aggregation="max",
        start_date=observed_period["start_date"],
        end_date=observed_period["end_date"],
        dpi=50,
        title=True,
        titlevariable="Static Map",
    )

@pytest.fixture(scope="session")
def tms_payload(forecast_day):
    return TmsPayload(
        istep=_start_of_day(forecast_day),
        fstep=_end_of_day(forecast_day),
        server="a",
        mode="forecast",
        variable="precipitation",
        aggregation="sum",
        x=5,
        y=8,
        z=4,
    )

@pytest.fixture(scope="session")
def pmtiles_payload(forecast_day):
    return PmtilesPayload(
        mode="forecast",
        model="ct2w15_as",
        variable="precipitation",
        aggregation="sum",
        istep=_start_of_day(forecast_day),
        fstep=(datetime.combine(forecast_day, time.min) + timedelta(hours=1)).strftime("%Y-%m-%d %H:%M:%S"),
        max_zoom=4,
    )

@pytest.fixture(scope="session")
def schema_definition_payload():
    return {
        "identifier": "schemaIdentifier",
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
        },
    }

@pytest.mark.parametrize("method_name", ["get_daily", "get_hourly", "get_period"])
def test_forecast_routes(advisor, weather_payload, method_name):
    response = getattr(advisor.forecast, method_name)(weather_payload)
    _assert_json_success(response)

@pytest.mark.parametrize("method_name", ["get_daily", "get_hourly", "get_period"])
def test_observed_weather_routes(advisor, weather_payload, method_name):
    response = getattr(advisor.observed, method_name)(weather_payload)
    _assert_json_success(response)

def test_observed_station_data(advisor, station_payload):
    response = advisor.observed.get_station_data(station_payload)
    _assert_json_success(response)

@pytest.mark.parametrize("method_name", ["get_fire_focus", "get_lightning", "get_lightning_details"])
def test_observed_radius_routes(advisor, radius_payload, method_name):
    response = getattr(advisor.observed, method_name)(radius_payload)
    _assert_json_success(response)

@pytest.mark.parametrize(
    "method_name",
    ["get_fire_focus_by_geometry", "get_lightning_by_geometry"],
)
def test_observed_geometry_routes(advisor, geometry_payload, method_name):
    response = getattr(advisor.observed, method_name)(geometry_payload)
    _assert_json_success(response)

def test_observed_lightning_lite(advisor, lightning_lite_payload):
    response = advisor.observed.get_lightning_lite(lightning_lite_payload)
    _assert_json_success(response)

def test_current_weather(advisor, current_weather_payload):
    response = advisor.current_weather.get(current_weather_payload)
    _assert_json_success(response)

@pytest.mark.parametrize("method_name", ["get_daily", "get_monthly"])
def test_climatology_routes(advisor, climatology_payload, method_name):
    response = getattr(advisor.climatology, method_name)(climatology_payload)
    _assert_json_success(response)

def test_monitoring_alerts(advisor):
    response = advisor.monitoring.get_alerts()
    _assert_json_success(response)

def test_stations_last_data(advisor, stations_last_data_payload):
    response = advisor.stations.get_last_data(stations_last_data_payload)
    _assert_json_success(response)

def test_plan_info(advisor, plan_info_payload):
    response = advisor.plan.get_info(plan_info_payload)
    _assert_json_success(response)

def test_plan_request_details(advisor, request_details_payload):
    response = advisor.plan.get_request_details(request_details_payload)
    _assert_json_success(response)

def test_plan_locale(advisor, plan_locale_payload):
    response = advisor.plan.get_locale(plan_locale_payload)
    _assert_json_success(response)

@pytest.mark.parametrize(
    "method_name",
    ["get_forecast_daily", "get_forecast_hourly", "get_observed_daily", "get_observed_hourly"],
)
def test_chart_routes(advisor, weather_chart_payload, method_name):
    response = getattr(advisor.chart, method_name)(weather_chart_payload)
    _assert_binary_success(response)

def test_storage_list_files(advisor, storage_list_payload):
    response = advisor.storage.list_files(storage_list_payload)
    _assert_json_success(response)

def test_storage_download_file(advisor, storage_download_payload):
    response = advisor.storage.download_file(storage_download_payload)
    _assert_binary_success(response)

def test_storage_download_file_by_stream(advisor, storage_download_payload):
    response = advisor.storage.download_file_by_stream(storage_download_payload)
    assert response["error"] is None, response["error"]
    chunks = list(response["data"])
    assert chunks
    assert b"".join(chunks)

def test_static_map(advisor, static_map_payload):
    response = advisor.static_map.get_static_map(static_map_payload)
    _assert_binary_success(response)

def test_tms(advisor, tms_payload):
    response = advisor.tms.get(tms_payload)
    _assert_binary_success(response)

def test_pmtiles(advisor, pmtiles_payload):
    response = advisor.pmtiles.get(pmtiles_payload)
    _assert_binary_success(response)

def test_schema_get_definition(advisor):
    response = advisor.schema.get_definition()
    _assert_json_success(response)

def test_schema_post_definition(advisor, schema_definition_payload):
    response = advisor.schema.post_definition(schema_definition_payload)
    _assert_json_success(response)
