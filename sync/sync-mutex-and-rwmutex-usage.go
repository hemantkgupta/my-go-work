package main

import (
	"fmt"
	"sync"
)

var account Account
var wg sync.WaitGroup

func main() {
	account.Name = "Test account"

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go account.Deposit(100)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go account.Withdraw(100)
	}

	wg.Wait()
	fmt.Printf("Balance: %d\n", account.GetBalance())
}

type Account struct {
	balance int
	Name    string
	lock    sync.RWMutex
}

func (a *Account) Withdraw(amount int) {
	defer wg.Done()
	a.lock.Lock()
	a.balance -= amount
	a.lock.Unlock()
}

func (a *Account) Deposit(amount int) {
	defer wg.Done()
	a.lock.Lock()
	a.balance += amount
	a.lock.Unlock()
}

func (a *Account) GetBalance() int {
	a.lock.RLock()
	defer a.lock.RUnlock()

	return a.balance
}
