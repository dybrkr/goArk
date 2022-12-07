package request

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type RestAPI struct {
	httpClient *http.Client
	EndPoint   string
}

func CreateAPI(endPoint string) *RestAPI {
	return &RestAPI{
		httpClient: &http.Client{Timeout: time.Second},
		EndPoint:   endPoint,
	}
}

func (r *RestAPI) SendAuthRequest(method string, subUrl string, headers map[string]string, data interface{}) (string, error) {
	var body io.Reader

	body = nil
	if data != nil {
		encoded, err := json.Marshal(data)
		if err != nil {
			return err.Error(), err
		}
		body = bytes.NewReader(encoded)
	}

	url := r.EndPoint + subUrl
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return err.Error(), err
	}

	for k, v := range headers {
		request.Header.Add(k, v)
	}

	response, err := r.httpClient.Do(request)
	if err != nil {
		return err.Error(), err
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err.Error(), err
	}
	return string(content), err
}

func EncodeURLValues(urlPath string, values url.Values) string {
	u := urlPath
	if len(values) > 0 {
		u += "?" + values.Encode()
	}
	return u
}
