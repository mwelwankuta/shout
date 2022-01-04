package database

import "testing"

func TestConnect(t *testing.T) {
	if err := Connect(); err != nil {
		t.Fatalf("Could not connect to database")
	}
}
