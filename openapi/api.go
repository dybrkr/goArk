package openapi

import (
	"github.com/dybrkr/goArk/request"
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
