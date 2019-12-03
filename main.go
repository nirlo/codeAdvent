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
  scanner.Split(bufio.ScanLines)
  var modules []int

  for scanner.Scan() {
    num, err := strconv.Atoi(scanner.Text())
    if err != nil {
      panic(err) // at the disco
    }
    modules = append(modules, num)
  }

  file.Close()

  fmt.Printf("%d\n", calculateFuel(modules))
}

// 
func calculateFuel(modules []int) int {
  var total int
  for _, number := range modules {
    fuel := calcFuel(number)
    for (fuel >= 0) {
      total += fuel
      fuel = calcFuel(fuel)
    }
  }
  return total
}

func calcFuel(f int) int {
  return (f/3)-2
}

