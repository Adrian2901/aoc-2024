package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func verifyReport(report []int) bool {
	var isDecreasing bool

	for i := 0; i < len(report) - 1; i++ {
		diff := report[i] - report[i+1]
		if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
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
func checkReports() int {
	var safe int	// safe report counter

    // I'm assuming the input file is in the same directory and is named input.txt
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Could not read input.txt")
        return 0
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
			fmt.Println("Safe report:", report)
			safe++
		}
    }
    return safe
}

func main() {
    safe := checkReports()
	fmt.Printf("Safe: %d\n", safe)
}