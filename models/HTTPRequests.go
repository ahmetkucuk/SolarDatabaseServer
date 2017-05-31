package models
import (
	"github.com/astaxie/beego/context"
	"strings"
	"net/url"
)



type SearchByIdRequest struct {
	Id	string
	EventType	string
}

type ImageRequest struct {
	ImageDate	string
	Size	string
	Wavelength	string
}

type GenerateVideoRequest struct {
	EventID	string
	Email	string
	Institute string
}

type PreviewVideoRequest struct {
	EventID	string
	EventType string
}

type SDOImageRequest struct {
	Resolution string
	Wavelength	string
	StartTime string
}

type ParameterImageRequest struct {
	Resolution string
	Wavelength	string
	StartTime string
	ParameterID string
}

type ImageParameterRequest struct {
	Wavelength	string
	StartTime string
}

type TrackIDRequest struct {
	EventID	string
	EventType string
}

func CreateGenerateVideoRequest(input url.Values)  (request GenerateVideoRequest, err error) {
	r := GenerateVideoRequest{}

	r.EventID = input.Get("eventid")
	r.Email = input.Get("email")
	r.Institute = input.Get("institude")

	return r, nil
}

func CreatePreviewVideoRequest(input url.Values)  (request PreviewVideoRequest, err error) {
	r := PreviewVideoRequest{}

	r.EventID = input.Get("eventid")
	r.EventType = input.Get("eventtype")

	return r, nil
}

func CreateImageRequest(input *context.BeegoInput)  (request ImageRequest, err error) {
	r := ImageRequest{}

	ImageDate := input.Param(":ImageDate")
	r.ImageDate = strings.Split(ImageDate, "=")[1]


	Size := input.Param(":Size")
	r.Size = strings.Split(Size, "=")[1]


	Wavelengths := input.Param(":Wavelength")
	r.Wavelength = strings.Split(Wavelengths, "=")[1]


	return r, nil
}

func CreateSearchByIdRequest(input url.Values)  (SearchByIdRequest, error) {
	r := SearchByIdRequest{}

	r.Id = input.Get("id")
	r.EventType = input.Get("eventtype")

	return r, nil
}

func CreateSDOImageRequest(input url.Values)  (SDOImageRequest, error) {
	r := SDOImageRequest{}

	r.Resolution = input.Get("resolution")
	r.Wavelength = input.Get("wavelength")
	r.StartTime = input.Get("starttime")

	return r, nil
}

func CreateParameterImageRequest(input url.Values)  (ParameterImageRequest, error) {
	r := ParameterImageRequest{}

	r.Resolution = input.Get("resolution")
	r.Wavelength = input.Get("wavelength")
	r.StartTime = input.Get("starttime")
	r.ParameterID = input.Get("paramid")

	return r, nil
}

func CreateImageParameterRequest(input url.Values)  (ImageParameterRequest, error) {
	r := ImageParameterRequest{}

	r.Wavelength = input.Get("wavelength")
	r.StartTime = input.Get("starttime")

	return r, nil
}

func CreateTrackIDRequest(input url.Values)  (TrackIDRequest, error) {
	r := TrackIDRequest{}

	r.EventID = input.Get("id")
	r.EventType = input.Get("eventtype")

	return r, nil
}


