package advisorsdk

import (
	"encoding/json"
)

type AdvisorResponse interface{}

type ClimatologyPayload struct {
	LocaleId  uint32
	Latitude  string
	Longitude string
	StationId string
	Variables []string
}

type CurrentWeatherPayload struct {
	LocaleId  uint32
	Latitude  string
	Longitude string
	StationId string
	Variables []string
	Timezone  int8
}

type GeometryPayload struct {
	StartDate string
	EndDate   string
	Radius    uint32
	Geometry  string
}

type StationPayload struct {
	StationId string
	Layer     string
	Variables []string
	StartDate string
	EndDate   string
}

type RadiusPayload struct {
	LocaleId  uint32
	Latitude  string
	Longitude string
	StartDate string
	EndDate   string
	Radius    uint32
}

type SchemaPayload map[string]interface{}

type TmsPayload struct {
	Istep       string
	Fstep       string
	Server      string
	Mode        string
	Variable    string
	Aggregation string
	X           uint16
	Y           uint16
	Z           uint16
}

type WeatherPayload struct {
	LocaleId  uint32
	Latitude  string
	Longitude string
	StationId string
	StartDate string
	EndDate   string
	Variables []string
	Timezone  int8
}

func (b WeatherPayload) toQueryParams() string {
	builder := queryParamsBuilder{}

	return builder.
		addLocaleId(b.LocaleId).
		addLatLon(b.Latitude, b.Longitude).
		addStationId(b.StationId).
		addVariables(b.Variables).
		addStartDate(b.StartDate).
		addEndDate(b.EndDate).
		addTimezone(b.Timezone).
		build()
}

func (c ClimatologyPayload) toQueryParams() string {
	builder := queryParamsBuilder{}

	return builder.
		addLocaleId(c.LocaleId).
		addLatLon(c.Latitude, c.Longitude).
		addStationId(c.StationId).
		addVariables(c.Variables).
		build()
}

func (c CurrentWeatherPayload) toQueryParams() string {
	builder := queryParamsBuilder{}

	return builder.
		addLocaleId(c.LocaleId).
		addLatLon(c.Latitude, c.Longitude).
		addStationId(c.StationId).
		addVariables(c.Variables).
		addTimezone(c.Timezone).
		build()
}

func (g GeometryPayload) toQueryParams() string {
	builder := queryParamsBuilder{}

	return builder.
		addStartDate(g.StartDate).
		addEndDate(g.EndDate).
		addRadius(g.Radius).
		build()
}

func (g GeometryPayload) toBodyBytes() []byte {
	body, _ := json.Marshal(map[string]string{
		"geometry": g.Geometry,
	})

	return body
}

func (s SchemaPayload) toBodyBytes() []byte {
	body, _ := json.Marshal(s)
	return body
}

func (s StationPayload) toQueryParams() string {
	builder := queryParamsBuilder{}

	return builder.
		addStationId(s.StationId).
		addLayer(s.Layer).
		addVariables(s.Variables).
		addStartDate(s.StartDate).
		addEndDate(s.EndDate).
		build()
}

func (r RadiusPayload) toQueryParams() string {
	builder := queryParamsBuilder{}

	return builder.
		addLocaleId(r.LocaleId).
		addLatLon(r.Latitude, r.Longitude).
		addStartDate(r.StartDate).
		addEndDate(r.EndDate).
		addRadius(r.Radius).
		build()
}
