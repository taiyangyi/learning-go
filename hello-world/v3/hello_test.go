package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, You!"
	got := hello("You!")

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

// PASS
// ok      github.com/taiyangyi/learn-go-with-tests/hello-world/v3 0.036s
