package memo

import "testing"

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
