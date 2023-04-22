package teststore

import (
	"github.com/sletkov/go-http-server/internal/app/models"
	"github.com/sletkov/go-http-server/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
	users map[int]*models.User
}

// Create ...
func (r *UserRepository) Create(u *models.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	u.Id = len(r.users) + 1
	r.users[u.Id] = u

	return nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

// Find ...
func (r *UserRepository) Find(id int) (*models.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}
