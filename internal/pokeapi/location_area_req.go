package pokeapi

import (
    "net/http"
    "fmt"
    "io"
    "encoding/json"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
    endpoint := "/location-area"
    fullURL := baseURL + endpoint
    if pageURL != nil {
        fullURL = *pageURL
    }

    // check the cache
    data, ok := c.cache.Get(fullURL)
    if ok {
        // cache hit   
        fmt.Println("cache hit!")
        locationAreasResp := LocationAreasResp{}
        err := json.Unmarshal(data, &locationAreasResp)
        if err != nil {
            return LocationAreasResp{}, err
        }

        return locationAreasResp, nil
    }
    fmt.Println("cache miss!")

    req, err := http.NewRequest("GET", fullURL, nil)
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

    data, err = io.ReadAll(resp.Body)
    if err != nil {
        return LocationAreasResp{}, err
    }

    locationAreasResp := LocationAreasResp{}
    err = json.Unmarshal(data, &locationAreasResp)
    if err != nil {
        return LocationAreasResp{}, err
    }
    
    c.cache.Add(fullURL, data)

    return locationAreasResp, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
    endpoint := "/location-area/" + locationAreaName
    fullURL := baseURL + endpoint

    // check the cache
    data, ok := c.cache.Get(fullURL)
    if ok {
        // cache hit   
        fmt.Println("cache hit!")
        locationArea := LocationArea{}
        err := json.Unmarshal(data, &locationArea)
        if err != nil {
            return LocationArea{}, err
        }

        return locationArea, nil
    }
    fmt.Println("cache miss!")

    req, err := http.NewRequest("GET", fullURL, nil)
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

    data, err = io.ReadAll(resp.Body)
    if err != nil {
        return LocationArea{}, err
    }

    locationArea := LocationArea{}
    err = json.Unmarshal(data, &locationArea)
    if err != nil {
        return LocationArea{}, err
    }
    
    c.cache.Add(fullURL, data)

    return locationArea, nil
}
