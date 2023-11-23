package config

import "testing"

func TestConn(t *testing.T) {
	NewDBConnection("host=localhost port=5432 user=root password=root sslmode=disable dbname=notes")
}
