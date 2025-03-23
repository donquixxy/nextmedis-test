package main

import (
	"errors"
	"log"
	"math/rand"
	"sync"
	"time"
)

type User struct {
	ID      int
	Name    string
	Balance float64
	m       sync.Mutex
}

func (u *User) Deposit(amount float64) {
	u.m.Lock()
	defer u.m.Unlock()
	prevBalance := u.Balance
	u.Balance += amount
	log.Printf("[Deposit] - User id %v | %v + %v =  %v", u.ID, prevBalance, amount, u.Balance)
}

func (u *User) Withdraw(amount float64) error {
	u.m.Lock()
	defer u.m.Unlock()
	prevBalance := u.Balance
	if u.Balance < amount {
		log.Printf("[Withdraw] - User id %v |  Withdraw %v | Current balance %v | Failed", u.ID, amount, u.Balance)
		return errors.New("not enough balance to deduct")
	}

	u.Balance -= amount
	log.Printf("[Withdraw] - User id %v | %v - %v = %v", u.ID, prevBalance, amount, u.Balance)
	return nil
}

func main() {
	u := &User{
		ID:      1,
		Name:    "Ari",
		Balance: 0,
		m:       sync.Mutex{},
	}

	var wg sync.WaitGroup
	rand.NewSource(time.Now().UnixNano())
	errCh := make(chan error, 1)
	stopLoop := false
	count := 0
	for i := 0; i < 10; i++ {

		if stopLoop {
			break
		}

		wg.Add(2) // Deposit and withdraw;

		go func() {
			depositAmount := rand.Float64() * 100
			defer wg.Done()
			u.Deposit(depositAmount)
		}()

		go func() {
			defer wg.Done()
			wdAmount := rand.Float64() * 120
			if err := u.Withdraw(wdAmount); err != nil {
				errCh <- err
				return
			}
		}()

		wg.Wait()

		select {
		case e := <-errCh:
			log.Printf("Failed to withdraw, stopping loop | %v", e)
			stopLoop = true
		default:
			count++
		}
	}

	close(errCh)
	log.Print(count)
	log.Printf("Final User balance: %v", u.Balance)
}
