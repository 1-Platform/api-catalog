package api

import "time"

var (
	now = time.Now()
)

type API struct {
	Version string
}

func (a *API) Health() (map[string]any, error) {
	return map[string]any{
		"version":    a.Version,
		"startedAt":  now.String(),
		"releasedOn": now.String(),
	}, nil
}

func New(ver string) (*API, error) {
	return &API{Version: ver}, nil
}
