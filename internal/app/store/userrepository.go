package store

import "github.com/sletkov/go-http-server/internal/app/models"

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *models.User) (*models.User, error) {

	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	if err := r.store.db.QueryRow("INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.Id); err != nil {
		return nil, err
	}
	return u, nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	u := &models.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.Id,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		return nil, err
	}
	return u, nil
}
