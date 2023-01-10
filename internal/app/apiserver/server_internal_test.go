package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smgtnk/http-rest-api/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {

	s := NewServer(teststore.New())

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)

	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}
