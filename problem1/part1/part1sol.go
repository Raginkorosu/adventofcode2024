package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
  "strconv"
  "sort"
)

func main() {

  var list1 []int
  var list2 []int

	inputFile, OpenFileErr := os.Open("../input")
	check(OpenFileErr)

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), "   ")
    number1, parseErr1 := strconv.Atoi(numbers[0])
    number2, parseErr2 := strconv.Atoi(numbers[1])
    check(parseErr1)
    check(parseErr2)
    list1 = append(list1, int(number1))
    list2 = append(list2, int(number2))
	}

  fmt.Println(len(list1))
  fmt.Println(len(list2))
  
  sort.Slice(list1, func(i, j int) bool {
    return list1[i] < list1[j]
  })
  
  sort.Slice(list2, func(i, j int) bool {
    return list2[i] < list2[j]
  })

  var totalDistance uint = 0

  for i := range list1 {
    totalDistance += getDistance(list1[i], list2[i])
    fmt.Printf("Number1: %d, Number2: %d, total Distance:  %d \n", list1[i], list2[i], totalDistance)
  }
  
  fmt.Printf("Total distance: %d \n", totalDistance)

  inputFile.Close()
}

func getDistance(number1 int, number2 int) uint{
    var distance uint = 0
    if(number1 < number2) {
      distance = uint((number2 - number1))
    }else{
      distance = uint((number1 - number2))
    }
    return distance
}

func check(e error) {
	if e != nil {
		fmt.Println("Couldn't read from file")
		panic(e)
	}
}

