package main

import (
	"fmt"
	"sync"
	"time"
)

var println = fmt.Println
var printf = fmt.Printf

type Account struct {
	balance int
	lock    sync.Mutex
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) WithDraw(amount int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if amount > a.balance {
		println("Oops not enough balance")
	} else {
		printf("new balance %d \n", (a.balance - amount))
		a.balance -= amount
	}
}

func main() {
	var account Account
	account.balance = 100
	printf("balance %d \n", account.balance)
	for i := 0; i < 12; i++ {
		go account.WithDraw(10)
	}
	time.Sleep(2 * time.Second)
}
