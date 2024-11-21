package model

import "github.com/ethereum/go-ethereum/common"

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

type BlockchainHistory struct {
	SmartContractAddress common.Address
	Player1              common.Address
	Player2              common.Address
	Player1Choice        uint8
	Player2Choice        uint8
	ResultValue          uint8
}

type EnvironmentVariables struct {
	PrivateKeyHex string
}
