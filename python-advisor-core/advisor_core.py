from request_handler import RequestHandler
from grouped_routes import (
    ForecastAPI,
    ObservedAPI,
)

class AdvisorCore:
    def __init__(self, token, base_url="https://advisor-core.climatempo.io/api/v1/", retries=2, delay=1):
        self.request_handler = RequestHandler(base_url, token, retries, delay)
        self.forecast = ForecastAPI(self.request_handler)
        self.observed = ObservedAPI(self.request_handler)

# Usage Example
if __name__ == "__main__":
    token = "devteam-7ad97eb5fd9bfa06d55bdd3afb1a1094bc787944"
    advisor = AdvisorCore(token)

    daily_forecast = advisor.forecast.get_daily_forecast(
        localeId="",
        startDate="2024-12-12",
        endDate="2024-12-15",
    )
    print(daily_forecast)

    hourly_forecast = advisor.forecast.get_hourly_forecast(
        localeId="3477",
        startDate="2024-12-12 00:00:00",
        endDate="2024-12-13 00:00:00",
    )
    print(hourly_forecast)

    daily_observed = advisor.observed.get_daily_observed(
        localeId="3477",
        startDate="2024-12-01",
        endDate="2024-12-10",
    )
    print(daily_observed)

    hourly_observed = advisor.observed.get_hourly_observed(
        localeId="3477",
        startDate="2024-12-10 00:00:00",
        endDate="2024-12-11 00:00:00",
    )
    print(hourly_observed)
