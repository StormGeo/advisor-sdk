from request_handler import RequestHandler
from grouped_routes import (
    ForecastAPI,
    ObservedAPI,
    ClimatologyAPI,
    CurrentWeatherAPI,
    MonitoringAlertsAPI,
    PlanAPI,
    ChartAPI,
    TmsAPI,
    SchemaAPI
)

class AdvisorCore:
    def __init__(self, token, retries=5, delay=5):
        base_url="https://advisor-core.climatempo.io/api/"
        request_handler = RequestHandler(base_url, token, retries, delay)
        self.forecast = ForecastAPI(request_handler)
        self.observed = ObservedAPI(request_handler)
        self.current_weather = CurrentWeatherAPI(request_handler)
        self.climatology = ClimatologyAPI(request_handler)
        self.monitoring = MonitoringAlertsAPI(request_handler)
        self.plan = PlanAPI(request_handler)
        self.chart = ChartAPI(request_handler)
        self.tms = TmsAPI(request_handler)
        self.schema = SchemaAPI(request_handler)
