package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	total := 0
	f, err := os.Open("../input")
	if err != nil {
		log.Fatal("Input file isn't there?", err.Error())
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		total = total + getCalVal(s.Text())
	}
	fmt.Println(total)
}

func getCalVal(line string) int {
	// On each line, the calibration value can be found by combining the first digit
	// and the last digit (in that order) to form a single two-digit number
	digits := []int{}
	acc := ""
	for _, c := range line {
		if c >= 48 && c <= 57 {
			acc = ""
			digits = append(digits, int(c)-48)
		} else {
			acc += string(c)
			if ok, val := containsNumber(acc); ok {
				digits = append(digits, val)
				acc = string(acc[len(acc)-1]) // keep the last char in case of 'oneight' etc.
			}
		}
	}
	l := len(digits)
	if l > 0 {
		return int(digits[0]*10) + int(digits[l-1])
	}
	return 0
}

func containsNumber(s string) (bool, int) {
	numbers := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	for i, n := range numbers {
		if strings.Index(s, n) >= 0 {
			return true, i
		}
	}
	return false, -1
}
