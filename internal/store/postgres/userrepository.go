package postgres

import (
	"database/sql"

	"gitlab.toledo24.ru/web/ms_layout/internal/entities"
	"gitlab.toledo24.ru/web/ms_layout/internal/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(com *entities.User) error {
	return r.store.db.QueryRow(
		"INSERT INTO user (email, data) VALUES ($1, $2) RETURNING id",
		&com.Email,
		&com.Data,
	).Scan(&com.ID)
}

func (r *UserRepository) Find(email string) (*entities.User, error) {
	com := &entities.User{}
	if err := r.store.db.QueryRow(
		"SELECT * FROM user WHERE email = $1",
		email,
	).Scan(
		&com.ID,
		&com.Email,
		&com.Data,
		&com.CreatedAt,
		&com.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return com, nil

}

func (r *UserRepository) Remove(email string) error {
	result, err := r.store.db.Exec(
		"DELETE FROM user WHERE email = $1",
		email,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return store.ErrRecordNotFound
	}

	return nil
}

func (r *UserRepository) Update(com *entities.User) error {
	result, err := r.store.db.Exec(
		"UPDATE user SET data = $1, updated_at = NOW() WHERE email = $2",
		&com.Data,
		&com.Email,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return store.ErrRecordNotFound
	}

	return nil
}
