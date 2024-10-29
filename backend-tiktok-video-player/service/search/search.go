package search

import (
	"encoding/json"
	"time"

	"github.com/RetakerMZ/backend-tiktok-video-player/config"
	"github.com/RetakerMZ/backend-tiktok-video-player/domain"
	"github.com/go-resty/resty/v2"
)

type searchService struct {
	contextTimeout time.Duration
	resty          *resty.Client
}

func NewSearchService(t time.Duration, resty *resty.Client) domain.SearchService {
	return &searchService{
		contextTimeout: t,
		resty:          resty,
	}
}

func (s *searchService) GetSearch(query string, cursor string) (domain.SearchResult, error) {
	url := "https://" + config.GetRapidApiUrl() + "/feed/search?keywords=" + query + "&count=30" + "&cursor=" + cursor + "&publish_time=7&sort_type=1"
	response, err := s.resty.
		R().
		SetHeader("Content-Type", "application/json").
		SetHeader("x-rapidapi-key", config.GetRapidApiKey()).
		SetHeader("x-rapidapi-host", config.GetRapidApiUrl()).
		Get(url)
	if err != nil {
		return domain.SearchResult{}, err
	}
	var res domain.SearchResult
	err = json.Unmarshal(response.Body(), &res)
	if err != nil {
		return domain.SearchResult{}, err
	}
	return res, nil
}
