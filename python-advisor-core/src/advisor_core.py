from request_handler import RequestHandler
from grouped_routes import (
    ForecastAPI,
    ObservedAPI,
    CurrentWeatherAPI,
    ClimatologyAPI,
    MonitoringAlertsAPI,
    PlanAPI,
    ChartAPI,
    TmsAPI,
    SchemaAPI
)

class AdvisorCore:
    """
    Central class that encapsulates access to various routes of the Advisor API.
    """
    def __init__(self, token, retries=5, delay=5):
        base_url="https://advisor-core.climatempo.io/api/"
        request_handler = RequestHandler(base_url, token, retries, delay)
        self.forecast = ForecastAPI(request_handler) 
        """Fetch weather forecast."""
        self.observed = ObservedAPI(request_handler)
        """Fetch observed weather."""
        self.current_weather = CurrentWeatherAPI(request_handler)
        """Fetch current weather."""
        self.climatology = ClimatologyAPI(request_handler)
        """Fetch climatology weather."""
        self.monitoring = MonitoringAlertsAPI(request_handler)
        """Fetch alerts."""
        self.plan = PlanAPI(request_handler)
        """Fetch plan information."""
        self.chart = ChartAPI(request_handler)
        """Fetch weather data charts"""
        self.tms = TmsAPI(request_handler)
        """Fetch tiles map service."""
        self.schema = SchemaAPI(request_handler)
        """Get and set schema/parameters."""
