package sync

import (
	"fmt"
	"sync"
)

var balance int = 100

func SyncExample() {
	var wg sync.WaitGroup
	var mutex sync.RWMutex
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &mutex)
		wg.Add(1)
		go Balance(&wg, &mutex)
	}

	wg.Add(1)
	go WithDraw(500, &wg, &mutex)

	wg.Wait()
	fmt.Println("Finished with balance", balance)
}

func Deposit(amount int, wg *sync.WaitGroup, mutex *sync.RWMutex) {
	mutex.Lock()
	b := balance
	balance = b + amount
	defer mutex.Unlock()
	defer wg.Done()
}

func WithDraw(amount int, wg *sync.WaitGroup, mutex *sync.RWMutex) {
	mutex.Lock()
	if balance < amount {
		fmt.Println("Not enough balance")
		mutex.Unlock()
		wg.Done()
		return
	}
	balance -= amount
	mutex.Unlock()
	defer wg.Done()
}

func Balance(wg *sync.WaitGroup, mutex *sync.RWMutex) {
	defer wg.Done()
	defer mutex.RUnlock()
	mutex.RLock()
	b := balance
	fmt.Println("Current balance is", b)
}
