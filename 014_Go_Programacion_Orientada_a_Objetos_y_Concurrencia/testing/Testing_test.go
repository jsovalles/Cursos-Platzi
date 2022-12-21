package testing

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum1(t *testing.T) {
	fmt.Println()
	total := Sum(5, 5)

	if total != 10 {
		t.Errorf("Wrong output, expected 10, go %d", total)
	}
}

func Test_Sum(t *testing.T) {
	tables := []struct {
		description string
		a           int
		b           int
		total       int
	}{
		{
			description: "adding 2 digit numbers",
			a:           12,
			b:           20,
			total:       32,
		},
		{
			description: "adding 1 digit numbers",
			a:           2,
			b:           2,
			total:       4,
		},
	}

	for _, item := range tables {
		t.Run(item.description, func(t *testing.T) {
			if totalSum := Sum(item.a, item.b); totalSum != item.total {
				t.Errorf("error: incorrect sum. got %d, expected %d", totalSum, item.total)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {

	t.Run("Test", func(t *testing.T) {

	})

	tables := []struct {
		description string
		a           int
		n           int
	}{
		{
			description: "Fibonacci with 1 number",
			a:           1,
			n:           1,
		},
		{
			description: "Fibonacci with 8 numbers",
			a:           8,
			n:           21,
		},
		{
			description: "Fibonacci with 50 numbers",
			a:           50,
			n:           12586269025,
		},
	}
	for _, item := range tables {
		t.Run(item.description, func(t *testing.T) {
			fib := Fibonacci(item.a)
			assert.Equal(t, fib, item.n)
		})
	}
}
