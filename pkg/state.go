package pkg

import (
	"math"
	"math/rand"
)

func (state State) FreeBox() []int {
	fb := make([]int, 0, 9)
	for index, box := range state {
		if box == BoxEmpty {
			fb = append(fb, index)
		}
	}
	return fb
}

func (state State) CheckWinner() (Player, bool) {
	for _, winState := range winStates {
		b0 := state[winState[0]]
		b1 := state[winState[1]]
		b2 := state[winState[2]]
		if b0 == b1 && b1 == b2 {
			switch b0 {
			case BoxEmpty: break
			case BoxX: return PlayerX, true
			case BoxO: return PlayerO, true
			default: panic("check winner weird....")
			}
		}
	}
	return PlayerNil, false
}

func (state State) CurrentPlayer() Player {
	switch len(state.FreeBox())%2 {
	case 0: return PlayerO
	case 1: return PlayerX
	default: return PlayerNil
	}
}

func (state State) IsOver() (Player, bool) {
	winner, won := state.CheckWinner()
	if won {
		return winner, won
	}
	return winner, len(state.FreeBox())==0
}

func (state State) Play(index int) (State, error) {
	if winner, over := state.IsOver(); over {
		return state, &GameOverError{winner: winner}
	} else if state[index] != BoxEmpty {
		return state, &BoxChosenError{index: index, chosenSquare: state[index].Val}
	} else {
		newState := state
		switch state.CurrentPlayer() {
		case PlayerO: newState[index] = BoxO
		case PlayerX: newState[index] = BoxX
		default: return state, &InvalidPlayerError{player: state.CurrentPlayer()}
		}
		return newState, nil
	}
}

func (state State) NextStates() []State {
	states := make([]State, 0, 9)
	for _, index := range state.FreeBox() {
		nextState, err := state.Play(index)
		if err == nil {
			states = append(states, nextState)
		}
	}
	return states
}

func (state State) Hash() int {
	var hash int
	for i, box := range state {
		hash += box.Val * int(math.Pow(3, float64(i)))
	}
	return hash
}

func (state State) Score() Score {
	if score, ok := AllScores[state.Hash()]; ok {
		return score
	}
	panic("Game hasn't initialized allscores")
}

func (state State) OptimalMoves() []State {
	score :=  state.Score()
	var optimalMoves []State
	for _, nextState := range state.NextStates() {
		if nextState.Score() == score {
			optimalMoves = append(optimalMoves, nextState)
		}
	}
	return optimalMoves
}

func (state State) NextBestMove() (State, error) {
	score, ok := AllScores[state.Hash()]
	if !ok {
		return state, &ScoresNotInitializedError{state: state}
	}

	minFinish, maxFinish, err := minMaxFinishMove(state)
	if err != nil {
		return state, err
	}

	switch state.CurrentPlayer() {
	case PlayerX:
		if score == ScoreX {
			return minFinish, nil
		} else {
			return maxFinish, nil
		}
	case PlayerO:
		if score == ScoreO {
			return minFinish, nil
		} else {
			return maxFinish, nil
		}
	default:
		return state, &InvalidPlayerError{state.CurrentPlayer()}
	}
}

//func optimalMoveForScore0(state State)

func minMaxFinishMove(state State) (minState, maxState State, err error) {
	min := 10
	max := -1
	optimalMoves := state.OptimalMoves()

	if len(optimalMoves) == 0 {
		winner, _ := state.CheckWinner()
		return state, state, &GameOverError{winner: winner}
	}

	for _, move := range optimalMoves {
		if minDepth, _ := AllMinDepths[move.Hash()]; minDepth < min {
			min = minDepth
		}
		if maxDepth, _ := AllMaxDepths[move.Hash()]; maxDepth > max {
			max = maxDepth
		}
	}

	var minMoves, maxMoves []State
	for _, nextState := range optimalMoves {
		if AllMinDepths[nextState.Hash()] == min {
			minMoves = append(minMoves, nextState)
		}
		if AllMaxDepths[nextState.Hash()] == max {
			maxMoves = append(maxMoves, nextState)
		}
	}

	if len(minMoves) == 0 || len(maxMoves) == 0 {
		err = &InvalidStateError{state }
		return
	}
	minState = minMoves[rand.Intn(len(minMoves))]
	maxState = maxMoves[rand.Intn(len(maxMoves))]
	return
}



