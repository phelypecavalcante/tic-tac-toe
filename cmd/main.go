package main

import (
	"github.com/phelypecavalcante/tic-tac-toe/internal/models"
	"github.com/phelypecavalcante/tic-tac-toe/internal/service"
)

func main() {
	p1 := models.NewPlayer("Player One", "X")
	p2 := models.NewPlayer("Player Two", "O")
	m := service.NewMatch(p1, p2)
	m.Start()
}
