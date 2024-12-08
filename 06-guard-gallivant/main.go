package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Read the input data file
func readData() ([][]string, [2]int) {
	initMap := make([][]string, 0)
	initPos := [2]int{}
    // I'm assuming the input file is in the same directory and is named input.txt
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Could not read input.txt")
    }
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var i int
    for scanner.Scan() {
		line := scanner.Text()
		// Find the guard (and their initial position)
		if strings.Contains(line, "^") {
			j := strings.Index(line, "^")
			initPos = [2]int{i, j}
		} else {
			i++
		}
		chars := strings.Split(line, "")
		initMap = append(initMap, chars)
		
	}
	return initMap, initPos
}

func guardPatrol(initMap [][]string, pos [2]int) {
	// Ordered list of all possible directions (up, right, down, left)
	var directions = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var dir = 0	// Default option is to go up
	// Get the mp boundaries
	maxY := len(initMap) - 1
	maxX := len(initMap[0]) - 1
	// "Set" of visited coordinates
	visited := map[[2]int]bool{}
	visited[pos] = true
	

	// This is a horrible for loop, but it works and is fairly concise
	for next := [2]int{pos[0]+directions[dir][0], pos[1]+directions[dir][1]};
		next[0] >= 0 && next[0] <= maxY && next[1] >= 0 && next[1] <= maxX; 	// Check if the guard is within the map boundaries
		next = [2]int{pos[0]+directions[dir][0], pos[1]+directions[dir][1]} {	// Update the next tile for the guard to go to
			// Check if the next tile is an obstacle
			if initMap[next[0]][next[1]] == "#" {
				// Turn the guard 90 degrees
				dir++
				if dir > len(directions)-1{
					dir = 0
				}
			} else {
				// Add the next tile to the visited set, and move the guard there
				visited[[2]int {next[0], next[1]}] = true
				pos = [2]int {next[0], next[1]}
			}
	}
	fmt.Printf("Part 1: %d\n", len(visited))
}


func main() {
	initMap, initPos := readData()
	guardPatrol(initMap, initPos)
}