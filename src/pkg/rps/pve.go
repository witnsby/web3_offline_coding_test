package rps

import (
	"bufio"
	"github.com/witnsby/web3_offline_coding_test/src/pkg/model"
	"os"
)

func newResponsesBot() *ResponsesBotWrapper {
	return &ResponsesBotWrapper{
		&model.ResponsesBot{},
	}
}

func PlayPvBot() {
	choices := newChoices()
	responsesWithBot := newResponsesBot()

	reader := bufio.NewReader(os.Stdin)

	getYourChoice()
	responsesWithBot.Player = getPlayerChoice(reader, choices)
	responsesWithBot.Bot = getBotChoiceHarmonyVRF(choices.Option)

	getWinner(responsesWithBot)

}
