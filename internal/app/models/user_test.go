package models_test

import (
	"github.com/sletkov/go-http-server/internal/app/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := models.TestUser(t)

	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *models.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *models.User {
				return models.TestUser(t)
			},
			isValid: true,
		},

		{
			name: "empty email",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},

		{
			name: "invalid email",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Email = "invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "with encrypted password",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encrypted password"
				return u
			},
			isValid: true,
		},
		{
			name: "empty password",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},

		{
			name: "short password",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Password = "1234"
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
	u := models.TestUser(t)
	assert.NoError(t, u.Validate())
}
