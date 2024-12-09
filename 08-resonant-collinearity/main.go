package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var maxY, maxX int

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Read the input data file
func readData() map[string][][2]int {
	antennas := map[string][][2]int{}
    // I'm assuming the input file is in the same directory and is named input.txt
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Could not read input.txt")
    }
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var i int // rows
    for scanner.Scan() {
		line := scanner.Text()
		maxX = len(line)
		tiles := strings.Split(line, "")
		for j, v := range tiles {
			if v != "." {
				antennas[v] = append(antennas[v], [2]int{i, j})
			}
		}
		i++
	}
	maxY = i

	return antennas
}

func placeAntinodes(antennas map[string] [][2]int) int {
	potentialAntinodes := make([][2]int, 0)
	antinodes := map[[2]int]bool{}
	
	for _, v := range antennas {
		for i := 0; i < len(v) - 1; i++ {
			for j := i+1; j < len(v); j++ {
				xDiff := abs(v[i][0] - v[j][0])
				yDiff := abs(v[i][1] - v[j][1])
				// TODO: I will refactor this later, I was just tesing if it works (it does)
				var antinode1X, antinode1Y, antinode2X, antinode2Y int
				if v[i][0] < v[j][0] {
					antinode1X = v[i][0] - xDiff
					antinode2X = v[j][0] + xDiff
				} else {
					antinode1X = v[i][0] + xDiff
					antinode2X = v[j][0] - xDiff
				}

				if v[i][1] < v[j][1] {
					antinode1Y = v[i][1] - yDiff
					antinode2Y = v[j][1] + yDiff
				} else {
					antinode1Y = v[i][1] + yDiff
					antinode2Y = v[j][1] - yDiff
				}
				potentialAntinodes = append(potentialAntinodes, [2]int{antinode1X, antinode1Y}, [2]int{antinode2X, antinode2Y})
			}
		}
	}
	for _, v := range potentialAntinodes {
		if v[0] >= 0 && v[0] < maxY && v[1] >=0 && v[1] < maxX {
			antinodes[v] = true
			fmt.Println(v)
		}
	}
	return len(antinodes)
}


func main() {
	antennas := readData()
	fmt.Println(placeAntinodes(antennas))
}