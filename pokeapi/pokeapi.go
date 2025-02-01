package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type PokeClient struct {
	httpClient *http.Client
	baseURL    string
}

func NewPokeClient() *PokeClient {
	pokeClient := PokeClient{
		httpClient: &http.Client{},
		baseURL:    "https://pokeapi.co/api/v2/",
	}
	return &pokeClient
}

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationData struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

func (pc PokeClient) LocationAreas(url *string) (LocationData, error) {
	fullURL := pc.baseURL + "location-area"
	if url != nil {
		fullURL = *url
	}
	res, err := pc.httpClient.Get(fullURL)
	if err != nil {
		return LocationData{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationData{}, err
	}
	var locationAreas LocationData
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return LocationData{}, err
	}
	return locationAreas, nil
}
