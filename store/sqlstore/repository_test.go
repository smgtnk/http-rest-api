package sqlstore_test

import (
	"testing"

	"github.com/smgtnk/http-rest-api/internal/app/model"
	"github.com/smgtnk/http-rest-api/store"
	"github.com/smgtnk/http-rest-api/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {

	db, teardown := sqlstore.TestDB(t, databaseURL)
	s := sqlstore.New(db)

	defer teardown("users")

	u := model.TestUser(t)
	err := s.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {

	db, teardown := sqlstore.TestDB(t, databaseURL)
	s := sqlstore.New(db)

	defer teardown("users")

	email := "user@example.org"

	_, err := s.User().FindByEmail(email)

	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = "user@example.org"
	s.User().Create(u)

	u, err = s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {

	db, teardown := sqlstore.TestDB(t, databaseURL)
	s := sqlstore.New(db)

	defer teardown("users")

	u := model.TestUser(t)

	s.User().Create(u)

	u, err := s.User().Find(u.ID)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
