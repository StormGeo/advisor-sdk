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

func (a *queryParamsBuilder) addFileTypes(fileTypes []string) *queryParamsBuilder {
	if len(fileTypes) != 0 {
		for _, fileType := range fileTypes {
			a.params = append(a.params, "fileTypes[]="+fileType)
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

func (a *queryParamsBuilder) addAccessKey(accessKey string) *queryParamsBuilder {
	if accessKey != "" {
		a.params = append(a.params, "accessKey="+url.QueryEscape(accessKey))
	}

	return a
}

func (a *queryParamsBuilder) addFileName(fileName string) *queryParamsBuilder {
	if fileName != "" {
		a.params = append(a.params, "fileName="+url.QueryEscape(fileName))
	}

	return a
}

func (a *queryParamsBuilder) addFileExtension(fileExtension string) *queryParamsBuilder {
	if fileExtension != "" {
		a.params = append(a.params, "fileExtension="+url.QueryEscape(fileExtension))
	}

	return a
}

func (a *queryParamsBuilder) addAggregation(aggregation string) *queryParamsBuilder {
	if aggregation != "" {
		a.params = append(a.params, "aggregation="+url.QueryEscape(aggregation))
	}

	return a
}

func (a *queryParamsBuilder) addModel(model string) *queryParamsBuilder {
	if model != "" {
		a.params = append(a.params, "model="+url.QueryEscape(model))
	}

	return a
}

func (a *queryParamsBuilder) addBBox(lonmin, lonmax, latmin, latmax string) *queryParamsBuilder {
	if lonmin != "" || lonmax != "" || latmin != "" || latmax != "" {
		if lonmin != "" {
			a.params = append(a.params, "lonmin="+lonmin)
		}
		if lonmax != "" {
			a.params = append(a.params, "lonmax="+lonmax)
		}
		if latmin != "" {
			a.params = append(a.params, "latmin="+latmin)
		}
		if latmax != "" {
			a.params = append(a.params, "latmax="+latmax)
		}
	}

	return a
}

func (a *queryParamsBuilder) addDpi(dpi int32) *queryParamsBuilder {
	if dpi > 0 {
		a.params = append(a.params, fmt.Sprintf("dpi=%d", dpi))
	}

	return a
}

func (a *queryParamsBuilder) addTitle(title bool) *queryParamsBuilder {
	if title {
		a.params = append(a.params, "title=true")
	} else {
		a.params = append(a.params, "title=false")
	}

	return a
}

func (a *queryParamsBuilder) addTitleVariable(titleVariable string) *queryParamsBuilder {
	if titleVariable != "" {
		a.params = append(a.params, "titlevariable="+url.QueryEscape(titleVariable))
	}

	return a
}

func (a *queryParamsBuilder) addHours(hours int32) *queryParamsBuilder {
	if hours > 0 {
		a.params = append(a.params, fmt.Sprintf("hours=%d", hours))
	}

	return a
}

func (a queryParamsBuilder) build() string {
	return strings.Join(a.params, "&")
}
