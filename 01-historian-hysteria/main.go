package main

import (
	"fmt"
	"math"
	"os"
)

// max number of locations on the lists
const Locations = 1000

// Read the input data file
func readInput() (n [Locations]int, m [Locations]int) {
    // I'm assuming the input file is in the same directory and is named input.txt
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Could not read input.txt")
        return
    }
    for i := 0; i < 1000; i++ {
        // Get formated input (x x\n)
        fmt.Fscanf(file, "%d %d\n", &n[i], &m[i])
    }
    return
}

// Find the maximum number in the array
func maxNumber(numbers [Locations]int) int {
    max := numbers[0]
    for _, num := range numbers {
        if num > max {
            max = num
        }
    }
    return max
}

// I used radix sort here mostly just to learn how to implement it
func radixSort(numbers *[Locations]int, n int) {
    max := maxNumber(*numbers)
    for exp := 1; max/exp > 0; exp *= 10 {
        countSort(numbers, n, exp)
    }
}

// Counting sort of numbers based on the digit represented by exp
func countSort(numbers *[Locations]int, n int, exp int){
    var output [Locations]int
    var count [10]int

    for i := 0; i < n; i++ {
        count[((*numbers)[i]/exp)%10]++
    }

    for i := 1; i < 10; i++ {
        count[i] += count[i-1]
    }

    for i := n - 1; i >= 0; i-- {
        digit := ((*numbers)[i] / exp) % 10
        output[count[digit]-1] = (*numbers)[i]
        count[digit]--
    }

    for i := 0; i < n; i++ {
        (*numbers)[i] = output[i]
    }
}

// Compare two arrays of sorted locations, get the sum of differences
func compare(n, m [Locations]int) int{
    var diff int
    for i := 0; i < Locations; i++ {
        diff += int(math.Abs(float64(n[i] - m[i])))
    }
    return diff
}

func main() {
    var n, m [Locations]int = readInput()
    radixSort(&n, Locations)
    radixSort(&m, Locations)
    fmt.Print(compare(n, m))
}