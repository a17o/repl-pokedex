package utils

import (
  "fmt"
  "strings"
)

func cleanInput(text string) []string {
  pokemons := strings.Fields(text)
  return pokemons
}
