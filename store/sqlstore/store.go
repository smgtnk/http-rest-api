package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/smgtnk/http-rest-api/store"
	// _ "github.com/denisenkom/go-mssqldb"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
	// store.User().Create()
}
