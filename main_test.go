package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHomepageHandler tests the home page endpoint.
func TestHomeHandler(t *testing.T) {
	mockResponse := `{"message":"@cloudvesna Task Management API with Golang and Gin"}`
	r := SetUpRouter()
	r.GET("/", HomeHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
