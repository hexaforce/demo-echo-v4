package schemas

import (
	"errors"
	"fmt"

	uuid "github.com/google/uuid"
)

// Account example
type Account struct {
	ID   int       `json:"id" example:"1" format:"int64"`
	Name string    `json:"name" example:"account name"`
	UUID uuid.UUID `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
}

var (
	ErrNameInvalid = errors.New("name is empty")
)

// AddAccount
type AddAccount struct {
	Name string `json:"name" example:"account name"`
}

// Validation
func (a AddAccount) Validation() error {
	switch {
	case len(a.Name) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

// UpdateAccount
type UpdateAccount struct {
	Name string `json:"name" example:"account name"`
}

// Validation
func (a UpdateAccount) Validation() error {
	switch {
	case len(a.Name) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

// AccountsAll
func AccountsAll(q string) ([]Account, error) {
	if q == "" {
		return accounts, nil
	}
	as := []Account{}
	for k, v := range accounts {
		if q == v.Name {
			as = append(as, accounts[k])
		}
	}
	return as, nil
}

// AccountOne
func AccountOne(id int) (Account, error) {
	for _, v := range accounts {
		if id == v.ID {
			return v, nil
		}
	}
	return Account{}, ErrNoRow
}

// Insert
func (a Account) Insert() (int, error) {
	accountMaxID++
	a.ID = accountMaxID
	a.Name = fmt.Sprintf("account_%d", accountMaxID)
	accounts = append(accounts, a)
	return accountMaxID, nil
}

// Delete
func Delete(id int) error {
	for k, v := range accounts {
		if id == v.ID {
			accounts = append(accounts[:k], accounts[k+1:]...)
			return nil
		}
	}
	return fmt.Errorf("account id=%d is not found", id)
}

// Update
func (a Account) Update() error {
	for k, v := range accounts {
		if a.ID == v.ID {
			accounts[k].Name = a.Name
			return nil
		}
	}
	return fmt.Errorf("account id=%d is not found", a.ID)
}

var accountMaxID = 3
var accounts = []Account{
	{ID: 1, Name: "account_1"},
	{ID: 2, Name: "account_2"},
	{ID: 3, Name: "account_3"},
}
