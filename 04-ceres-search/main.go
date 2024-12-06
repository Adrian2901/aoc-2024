package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Table of all 8 valid directions
// (-1, -1) -> up-left, (1, 1) -> down-right, etc.
var Directions = [][]int{{-1, -1}, {-1, 0}, {0, -1}, {-1, 1}, {1, -1}, {0, 1}, {1, 0}, {1, 1}}

// Read the input data file
func readData() [][]string {
	table := make([][]string, 0)
    // I'm assuming the input file is in the same directory and is named input.txt
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Could not read input.txt")
    }
	defer file.Close()

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		reportText := scanner.Text()
		report := strings.Split(reportText, "")
		table = append(table, report)
	}
	return table
}

// Search for a keyword in a word search table
func search(data [][]string, word string) int {
	count := 0	// Count of found words

	// Iterate through the rows
	for i := 0; i < len(data); i++ {
		// Iterate through colums in a row
		for j := 0; j < len(data[i]); j++ {
			// If the letter matches the first letter of the keyword
			if data[i][j] == string(word[0]) {
				// Iterate through all valid directions
				for _, dir := range Directions {
					correct := 0	// Counter of correct letters in the current direction
					for letter := 0; letter < len(word); letter++ {
						// Move indices in the current direction
						x := j + (letter * dir[1])
						y := i + (letter * dir[0])

						// Break if indices are out of bounds
						if x < 0 || y < 0 || x > len(data[0])-1 || y > len(data)-1 {
							break
						}
						// Check if the processed letter matches the word
						if string(word[letter]) == data[y][x] {
							correct++
						} else {
							break
						}
					}
					// Update the counter if the full word was found
					if correct == len(word) {
						count++
					}
				}
			}
		}
	}
	return count
}


func main() {
	data := readData()
    fmt.Print(search(data, "XMAS"))
}