package main

import(
	"fmt"
	"github.com/StanislavKH/evtsource/account"
	"github.com/StanislavKH/evtsource/storage"
)

func main() {
	var err error

	s := storage.New()
	a, err := s.GetAccount(321, "Bob")
	onError(err)

	err = a.Deposit(100)
	onError(err)

	err = a.Deposit(320)
	onError(err)

	err = a.Withdrawal(20)
	onError(err)

	err = a.Withdrawal(500)
	onError(err)

	err = a.UpdateOwner("Wilma")
	onError(err)

	s.SetAccount(a)
	fmt.Printf("ID: %d, Owner: %s, Balance: %d\r\n", a.ID(), a.Owner(), a.Balance())

	_, err = s.GetAccount(321, "Stan")
	onError(err)

	// Create new account from events list
	b := account.NewFromEvents(a.EventsList())
	fmt.Printf("ID: %d, Owner: %s, Balance: %d\r\n", b.ID(), b.Owner(), b.Balance())
}

func onError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}