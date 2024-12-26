package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
func readData() []int {
	data := []int{}
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
		tiles := strings.Split(line, " ")
		for _, v := range tiles {
			num, _ := strconv.Atoi(v)
			data = append(data, num)
		}
		i++
	}

	return data
}

func main() {
	stones := readData()
	for i := 0; i < 25; i++ {
		temp := []int{}
		copy(temp, stones)
		for _, v := range stones {
			if v == 0 {
				temp = append(temp, 1)
			} else {
				str := strconv.Itoa(v)
				len := len(str)
				if len % 2 == 0 {
					n, _ := strconv.Atoi(str[len/2:])
					m, _ := strconv.Atoi(str[:len/2])
					temp = append(temp, n, m)
				} else {
					temp = append(temp, v * 2024)
				}
			}
		}
		stones = temp
	}
	fmt.Println(len(stones))
}