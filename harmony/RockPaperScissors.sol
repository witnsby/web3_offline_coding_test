// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract RockPaperScissors {
    enum Choice { Rock, Paper, Scissors }
    enum Result { Draw, Player1Wins, Player2Wins }

    struct GameResult {
        address player1;
        address player2;
        Choice player1Choice;
        Choice player2Choice;
        Result result;
        uint256 timestamp;
    }

    GameResult[] public gameResults;

    function addGameResult(
        address _player1,
        address _player2,
        Choice _player1Choice,
        Choice _player2Choice,
        Result _result
    ) public {
        GameResult memory newResult = GameResult({
            player1: _player1,
            player2: _player2,
            player1Choice: _player1Choice,
            player2Choice: _player2Choice,
            result: _result,
            timestamp: block.timestamp
        });
        gameResults.push(newResult);
    }

    function getGameResult(uint256 index) public view returns (
        address player1,
        address player2,
        Choice player1Choice,
        Choice player2Choice,
        Result result,
        uint256 timestamp
    ) {
        GameResult storage game = gameResults[index];
        return (
            game.player1,
            game.player2,
            game.player1Choice,
            game.player2Choice,
            game.result,
            game.timestamp
        );
    }

    function getGameCount() public view returns (uint256) {
        return gameResults.length;
    }
}
