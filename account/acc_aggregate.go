package account

// Account aggregate.
type Account struct {
	id      ID
	owner   Owner
	balance Amount
	updates []Event
	version int
}

// NewFromEvents is a helper method that creates a new account
// from a series of events.
func NewFromEvents(events []Event) *Account {
	a := &Account{}
	for _, event := range events {
		a.On(event, false)
	}
	return a
}

// New creates a new Account from id and owner.
func New(id ID, owner Owner) *Account {
	a := &Account{}
	a.raise(&AccountCreated{
		ID:    id,
		Owner: owner,
	})
	return a
}

// ID returns the account id.
func (a Account) ID() ID {
	return a.id
}

// Owner returns the account owner.
func (a Account) Owner() Owner {
	return a.owner
}

// Balance returns the account balance.
func (a Account) Balance() Amount {
	return a.balance
}

// UpdateOwner update owner.
func (a *Account) UpdateOwner(owner Owner) error {
	a.raise(&UpdateOwner{
		Owner: owner,
	})
	return nil
}

// Deposit amount into Account balance
func (a *Account) Deposit(amount Amount) error {
	if amount < 0 {
		return ErrAmountDepositErr
	}
	a.raise(&Deposit{
		ID:     a.id,
		Amount: amount,
	})
	return nil
}

// Withdraw amount from Account balance
func (a *Account) Withdrawal(amount Amount) error {
	if a.balance < amount {
		return ErrAmountNotAllowed
	}
	a.raise(&Withdrawal{
		ID:     a.id,
		Amount: amount,
	})
	return nil
}

// On handles account events.
func (a *Account) On(event Event, new bool) {
	switch e := event.(type) {
	case *AccountCreated:
		a.id = e.ID
		a.owner = e.Owner
		a.balance = 0
	case *UpdateOwner:
		a.owner = e.Owner
	case *Deposit:
		a.balance += e.Amount
	case *Withdrawal:
		a.balance -= e.Amount
	}
	if !new {
		a.version++
	}
}

// EventsList returns events from the Account aggregate.
func (a Account) EventsList() []Event {
	return a.updates
}

// Version returns the last version of the aggregate before changes.
func (a Account) Version() int {
	return a.version
}

func (a *Account) raise(event Event) {
	a.updates = append(a.updates, event)
	a.On(event, true)
}
