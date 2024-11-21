# Tests

---
## Annotation:

- **[MAIN](../README.md)**: The main documentation file providing an overview of the project.
- **[DELIVERED_FUNCTIONALITY](./DELIVERED_FUNCTIONALITY.md)**: Details the core functionalities delivered by the application.
- **[HOW_TO_RUN](./HOW_TO_RUN.md)**: Step-by-step instructions for setting up and running the application.
- **[CONTRACT_LOGIC](./CONTRACT_LOGIC.md)**: Explains the Solidity smart contract logic deployed on the blockchain.
- **[TESTS](./TESTS.md)**: Provides information about the testing strategy, test coverage, and commands for running tests.
- **[ASSUMPTIONS](./ASSUMPTIONS.md)**: Outlines the assumptions made during the development of the application.
---

The application includes comprehensive tests to ensure reliability and correctness. These tests cover:

- Core functionality of the application
- CLI logic
- Smart contract interactions
- Edge cases

## Running Tests

To execute all tests, use the following command:

```bash
go test ./... -v
```

### Running Tests: Options and Commands

#### Test Command Breakdown
- **`./...`**: Recursively runs tests in all packages of the project.
- **`-v`**: Enables verbose output to display detailed test results.

---

## Makefile Target

To simplify the testing process, the repository includes a `make` target:

```bash
make tests
```

This command executes all tests in the project and reports any failures.

## Test Coverage

Tests are designed to:

-   Verify game logic, such as valid and invalid moves.
-   Ensure proper interactions with the Harmony blockchain.
-   Handle edge cases, such as invalid inputs or transaction failures.

By running the tests regularly, you can ensure that the application remains robust and meets its expected functionality.
