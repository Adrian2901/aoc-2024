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

func moveFiles(diskMap map[int]int, max int) {
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


func main() {
	diskMap, max := readData()
	moveFiles(diskMap, max)

	var sum int
	for k, v := range diskMap {
		sum += k*v
	}
	fmt.Println("Part 1: ", sum)
}