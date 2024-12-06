package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part1() int {
	result := 0

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			result += a * b
		}
    }

	return result
}

func part2() int {
	result := 0
	isEnabled := true

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
		matches := re.FindAllStringSubmatch(line, -1)
		
		for _, match := range matches {
			if match[0] == "do()" {
				isEnabled = true
			} else if match[0] == "don't()" {
				isEnabled = false
			} else if isEnabled {
				a, _ := strconv.Atoi(match[1])
				b, _ := strconv.Atoi(match[2])
				result += a * b
			}
		}
    }

	return result
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}