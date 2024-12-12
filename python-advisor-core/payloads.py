from typing import List

class WeatherPayload:
    def __init__(
        self,
        localeId: str = None,
        stationId: str = None,
        latitude: float = None,
        longitude: float = None,
        timezone: float = None,
        variables: List[str] = None,
        startDate: str = None,
        endDate: str = None,
    ):
        self.localeId = localeId
        self.stationId = stationId 
        self.latitude = latitude
        self.longitude = longitude
        self.timezone = timezone
        self.variables = variables
        self.startDate = startDate
        self.endDate = endDate

    def getDict(self) -> dict:
        return {
            "localeId": self.localeId,
            "stationId": self.stationId,
            "latitude": self.latitude,
            "longitude": self.longitude,
            "timezone": self.timezone,
            "variables": self.variables,
            "startDate": self.startDate,
            "endDate": self.endDate,
        }

class CurrentWeatherPayload:
    def __init__(
        self,
        localeId: str = None,
        stationId: str = None,
        latitude: float = None,
        longitude: float = None,
        timezone: float = None,
        variables: List[str] = None,
    ):
        self.localeId = localeId
        self.stationId = stationId 
        self.latitude = latitude
        self.longitude = longitude
        self.timezone = timezone
        self.variables = variables
    
    def getDict(self) -> dict:
        return {
            "localeId": self.localeId,
            "stationId": self.stationId,
            "latitude": self.latitude,
            "longitude": self.longitude,
            "timezone": self.timezone,
            "variables": self.variables,
        }

class ClimatologyPayload:
    def __init__(
        self,
        stationId: str = None,
        localeId: str = None,
        latitude: float = None,
        longitude: float = None,
        variables: List[str] = None,
    ):
        self.stationId = stationId
        self.localeId = localeId
        self.latitude = latitude
        self.longitude = longitude
        self.variables = variables

    def getDict(self) -> dict:
        return {
            "localeId": self.localeId,
            "stationId": self.stationId,
            "latitude": self.latitude,
            "longitude": self.longitude,
            "variables": self.variables,
        }

class SpecificObservedPayload:
    def __init__(
        self,
        localeId: str = None,
        stationId: str = None,
        latitude: float = None,
        longitude: float = None,
        startDate: str = None,
        endDate: str = None,
        radius: int = None,
    ):
        self.localeId = localeId
        self.stationId = stationId 
        self.latitude = latitude
        self.longitude = longitude
        self.startDate = startDate
        self.endDate = endDate
        self.radius = radius

    def getDict(self) -> dict:
        return {
            "localeId": self.localeId,
            "stationId": self.stationId,
            "latitude": self.latitude,
            "longitude": self.longitude,
            "startDate": self.startDate,
            "endDate": self.endDate,
            "radius": self.radius,
        }
