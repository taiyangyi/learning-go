package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, World!"
	got := hello()

	if got != want {
		t.Errorf("got %q want %q", got, want)

	}
}

// go test

// PASS
// ok      github.com/taiyangyi/learn-go-with-tests/hello-world/v2 0.076s

// --- FAIL: TestHello (0.00s)
//     hello_test.go:10: got "Hello, World!" want "Hello, World"
