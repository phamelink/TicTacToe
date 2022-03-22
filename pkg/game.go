package pkg

type Player int
type State [9]Box
type winning []int
type Score int

type Box struct {
	Val     int    `json:"val`
	Content string `json:"content"`
}

type Game struct {
	CurrentPlayer Player `json:"current_player"`
	Winner Player `json:"winner,omitempty"`
	GameState State `json:"game_state"`
}

var winStates = []winning{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},

	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},

	{0, 4, 8},
	{2, 4, 6},
}

const (
	PlayerNil = iota
	PlayerX
	PlayerO
)

const (
	ScoreO = iota -1
	ScoreNil
	ScoreX
)

var (
	BoxEmpty = Box{Val: 0, Content: " "}
	BoxX = Box{Val: 1, Content: "X"}
	BoxO = Box{Val: 2, Content: "O"}
)

var InitGame = State{BoxEmpty, BoxEmpty, BoxEmpty, BoxEmpty, BoxEmpty, BoxEmpty, BoxEmpty, BoxEmpty, BoxEmpty}

var AllStates = allStates()

var AllScores = allScores()

var AllMinDepths, AllMaxDepths = allDepths()



func (player Player) String() string {
	switch player {
	case PlayerO: return "O"
	case PlayerX: return "X"
	default: return " "
	}
}


func (state State) String() string {
	str := "\n "
	for i := 0; i < 9; i++ {
		box := ""
		switch state[i] {
		case BoxX: box = "X"
		case BoxO: box = "O"
		case BoxEmpty: box = " "
		}
		str += box
		if (i+1)%3!=0 {
			str += " | "
		}
		if i == 2  || i == 5 {
			str += "\n - + - + -\n "
		}
	}
	return str
}





