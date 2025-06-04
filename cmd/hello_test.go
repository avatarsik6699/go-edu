package main

import "testing"

func TestDomainHelloWorld(t *testing.T) {
	got := DomainHelloWorld("Jhon")
	want := "Hello, Jhon"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
