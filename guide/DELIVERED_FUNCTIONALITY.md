Here's the reorganized and formatted markdown:

---
## Annotation:

- **[MAIN](../README.md)**: The main documentation file providing an overview of the project.
- **[DELIVERED_FUNCTIONALITY](./DELIVERED_FUNCTIONALITY.md)**: Details the core functionalities delivered by the application.
- **[HOW_TO_RUN](./HOW_TO_RUN.md)**: Step-by-step instructions for setting up and running the application.
- **[CONTRACT_LOGIC](./CONTRACT_LOGIC.md)**: Explains the Solidity smart contract logic deployed on the blockchain.
- **[TESTS](./TESTS.md)**: Provides information about the testing strategy, test coverage, and commands for running tests.
- **[ASSUMPTIONS](./ASSUMPTIONS.md)**: Outlines the assumptions made during the development of the application.

---

# Delivered Functionality

The application delivers the following core functionalities, meeting the specified requirements:

## 1. Application Implementation

- The CLI application is developed in **Go**, ensuring high performance and simplicity for user interactions.
- It provides two game modes:
  - **Player vs Player**: Two players can compete in a Rock-Paper-Scissors game, with moves validated and results stored on-chain.
  - **Player vs Bot**: A player can challenge a bot that utilizes randomness to make its moves.

## 2. VRF Randomness Integration

- The application implements **Harmony VRF (Verifiable Random Function)** logic to ensure unbiased and unpredictable bot moves.
- A fallback randomness method is included for scenarios where VRF is unavailable.

## 3. Smart Contract Deployment

- A **Solidity smart contract** is deployed to the **Harmony Testnet**, enabling secure and transparent storage of game results.
- The smart contract handles:
  - Recording game results (players, moves, outcomes, and timestamps).
  - Retrieving stored results and metadata.

## 4. Integration Between Application and Blockchain

- The Go application seamlessly interacts with the deployed smart contract on the Harmony blockchain.
- Features include:
  - Submitting game results to the blockchain.
  - Retrieving game results from the blockchain for auditability and transparency.

---

This application delivers a robust and blockchain-integrated gaming experience, demonstrating the power of decentralized technologies in practical use cases.

