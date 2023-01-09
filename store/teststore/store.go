package teststore

import (
	"github.com/smgtnk/http-rest-api/internal/app/model"
	"github.com/smgtnk/http-rest-api/store"
	// _ "github.com/denisenkom/go-mssqldb"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
