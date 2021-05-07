package storage_test

import (
	"testing"

	"github.com/bugbounce/totp-cli/storage"
	"github.com/stretchr/testify/assert"
)

func TestFindAccount(t *testing.T) {
	namespace := &storage.Namespace{
		Name: "mynamespace",
		Accounts: []*storage.Account{
			{Name: "Account1", Token: "token1"},
			{Name: "Account2", Token: "token2"},
			{Name: "Account3", Token: "token3"},
		},
	}

	account, err := namespace.FindAccount("Account1")

	assert.Equal(t, err, nil, "Error should be nil")
	assert.Equal(t, account.Name, "Account1", "Found account name should be Account1")
}

func TestFindAccount_NotFound(t *testing.T) {
	namespace := &storage.Namespace{
		Name: "mynamespace",
		Accounts: []*storage.Account{
			{Name: "Account1", Token: "token1"},
			{Name: "Account2", Token: "token2"},
			{Name: "Account3", Token: "token3"},
		},
	}

	account, err := namespace.FindAccount("AccountNotFound")

	assert.EqualError(t, err, "Account not found", "Error should be 'Account not found'")
	assert.Equal(t, account, &storage.Account{}, "Account should be nil")
}

func TestDeleteAccount(t *testing.T) {
	var account *storage.Account
	var err error

	namespace := &storage.Namespace{
		Name: "mynamespace",
		Accounts: []*storage.Account{
			{Name: "Account1", Token: "token1"},
			{Name: "Account2", Token: "token2"},
			{Name: "Account3", Token: "token3"},
		},
	}

	assert.Equal(t, len(namespace.Accounts), 3)
	account, err = namespace.FindAccount("Account1")
	assert.Equal(t, err, nil, "Error should be nil")

	namespace.DeleteAccount(account)
	assert.Equal(t, len(namespace.Accounts), 2)
	account, err = namespace.FindAccount("Account1")
	assert.EqualError(t, err, "Account not found", "Error should be 'Account not found'")
	// Delete again :D
	namespace.DeleteAccount(account)
	assert.Equal(t, len(namespace.Accounts), 2)
}
