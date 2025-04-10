package main 

import (
  "strings"
)

func cleanInput(text string) []string {
  pokemons := strings.Fields(strings.ToLower(text))
  return pokemons
}
