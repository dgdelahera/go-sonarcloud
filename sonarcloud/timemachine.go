package sonarcloud

import (
	"encoding/json"
	"fmt"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/timemachine"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Timemachine service

func (s *Timemachine) Index(r timemachine.IndexRequest) (*timemachine.IndexResponse, error) {
	params := paramsFrom(r)

	req, err := s.client.NewRequestWithParameters("GET", fmt.Sprintf("%s/timemachine/index", API), params...)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		// TODO: parse error message
		return nil, fmt.Errorf("received non 2xx status code: %d", resp.StatusCode)
	}

	response := &timemachine.IndexResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}
