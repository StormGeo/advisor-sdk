package advisorsdk

import (
	"fmt"
	"net/url"
	"strings"
)

type queryParamsBuilder struct {
	params []string
}

func (a *queryParamsBuilder) addLocaleId(localeId uint32) *queryParamsBuilder {
	if localeId != 0 {
		a.params = append(a.params, fmt.Sprintf("localeId=%d", localeId))
	}

	return a
}

func (a *queryParamsBuilder) addLatLon(latitude string, longitude string) *queryParamsBuilder {
	if latitude != "" {
		a.params = append(a.params, "latitude="+latitude)
	}

	if longitude != "" {
		a.params = append(a.params, "longitude="+longitude)
	}

	return a
}

func (a *queryParamsBuilder) addStationId(stationId string) *queryParamsBuilder {
	if stationId != "" {
		a.params = append(a.params, "stationId="+stationId)
	}

	return a
}

func (a *queryParamsBuilder) addLayer(layer string) *queryParamsBuilder {
	if layer != "" {
		a.params = append(a.params, "layer="+layer)
	}

	return a
}

func (a *queryParamsBuilder) addVariables(variables []string) *queryParamsBuilder {
	if len(variables) != 0 {
		for _, variable := range variables {
			a.params = append(a.params, "variables[]="+variable)
		}
	}

	return a
}

func (a *queryParamsBuilder) addStartDate(startDate string) *queryParamsBuilder {
	if startDate != "" {
		a.params = append(a.params, "startDate="+startDate)
	}

	return a
}

func (a *queryParamsBuilder) addEndDate(endDate string) *queryParamsBuilder {
	if endDate != "" {
		a.params = append(a.params, "endDate="+endDate)
	}

	return a
}

func (a *queryParamsBuilder) addRadius(radius uint32) *queryParamsBuilder {
	if radius > 0 {
		a.params = append(a.params, fmt.Sprintf("radius=%d", radius))
	}

	return a
}

func (a *queryParamsBuilder) addTimezone(timezone int8) *queryParamsBuilder {
	if timezone >= -12 && timezone <= 12 {
		a.params = append(a.params, fmt.Sprintf("timezone=%d", timezone))
	}

	return a
}

func (a queryParamsBuilder) build() string {
	return url.QueryEscape(strings.Join(a.params, "&"))
}
