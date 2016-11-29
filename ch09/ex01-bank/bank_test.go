package bank

import (
	"math/rand"
	"sync"
	"testing"
)

func TestBankAsync(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		amount := rand.Int()%1000 + 1

		go func() {
			Deposit(amount)
			ok := Withdraw(amount - 1)
			if !ok {
				t.Error("not ok")
			}
			wg.Done()
		}()
	}
	wg.Wait()

	want := 100
	got := Balance()
	if got != want {
		t.Errorf("got:%d want:%d", got, want)
	}

	var expected bool

	expected = true
	if ok := Withdraw(100); ok != expected {
		t.Errorf("expected:%t", expected)
	}

	expected = false
	if ok := Withdraw(1); ok != expected {
		t.Errorf("expected:%t", expected)
	}
}
