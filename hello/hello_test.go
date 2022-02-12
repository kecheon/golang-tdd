package main

import "testing"


func TestHello(t *testing.T) {
  assertCorrectMsg := func(t testing.TB, got string, want string) {
    t.Helper()
    if got != want {
      t.Errorf("got %q want %q", got, want)
    }
  }

  t.Run("say hello to somebody", func(t *testing.T) {
    got := Hello("Chris", "English")
    want := "Hello Chris"

    assertCorrectMsg(t, got, want)
  })

  t.Run("say hello world when empty name", func(t *testing.T) {
    got := Hello("", "English")
    want := "Hello world"

    assertCorrectMsg(t, got, want)
  })

  t.Run("in Korean", func(t *testing.T) {
    got := Hello("제육", "Korean")
    want := "안녕 제육"

    assertCorrectMsg(t, got, want)
  })

  t.Run("in French", func(t *testing.T) {
    got := Hello("제육", "French")
    want := "Bonjour 제육"

    assertCorrectMsg(t, got, want)
  })
}
