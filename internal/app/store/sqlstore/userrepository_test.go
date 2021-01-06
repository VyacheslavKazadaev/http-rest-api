package sqlstore_test

import (
	"testing"

	"github.com/gopherschool/http-rest-api/internal/app/model"
	"github.com/gopherschool/http-rest-api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	u := model.TestUser(t)
	err := s.User().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u1 := model.TestUser(t)
	_, err := s.User().FindByEmail(u1.Email)
	assert.Error(t, err)

	s.User().Create(u1)
	u, err := s.User().FindByEmail(u1.Email)
	assert.NotNil(t, u)
	assert.NoError(t, err)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u1 := model.TestUser(t)
	_, err := s.User().Find(u1.ID)
	assert.Error(t, err)

	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	assert.NotNil(t, u2)
	assert.NoError(t, err)
}
