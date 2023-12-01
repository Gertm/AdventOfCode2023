package main

import (
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_readInput(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processInput(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
*/

func Test_getCalVal(t *testing.T) {
	tests := []struct {
		args string
		want int
	}{
		{
			"1abc2",
			12,
		},
		{
			"pqr3stu8vwx",
			38,
		},
		{
			"a1b2c3d4e5f",
			15,
		},
		{
			"treb7uchet",
			77,
		},
		{
			"two1nine",
			29,
		},
		{
			"eightwothree",
			83,
		},
		{
			"abcone2threexyz",
			13,
		},
		{
			"xtwone3four",
			24,
		},
		{
			"4nineeightseven2",
			42,
		},
		{
			"wdfwoneightowdi",
			18,
		},
	}
	for _, tt := range tests {
		t.Run("aoc1", func(t *testing.T) {
			if got := getCalVal(tt.args); got != tt.want {
				t.Errorf("getCalVal() = %v, want %v", got, tt.want)
			}
		})
	}
}
