package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreas, error) {
	endPoint := "/location-area"
	fullURL := baseURL + endPoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	body, ok := c.cache.Get(fullURL)
	if ok {
		locationAreas := LocationAreas{}
		err := json.Unmarshal(body, &locationAreas)
		if err != nil {
			return LocationAreas{}, err
		}

		return locationAreas, nil
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreas{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	locationAreas := LocationAreas{}
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return LocationAreas{}, err
	}

	c.cache.Add(fullURL, body)

	return locationAreas, nil
}

func (c *Client) GetLocationAreas(localtionAreaName string) (LocationArea, error) {
	endPoint := "/location-area/" + localtionAreaName
	fullURL := baseURL + endPoint

	body, ok := c.cache.Get(fullURL)
	if ok {
		locationArea := LocationArea{}
		err := json.Unmarshal(body, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, body)

	return locationArea, nil
}
