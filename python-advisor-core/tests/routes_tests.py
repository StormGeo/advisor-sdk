import os
import uuid

import pytest

from advisor_core import (
    ClimatologyPayload,
    CurrentWeatherPayload,
    GeometryPayload,
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

DEFAULT_STATION_ID = "bWV0b3M6MDM0MTMyRjM6LTIyLjIzMTQ1MjQ4MDg0NDU2Oi00NC4yNTEzNTMwMzgzMTcx"
DEFAULT_GEOMETRY = (
    '{"type":"Polygon","coordinates":[[[-47.09861059094109,-23.280351816702165],'
    '[-47.09861059094109,-23.895097240590488],[-46.12890390857018,-23.895097240590488],'
    '[-46.12890390857018,-23.280351816702165],[-47.09861059094109,-23.280351816702165]]]}'
)


def _assert_json_success(response):
    assert response["error"] is None, response["error"]
    assert response["data"] is not None


def _assert_binary_success(response):
    assert response["error"] is None, response["error"]
    assert isinstance(response["data"], bytes)
    assert response["data"]


def _flatten_dicts(value):
    if isinstance(value, dict):
        yield value
        for nested in value.values():
            yield from _flatten_dicts(nested)
    elif isinstance(value, list):
        for nested in value:
            yield from _flatten_dicts(nested)


def _extract_storage_download_payload(data):
    for item in _flatten_dicts(data):
        file_name = item.get("fileName") or item.get("file_name") or item.get("name")
        access_key = item.get("accessKey") or item.get("access_key")
        if file_name and access_key:
            return StorageDownloadPayload(file_name=file_name, access_key=access_key)
    return None


@pytest.fixture(scope="session")
def station_id():
    return os.getenv("ADVISOR_STATION_ID", DEFAULT_STATION_ID)


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
def geometry_payload():
    return GeometryPayload(
        geometry=DEFAULT_GEOMETRY,
        start_date="2025-04-20 00:00:00",
        end_date="2025-04-20 23:59:59",
        radius=10000,
    )


@pytest.fixture(scope="session")
def storage_list_payload():
    return StorageListPayload(page=1, page_size=10)


@pytest.fixture(scope="session")
def storage_download_payload(advisor, storage_list_payload):
    file_name = os.getenv("ADVISOR_STORAGE_FILE_NAME")
    access_key = os.getenv("ADVISOR_STORAGE_ACCESS_KEY")

    if file_name or access_key:
        if not file_name or not access_key:
            raise pytest.UsageError(
                "Set both ADVISOR_STORAGE_FILE_NAME and ADVISOR_STORAGE_ACCESS_KEY, or neither."
            )
        return StorageDownloadPayload(file_name=file_name, access_key=access_key)

    response = advisor.storage.list_files(storage_list_payload)
    _assert_json_success(response)

    payload = _extract_storage_download_payload(response["data"])
    assert payload is not None, (
        "Storage list response did not include fileName/accessKey. "
        "Set ADVISOR_STORAGE_FILE_NAME and ADVISOR_STORAGE_ACCESS_KEY explicitly."
    )
    return payload


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
def static_map_payload():
    return StaticMapPayload(
        type="periods",
        category="observed",
        variable="temperature",
        aggregation="max",
        start_date="2026-03-01 00:00:00",
        end_date="2026-03-05 23:59:59",
        dpi=50,
        title=True,
        titlevariable="Static Map",
    )


@pytest.fixture(scope="session")
def tms_payload():
    return TmsPayload(
        istep="2025-08-01 00:00:00",
        fstep="2025-08-01 23:59:59",
        server="a",
        mode="forecast",
        variable="precipitation",
        aggregation="sum",
        x=5,
        y=8,
        z=4,
    )


@pytest.fixture(scope="session")
def pmtiles_payload():
    return PmtilesPayload(
        mode="forecast",
        model="ct2w15_as",
        variable="precipitation",
        aggregation="sum",
        istep="2026-03-02 00:00:00",
        fstep="2026-03-02 01:00:00",
        max_zoom=4,
    )


@pytest.fixture
def schema_identifier():
    return f"pytest-live-{uuid.uuid4().hex[:12]}"


@pytest.fixture
def schema_definition_payload(schema_identifier):
    return {
        "identifier": schema_identifier,
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


@pytest.fixture
def schema_parameters_payload(schema_identifier):
    return {
        "identifier": schema_identifier,
        "arbitraryField1": True,
        "arbitraryField2": 15,
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


@pytest.mark.parametrize("method_name", ["get_fire_focus", "get_lightning"])
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


def test_schema_post_parameters(advisor, schema_definition_payload, schema_parameters_payload):
    definition_response = advisor.schema.post_definition(schema_definition_payload)
    _assert_json_success(definition_response)

    response = advisor.schema.post_parameters(schema_parameters_payload)
    _assert_json_success(response)
