package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleHallo(t *testing.T) {
	s := New(Newconfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/Hello", nil)

	s.handleHallo().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}
