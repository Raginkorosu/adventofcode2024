package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {

  start := time.Now()

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

  var totalOccurances int = 0

  for i := range list1 {
    totalOccurances += (list1[i] * getOccurance(list1[i], list2))
  }
  
  fmt.Printf("Similarity Score: %d \n", totalOccurances);
  elasped := (time.Now().Sub(start))
  fmt.Printf("Total elasped time: %d \n", elasped/1000);

  inputFile.Close()

}

func getOccurance(number int, list []int) int {
  var occurances int = 0;
  
  for i := range list{
    if(number == list[i]){
      occurances++
    }
  }

  return occurances
}

func check(e error) {
	if e != nil {
		fmt.Println("Couldn't read from file")
		panic(e)
	}
}

