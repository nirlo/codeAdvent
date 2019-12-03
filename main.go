package main

import (
  "os"
  "fmt"
  "bufio"
  "strconv"
)

func main() {
  file, err := os.Open("fuel.txt")
  if err != nil {
    // cause undo panic cause this is just a small program
    panic(err)
  }

  scanner := bufio.NewScanner(file)
  scanner.split(bufio.ScanLines)
  var numbers []int

  for scanner.Scan() {
    num, err := strconv.Atoi(scanner.Text())
    if err != nil {
      panic(err) // at the disco
    }
    numbers = append(numbers, num)
  }

  file.Close()

  fmt.Printf("%d\n", calculateFuel(numbers))
}

func calculateFuel([]int numbers) int {
  var fuel
  for _, number := range numbers {
    fuel = fuel + (number / 3 - 2)
  }
  return fuel
}
