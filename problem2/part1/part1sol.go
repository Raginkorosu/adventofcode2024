package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	totalValidReports := 0
	inputFile, OpenFileErr := os.Open("../input")
	check(OpenFileErr)

	list := readLineByLine(inputFile)

	fmt.Printf("Total number of element in the list: %d \n", len(list))

	//Remove all reports with duplicates
	for index, element := range list {
		hasDuplicates := checkDuplicates(element)
		if hasDuplicates {
			list = slices.Delete(list, index, index+1)
		}
	}
	fmt.Printf("Total number of element in the list: %d \n", len(list))

	fmt.Println(totalValidReports)

}

func checkDuplicates(report []int) bool {
	hasDuplicates := false

	for _, element := range report {
		result := slices.Contains(report, element)
		if result {
			return true
		}

	}

	return hasDuplicates
}

func checkIfReportIsIncreasing(report []int) bool {
	isValid := false
	currentIndex := 0

	for index := range report {
		if index < len(report)-2 {
			if report[index+1] > report[index] {
				result := report[index+1] - report[index]
				if result <= 3 {
					currentIndex = index
				}
			} else {
				return false
			}
		}
	}

	if currentIndex == len(report)-2 {
		isValid = true
	}

	return isValid
}

func readLineByLine(inputFile *os.File) [][]int {
	var reportList [][]int

	scanner := bufio.NewScanner(inputFile)
	listIndex := 0

	for scanner.Scan() {
		var numberList []int
		numbers := strings.Split(scanner.Text(), " ")
		for _, element := range numbers {

			number, parseErr1 := strconv.Atoi(element)
			numberList = append(numberList, number)
			check(parseErr1)
		}
		reportList = append(reportList, numberList)
		listIndex++
	}

	return reportList
}

func check(e error) {

	if e != nil {
		fmt.Println("Couldn't read from file")
		panic(e)

	}
}
