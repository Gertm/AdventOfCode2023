package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	RED   = 12
	GREEN = 13
	BLUE  = 14
)

func main() {
	total := 0
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("Input file isn't there?", err.Error())
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		total = total + lineID(s.Text())
	}
	fmt.Println(total)
}

type Reveal struct {
	Red   int
	Green int
	Blue  int
}

func RevealFromString(r string) Reveal {
	result := Reveal{}
	cubes := strings.Split(r, ",")
	for _, cube := range cubes {
		parts := strings.Split(strings.TrimSpace(cube), " ")
		number, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println(err)
		}
		switch strings.TrimSpace(parts[1]) {
		case "green":
			result.Green = number
		case "red":
			result.Red = number
		case "blue":
			result.Blue = number
		}
	}
	return result
}

func (r Reveal) Valid() bool {
	if r.Red <= RED && r.Blue <= BLUE && r.Green <= GREEN {
		return true
	}
	return false
}

func lineID(line string) int {
	gameID, err := strconv.Atoi(line[5:strings.Index(line, ":")])
	if err != nil {
		log.Fatal(err)
	}

	reveals := strings.Split(line[strings.Index(line, ":")+1:], ";")
	for _, r := range reveals {
		if !RevealFromString(r).Valid() {
			return 0
		}
	}

	return gameID
}
