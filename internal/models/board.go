package models

import (
	"fmt"
)

type Board struct {
	matrix [3][3]string
}

func NewBoard() Board {
	return Board{[3][3]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}}
}

func (b *Board) Print() {
	fmt.Printf(
		"\n %s | %s | %s\n %s | %s | %s\n %s | %s | %s\n",
		b.matrix[0][0], b.matrix[0][1], b.matrix[0][2],
		b.matrix[1][0], b.matrix[1][1], b.matrix[1][2],
		b.matrix[2][0], b.matrix[2][1], b.matrix[2][2],
	)
}

func (b *Board) Insert(player Player, row int, column int) {
	b.matrix[row][column] = player.symbol
}

func (b *Board) IsEqual(a [2]int, coords ...[2]int) bool {
	if len(coords) > 1 {
		for _, coord := range coords {
			if b.matrix[a[0]][a[1]] != b.matrix[coord[0]][coord[1]] || b.matrix[a[0]][a[1]] == "_" {
				return false
			}
		}
	}
	return true
}
