package models

import (
	"net/url"
	"strings"
	"github.com/go-ozzo/ozzo-validation"
	"solardatabase/utils"
	"strconv"
)

type TemporalRequest struct {
	StartTime	string
	EndTime	string
	TableNames	[]string
	SortBy string
	Limit int
	Offset int
	Interpolated bool
}


func CreateTemporalRequest(input url.Values)  (r TemporalRequest, err error) {

	r.StartTime = input.Get("starttime")
	r.EndTime = input.Get("endtime")
	r.TableNames = strings.Split(input.Get("tablenames"), ",")
	r.SortBy = input.Get("sortby")

	r.Limit, err = strconv.Atoi(input.Get("limit"))

	if err != nil {
		return r, err
	}

	r.Offset, err = strconv.Atoi(input.Get("offset"))

	if err != nil {
		return r, err
	}

	interpolated := input.Get("interpolated")

	if interpolated == "True" {
		r.Interpolated = true
	} else {
		r.Interpolated = false
	}
	err = r.Validate()
	return r, err
}



func (r TemporalRequest) Validate() error {

	var sTime, err1 = utils.ParseDatetime(r.StartTime)

	if err1 != nil {
		return err1
	}
	var eTime, err2 = utils.ParseDatetime(r.EndTime)

	if err2 != nil {
		return err2
	}

	return validation.ValidateStruct(&r,
		validation.Field(&r.StartTime, &dateRule{message: "Start Time should be before end time", timeRange:TimeRange{start:sTime, end:eTime}}),
	)
}