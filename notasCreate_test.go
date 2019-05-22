package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNoteCreateReturns200Status(t *testing.T) {
	db := ConnectToMongo()
	r := SetupRouter(db)

	w := httptest.NewRecorder()

	bodyReader := strings.NewReader(`{"title":"test"}`)

	req, _ := http.NewRequest("POST", "/note", bodyReader)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), `"title":"test"`)
}
