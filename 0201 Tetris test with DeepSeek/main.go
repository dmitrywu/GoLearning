package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 10
	height = 20
)

var (
	field       [height][width]rune
	currentTetromino [4][4]rune
	tetrominoX, tetrominoY int
)

func main() {
	rand.Seed(time.Now().UnixNano())
	initField()
	spawnTetromino()

	for {
		drawField()
		time.Sleep(500 * time.Millisecond)
		if !moveTetromino(0, 1) {
			mergeTetromino()
			spawnTetromino()
		}
	}
}

func initField() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			field[y][x] = ' '
		}
	}
}

func drawField() {
	fmt.Print("\033[H\033[2J") // Очистка экрана
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Printf("%c ", field[y][x])
		}
		fmt.Println()
	}
}

func spawnTetromino() {
	tetrominoes := [][4][4]rune{
		{
			{' ', ' ', ' ', ' '},
			{' ', '■', '■', ' '},
			{' ', '■', '■', ' '},
			{' ', ' ', ' ', ' '},
		},
		{
			{' ', ' ', ' ', ' '},
			{'■', '■', '■', '■'},
			{' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' '},
		},
		// Добавьте другие тетромино по аналогии
	}

	currentTetromino = tetrominoes[rand.Intn(len(tetrominoes))]
	tetrominoX = width/2 - 2
	tetrominoY = 0
}

func moveTetromino(dx, dy int) bool {
	newX := tetrominoX + dx
	newY := tetrominoY + dy

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if currentTetromino[y][x] != ' ' {
				if newX+x < 0 || newX+x >= width || newY+y >= height || field[newY+y][newX+x] != ' ' {
					return false
				}
			}
		}
	}

	tetrominoX = newX
	tetrominoY = newY
	return true
}

func mergeTetromino() {
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if currentTetromino[y][x] != ' ' {
				field[tetrominoY+y][tetrominoX+x] = currentTetromino[y][x]
			}
		}
	}
}