from request_handler import RequestHandler
from payloads import (
    WeatherPayload,
    ClimatologyPayload,
    CurrentWeatherPayload,
    SpecificObservedPayload
)
from grouped_routes import (
    ForecastAPI,
    ObservedAPI,
    ClimatologyAPI,
    CurrentWeatherAPI,
)

class AdvisorCore:
    def __init__(self, token, base_url="https://advisor-core.climatempo.io/api/v1/", retries=2, delay=1):
        self.request_handler = RequestHandler(base_url, token, retries, delay)
        self.forecast = ForecastAPI(self.request_handler)
        self.observed = ObservedAPI(self.request_handler)
        self.current_weather = CurrentWeatherAPI(self.request_handler)
        self.climatology = ClimatologyAPI(self.request_handler)

# Usage Example
if __name__ == "__main__":
    token = "devteam-7ad97eb5fd9bfa06d55bdd3afb1a1094bc787944"
    advisor = AdvisorCore(token)

    specific_observed = SpecificObservedPayload(
        localeId="3477",
        radius="1000"
    )

    daily_forecast = advisor.forecast.get_daily_forecast(WeatherPayload(localeId="3477"))
    print(daily_forecast)

    hourly_forecast = advisor.forecast.get_hourly_forecast(WeatherPayload(localeId="3477"))
    print(hourly_forecast)

    period_forecast = advisor.forecast.get_period_forecast(WeatherPayload(localeId="3477"))
    print(period_forecast)

    daily_observed = advisor.observed.get_daily_observed(WeatherPayload(localeId="3477"))
    print(daily_observed)

    hourly_observed = advisor.observed.get_hourly_observed(WeatherPayload(localeId="3477"))
    print(hourly_observed)

    period_observed = advisor.observed.get_period_observed(WeatherPayload(localeId="3477"))
    print(period_observed)

    current_weather = advisor.current_weather.get_current_weather(CurrentWeatherPayload(localeId="3477"))
    print(current_weather)

    climatology_daily = advisor.climatology.get_daily(ClimatologyPayload(localeId="3477"))
    print(climatology_daily)
    
    climatology_monthly = advisor.climatology.get_monthly(ClimatologyPayload(localeId="3477"))
    print(climatology_monthly)

    observed_fire_focus = advisor.observed.get_observed_fire_focus(specific_observed)
    print(observed_fire_focus)

    observed_lightning = advisor.observed.get_observed_lightning(specific_observed)
    print(observed_lightning)