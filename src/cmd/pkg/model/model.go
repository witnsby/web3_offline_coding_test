package model

type Choices struct {
	Option []string
}

type ResponsesPlayers struct {
	Player1 string
	Player2 string
}

type ResponsesBot struct {
	Player string
	Bot    string
}

type WinningCases struct {
	Options map[string]string
}
