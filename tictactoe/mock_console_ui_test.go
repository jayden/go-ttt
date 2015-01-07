package tictactoe

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestMockConsoleIsAUI(t *testing.T) {
  mockConsole := new(MockConsoleUI)
  var ui UI = mockConsole

  t.Log("Mock Console UI is a kind of UI")
  assert.Equal(t, mockConsole, ui)
}
