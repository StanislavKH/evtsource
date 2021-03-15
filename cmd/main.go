package main

import(
	"fmt"
	"github.com/StanislavKH/evtsource/account"
)

func main() {
	var err error

	a := account.New(321, "Bob")

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

	fmt.Printf("ID: %d, Owner: %s, Balance: %d\r\n", a.ID(), a.Owner(), a.Balance())

	// Create new account from events list
	b := account.NewFromEvents(a.EventsList())
	fmt.Printf("ID: %d, Owner: %s, Balance: %d\r\n", b.ID(), b.Owner(), b.Balance())
}

func onError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}