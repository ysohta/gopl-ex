package bank

import (
	"math/rand"
	"sync"
	"testing"
)

func TestBankAsync(t *testing.T) {
	var wg sync.WaitGroup
	var got, want int
	var expected bool

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

	want = 100
	got = Balance()
	if got != want {
		t.Errorf("got:%d want:%d", got, want)
	}

	expected = true
	if ok := Withdraw(100); ok != expected {
		t.Errorf("expected:%t", expected)
	}
	want = 0
	got = Balance()
	if got != want {
		t.Errorf("got:%d want:%d", got, want)
	}

	expected = false
	if ok := Withdraw(1); ok != expected {
		t.Errorf("expected:%t", expected)
	}
	want = 0
	got = Balance()
	if got != want {
		t.Errorf("got:%d want:%d", got, want)
	}
}
