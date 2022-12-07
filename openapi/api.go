package openapi

import (
	"encoding/json"
	"github.com/dybrkr/goArk/openapi/news"
	"github.com/dybrkr/goArk/request"
	"net/http"
)

type OpenAPI struct {
	*request.RestAPI
	AccessKey string
}

func CreateAPI(endPoint string, accessKey string) *OpenAPI {
	return &OpenAPI{
		RestAPI:   request.CreateAPI(endPoint),
		AccessKey: accessKey,
	}
}

func (o *OpenAPI) GetEvents() (news.EventResponse, error) {
	header := map[string]string{}
	header["accept"] = "application/json"
	header["authorization"] = "bearer " + o.AccessKey

	events := news.EventResponse{}

	resp, err := o.SendAuthRequest(http.MethodGet, "/news/events", header, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(resp), &events)
	if err != nil {
		return nil, err
	}

	return events, nil
}
