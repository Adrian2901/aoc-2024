package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Read the input data file
func readData() (map[string][]string, [][]string) {
	rules := make(map[string][]string)
	updates := make([][]string, 0)

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Could not read input.txt")
    }
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Read the rules until an empty line (divider)
    for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rule := strings.Split(line, "|")
		rules[rule[0]] = append(rules[rule[0]], rule[1])
	}
	// Read update pages orders
	for scanner.Scan() {
		line := scanner.Text()
		order := strings.Split(line, ",")
		updates = append(updates, order)
	}

	return rules, updates
}

// Part 1: Verify whether the pages in updates are correctly ordered
func verifyOrder(rules map[string][]string, updates [][]string) int {
	var score int
	for _, update := range updates {
		for i := len(update)-1; i > 0; i-- {
			page := update[i]
			if !isOrdered(update[:i], rules[page]){
				break
			} else if i == 1 {
				// If an order is valid, add its middle value to the score
				middleNumber, _ := strconv.Atoi(update[(len(update)-1)/2])
				score += middleNumber
			}
		}
	}
	return score
}

// Check if the update pages don't break any rules
func isOrdered(update []string, priorities []string) bool {
	for _, page := range update {
		if slices.Contains(priorities, page) {
			return false
		}
	}
	return true
}


func main() {
	rules, updates := readData()
	result := verifyOrder(rules, updates)
	fmt.Printf("Part 1: %d\n", result)
}