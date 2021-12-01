package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	increases := 0
	var prev *int
	for scanner.Scan() {
		measurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if prev != nil && measurement > *prev {
			increases++
		}
		prev = &measurement
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	print(increases)
}
