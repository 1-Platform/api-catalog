package teams

import "context"

type Teams struct {
	store Store
}

type Store interface {
}

func New(store Store) *Teams {
	return &Teams{store: store}
}

func ListTeams(ctx context.Context) {}
