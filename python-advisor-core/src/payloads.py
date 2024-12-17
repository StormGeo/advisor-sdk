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
        """
        Initializes the WeatherPayload object with optional parameters.
        """
        self.localeId = localeId
        self.stationId = stationId 
        self.latitude = latitude
        self.longitude = longitude
        self.timezone = timezone
        self.variables = variables
        self.startDate = startDate
        self.endDate = endDate

    def getParams(self) -> dict:
        """
        Returns the parameters as a dictionary for API requests.
        """
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
        """
        Initializes the CurrentWeatherPayload object with optional parameters.
        """
        self.localeId = localeId
        self.stationId = stationId 
        self.latitude = latitude
        self.longitude = longitude
        self.timezone = timezone
        self.variables = variables
    
    def getParams(self) -> dict:
        """
        Returns the parameters as a dictionary for API requests.
        """
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
        """
        Initializes the ClimatologyPayload object with optional parameters.
        """
        self.stationId = stationId
        self.localeId = localeId
        self.latitude = latitude
        self.longitude = longitude
        self.variables = variables

    def getParams(self) -> dict:
        """
        Returns the parameters as a dictionary for API requests.
        """
        return {
            "localeId": self.localeId,
            "stationId": self.stationId,
            "latitude": self.latitude,
            "longitude": self.longitude,
            "variables": self.variables,
        }

class RadiusPayload:
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
        """
        Initializes the RadiusPayload object with optional parameters.
        """
        self.localeId = localeId
        self.stationId = stationId 
        self.latitude = latitude
        self.longitude = longitude
        self.startDate = startDate
        self.endDate = endDate
        self.radius = radius

    def getParams(self) -> dict:
        """
        Returns the parameters as a dictionary for API requests.
        """
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
        """
        Initializes the StationPayload object with optional parameters.
        """
        self.stationId = stationId 
        self.layer = layer
        self.variables = variables
        self.startDate = startDate
        self.endDate = endDate

    def getParams(self) -> dict:
        """
        Returns the parameters as a dictionary for API requests.
        """
        return {
            "stationId": self.stationId,
            "layer": self.layer,
            "variables": self.variables,
            "startDate": self.startDate,
            "endDate": self.endDate,
        }

class GeometryPayload:
    def __init__(
        self,
        radius: int = None,
        startDate: str = None,
        endDate: str = None,
        geometry: str = None
    ):
        """
        Initializes the GeometryPayload object with optional parameters.
        """
        self.radius = radius
        self.startDate = startDate
        self.endDate = endDate
        self.geometry = geometry

    def getParams(self) -> dict:
        """
        Returns the parameters as a dictionary for API requests.
        """
        return {
            "radius": self.radius,
            "startDate": self.startDate,
            "endDate": self.endDate,
        }

    def getBody(self) -> dict:
        """
        Returns the body of the request with the geometry information.
        """
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
        """
        Initializes the TmsPayload object with optional parameters.
        """
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
        """
        Returns the parameters as a dictionary for API requests.
        """
        return {
            "istep": self.istep,
            "fstep": self.fstep,
        }
