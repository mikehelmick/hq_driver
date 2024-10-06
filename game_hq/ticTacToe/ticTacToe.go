package tictactoe

import (
	"interfaces"
	"sync"
)

const X, O = 1, 2

type ticTacToe struct {
	mut     *sync.Mutex
	field   [3][3]int
	turn    int
	players [2]*interfaces.Player
	scores  [2]int
}

type ticTacToeClientState struct {
	Field  [3][3]int `json:"Field"`
	Turn   int       `json:"Turn"`
	Scores [2]int    `json:"Scores"`
}

func NewGameTicTacToe() *ticTacToe {
	gState := &ticTacToe{
		field:  [3][3]int{},
		turn:   0,
		mut:    &sync.Mutex{},
		scores: [2]int{},
	}
	return gState
}

func (gState *ticTacToe) reset() {
	gState.field = [3][3]int{}
	gState.turn = 1
}

func (gState *ticTacToe) move(x, y, team int) {
	gState.mut.Lock()
	defer gState.mut.Unlock()
	if gState.turn != team || gState.field[x][y] != 0 {
		return
	}
	gState.field[x][y] = team
	if gState.scan() {
		gState.reset()
		gState.scores[team]++
	} else {
		gState.turn = (gState.turn % 2) + 1
	}
}

func (gState *ticTacToe) scan() bool {
	//horizontal
	//00,01,02
	//10,11,12
	//20,21,22
	possibleTicTacToes := [][3][2]int{
		{
			{0, 0}, {0, 1}, {0, 2},
		}, {
			{1, 0}, {1, 1}, {1, 2},
		}, {
			{2, 0}, {2, 1}, {2, 2},
		}, {
			{0, 0}, {1, 0}, {2, 0},
		}, {
			{0, 1}, {1, 1}, {2, 1},
		}, {
			{0, 2}, {1, 2}, {2, 2},
		}, {
			{0, 0}, {1, 1}, {2, 2},
		}, {
			{0, 2}, {1, 1}, {2, 0},
		},
	}
	for _, streak := range possibleTicTacToes {
		if gState.field[streak[0][0]][streak[0][1]] != 0 &&
			gState.field[streak[0][0]][streak[0][1]] == gState.field[streak[1][0]][streak[1][1]] &&
			gState.field[streak[1][0]][streak[1][1]] == gState.field[streak[2][0]][streak[2][1]] {
			return true
		}
	}
	return false
}

func (gState *ticTacToe) JSON() interfaces.ClientState {
	cState := ticTacToeClientState{
		Field: gState.field,
		Turn:  gState.turn,
	}
	return cState
}

func (gState *ticTacToe) Players() []*interfaces.Player {
	return append([]*interfaces.Player{}, gState.players[0], gState.players[1])
}