package cache

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var maxChannels = 2

func CacheExample() {
	startRun := time.Now()
	cache := NewCache(GetFibonacci)
	fibo := []int{42, 40, 41, 42, 38}
	var wg sync.WaitGroup
	channels := make(chan int, maxChannels)
	for _, n := range fibo {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			channels <- 1
			start := time.Now()
			value, err := cache.Get(index)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("Calculate: %d, Time: %s, Result: %d\n", index, time.Since(start), value)
			<-channels
		}(n)

	}
	wg.Wait()

	fmt.Printf("Process completed in %s\n", time.Since(startRun))
}

// Function to calculate fibonacci
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Memory holds a function and a map of results
type Memory struct {
	f     Function               // Function to be used
	cache map[int]FunctionResult // Map of results for a given key
	lock  sync.RWMutex             // Lock to protect the cache
}

// A function has to receive a value and return a value and an error
type Function func(key int) (interface{}, error)

// The result of a function
type FunctionResult struct {
	value interface{}
	err   error
}

// NewCache creates a new cache
func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

// Get returns the value for a given key
func (m *Memory) Get(key int) (interface{}, error) {
	m.lock.Lock()
	// Check if the value is in the cache
	result, exists := m.cache[key]
	m.lock.
	// If the value is not in the cache, calculate it
	if !exists {
		m.lock.Lock()
		result.value, result.err = m.f(key) // Calculate the value
		m.cache[key] = result               // Store the value in the cache
		m.lock.Unlock()
	}
	return result.value, result.err
}

// Function to be used in the cache
func GetFibonacci(key int) (interface{}, error) {
	return Fibonacci(key), nil
}
