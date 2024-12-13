from query_builder import QueryParamsBuilder
from payloads import (
    WeatherPayload,
    CurrentWeatherPayload,
    ClimatologyPayload,
    SpecificObservedPayload,
    StationPayload,
    ObservedGeometryPayload,
    TmsPayload
)

class ForecastAPI:
    def __init__(self, request_handler):
        self.request_handler = request_handler

    def get_daily(self, payload: WeatherPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getParams())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "v1/forecast/daily", params=params)

    def get_hourly(self, payload: WeatherPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getParams())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "v1/forecast/hourly", params=params)
    
    def get_period(self, payload: WeatherPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getParams())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "v1/forecast/period", params=params)

class ObservedAPI:
    def __init__(self, request_handler):
        self.request_handler = request_handler

    def get_daily(self, payload: WeatherPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getParams())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "v1/observed/daily", params=params)

    def get_hourly(self, payload: WeatherPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getParams())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "v1/observed/hourly", params=params)
    
    def get_period(self, payload: WeatherPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getParams())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "v1/observed/period", params=params)

    def get_station_data(self, payload: StationPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getParams())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "v1/station", params=params)

    def get_fire_focus(self, payload: SpecificObservedPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getParams())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "v1/observed/fire-focus", params=params)
    
    def get_lightning(self, payload: SpecificObservedPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getParams())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "v1/observed/lightning", params=params)
    
    def post_fire_focus(self, payload: ObservedGeometryPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getParams())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("POST", "v1/observed/fire-focus", params=params, json_data=payload.getBody())
    
    def post_lightning(self, payload: ObservedGeometryPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getParams())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("POST", "v1/observed/lightning", params=params, json_data=payload.getBody())

class CurrentWeatherAPI:
    def __init__(self, request_handler):
        self.request_handler = request_handler

    def get(self, payload: CurrentWeatherPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getParams())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "v1/forecast/daily", params=params)

class ClimatologyAPI:
    def __init__(self, request_handler):
        self.request_handler = request_handler

    def get_daily(self, payload: ClimatologyPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getParams())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "v1/climatology/daily", params=params)

    def get_monthly(self, payload: ClimatologyPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getParams())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "v1/climatology/monthly", params=params)

class MonitoringAlertsAPI:
    def __init__(self, request_handler):
        self.request_handler = request_handler

    def get_alerts(self):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "v1/monitoring/alerts", params=params)

class PlanAPI:
    def __init__(self, request_handler):
        self.request_handler = request_handler

    def get_info(self):
        return self.request_handler.make_request("GET", f"v1/plan/{self.request_handler.token}")

class ChartAPI:
    def __init__(self, request_handler):
        self.request_handler = request_handler

    def get_forecast_daily(self, payload: WeatherPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getParams())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "v1/forecast/daily/chart", params=params)
    
    def get_forecast_hourly(self, payload: WeatherPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getParams())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "v1/forecast/hourly/chart", params=params)
    
    def get_observed_daily(self, payload: WeatherPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getParams())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "v1/observed/daily/chart", params=params)
    
    def get_observed_hourly(self, payload: WeatherPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getParams())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "v1/observed/hourly/chart", params=params)

class TmsAPI:
    def __init__(self, request_handler):
        self.request_handler = request_handler

    def get(self, payload: TmsPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getParams())
            .add_token("token", self.request_handler.token)
            .build()
        )
        path = f"tms/{payload.server}/{payload.mode}/{payload.variable}/{payload.aggregation}/{payload.x}/{payload.y}/{payload.z}.png"
        return self.request_handler.make_request("GET", path, params=params)

class SchemaAPI:
    def __init__(self, request_handler):
        self.request_handler = request_handler

    def get_definition(self):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "v1/schema/definition", params=params)
    
    def post_definition(self, payload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("POST", "v1/schema/definition", params=params, json_data=payload)
    
    def post_parameters(self, payload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("POST", "v1/schema/parameters", params=params, json_data=payload)
