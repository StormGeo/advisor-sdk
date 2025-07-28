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
		a.params = append(a.params, "startDate="+url.QueryEscape(startDate))
	}

	return a
}

func (a *queryParamsBuilder) addEndDate(endDate string) *queryParamsBuilder {
	if endDate != "" {
		a.params = append(a.params, "endDate="+url.QueryEscape(endDate))
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

func (a *queryParamsBuilder) addPage(page uint32) *queryParamsBuilder {
	if page > 0 {
		a.params = append(a.params, fmt.Sprintf("page=%d", page))
	}

	return a
}

func (a *queryParamsBuilder) addPageSize(pageSize uint32) *queryParamsBuilder {
	if pageSize > 0 {
		a.params = append(a.params, fmt.Sprintf("pageSize=%d", pageSize))
	}

	return a
}

func (a *queryParamsBuilder) addPath(path string) *queryParamsBuilder {
	if path != "" {
		a.params = append(a.params, "path="+url.QueryEscape(path))
	}

	return a
}

func (a *queryParamsBuilder) addStatus(status uint32) *queryParamsBuilder {
	if status > 0 {
		a.params = append(a.params, fmt.Sprintf("status=%d", status))
	}

	return a
}

func (a queryParamsBuilder) build() string {
	return strings.Join(a.params, "&")
}
