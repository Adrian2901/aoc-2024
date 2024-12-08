package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Read the input data file
func readData() map[int][]int {
	equations := map[int][]int{}
    // I'm assuming the input file is in the same directory and is named input.txt
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Could not read input.txt")
    }
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
    for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, " ")
		var result int
		for i, v := range numbers {
			if i == 0 {
				result, _ = strconv.Atoi(strings.Trim(v, ":"))
			} else {
				number, _ := strconv.Atoi(v)
				equations[result] = append(equations[result], number)
			}
		}
	}

	return equations
}

// Recursive helper function to check all possibilities
func helper(current int, result int, numbers []int) bool {
	// Base case
	if len(numbers) == 0 {
		return current == result
	} 

	n := numbers[0]
	addition := helper(current+n, result, numbers[1:])
	multiplication := helper(current*n, result, numbers[1:])
	// Part 2: Check concatenated numbers as well
	concatenation := helper(concatenate(current, n), result, numbers[1:])

	return addition || multiplication || concatenation
}

// Part 2: Concatenate two numbers
func concatenate(n int, m int) int {
	s1 := strconv.Itoa(n)
	s2 := strconv.Itoa(m)
	result, _ := strconv.Atoi(s1+s2)
	return result
}


func main() {
	equations := readData()
	var sum int
	for k, v := range equations {
		if helper(0, k, v) {
			sum += k
		}
	}
	// Just printing Part 2 for this one, it's similar enough
	fmt.Printf("Result: %d\n", sum)
}