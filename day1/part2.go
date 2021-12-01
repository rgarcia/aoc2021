package main

import (
	"bufio"
	"container/ring"
	"log"
	"os"
	"strconv"
)

type window struct {
	r *ring.Ring // because why not
}

func newWindow() window {
	r := ring.New(3)
	return window{r}
}

func (w window) isEmpty() bool {
	empty := true
	w.r.Do(func(p interface{}) {
		if p != nil {
			empty = false
		}
	})
	return empty
}

func (w window) isFull() bool {
	ready := true
	w.r.Do(func(p interface{}) {
		if p == nil {
			ready = false
		}
	})
	return ready
}

func (w *window) insert(value int) {
	w.r.Value = value
	w.r = w.r.Next()
}

func (w *window) sum() int {
	sum := 0
	w.r.Do(func(p interface{}) {
		if p != nil {
			sum += p.(int)
		}
	})
	return sum
}

func main() {
	inputFile, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	increases := 0
	prevWindow := newWindow()
	currWindow := newWindow()
	for scanner.Scan() {
		measurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if prevWindow.isEmpty() && currWindow.isEmpty() {
			// first measurement: insert into prevWindow only
			prevWindow.insert(measurement)
		} else {
			currWindow.insert(measurement)
			if prevWindow.isFull() && currWindow.sum() > prevWindow.sum() {
				increases++
			}
			prevWindow.insert(measurement)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	print(increases)
}
