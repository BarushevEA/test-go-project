package main

import "testing"

func TestExample(t *testing.T) {
	want := "Hello, GitHub!"
	got := "Hello, GitHub!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
