package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endPoint := "/location-area"
	fullURL := baseURL + endPoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	body, ok := c.cache.Get(fullURL)
	if ok {
		locationAreas := LocationAreasResp{}
		err := json.Unmarshal(body, &locationAreas)
		if err != nil {
			return LocationAreasResp{}, err
		}

		return locationAreas, nil
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreas := LocationAreasResp{}
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullURL, body)

	return locationAreas, nil
}
