package postgres

import (
	"database/sql"

	"gitlab.toledo24.ru/web/ms_layout/internal/store"

	_ "github.com/lib/pq"
)

type Store struct {
	db               *sql.DB
	userRepository   *UserRepository
	streetRepository *StreetRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
func (s *Store) Street() store.StreetRepository {
	if s.streetRepository != nil {
		return s.streetRepository
	}

	s.streetRepository = &StreetRepository{
		store: s,
	}

	return s.streetRepository
}
