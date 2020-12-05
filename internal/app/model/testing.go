package model

import "testing"

// TestUser ...
func TestUser(t *testing.T) *User {
	t.Helper()
	return &User{
		Email:    "users@example.org",
		Password: "password",
	}
}
