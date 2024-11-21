# Assumptions

---
## Annotation:

- **[MAIN](../README.md)**: The main documentation file providing an overview of the project.
- **[DELIVERED_FUNCTIONALITY](./DELIVERED_FUNCTIONALITY.md)**: Details the core functionalities delivered by the application.
- **[HOW_TO_RUN](./HOW_TO_RUN.md)**: Step-by-step instructions for setting up and running the application.
- **[CONTRACT_LOGIC](./CONTRACT_LOGIC.md)**: Explains the Solidity smart contract logic deployed on the blockchain.
- **[TESTS](./TESTS.md)**: Provides information about the testing strategy, test coverage, and commands for running tests.
- **[ASSUMPTIONS](./ASSUMPTIONS.md)**: Outlines the assumptions made during the development of the application.
---


The development of this application is based on the following assumptions:

## 1. Blockchain Environment

- The **Harmony Testnet** (https://api.s0.b.hmny.io) is available and operational for deploying and interacting with the smart contract.
- Network latency and transaction costs on the Testnet are minimal and suitable for the use case.

## 2. User Interaction

- Players must have the required setup to interact with the CLI application, such as:
  - **Go** installed for running the application binary.
- Users playing the game must have access to a wallet address for transactions on the Harmony blockchain.

## 3. VRF Randomness

- Harmony's **VRF (Verifiable Random Function)** service is operational and accessible for generating randomness in the **Player vs Bot** mode.
- If VRF is unavailable, the fallback randomness method provides sufficient entropy for bot moves.

## 4. Smart Contract

- The deployed smart contract assumes:
  - Both players' moves are provided correctly and securely via the CLI.
  - The game results passed to the contract are already validated in the Go application.
- The smart contract does not include logic for determining the winner; this is handled by the Go application before submitting results.

## 5. Testing and Deployment

- The **Harmony Testnet** is used for development, testing, and demonstration purposes.
- For production, the application would require deployment to the Harmony Mainnet or another suitable network.
- Deployment assumes the availability of required tools, such as **MetaMask** and **Remix** for deploying the smart contract.

## 6. Game Rules

- The rules of Rock-Paper-Scissors are universally understood and implemented as follows:
  - **Rock beats Scissors.**
  - **Scissors beats Paper.**
  - **Paper beats Rock.**
- No additional game rules or variations are included.

## 7. Error Handling

- The application assumes:
  - Most errors, such as invalid user inputs or transaction failures, are handled gracefully.
  - Network outages or blockchain service interruptions may require manual intervention or retries.

## 8. CLI Usability

- The CLI application is designed for developers or users with basic command-line experience.
- Advanced graphical interfaces or non-CLI interaction methods are out of scope for this implementation.

