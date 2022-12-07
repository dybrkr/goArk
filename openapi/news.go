package openapi

import (
	"github.com/dybrkr/goArk/openapi/model"
	"net/http"
)

const (
	GetEvent = "/news/events"
)

func (o *OpenAPI) GetEvents() ([]model.Event, error) {
	var resp []model.Event

	err := SendRequest(o, http.MethodGet, GetEvent, nil, &resp)
	return resp, err
}
