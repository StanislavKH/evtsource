package account

import(
	"errors"
)

var(
	ErrAmountNotAllowed = errors.New("Current amount not allowed to withdrawal due to account balance.")
	ErrAmountDepositErr = errors.New("Deposit amount should be greater then zero.")
)

type ID int
type Owner string
type Amount int

// Event event marker.
type Event interface {
	isEvent()
}

func (e AccountCreated) isEvent() {}
func (e UpdateOwner) isEvent()    {}
func (e Deposit) isEvent()        {}
func (e Withdrawal) isEvent()     {}

// AccountCreated -
type AccountCreated struct {
	ID    ID
	Owner Owner
}

// UpdateOwner -
type UpdateOwner struct {
	ID    ID
	Owner Owner
}

// Deposit -
type Deposit struct {
	ID     ID
	Amount Amount
}

// Withdrawal -
type Withdrawal struct {
	ID     ID
	Amount Amount
}
