package storage

import (
	"fmt"
	"errors"
	"github.com/StanislavKH/evtsource/account"
)

// Storage is fake account storage structure
type Storage struct {
	Accounts map[account.ID]*account.Account
}

// New storage initialization
func New() *Storage {
	return &Storage{Accounts: make(map[account.ID]*account.Account)}
}

// GetAccount create and return new if not exists, return error if already exists
func (s *Storage) GetAccount(id account.ID, owner account.Owner) (*account.Account, error) {
	if _, ok := s.Accounts[id]; !ok {
		a := account.New(id, owner)
		return a, nil
	}
	return nil, errors.New(fmt.Sprintf("Account with id: %d already exists!", id))
}

// SetAccount add / udate account into storage
func (s *Storage) SetAccount(a *account.Account) {
	s.Accounts[a.ID()] = a
}
