# Rock-Paper-Scissors

---
## Annotation:

- **[MAIN](./README.md)**: The main documentation file providing an overview of the project.
- **[DELIVERED_FUNCTIONALITY](./guide/DELIVERED_FUNCTIONALITY.md)**: Details the core functionalities delivered by the application.
- **[HOW_TO_RUN](./guide/HOW_TO_RUN.md)**: Step-by-step instructions for setting up and running the application.
- **[CONTRACT_LOGIC](./guide/CONTRACT_LOGIC.md)**: Explains the Solidity smart contract logic deployed on the blockchain.
- **[TESTS](./guide/TESTS.md)**: Provides information about the testing strategy, test coverage, and commands for running tests.
- **[ASSUMPTIONS](./guide/ASSUMPTIONS.md)**: Outlines the assumptions made during the development of the application.
---


## Application Description

The application is an on-chain Rock-Paper-Scissors CLI game, designed to demonstrate blockchain integration with a simple and interactive command-line interface (CLI). It allows two players to compete in a traditional Rock-Paper-Scissors game, with game results stored securely on-chain using a smart contract deployed on the Harmony Testnet ([https://api.s0.b.hmny.io](https://api.s0.b.hmny.io)).

Also, the application includes a "user vs bot" mode, where players can challenge an AI opponent. The bot leverages Harmony VRF randomness to generate its moves, ensuring fairness and unpredictability. If VRF is unavailable, the application defaults to an alternative randomness method.

## Key Features delivered:

### Two-Player Game Mode

- Two players can compete in a Rock-Paper-Scissors match.
- Moves are validated to prevent invalid inputs.
- Results are stored on the Harmony Testnet using a Solidity smart contract.

### User vs Bot Mode

- Players can challenge a bot opponent.
- The bot's move generation leverages `Harmony VRF randomness`, ensuring unbiased gameplay.
- A fallback randomness method is implemented in case VRF is unavailable.

### Blockchain Integration

- A Solidity smart contract interacts with the CLI application to store and retrieve game outcomes.
- The contract tracks game results and provides an auditable record of gameplay.

### Error Handling

- Invalid inputs (e.g., non-existent moves) are detected and rejected.
- Blockchain transaction failures are captured, logged, and presented to the user with appropriate feedback.

## Technology Stack

### Programming Languages

- **Go**: For the CLI application.
- **Solidity**: For the smart contract.

### Blockchain

- **Harmony Testnet** ([https://api.s0.b.hmny.io](https://api.s0.b.hmny.io)) for deploying and interacting with smart contracts.

---

This application demonstrates the seamless integration of blockchain technology into a CLI game, showcasing the potential for secure and transparent storage of user interactions.


# Additional Information

| Item                                      | Details                                                                                                                                                       |
|-------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **Contract Address**                      | `0x67925fac2485a8c71eff765b199ed91eeb3596da`                                                                                                                  |
| **Harmony Faucet**                        | [https://faucet.pops.one/](https://faucet.pops.one/)                                                                                                          |
| **Harmony Explorer Testnet**              | [https://explorer.testnet.harmony.one/](https://explorer.testnet.harmony.one/)                                                                                 |
| **Contract on Explorer**                  | [View Contract](https://explorer.testnet.harmony.one/address/one1v7f9ltpysk5vw8hlwed3n8kerm4nt9k6z86hh7)                                                       |
| **Environment Variable for Private Key**  | The private key should be added to the environment variable as `PRIVATEKEYHEX`.                                                                                |
| **Blockchain Transaction Result Example** | Successfully added a new game result. Transaction hash: `0x6429c1d2099386c20e9467badb3d2d86f6ff06033e9cb1380f033c6afdd12930`                                   |
