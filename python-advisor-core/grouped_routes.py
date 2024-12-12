from query_builder import QueryParamsBuilder

class ForecastAPI:
    def __init__(self, request_handler):
        self.request_handler = request_handler

    def get_daily_forecast(self, localeId, startDate=None, endDate=None):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add("localeId", localeId)
            .add("startDate", startDate)
            .add("endDate", endDate)
            .add("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "forecast/daily", params=params)

    def get_hourly_forecast(self, localeId, startDate=None, endDate=None):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add("localeId", localeId)
            .add("startDate", startDate)
            .add("endDate", endDate)
            .add("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "forecast/hourly", params=params)

class ObservedAPI:
    def __init__(self, request_handler):
        self.request_handler = request_handler

    def get_daily_observed(self, localeId, startDate=None, endDate=None):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add("localeId", localeId)
            .add("startDate", startDate)
            .add("endDate", endDate)
            .add("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "observed/daily", params=params)

    def get_hourly_observed(self, localeId, startDate=None, endDate=None):
        builder = QueryParamsBuilder()
        params = (
            builder
            .add("localeId", localeId)
            .add("startDate", startDate)
            .add("endDate", endDate)
            .add("token", self.request_handler.token)
            .build()
        )
        return self.request_handler.make_request("GET", "observed/hourly", params=params)
