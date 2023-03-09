package api

import (
	"github.com/1-platform/api-catalog/internal/teams"
	"time"
)

var (
	now = time.Now()
)

type API struct {
	Version string
	teams   *teams.Teams
}

func (a *API) Health() (map[string]any, error) {
	return map[string]any{
		"version":    a.Version,
		"startedAt":  now.String(),
		"releasedOn": now.String(),
	}, nil
}

func New(ver string, tm *teams.Teams) (*API, error) {
	return &API{
		Version: ver,
		teams:   tm,
	}, nil
}
