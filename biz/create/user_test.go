package create_test

import (
	"app/biz/create"
	"testing"
)

func TestCreateUser(t *testing.T) {
	testcases := []struct {
		name     string
		username string
		password string
		err      error
	}{
		{
			name:     "Correct username and password length",
			username: "username",
			password: "password",
			err:      nil,
		},
		{
			name:     "Username too short ie less than 4 chars",
			username: "abc",
			password: "12345678",
			err:      create.ErrUsernameTooShort,
		},
		{
			name:     "Password too short ie less than 4 chars",
			username: "abcdefgh",
			password: "123",
			err:      create.ErrPasswordTooShort,
		},
	}

	for _, test := range testcases {
		_, err := create.NewUser(test.username, test.password)
		if err != test.err {
			t.Fatalf("Waiting %v but got %v", test.err, err)
		}
	}
}