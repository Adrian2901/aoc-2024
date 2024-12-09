package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Read the input data file
func readData() (map[int]int, int) {
	diskMap := map[int]int{}
    // I'm assuming the input file is in the same directory and is named input.txt
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Could not read input.txt")
    }
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var i, j int
    for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, "")
		for k, v := range data {
			number, _ := strconv.Atoi(v)
			if k & 1 == 0 {
				for range number {
					diskMap[i] = j
					i++
				}
				j++
			} else {
				i += number
			}
		}
	}

	return diskMap, i
}

// Part 1: Move file blocks to fill all gaps
func packFiles(diskMap map[int]int, max int) {
	l, r := 0, max-1
	for l < r {
		if _, ok := diskMap[l]; !ok {
			if _, ok := diskMap[r]; ok {
				diskMap[l] = diskMap[r]
				delete(diskMap, r)
				l++
			}
			r--
		} else {
			l++
		}
	}
}

// Part 2: Move whole files if there's enough space
func shiftFiles(diskMap map[int]int, max int) {
	l, r := 0, max-1
	for 1 < r {
		// Move the right pointer until we find a file
		if _, ok := diskMap[r]; !ok {
			r--
			continue
		}

		// Count size of the current file
		size := 1
		for diskMap[r] == diskMap[r-1] {
			size++
			r--
		}

		// Find a gap large enough to move the file
		gap := 0
		l = 0
		for l < r && gap < size {
			if _, ok := diskMap[l]; !ok {
				gap++
			} else {
				gap = 0
			}
			l++
		}

		// Move the file if a large enough gap was found
		if gap >= size {
			for i := 0; i < size; i++ {
				diskMap[l-i-1] = diskMap[r+i]
				delete(diskMap, r+i)
			}
		}
		r--
	}
}





func main() {
	diskMap, max := readData()

	// Part 1:
	// packFiles(diskMap, max)

	// Part 2:
	shiftFiles(diskMap, max)

	// Calculate checksum
	var sum int
	for k, v := range diskMap {
		sum += k*v
	}

	fmt.Println("Result: ", sum)
}