package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type position struct {
	horizontal, depth, aim int
}

func (c *position) forward(i int) {
	c.horizontal += i
	c.depth += i * c.aim
}

func (c *position) down(i int) {
	c.aim += i
}

func (c *position) up(i int) {
	c.aim -= i
}

func main() {
	inputFile, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	var pos position
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if len(parts) != 2 {
			continue
		}
		dir := parts[0]
		mag, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		switch dir {
		case "up":
			pos.up(mag)
		case "down":
			pos.down(mag)
		case "forward":
			pos.forward(mag)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	print(pos.horizontal * pos.depth)
}
