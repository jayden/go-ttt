package tictactoe

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestMakeEmptyBoard(t *testing.T) {
  t.Log("Returns a blank board.")
  assert.Equal(t, MakeBoard(), make([]byte, 9))
}
