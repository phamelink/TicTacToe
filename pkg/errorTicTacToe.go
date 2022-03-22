package pkg

import "fmt"

type GameOverError struct {
	winner Player
}

type BoxChosenError struct {
	index           int
	chosenSquare 	int
}

type InvalidStateError struct {
	state State
}

type InvalidPlayerError struct {
	player Player
}

type ScoresNotInitializedError struct {
	state State
}

func (e *GameOverError) Error() string {
	return fmt.Sprintf("Game is already over, %v won!", e.winner)
}

func (e *BoxChosenError) Error() string {
	return fmt.Sprintf("Square i:%d already selected by %s", e.index, e.chosenSquare)
}

func (e *InvalidStateError) Error() string {
	return fmt.Sprintf("Chosen state %v is invalid", e.state)
}

func (e *InvalidPlayerError) Error() string {
	return fmt.Sprintf("Player %v is invalid", e.player)
}

func (e *ScoresNotInitializedError) Error() string {
	return fmt.Sprintf("Score not initialized for state %v", e.state)
}