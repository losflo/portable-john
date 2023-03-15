package heremaps

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Service struct {
	endpoint string
	apiKey   string
}

type Config struct {
	Endpoint string
	ApiKey   string
}

func NewService(cfg Config) *Service {
	return &Service{
		endpoint: cfg.Endpoint,
		apiKey:   cfg.ApiKey,
	}
} // ./NewService

func (s Service) Geocode(address string) (*Place, error) {
	endpoint := s.endpoint + "/geocode"
	params := url.Values{}
	params.Add("apiKey", s.apiKey)
	params.Add("q", address)
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(endpoint + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	type response struct {
		Items []Place `json:"items"`
	}
	var rs response
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &rs)
	if err != nil {
		return nil, err
	}
	if len(rs.Items) > 0 {
		return &rs.Items[0], nil
	}
	return nil, nil
} // ./Geocode
