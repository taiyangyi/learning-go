package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Bob")
		want := "Hello,Bob"
		assertMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("")
		want := "Hello,Worl"
		assertMessage(t, got, want)

	})

}

func assertMessage(t *testing.T, got, want string) {

	// 告诉测试套件这个方法是辅助函数。
	// 通过这样做，当测试失败时所报告的行号将在 函数调用中 而不是在辅助函数内部。
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
