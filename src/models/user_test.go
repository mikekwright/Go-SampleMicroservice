package models

import (
	"testing"
)

func TestUserCreation_IsSuccessful(t *testing.T) {
	expectedName := "My Name"

	testModel := User {UserId: 1, Name: expectedName, Email: "test@test.com" }

	if testModel.Name != expectedName {
		t.Fatalf("Name should match, but go unexpected name %s", testModel.Name)
	}
}