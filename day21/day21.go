package main

import (
	. "../util"
	"fmt"
	"strings"
)

type DataType [2]int

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	result := [2]int{}
	for i, line := range dataSplit {
		result[i] = ParseInt(line[28:])
	}

	return result
}

type Player struct {
	position int
	score int
}

type Game struct {
	players [2]*Player
	turn int
	winValue int
	scoreToWin int
}

func NewGame(data DataType, scoreToWin int) Game {
	return Game{
		players: [2]*Player{{position: data[0]}, {position: data[1]}},
		turn: 0,
		winValue: 1,
		scoreToWin: scoreToWin,
	}
}

func (game *Game) PlayOneTurn(diceSum int) bool {
	currentPlayer := game.players[game.turn % 2]

	currentPlayer.position = (((currentPlayer.position + diceSum) - 1) % 10) + 1
	currentPlayer.score += currentPlayer.position
	game.turn++

	return currentPlayer.score >= game.scoreToWin
}

func (game *Game) GetWinner() *Player {
	for _, p := range game.players {
		if p.score >= game.scoreToWin {
			return p
		}
	}

	return nil
}

func (game *Game) GetLooser() *Player {
	for _, p := range game.players {
		if p.score < game.scoreToWin {
			return p
		}
	}

	return nil
}

func (game *Game) GetPlayerID(player *Player) int {
	for i, p := range game.players {
		if p == player {
			return i
		}
	}

	return -1
}

func (game *Game) Copy() Game {
	newGame := Game{
		players:    [2]*Player{},
		turn:       game.turn,
		winValue:   game.winValue,
		scoreToWin: game.scoreToWin,
	}

	for i, p := range game.players {
		newGame.players[i] = &Player{
			position: p.position,
			score:    p.score,
		}
	}

	return newGame
}

func solvePart1(data DataType) (rc int) {
	game := NewGame(data, 1000)

	dice := 1
	for true {
		diceSum := 3 * dice + 3
		dice += 3
		if game.PlayOneTurn(diceSum) {
			return game.GetLooser().score * (dice - 1)
		}
	}

	return -1
}

func solvePart2(data DataType) (rc int) {
	game := NewGame(data, 21)

	quantumDiceSum := make(map[int]int)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				quantumDiceSum[i+j+k]++
			}
		}
	}

	winCount := [2]int{}

	var play func(Game)
	play = func(game Game) {
		for diceSum, diceSumWinValue := range quantumDiceSum {
			gameCopy := game.Copy()
			gameCopy.winValue *= diceSumWinValue
			if gameCopy.PlayOneTurn(diceSum) {
				winCount[gameCopy.GetPlayerID(gameCopy.GetWinner())] += gameCopy.winValue
			} else {
				play(gameCopy)
			}
		}
	}

	play(game)

	return Max(winCount[0], winCount[1])
}

func main() {
	data := parseData(FetchInputData(21))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
