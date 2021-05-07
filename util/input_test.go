package util_test

import (
	"fmt"
	"testing"

	"github.com/bugbounce/totp-cli/util"
	"github.com/stretchr/testify/assert"
)

type PasswordPair struct {
	Password []byte
	Confirm  []byte
	Correct  bool
}

func TestCheckPasswordConfirm(t *testing.T) {
	var passwordPairs []PasswordPair = []PasswordPair{
		{[]byte("asdf"), []byte("asdf"), true},
		{[]byte("asdfg"), []byte("asdf"), false},
		{[]byte("asdfg"), []byte("asdfh"), false},
		{[]byte("asdf"), []byte("asdfh"), false},
		{[]byte("asdf"), nil, false},
		{nil, []byte("asdf"), false},
		{nil, nil, true},
	}

	for _, pair := range passwordPairs {
		assert.Equal(
			t,
			util.CheckPasswordConfirm(pair.Password, pair.Confirm),
			pair.Correct,
			fmt.Sprintf("%s == %s -> %t", pair.Password, pair.Confirm, pair.Correct),
		)
	}
}
