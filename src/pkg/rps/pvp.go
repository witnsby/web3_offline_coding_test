package rps

import (
	"bufio"
	"github.com/witnsby/web3_offline_coding_test/src/pkg/model"
	"os"
)

func newResponses() *ResponsesPlayersWrapper {
	return &ResponsesPlayersWrapper{
		&model.ResponsesPlayers{},
	}
}

func PlayPvP() {
	reader := bufio.NewReader(os.Stdin)
	choices := newChoices()
	responses := newResponses()

	getYourChoice()
	responses.Player1 = getPlayerChoice(reader, choices)

	clearConsole()

	getYourChoice()
	responses.Player2 = getPlayerChoice(reader, choices)
	getWinner(responses)
}
