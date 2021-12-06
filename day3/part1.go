package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type power struct {
	ones  []int
	zeros []int
}

func (p *power) update(s string) {
	if len(p.ones) == 0 {
		p.ones = make([]int, len(s))
		p.zeros = make([]int, len(s))
	}
	for i := range s {
		if s[i] == '1' {
			p.ones[i]++
		} else if s[i] == '0' {
			p.zeros[i]++
		}
	}
}

func (p power) value() int64 {
	gamma := ""
	epsilon := ""
	for i := 0; i < len(p.ones); i++ {
		if p.ones[i] >= p.zeros[i] {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	g, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		panic(err)
	}
	e, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		panic(err)
	}
	return g * e
}

func main() {
	inputFile, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	var p power
	for scanner.Scan() {
		p.update(strings.TrimSpace(scanner.Text()))
	}
	print(p.value())
}
