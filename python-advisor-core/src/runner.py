from advisor_core import AdvisorCore
from payloads import *

if __name__ == "__main__":
    token = "devteam-7ad97eb5fd9bfa06d55bdd3afb1a1094bc787944"
    advisor = AdvisorCore(token)

    #advisor.setHeaderAccept('application/json')
    #advisor.setHeaderAccept('application/xml')
    #advisor.setHeaderAccept('text/csv')

    #advisor.setHeaderAcceptLanguage('pt-BR')
    #advisor.setHeaderAcceptLanguage('es-ES')

    specific_observed = RadiusPayload(
        locale_id="3477",
        radius=1000
    )

    observed_geometry = GeometryPayload(
        geometry="{\"type\": \"MultiPoint\", \"coordinates\": [[-41.88, -22.74]]}",
        radius=10000
    )

    tms_payload = TmsPayload(
        istep="2024-12-20 10:00:00",
        fstep="2024-12-20 12:00:00",
        server="a",
        mode="forecast",
        variable="precipitation",
        aggregation="sum",
        x=5,
        y=7,
        z=4
    )

    schema_definition_payload = {
        "identifier": "arbitraryIdentifier",
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
        }
    }

    schema_parameters_payload = {
        "identifier": "arbitraryIdentifier",
        "arbitraryField1": True,
        "arbitraryField2": 15
    }

    # # CURRENT WEATHER
    # current_weather = advisor.current_weather.get(CurrentWeatherPayload(locale_id="3477"))
    # #print(current_weather)
    # if current_weather['error']:
    #     print('Error trying to get data!')
    #     print(current_weather['error'])
    # else:
    #     print(current_weather['data'])


    # FORECAST
    daily_forecast = advisor.forecast.get_daily(WeatherPayload(locale_id=3477, variables=["temperature", "precipitation"]))
    print(daily_forecast)

    # hourly_forecast = advisor.forecast.get_hourly(WeatherPayload(locale_id="3477"))
    # # print(hourly_forecast)

    # period_forecast = advisor.forecast.get_period(WeatherPayload(locale_id="3477"))
    # # print(period_forecast)


    # # OBSERVED
    # daily_observed = advisor.observed.get_daily(WeatherPayload(locale_id="3477"))
    # # print(daily_observed)
    
    # hourly_observed = advisor.observed.get_hourly(WeatherPayload(locale_id="3477"))
    # # print(hourly_observed)

    # period_observed = advisor.observed.get_period(WeatherPayload(locale_id="3477"))
    # # print(period_observed)
    
    # get_fire_focus = advisor.observed.get_fire_focus(specific_observed)
    # # print(get_fire_focus)

    # get_fire_focus_by_geometry = advisor.observed.get_fire_focus_by_geometry(observed_geometry)
    # # print(get_fire_focus_by_geometry)

    # get_lightning = advisor.observed.get_lightning(specific_observed)
    # # print(get_lightning)

    # get_lightning_by_geometry = advisor.observed.get_lightning_by_geometry(observed_geometry)
    # # print(get_lightning_by_geometry)

    # station = advisor.observed.get_station_data(StationPayload(station_id="bWV0b3M6MDM0MTMyRjM6LTIyLjIzMTQ1MjQ4MDg0NDU2Oi00NC4yNTEzNTMwMzgzMTcx"))
    # # print(station)


    # # CLIMATOLOGY
    # climatology_daily = advisor.climatology.get_daily(ClimatologyPayload(locale_id="3477"))
    # # print(climatology_daily)
    
    # climatology_monthly = advisor.climatology.get_monthly(ClimatologyPayload(locale_id="3477"))
    # # print(climatology_monthly)


    # # MONITORING
    # alerts = advisor.monitoring.get_alerts()
    # # print(alerts)


    # # PLAN INFORMATION
    # plan_info = advisor.plan.get_info()
    # # print(plan_info)


    # # CHART
    # forecast_daily_chart = advisor.chart.get_forecast_daily(WeatherPayload(locale_id="3477"))
    # if forecast_daily_chart['error']:
    #     print('Error trying to get data!')
    #     print(forecast_daily_chart['error'])
    # else:
    #     with open("forecast_daily_chart.png", "wb") as f:
    #         f.write(forecast_daily_chart["data"])

    # forecast_hourly_chart = advisor.chart.get_forecast_hourly(WeatherPayload(locale_id="3477"))
    # if forecast_hourly_chart['error']:
    #     print('Error trying to get data!')
    #     print(forecast_hourly_chart['error'])
    # else:
    #     with open("forecast_hourly_chart.png", "wb") as f:
    #         f.write(forecast_hourly_chart["data"])

    # observed_daily_chart = advisor.chart.get_observed_daily(WeatherPayload(locale_id="3477"))
    # if observed_daily_chart['error']:
    #     print('Error trying to get data!')
    #     print(observed_daily_chart['error'])
    # else:
    #     with open("observed_daily_chart.png", "wb") as f:
    #         f.write(observed_daily_chart["data"])

    # observed_hourly_chart = advisor.chart.get_observed_hourly(WeatherPayload(locale_id="3477"))
    # if observed_hourly_chart['error']:
    #     print('Error trying to get data!')
    #     print(observed_hourly_chart['error'])
    # else:
    #     with open("observed_hourly_chart.png", "wb") as f:
    #         f.write(observed_hourly_chart["data"])


    # # TMS
    # tms_image = advisor.tms.get(tms_payload)
    # if tms_image['error']:
    #     print('Error trying to get data!')
    #     print(tms_image['error'])
    # else:
    #     with open("tms_image.png", "wb") as f:
    #         f.write(tms_image["data"])


    # # SCHEMA/PARAMETER
    # get_schema_definition = advisor.schema.get_definition()
    # print(get_schema_definition)

    # post_schema_definition = advisor.schema.post_definition(schema_definition_payload)
    # print(post_schema_definition)

    # post_schema_parameters = advisor.schema.post_parameters(schema_parameters_payload)
    # print(post_schema_parameters)
