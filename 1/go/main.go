package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("AoC 2023 - exercise 1")
	input := readInput()
	total := 0
	for _, line := range input {
		total = total + getCalVal(line)
	}
	fmt.Println(total)
}

func readInput() []string {
	f, err := os.Open("../input")
	if err != nil {
		log.Fatal("Input file isn't there?", err.Error())
	}
	defer f.Close()
	result := make([]string, 1000)
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		result = append(result, s.Text())
	}
	return result
}

func getCalVal(line string) int {
	// On each line, the calibration value can be found by combining the first digit
	// and the last digit (in that order) to form a single two-digit number
	digits := []int{}
	for _, c := range line {
		if c >= 48 && c <= 57 {
			digits = append(digits, int(c)-48)
		}
	}
	l := len(digits)
	if l > 0 {
		return int(digits[0]*10) + int(digits[l-1])
	}
	return 0

}
