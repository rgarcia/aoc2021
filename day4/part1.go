package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type board struct {
	b      [5][5]int
	rowIdx int
	marked [5][5]bool
}

var draws []int
var boards []board

func (b *board) addRow(txt string) {
	row := [5]int{}
	for i, part := range strings.Fields(txt) {
		num, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		row[i] = num
	}
	b.b[b.rowIdx] = row
	b.rowIdx++
}

func (b *board) draw(num int) {
	for i := range b.b {
		for j := range b.b[i] {
			if b.b[i][j] == num {
				b.marked[i][j] = true
			}
		}
	}
}

func (b board) wins() (bool, int) {
	for i := 0; i < 5; i++ {
		if b.marked[i][0] && b.marked[i][1] && b.marked[i][2] && b.marked[i][3] && b.marked[i][4] {
			return true, b.score()
		}
		if b.marked[0][i] && b.marked[1][i] && b.marked[2][i] && b.marked[3][i] && b.marked[4][i] {
			return true, b.score()
		}
	}
	return false, 0
}

func (b board) score() int {
	score := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.marked[i][j] {
				score += b.b[i][j]
			}
		}
	}
	return score
}

func main() {
	inputFile, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	boardIdx := -1
	var b *board
	for scanner.Scan() {
		txt := strings.TrimSpace(scanner.Text())
		if draws == nil {
			for _, draw := range strings.Split(txt, ",") {
				num, err := strconv.Atoi(draw)
				if err != nil {
					panic(err)
				}
				draws = append(draws, num)
			}
			continue
		}
		if txt == "" {
			boardIdx++
			if b != nil {
				boards = append(boards, *b)
			}
			b = &board{}
			continue
		}
		b.addRow(txt)
	}
	boards = append(boards, *b)

	for _, draw := range draws {
		for i := range boards {
			b := &boards[i]
			b.draw(draw)
			if winner, score := b.wins(); winner == true {
				fmt.Println(score * draw)
				os.Exit(0)
			}
		}
	}
	fmt.Println("no winner")
}
