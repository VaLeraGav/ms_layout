package postgres

import (
	"gitlab.toledo24.ru/web/ms_layout/internal/entities"
)

type StreetRepository struct {
	store *Store
}

func (r *StreetRepository) Create(com *entities.Street) error {
	return nil
}

func (r *StreetRepository) Find(email string) (*entities.Street, error) {
	com := &entities.Street{}
	return com, nil
}

func (r *StreetRepository) Remove(email string) error {
	return nil
}

func (r *StreetRepository) Update(com *entities.Street) error {
	return nil
}
