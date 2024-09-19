package store

type Store interface {
	User() UserRepository
	Street() StreetRepository
}
