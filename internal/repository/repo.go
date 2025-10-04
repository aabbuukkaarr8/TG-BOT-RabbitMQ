package repository

import "github.com/aabbuukkaarr8/TG-BOT/internal/store"

type Repository struct {
	store *store.Store
}

func NewRepository(store *store.Store) *Repository {
	return &Repository{
		store: store,
	}
}
