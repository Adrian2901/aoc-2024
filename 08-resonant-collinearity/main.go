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
	antinodes := map[[2]int]bool{}
	
	for _, v := range antennas {
		for i := 0; i < len(v) - 1; i++ {
			for j := i+1; j < len(v); j++ {
				// Calculate the distances between antennas
				yDiff := abs(v[i][0] - v[j][0])
				xDiff := abs(v[i][1] - v[j][1])
				// Check directions in relation to antenna i
				var yDir, xDir int
				if v[i][0] - v[j][0] > 0 {
					yDir = -1
				} else {
					yDir = 1
				}
				if v[i][1] - v[j][1] > 0 {
					xDir = -1
				} else {
					xDir = 1
				}
				newY := v[i][0] 
				newX := v[i][1] 
				// Keep adding antennas in intervals until we hit a border
				for n := 1; newY >=0 && newY < maxY && newX >=0 && newX < maxX; n++ {
					antinodes[[2]int{newY, newX}] = true
					newY = v[i][0] + (yDiff*-yDir*n)
					newX = v[i][1] + (xDiff*-xDir*n)
				}
				
				// Do the same in the other direction
				newY = v[j][0] 
				newX = v[j][1]
				for n := 1; newY >=0 && newY < maxY && newX >=0 && newX < maxX; n++ {
					antinodes[[2]int{newY, newX}] = true
					newY = v[j][0] + (yDiff*yDir*n)
					newX = v[j][1] + (xDiff*xDir*n)
					
				}
			}
		}
	}
	return len(antinodes)
}


func main() {
	antennas := readData()
	fmt.Println(placeAntinodes(antennas))
}