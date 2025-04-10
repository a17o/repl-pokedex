package main
import (
  "testing"
)

func TestCleanInput(t *testing.T) {
  cases := []struct {
    input    string
    expected []string
  }{
    {
      input:    " hello world ",
      expected: []string{"hello", "world"},
    },
    {
      input:    "pickaCHU Bulbasaur CHARMANDER ",
      expected: []string{"pickachu", "bulbasaur", "charmander"},
    },
    {
      input:    "oryx",
      expected: []string{"oryx"},
    },
    {
      input:    "",
      expected: []string{},
    },
  }

  for _, c := range cases {
    actual := cleanInput(c.input)
    if len(actual) != len(c.expected) {
      t.Errorf("Test failed: length of the expected string (%s) differs from the actual (%s)", actual, c.expected)
    }
    for i := range actual {
      word := actual[i]
      expectedWord := c.expected[i]

      if word != expectedWord {
        t.Errorf("Test failed: expected %s, got %s", expectedWord, word)
      }
    }
  }
}

