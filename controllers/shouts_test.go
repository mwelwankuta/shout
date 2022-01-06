package controllers

import (
	"net/http/httptest"
	"testing"

	"github.com/mwelwankuta/shout/database"
)

func TestHome(t *testing.T) {
	database.Connect()

	handler := Home

	req := httptest.NewRequest("GET", "http://localhost:5000/api/shouts", nil)
	writer := httptest.NewRecorder()

	handler(writer, req)

	res := writer.Result()
	if res.StatusCode != 200 {
		t.Fatalf("Expected Status Code 200, Returned %v", res.StatusCode)
	}
}
