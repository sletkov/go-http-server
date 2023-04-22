package store

import "github.com/sletkov/go-http-server/internal/app/models"

// UserRepository ...
type UserRepository interface {
	Create(*models.User) error
	Find(int) (*models.User, error)
	FindByEmail(string) (*models.User, error)
}
