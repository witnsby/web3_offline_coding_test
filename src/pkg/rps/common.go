package rps

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/witnsby/web3_offline_coding_test/src/pkg/blockchain"
	"github.com/witnsby/web3_offline_coding_test/src/pkg/helper"
	"github.com/witnsby/web3_offline_coding_test/src/pkg/model"
	"math/big"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type client struct {
	payload *strings.Reader
}

type Responses interface {
	GetFirst() string
	GetSecond() string
	GetFirstName() string
	GetSecondName() string
}

type ResponsesPlayersWrapper struct {
	*model.ResponsesPlayers
}

type ResponsesBotWrapper struct {
	*model.ResponsesBot
}

func (r *ResponsesPlayersWrapper) GetFirst() string {
	return r.Player1
}

func (r *ResponsesPlayersWrapper) GetSecond() string {
	return r.Player2
}

func (r *ResponsesPlayersWrapper) GetFirstName() string {
	return "Player 1"
}

func (r *ResponsesPlayersWrapper) GetSecondName() string {
	return "Player 2"
}

func (r *ResponsesBotWrapper) GetFirst() string {
	return r.Player
}

func (r *ResponsesBotWrapper) GetSecond() string {
	return r.Bot
}

func (r *ResponsesBotWrapper) GetFirstName() string {
	return "Player"
}

func (r *ResponsesBotWrapper) GetSecondName() string {
	return "Bot"
}

// newChoices initializes and returns a pointer to a Choices struct
// containing the predefined options for the game: "rock", "paper", and "scissors".
func newChoices() *model.Choices {
	return &model.Choices{
		Option: []string{
			"rock",
			"paper",
			"scissors",
		},
	}
}

// newWinningCases defining the rules of the game.
func newWinningCases() *model.WinningCases {
	return &model.WinningCases{
		Options: map[string]string{
			"rock":     "scissors",
			"paper":    "rock",
			"scissors": "paper",
		},
	}
}

// newPayload initializes and returns a pointer to a client struct
// with a predefined payload for requesting the latest block information
func newPayload() *client {
	return &client{
		payload: strings.NewReader(`{
        "jsonrpc": "2.0",
        "method": "hmy_getBlockByNumber",
        "params": ["latest", true],
        "id": 1
    }`),
	}
}

// prepareChoice converts input to lowercase for consistent formatting of user choices
func prepareChoice(input string) string {
	return strings.TrimSpace(strings.ToLower(input))
}

// getPlayerChoice processes user input and
func getPlayerChoice(reader *bufio.Reader, validChoices *model.Choices) string {
	for {
		input, _ := reader.ReadString('\n')
		choice := prepareChoice(input)

		if isValidChoice(choice, validChoices) {
			return choice
		}
		fmt.Print(helper.InvalidChoice)
	}
}

// isValidChoice validate input data with allowed options
func isValidChoice(choice string, choices *model.Choices) bool {
	for _, validChoice := range choices.Option {
		if choice == validChoice {
			return true
		}
	}
	return false
}

// determineWinner determines the outcome of a game round based on the choices of player 1 and player 2.
func determineWinner(p1, p2 string, history *model.BlockchainHistory) string {

	if p1 == p2 {
		history.ResultValue = uint8(0)
		return helper.WonTie
	}

	winningCases := newWinningCases()

	if winningCases.Options[p1] == p2 {
		history.ResultValue = uint8(1)
		return helper.WonPlayer1
	} else {
		history.ResultValue = uint8(2)
		return helper.WonPlayer2
	}
}

// getBotChoice selects a random choice from the provided list of choices for the bot
func getBotChoice(choices []string) string {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(choices))))
	if err != nil {
		logrus.Error(err)
	}
	return choices[n.Int64()]
}

// getHarmonyData sends a POST request to the Harmony network and decodes the JSON response.
func getHarmonyData() (map[string]interface{}, error) {
	payload := newPayload()

	req, _ := http.NewRequest("POST", helper.HarmonyRPCUrl, payload.payload)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return result, nil
}

// getHarmonyRandomness retrieves randomness data.
// It calls getHarmonyData to fetch the block information and parses the "hash" field from the result.
// If any error occurs during data retrieval or parsing, it returns the error.
func getHarmonyRandomness() (string, error) {
	result, err := getHarmonyData()
	if err != nil {
		logrus.Warning(err)
	}

	blockData, ok := result["result"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("failed to parse block data")
	}

	// Extract block hash (or use randomness if available)
	blockHash, ok := blockData["hash"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get block hash")
	}

	return blockHash, nil
}

// getBotChoiceHarmonyVRF selects a random choice for the bot using Harmony blockchain's VRF (Verifiable Random Function).
// It retrieves randomness from the Harmony blockchain and uses it to determine a choice index.
// If there's an error in fetching or decoding the randomness, it falls back to the crypto/rand-based getBotChoice function.
func getBotChoiceHarmonyVRF(choices []string) string {
	randomnessStr, err := getHarmonyRandomness()
	if err != nil {
		logrus.Warningf("Failed to get randomness from Harmony VRF. Falling back to crypto/rand. %s", err)
		return getBotChoice(choices)
	}

	randomnessBytes, err := hex.DecodeString(strings.TrimPrefix(randomnessStr, "0x"))
	if err != nil {
		logrus.Warningf("Failed to decode randomness. Falling back to crypto/rand. %s", err)
		return getBotChoice(choices)
	}

	randomnessInt := new(big.Int).SetBytes(randomnessBytes)
	choiceIndex := new(big.Int).Mod(randomnessInt, big.NewInt(int64(len(choices)))).Int64()

	return choices[choiceIndex]
}

// getWinner determines the winner.
// It uses the determineWinner function to evaluate the result and prints the winner.
func getWinner(responses Responses) {

	history := blockchain.InitResults()

	first := responses.GetFirst()
	history.Player1Choice = choiceToUint(first)
	second := responses.GetSecond()
	history.Player1Choice = choiceToUint(second)

	winner := determineWinner(first, second, history)

	fmt.Printf("%s\n%s response: %s\n%s response: %s\n",
		winner,
		responses.GetFirstName(),
		first,
		responses.GetSecondName(),
		second,
	)

	blockchain.Run(history)
}

func choiceToUint(choice string) uint8 {
	switch fmt.Sprintf("%s", choice) { // Assuming choice.String holds the player's choice as a string
	case "rock":
		return 1 // Assuming "rock" corresponds to 1
	case "paper":
		return 2 // Assuming "paper" corresponds to 2
	case "scissors":
		return 3 // Assuming "scissors" corresponds to 3
	}
	return 1
}

// getYourChoice prompts the player".
func getYourChoice() {
	fmt.Print("Player, enter your choice (rock/paper/scissors): ")
}

// clearConsole clears the console screen based on the operating system.
func clearConsole() {
	cmd := exec.Command("clear") // Linux and MacOS
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // Windows
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
