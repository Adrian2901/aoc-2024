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
func verifyOrder(rules map[string][]string, updates [][]string) {
	var score int
	incorrect := make([][]string, 0)
	for _, update := range updates {
		for i := len(update)-1; i > 0; i-- {
			page := update[i]
			if !isOrdered(update[:i], rules[page]){
				incorrect = append(incorrect, update)
				break
			} else if i == 1 {
				// If an order is valid, add its middle value to the score
				middleNumber, _ := strconv.Atoi(update[(len(update)-1)/2])
				score += middleNumber
			}
		}
	}
	defer fixOrders(rules, incorrect)

	fmt.Printf("Part 1: %d\n", score)
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

// Part 2: Fix incorrect updates and calculate their score
func fixOrders(rules map[string][]string, updates [][]string) {
	var score int
	for _, update := range updates {
		// My first approach was unstable and missed some cases, so I had to stabilze it
		stable := false
		for !stable {
			stable = true
			for i := len(update) - 1; i > 0; i-- {
				page := update[i]
				for j := i - 1; j >= 0; j-- {
					if slices.Contains(rules[page], update[j]) {
						// Swap and mark as unstable
						update[i], update[j] = update[j], update[i]
						stable = false
					}
				}
			}
		}
		// Calculate the score (middle element as number)
		middleIndex := (len(update) - 1) / 2
		middleNumber, _ := strconv.Atoi(update[middleIndex])
		score += middleNumber
	}
	fmt.Printf("Part 2: %d\n", score)
}



func main() {
	rules, updates := readData()
	verifyOrder(rules, updates)
}