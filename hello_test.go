package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Garrison")
	want := "Hello, Garrison"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
