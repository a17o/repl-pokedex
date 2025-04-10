package utils
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
    for i := range actual {
      word := actual[i]
      expectedWord := c.expected[i]
    }
  }
}

