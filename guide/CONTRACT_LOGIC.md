# General Logic of the Contract

---
## Annotation:

- **[MAIN](../README.md)**: The main documentation file providing an overview of the project.
- **[DELIVERED_FUNCTIONALITY](./DELIVERED_FUNCTIONALITY.md)**: Details the core functionalities delivered by the application.
- **[HOW_TO_RUN](./HOW_TO_RUN.md)**: Step-by-step instructions for setting up and running the application.
- **[CONTRACT_LOGIC](./CONTRACT_LOGIC.md)**: Explains the Solidity smart contract logic deployed on the blockchain.
- **[TESTS](./TESTS.md)**: Provides information about the testing strategy, test coverage, and commands for running tests.
- **[ASSUMPTIONS](./ASSUMPTIONS.md)**: Outlines the assumptions made during the development of the application.
---

The **RockPaperScissors** smart contract is designed to record the results of Rock-Paper-Scissors games on the blockchain. It ensures transparency, immutability, and accessibility of game outcomes. Below is an overview of its components and functionality:

## Components and Functionality

### 1. Enumerations
- **Choice**: Represents the possible moves in the game — Rock, Paper, or Scissors.
- **Result**: Represents the outcome of the game — Draw, Player1Wins, or Player2Wins.

### 2. Data Structures
- **GameResult**: A structure storing the details of a single game:
  - Players involved (`player1` and `player2`).
  - Choices made by both players (`player1Choice` and `player2Choice`).
  - Result of the game (`result`).

### 3. Storage
- **gameResults**: A dynamic array storing all game results.

### 4. Functions
- **addGameResult**:
  - **Purpose**: Adds a new game result to the blockchain.
  - **Parameters**: Player addresses, choices, and the result.
  - **Functionality**: Appends a `GameResult` structure to the `gameResults` array.

- **getGameResult**:
  - **Purpose**: Retrieves a specific game result by index.
  - **Functionality**: Returns all details of the selected game.

- **getGameCount**:
  - **Purpose**: Returns the total number of games recorded.

---

# How to Deploy the Contract

The contract can be deployed to the **Harmony Testnet** using Remix, a browser-based Ethereum IDE. Follow the steps below, referring to the [Harmony Deployment Documentation](https://docs.harmony.one/home/developers/deploying-on-harmony/using-remix/ethereum-remix):

1. Open Remix in your browser.
2. Paste the **RockPaperScissors** smart contract code into a new file.
3. Connect Remix to the Harmony Testnet via the provided instructions in the documentation.
4. Compile the smart contract.
5. Deploy the contract to the Harmony Testnet using the connected wallet.

