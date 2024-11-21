# How to Run the Application

---
## Annotation:

- **[MAIN](../README.md)**: The main documentation file providing an overview of the project.
- **[DELIVERED_FUNCTIONALITY](./DELIVERED_FUNCTIONALITY.md)**: Details the core functionalities delivered by the application.
- **[HOW_TO_RUN](./HOW_TO_RUN.md)**: Step-by-step instructions for setting up and running the application.
- **[CONTRACT_LOGIC](./CONTRACT_LOGIC.md)**: Explains the Solidity smart contract logic deployed on the blockchain.
- **[TESTS](./TESTS.md)**: Provides information about the testing strategy, test coverage, and commands for running tests.
- **[ASSUMPTIONS](./ASSUMPTIONS.md)**: Outlines the assumptions made during the development of the application.

---

## Running the Application

The application is built and managed using a **Makefile** located in the root folder of the repository. This Makefile provides commands to test, build, and analyze the application. Below are the steps to run the application:

1. **Clone the Repository**:
    ```bash
    git clone <repository_url>
    cd <repository_folder>
    ```

2. **Build the Application**:  
   Use the `make build-app` command to build the application binary. The binary will be created in the `tmp` folder:
    ```bash
    make build-app
    ```

   - **Environment Variables**:
     - `GOOS`: Target operating system (default: `darwin`).
     - `GOARCH`: Target architecture (default: `arm64`).

   - **Example for Linux**:
     ```bash
     make build-app GOOS=linux GOARCH=amd64
     ```

3. **Run the Application**:  
   Once built, execute the binary:
    ```bash
    ./tmp/rock-paper-scissors
    ```
   It follows different game options like pvp or pve. then run application like this.
     ```bash
    ./tmp/rock-paper-scissors pve
    ```

## Testing the Application

1. **Run All Unit Tests**:
    ```bash
    make tests
    ```

2. **Generate a Code Coverage Report**:  
   Use the `make cover` target:
    ```bash
    make cover
    ```
   This will:
   - Execute the tests with race detection and generate a coverage profile.
   - Open an HTML coverage report.
   - Clean up the coverage profile file after use.

## Makefile Commands Overview

| **Command**      | **Description**                                                                                       |
|-------------------|-------------------------------------------------------------------------------------------------------|
| `make tests`     | Runs all unit tests in the project and outputs results.                                               |
| `make cover`     | Runs tests with coverage, opens an HTML coverage report, and cleans up.                               |
| `make build-app` | Builds the CLI application binary with the specified `GOOS` and `GOARCH`.                             |

The provided **Makefile** simplifies the build and test process, ensuring a consistent environment for running the application.

