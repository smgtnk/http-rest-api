package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/smgtnk/http-rest-api/store/sqlstore"
)

func Start(config *Config) error {

	db, err := newDB(config.DatabaseURL)

	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	serv := NewServer(store, sessionStore)

	return http.ListenAndServe(config.BindAddr, serv)

}

func newDB(databaseURL string) (*sql.DB, error) {

	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// func (s *APIServer) configureLogger() error {
// 	level, err := logrus.ParseLevel(s.config.LogLevel)

// 	if err != nil {
// 		return err
// 	}

// 	s.logger.SetLevel(level)

// 	return nil
// }

// func (s *APIServer) configureRouter() {

// 	s.router.HandleFunc("/hello", s.handleHallo())

// }

// func (s *APIServer) configureStore() error {

// 	st := store.New(s.config.Store)
// 	if err := st.Open(); err != nil {
// 		return err
// 	}

// 	s.store = st

// 	return nil

// }

// func (s *APIServer) handleHallo() http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, "Hello")
// 	}

// }
