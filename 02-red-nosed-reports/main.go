package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// small helper function, it didn't really make sense to import the math package just for this
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Part 2: Helper function to check reports with one level removed
func verifyDampenedReport(report []int) bool {
	var isDecreasing bool
	for i := 0; i < len(report) - 1; i++ {
		diff := report[i] - report[i+1]
		if abs(diff) < 1 || abs(diff) > 3 {
			return false
		}
		if i == 0 {
			isDecreasing = diff > 0
		} else if (isDecreasing && diff <= 0) || (!isDecreasing && diff >= 0) {
			return false
		}
	}
	return true
}

// Part 2: Modified function to accomodate for the "Problem dampener"
//         (one level per report can be removed)
//         This could be combined with the Part 1 solution, but I wanted to keep the original function intact
func verifyReport2(report []int) bool {
	var isDecreasing bool

	for i := 0; i < len(report) - 1; i++ {
		diff := report[i] - report[i+1]
		if abs(diff) < 1 || abs(diff) > 3 {
			slice1 := make([]int, len(report[:i]))
			copy(slice1, report[:i])
			slice1 = append(slice1, report[(i+1):]...)

			slice2 := make([]int, len(report[:i+1]))
			copy(slice2, report[:(i+1)])
			slice2 = append(slice2, report[(i+2):]...)
			return verifyDampenedReport(slice1) || verifyDampenedReport(slice2)
		}
		if i == 0 {
			isDecreasing = diff > 0
			if verifyDampenedReport(report[1:]) {
				return true
			}
		} else if (isDecreasing && diff <= 0) || (!isDecreasing && diff >= 0) {
			slice1 := make([]int, len(report[:i]))
			copy(slice1, report[:i])
			slice1 = append(slice1, report[(i+1):]...)

			slice2 := make([]int, len(report[:i+1]))
			copy(slice2, report[:(i+1)])
			slice2 = append(slice2, report[(i+2):]...)
			return verifyDampenedReport(slice1) || verifyDampenedReport(slice2)
		}
		
	}
	return true
}

// Part 1: Function to verify whether a report is safe
func verifyReport(report []int) bool {
	var isDecreasing bool

	for i := 0; i < len(report) - 1; i++ {
		diff := report[i] - report[i+1]
		if abs(diff) < 1 || abs(diff) > 3 {
			return false
		}
		if i == 0 {
			isDecreasing = diff > 0
		} else if (isDecreasing && diff <= 0) || (!isDecreasing && diff >= 0) {
			return false
		}
	}
	return true
}

// Read the input data file
func checkReports() (int, int) {
	var safe int	// safe report counter
	var safe2 int   // Part 2: safe report counter with dampener

    // I'm assuming the input file is in the same directory and is named input.txt
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Could not read input.txt")
        return 0, 0
    }
	defer file.Close()

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		reportText := scanner.Text()
		report := make([]int, 0)
		numbers := strings.Fields(reportText)
		for _, num := range numbers {
			value, _ := strconv.Atoi(num)
			report = append(report, value)
		}
		
		if verifyReport(report) {
			safe++
		}

		if verifyReport2(report) {
			safe2++
		}
    }
    return safe, safe2
}

func main() {
    safe, safe2 := checkReports()
	fmt.Printf("Part 1: %d safe\nPart 2: %d safe", safe, safe2)
}