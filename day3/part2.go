package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rating int64

const (
	Oxygen Rating = iota
	CO2
)

func rating(r Rating, input []string, bit int) int64 {
	if bit == len(input[0]) {
		panic("fail")
	}
	// determine most common at bit position
	ones, zeros := 0, 0
	for i := range input {
		if input[i][bit] == '1' {
			ones++
		} else {
			zeros++
		}
	}
	// apply bit criteria
	new := []string{}
	for i := range input {
		if r == Oxygen && (ones >= zeros && input[i][bit] == '1' ||
			zeros > ones && input[i][bit] == '0') {
			new = append(new, input[i])
		} else if r == CO2 && (ones >= zeros && input[i][bit] == '0' ||
			zeros > ones && input[i][bit] == '1') {
			new = append(new, input[i])
		}
	}
	if len(new) == 1 {
		v, err := strconv.ParseInt(new[0], 2, 64)
		if err != nil {
			panic(err)
		}
		return v
	}
	return rating(r, new, bit+1)
}

func oxygen(input []string) int64 {
	return rating(Oxygen, input, 0)
}

func c02(input []string) int64 {
	return rating(CO2, input, 0)
}

func main() {
	inputFile, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	var input []string
	for scanner.Scan() {
		input = append(input, strings.TrimSpace(scanner.Text()))
	}
	print(oxygen(input) * c02(input))
}
