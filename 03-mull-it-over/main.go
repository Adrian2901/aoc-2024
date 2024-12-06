package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
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

	fmt.Println(result)
}