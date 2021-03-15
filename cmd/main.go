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

	fmt.Println(a.ID(), a.Owner(), a.Balance())

	b := account.NewFromEvents(a.EventsList())
	fmt.Println(b.ID(), b.Owner(), b.Balance())
}

func onError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}