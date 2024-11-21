package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/witnsby/web3_offline_coding_test/src/cmd/pkg/helper"
	"github.com/witnsby/web3_offline_coding_test/src/cmd/pkg/model"
	"log"
	"math/big"
	"strings"
)

const (
	privateKeyHex = "f900ea7717dd77d38fecce3a8a48299bd4206f96e509942452a90074bd977746" // Replace with your private key (ensure this is secure)
	//abiJSON         = `[
	//	{"constant":true,"inputs":[],"name":"getGameCount","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},
	//	{"constant":false,"inputs":[{"name":"_player1","type":"address"},{"name":"_player2","type":"address"},{"name":"_player1Choice","type":"uint8"},{"name":"_player2Choice","type":"uint8"},{"name":"_result","type":"uint8"}],"name":"addGameResult","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}
	//]`
	abiJSON = `[
{"inputs":[{"internalType":"address","name":"_player1","type":"address"},{"internalType":"address","name":"_player2","type":"address"},{"internalType":"enum RockPaperScissors.Choice","name":"_player1Choice","type":"uint8"},{"internalType":"enum RockPaperScissors.Choice","name":"_player2Choice","type":"uint8"},{"internalType":"enum RockPaperScissors.Result","name":"_result","type":"uint8"}],"name":"addGameResult","outputs":[],"stateMutability":"nonpayable","type":"function"},
{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"gameResults","outputs":[{"internalType":"address","name":"player1","type":"address"},{"internalType":"address","name":"player2","type":"address"},{"internalType":"enum RockPaperScissors.Choice","name":"player1Choice","type":"uint8"},{"internalType":"enum RockPaperScissors.Choice","name":"player2Choice","type":"uint8"},{"internalType":"enum RockPaperScissors.Result","name":"result","type":"uint8"},{"internalType":"uint256","name":"timestamp","type":"uint256"}],"stateMutability":"view","type":"function"},
{"inputs":[],"name":"getGameCount","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"index","type":"uint256"}],"name":"getGameResult","outputs":[{"internalType":"address","name":"player1","type":"address"},{"internalType":"address","name":"player2","type":"address"},{"internalType":"enum RockPaperScissors.Choice","name":"player1Choice","type":"uint8"},{"internalType":"enum RockPaperScissors.Choice","name":"player2Choice","type":"uint8"},{"internalType":"enum RockPaperScissors.Result","name":"result","type":"uint8"},{"internalType":"uint256","name":"timestamp","type":"uint256"}],"stateMutability":"view","type":"function"}
]`
)

type Harmony struct {
	client *ethclient.Client
	abi    abi.ABI
}

type GameResult struct {
	Player1       common.Address `abi:"player1"`
	Player2       common.Address `abi:"player2"`
	Player1Choice uint8          `abi:"player1Choice"`
	Player2Choice uint8          `abi:"player2Choice"`
	Result        uint8          `abi:"result"`
	Timestamp     *big.Int       `abi:"timestamp"`
}

// initClientAndABI initializes the Harmony client and parses the contract ABI.
// It returns a pointer to a Harmony struct containing the initialized client and ABI.
func initClientAndABI() *Harmony {
	client, err := ethclient.Dial(helper.HarmonyRPCUrl)
	if err != nil {
		log.Fatalf("Failed to connect to Harmony RPC: %v", err)
	}

	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	return &Harmony{
		client: client,
		abi:    parsedABI,
	}
}

// InitResults function prepares the BlockchainHistory struct with the necessary addresses
// required for interacting with the smart contract and for recording game results.
func InitResults() *model.BlockchainHistory {
	return &model.BlockchainHistory{
		SmartContractAddress: common.HexToAddress(helper.ContractAddress),
		Player1:              common.HexToAddress(helper.Player1Hash),
		Player2:              common.HexToAddress(helper.Player2Hash),
	}
}

func Run() {
	clientHarmony := initClientAndABI()
	clientHarmony.verifyContractDeployment()

	clientHarmony.addGameResult()
}

// addGameResult adds a new game result to the smart contract by creating, signing, and sending a transaction.
// It initializes the game result data, packs the function call, signs the transaction with the private key, and submits it to the network.
func (h *Harmony) addGameResult() {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		logrus.Fatalf("Failed to parse private key: %v", err)
	}

	nonce, err := h.client.PendingNonceAt(context.Background(), getFromAddress(privateKey))
	if err != nil {
		logrus.Fatalf("Failed to get nonce: %v", err)
	}

	gasPrice, err := h.client.SuggestGasPrice(context.Background())
	if err != nil {
		logrus.Fatalf("Failed to suggest gas price: %v", err)
	}

	result := InitResults()
	result.Player1Choice = uint8(2)
	result.Player2Choice = uint8(0)
	result.ResultValue = uint8(1)

	addGameData, err := h.abi.Pack("addGameResult", result.Player1, result.Player2, result.Player1Choice, result.Player2Choice, result.ResultValue)
	if err != nil {
		logrus.Errorf("Failed to encode addGameResult function call: %v", err)
	}

	tx := types.NewTransaction(nonce, result.SmartContractAddress, big.NewInt(0), 300000, gasPrice, addGameData)
	chainID, err := h.client.NetworkID(context.Background())
	if err != nil {
		logrus.Errorf("Failed to get network ID: %v", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		logrus.Errorf("Failed to sign transaction: %v", err)
	}

	//send transaction
	err = h.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		logrus.Errorf("Failed to send transaction: %v", err)
	}

	fmt.Printf("Successfully added a new game result. Transaction hash: %s\n", signedTx.Hash().Hex())
}

// getFromAddress derives the Ethereum address from the given ECDSA private key.
// It retrieves the corresponding public key, casts it to ECDSA format, and computes the address.
func getFromAddress(privateKey *ecdsa.PrivateKey) common.Address {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Failed to cast public key to ECDSA")
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA)
}

// verifyContractDeployment checks if the smart contract is deployed at the specified address.
// It retrieves the bytecode from the blockchain for the given contract address and logs an error
// if no code is found, indicating the contract may not be deployed.
func (h *Harmony) verifyContractDeployment() {
	smartContractAddress := common.HexToAddress(helper.ContractAddress)
	code, err := h.client.CodeAt(context.Background(), smartContractAddress, nil)
	if err != nil {
		logrus.Fatalf("Failed to get code at address: %v", err)
	}
	if len(code) == 0 {
		logrus.Fatalf("No contract code found at address %s. The contract may not be deployed.", helper.ContractAddress)
	} else {
		logrus.Info("Contract code found at the specified address.")
	}
}

// checkAccountBalance retrieves and displays the Ether balance of the given Ethereum address.
// It queries the blockchain for the account's balance and prints it to the console.
// Logs a fatal error if the balance retrieval fails.
func (h *Harmony) checkAccountBalance(address common.Address) {
	balance, err := h.client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Failed to get account balance: %v", err)
	}
	fmt.Printf("Account Balance for %s: %s\n", address.Hex(), balance.String())
}

// getGameCount should get a cont for all games saved on blockchain
func (h *Harmony) getGameCount() {
	///TBD
}

// getGameResult should return result from blockchain for games
func (h *Harmony) getGameResult() {
	//TBD
}
