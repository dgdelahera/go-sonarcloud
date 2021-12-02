package sonarcloud

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/form/v4"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/favorites"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/paging"
	"strings"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Favorites service

func (s *Favorites) Add(r favorites.AddRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/favorites/add", API), strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		// TODO: parse error message
		return fmt.Errorf("received non 2xx status code: %d", resp.StatusCode)
	}

	return nil
}

func (s *Favorites) Remove(r favorites.RemoveRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/favorites/remove", API), strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		// TODO: parse error message
		return fmt.Errorf("received non 2xx status code: %d", resp.StatusCode)
	}

	return nil
}

func (s *Favorites) Search(r favorites.SearchRequest, p paging.PagingParams) (*favorites.SearchResponse, error) {
	params := paramsFrom(r, p)

	req, err := s.client.NewRequestWithParameters("GET", fmt.Sprintf("%s/favorites/search", API), params...)
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

	response := &favorites.SearchResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Favorites) SearchAll(r favorites.SearchRequest) (*favorites.SearchResponseAll, error) {
	p := paging.PagingParams{
		P:  1,
		Ps: 100,
	}
	response := &favorites.SearchResponseAll{}
	for {
		res, err := s.Search(r, p)
		if err != nil {
			return nil, fmt.Errorf("could not search all projects: %+v", err)
		}
		response.Favorites = append(response.Favorites, res.Favorites...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}
