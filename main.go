package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)

  for {
    fmt.Print("Pokedex > ")
    if !scanner.Scan() {
      fmt.Errorf("Something went wrong during scanning.", scanner.Err())
    }
    input := scanner.Text()
    if input == "exit" {
      fmt.Println("Goodbye!")
      break
    }

    fmt.Println("Your command was: ", input)
  }
}

