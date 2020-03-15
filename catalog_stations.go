package applemusic

import (
	"context"
	"fmt"
)

// StationAttributes represents the attributes of the resource.
type StationAttributes struct {
	URL              string          `json:"url"`
	Name             string          `json:"name"`
	Artwork          Artwork         `json:"artwork"`
	PlayParams       PlayParameters  `json:"playParams"` // Undocumented
	EditorialNotes   *EditorialNotes `json:"editorialNotes,omitempty"`
	IsLive           bool            `json:"isLive"`
	DurationInMillis int64           `json:"durationInMillis,omitempty"`
	EpisodeNumber    string          `json:"episodeNumber,omitempty"`
}

// Station represents a station.
type Station struct {
	Id         string            `json:"id"`
	Type       string            `json:"type"`
	Href       string            `json:"href"`
	Attributes StationAttributes `json:"attributes"`
}

// Stations represents a list of stations.
type Stations struct {
	Data []Station `json:"data"`
	Href string    `json:"href,omitempty"`
	Next string    `json:"next,omitempty"`
}

func (s *CatalogService) getStations(ctx context.Context, u string) (*Stations, *Response, error) {
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	stations := &Stations{}
	resp, err := s.client.Do(ctx, req, stations)
	if err != nil {
		return nil, resp, err
	}

	return stations, resp, nil
}

// GetStation fetches a station using its identifier.
func (s *CatalogService) GetStation(ctx context.Context, storefront, id string, opt *Options) (*Stations, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/stations/%s", storefront, id)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.getStations(ctx, u)
}

// GetStationsByIds fetches one or more stations using their identifiers.
func (s *CatalogService) GetStationsByIds(ctx context.Context, storefront string, ids []string, opt *Options) (*Stations, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/stations", storefront)
	u, err := addOptions(u, makeIdsOptions(ids, opt))
	if err != nil {
		return nil, nil, err
	}

	return s.getStations(ctx, u)
}
