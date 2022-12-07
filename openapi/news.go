package openapi

import (
	"encoding/json"
	"github.com/dybrkr/goArk/openapi/model"
	"net/http"
)

func (o *OpenAPI) GetEvents() ([]model.Event, error) {
	header := map[string]string{}
	header["accept"] = "application/json"
	header["authorization"] = "bearer " + o.AccessKey

	events := []model.Event{}

	resp, err := o.SendRequest(http.MethodGet, "/news/events", header, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(resp), &events)
	if err != nil {
		return nil, err
	}

	return events, nil
}
