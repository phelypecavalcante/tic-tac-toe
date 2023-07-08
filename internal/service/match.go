package service

import (
	"bufio"
	"fmt"
	"github.com/phelypecavalcante/tic-tac-toe/internal/models"
	"os"
	"strconv"
)

var mapper = map[string]string{
	"1": "00",
	"2": "01",
	"3": "02",
	"4": "10",
	"5": "11",
	"6": "12",
	"7": "20",
	"8": "21",
	"9": "22",
}

type Match struct {
	coordinate [3][3]*models.Player
	board      models.Board
	playerOne  models.Player
	playerTwo  models.Player
	reader     *bufio.Reader
}

func NewMatch(p1 models.Player, p2 models.Player) Match {
	return Match{
		board:     models.NewBoard(),
		playerOne: p1,
		playerTwo: p2,
		reader:    bufio.NewReader(os.Stdin),
	}
}

func (m *Match) Start() {
	fmt.Println("board is ready!")
	m.board.Print()
	for turn := 0; turn < 9; turn++ {
		m.move(turn)
		m.board.Print()
		if winner := m.getWinner(); winner != nil {
			fmt.Printf("\nCongrats! Player %s won if '%s' symbol\n", winner.GetName(), winner.GetSymbol())
			return
		}
	}
	fmt.Println("It was a draw!")
}

func (m *Match) move(turn int) {
	var player models.Player
	fmt.Printf("\n| --- starting turn %d --- |\n", turn+1)
	if turn%2 == 0 {
		player = m.playerOne
	} else {
		player = m.playerTwo
	}
	fmt.Printf("\n%d# %s insert a number between 1 and 9 to make a move!!\n", turn+1, player.GetName())
	coord, err := m.readMove()
	if err == nil {
		if err := m.registerMove(player, coord); err != nil {
			fmt.Printf("\nerr: %s, unable to register move\n", err)
			m.board.Print()
			m.move(turn)
		}
		return
	}
	fmt.Printf("\nerror: %s, restarting turn\n", err)
	m.board.Print()
	m.move(turn)

}

func (m *Match) readMove() (*[2]int, error) {
	place, _ := m.reader.ReadString('\n')
	fmt.Printf("you've selected " + place)
	if coordStr, ok := mapper[place[:len(place)-1]]; ok {
		i, _ := strconv.Atoi(string(coordStr[0]))
		j, _ := strconv.Atoi(string(coordStr[1]))
		return &[2]int{i, j}, nil
	}
	return nil, fmt.Errorf("invalid move")
}

func (m *Match) registerMove(player models.Player, coord *[2]int) error {
	if m.coordinate[coord[0]][coord[1]] != nil {
		return fmt.Errorf("position already taken")
	}
	m.coordinate[coord[0]][coord[1]] = &player
	m.board.Insert(player, coord[0], coord[1])
	return nil
}

func (m *Match) getWinner() *models.Player {
	for i := 0; i < 3; i++ {
		if m.board.IsEqual([2]int{i, 0}, [2]int{i, 1}, [2]int{i, 2}) {
			return m.coordinate[i][0]
		}

		if m.board.IsEqual([2]int{0, i}, [2]int{1, i}, [2]int{2, i}) {
			return m.coordinate[i][0]
		}
	}

	if m.board.IsEqual([2]int{0, 0}, [2]int{1, 1}, [2]int{2, 2}) {
		return m.coordinate[0][0]
	}

	if m.board.IsEqual([2]int{0, 2}, [2]int{1, 1}, [2]int{2, 0}) {
		return m.coordinate[0][2]
	}

	return nil
}
