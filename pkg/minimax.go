package pkg

func generateStates(state State, allStates map[int]State) {
	allStates[state.Hash()] = state
	if nextStates := state.NextStates(); len(nextStates) != 0 {
		for _, nextState := range nextStates {
			if _, ok := allStates[nextState.Hash()]; !ok {
				generateStates(nextState, allStates)
			}
		}
	}
}

func allStates() map[int]State {
	var allStates = make(map[int]State)
	generateStates(InitGame, allStates)
	return allStates
}

func generateScores(state State, allScores map[int]Score) {
	if _, ok := allScores[state.Hash()]; ok {
		return
	}
	if winner, over := state.IsOver(); over {
		switch winner {
		case PlayerX: allScores[state.Hash()] = ScoreX
		case PlayerO: allScores[state.Hash()] = ScoreO
		case PlayerNil: allScores[state.Hash()] = ScoreNil
		default: panic(&ScoresNotInitializedError{ state: state})
		}

	} else {
		nextStates := state.NextStates()
		for _, nextState := range nextStates {
			generateScores(nextState, allScores)
		}
		min, max := minMaxScore(nextStates, allScores)
		switch state.CurrentPlayer() {
		case PlayerX: allScores[state.Hash()] = max
		case PlayerO: allScores[state.Hash()] = min
		default: panic(&InvalidPlayerError{player: state.CurrentPlayer()})
		}
	}
}

func minMaxScore(nextStates []State, allScores map[int]Score) (min Score, max Score) {
	max = ScoreO
	min = ScoreX
	for _, nextState := range nextStates {
		score := allScores[nextState.Hash()]
		if max < score {
			max = score
		}
		if min > score {
			min = score
		}
	}
	return
}

func allScores() map[int]Score {
	var allScores = make(map[int]Score)
	generateScores(InitGame, allScores)
	return allScores
}

func generateDepths(state State, minAcc, maxAcc int, minDepths, maxDepths map[int]int) {
	if _, over := state.IsOver(); over {
		minDepths[state.Hash()] = 0
		maxDepths[state.Hash()] = 0
		return
	}
	min := 10
	max := -1
	minAcc++
	maxAcc++
	for _, nextState := range state.NextStates() {
		_, okMin := minDepths[nextState.Hash()]
		_, okMax := maxDepths[nextState.Hash()]
		if !okMin || !okMax {
			generateDepths(nextState, minAcc, maxAcc, minDepths, maxDepths)
		}
		minDepth, _ := minDepths[nextState.Hash()]
		maxDepth, _ := maxDepths[nextState.Hash()]
		if minDepth < min {
			min = minDepth
		}
		if maxDepth > max {
			max = maxDepth
		}
	}
	minDepths[state.Hash()] = min + 1
	maxDepths[state.Hash()] = max + 1
}

func allDepths() (minDepths, maxDepths map[int]int) {
	minDepths = make(map[int]int)
	maxDepths = make(map[int]int)
	generateDepths(InitGame, 0, 0, minDepths, maxDepths)
	return
}

