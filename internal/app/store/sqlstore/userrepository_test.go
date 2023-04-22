package sqlstore_test

import (
	"github.com/sletkov/go-http-server/internal/app/models"
	"github.com/sletkov/go-http-server/internal/app/store"
	"github.com/sletkov/go-http-server/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDb(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := models.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDb(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	email := "user@example.org"
	_, err := s.User().FindByEmail(email)

	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := models.TestUser(t)
	u.Email = email
	s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDb(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u1 := models.TestUser(t)
	s.User().Create(u1)
	u2, err := s.User().Find(u1.Id)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
