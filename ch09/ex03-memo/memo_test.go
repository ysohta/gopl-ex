package memo

import (
	"fmt"
	"testing"
	"time"
)

func TestCancel(t *testing.T) {
	tests := []struct {
		timeExecution time.Duration
		timeToCancel  time.Duration
		val           interface{}
		err           error
	}{
		{100, 10, nil, fmt.Errorf("canceled")},
		{10, 100, "success", nil},
	}
	for _, test := range tests {

		f := func(_ string) (interface{}, error) {
			time.Sleep(test.timeExecution * time.Millisecond)
			return "success", nil
		}
		m := New(f)
		done := make(chan struct{})
		m.Cancel = done
		defer m.Close()

		go func() {
			time.Sleep(test.timeToCancel * time.Millisecond)
			close(done)
		}()

		val, err := m.Get("")
		if val != test.val {
			t.Errorf("expected:%s actual:%s", test.val, val)
		}

		if err == nil || test.err == nil {
			if err != test.err {
				t.Errorf("expected:%s actual:%s", test.err, err)
			}
		} else if err.Error() != test.err.Error() {
			t.Errorf("expected:%s actual:%s", test.err, err)
		}
	}

}

func Test(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	Concurrent(t, m)
}
