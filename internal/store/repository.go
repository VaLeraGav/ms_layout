package store

import "gitlab.toledo24.ru/web/ms_layout/internal/entities"

type UserRepository interface {
	Create(*entities.User) error
	Find(string) (*entities.User, error)
	Remove(string) error
	Update(*entities.User) error
}

type StreetRepository interface {
	Create(*entities.Street) error
	Find(string) (*entities.Street, error)
	Remove(string) error
	Update(*entities.Street) error
}
