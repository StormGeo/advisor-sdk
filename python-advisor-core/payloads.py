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

    def getParams(self) -> dict:
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
    
    def getParams(self) -> dict:
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

    def getParams(self) -> dict:
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

    def getParams(self) -> dict:
        return {
            "localeId": self.localeId,
            "stationId": self.stationId,
            "latitude": self.latitude,
            "longitude": self.longitude,
            "startDate": self.startDate,
            "endDate": self.endDate,
            "radius": self.radius,
        }

class StationPayload:
    def __init__(
        self,
        stationId: str = None,
        layer: str = None,
        variables: List[str] = None,
        startDate: str = None,
        endDate: str = None,
    ):
        self.stationId = stationId 
        self.layer = layer
        self.variables = variables
        self.startDate = startDate
        self.endDate = endDate

    def getParams(self) -> dict:
        return {
            "stationId": self.stationId,
            "layer": self.layer,
            "variables": self.variables,
            "startDate": self.startDate,
            "endDate": self.endDate,
        }

class ObservedGeometryPayload:
    def __init__(
        self,
        radius: int = None,
        startDate: str = None,
        endDate: str = None,
        geometry: str = None
    ):
        self.radius = radius
        self.startDate = startDate
        self.endDate = endDate
        self.geometry = geometry

    def getParams(self) -> dict:
        return {
            "radius": self.radius,
            "startDate": self.startDate,
            "endDate": self.endDate,
        }

    def getBody(self) -> dict:
        return {
            "geometry": self.geometry,
        }

class TmsPayload:
    def __init__(
        self,
        istep: int = None,
        fstep: str = None,
        server: str = None,
        mode: str = None,
        variable: str = None,
        aggregation: str = None,
        x: int = None,
        y: int = None,
        z: int = None,
    ):
        self.istep = istep
        self.fstep = fstep
        self.server = server
        self.mode = mode
        self.variable = variable
        self.aggregation = aggregation
        self.x = x
        self.y = y
        self.z = z

    def getParams(self) -> dict:
        return {
            "istep": self.istep,
            "fstep": self.fstep,
        }
