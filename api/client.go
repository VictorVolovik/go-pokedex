package api

import (
	"net/http"
	"time"

	"VictorVolovik/go-pokedex/cache"
)

type Client struct {
	cache      *cache.Cache
	httpClient http.Client
}

// Create API client with cache
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: cache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
