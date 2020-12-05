package teststore_test

import (
	"testing"

	"github.com/gopherschool/http-rest-api/internal/app/model"
	"github.com/gopherschool/http-rest-api/internal/app/store"
	"github.com/gopherschool/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()

	u := model.TestUser(t)
	err := s.User().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)

	_, err := s.User().FindByEmail(u.Email)
	assert.EqualError(t, err, store.ErrRecordNoFound.Error())

	s.User().Create(u)
	u, err = s.User().FindByEmail(u.Email)

	assert.NotNil(t, u)
	assert.NoError(t, err)
}
