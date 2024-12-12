from query_builder import QueryParamsBuilder
from payloads import (
    WeatherPayload,
    CurrentWeatherPayload,
    ClimatologyPayload,
    SpecificObservedPayload
)

class ForecastAPI:
    def __init__(self, request_handler):
        self.request_handler = request_handler

    def get_daily_forecast(self, payload: WeatherPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getDict())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "forecast/daily", params=params)

    def get_hourly_forecast(self, payload: WeatherPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getDict())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "forecast/hourly", params=params)
    
    def get_period_forecast(self, payload: WeatherPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getDict())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "forecast/period", params=params)

class ObservedAPI:
    def __init__(self, request_handler):
        self.request_handler = request_handler

    def get_daily_observed(self, payload: WeatherPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getDict())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "observed/daily", params=params)

    def get_hourly_observed(self, payload: WeatherPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getDict())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "observed/hourly", params=params)
    
    def get_period_observed(self, payload: WeatherPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getDict())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "observed/period", params=params)

    def get_observed_fire_focus(self, payload: SpecificObservedPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getDict())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "observed/fire-focus", params=params)
    
    def get_observed_lightning(self, payload: SpecificObservedPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getDict())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "observed/lightning", params=params)

class CurrentWeatherAPI:
    def __init__(self, request_handler):
        self.request_handler = request_handler

    def get_current_weather(self, payload: CurrentWeatherPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getDict())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "forecast/daily", params=params)

class ClimatologyAPI:
    def __init__(self, request_handler):
        self.request_handler = request_handler

    def get_daily(self, payload: ClimatologyPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getDict())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "climatology/daily", params=params)

    def get_monthly(self, payload: ClimatologyPayload):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add_payload(payload.getDict())
            .add_token("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "climatology/monthly", params=params)
