package main

import "testing"

func TestMax(t *testing.T) {
	want := 5
	val, err := max(1, 5, -3)
	if err != nil {
		t.Error("not expected error: %v", err)
	}
	if val != want {
		t.Errorf("expected:%d actual:%d", want, val)
	}
}

func TestMaxOneArg(t *testing.T) {
	want := 1
	val, err := max(1)
	if err != nil {
		t.Error("not expected error: %v", err)
	}
	if val != want {
		t.Errorf("expected:%d actual:%d", want, val)
	}
}

func TestMaxNoArg(t *testing.T) {
	_, err := max()
	if err == nil {
		t.Error("should return error")
	}
}

func TestMin(t *testing.T) {
	want := -3
	val, err := min(1, 5, -3)
	if err != nil {
		t.Error("not expected error: %v", err)
	}
	if val != want {
		t.Errorf("expected:%d actual:%d", want, val)
	}
}

func TestMinOneArg(t *testing.T) {
	want := 1
	val, err := min(1)
	if err != nil {
		t.Error("not expected error: %v", err)
	}
	if val != want {
		t.Errorf("expected:%d actual:%d", want, val)
	}
}

func TestMinNoArg(t *testing.T) {
	_, err := min()
	if err == nil {
		t.Error("should return error")
	}
}
func TestMax2(t *testing.T) {
	want := 5
	val := max2(1, 5, -3)
	if val != want {
		t.Errorf("expected:%d actual:%d", want, val)
	}
}

func TestMax2OneArg(t *testing.T) {
	want := 1
	val := max2(1)
	if val != want {
		t.Errorf("expected:%d actual:%d", want, val)
	}
}

func TestMin3(t *testing.T) {
	want := -3
	val := min2(1, 5, -3)
	if val != want {
		t.Errorf("expected:%d actual:%d", want, val)
	}
}

func TestMin3OneArg(t *testing.T) {
	want := 1
	val := min2(1)
	if val != want {
		t.Errorf("expected:%d actual:%d", want, val)
	}
}
