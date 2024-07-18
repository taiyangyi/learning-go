package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, Yang"
	got := Hello("Yang")

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

// PASS
// ok      github.com/taiyangyi/learn-go-with-tests/hello-world/v4 0.053s
