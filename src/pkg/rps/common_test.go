package rps

import (
	"github.com/stretchr/testify/assert"
	"github.com/witnsby/web3_offline_coding_test/src/pkg/model"
	"testing"
)

func TestNewWinningCases(t *testing.T) {
	wc := newWinningCases()
	assert.NotNil(t, wc, "WinningCases should not be nil")
	assert.NotNil(t, wc.Options, "Options map should not be nil")

	type testCase struct {
		key         string
		expectedVal string
		shouldExist bool
		message     string
	}

	testCases := []testCase{
		{
			key:         "rock",
			expectedVal: "scissors",
			shouldExist: true,
			message:     "Rock should beat scissors",
		},
		{
			key:         "paper",
			expectedVal: "rock",
			shouldExist: true,
			message:     "Paper should beat rock",
		},
		{
			key:         "scissors",
			expectedVal: "paper",
			shouldExist: true,
			message:     "Scissors should beat paper",
		},
		{
			key:         "lizard",
			expectedVal: "",
			shouldExist: false,
			message:     "Lizard should not exist in options",
		},
		{
			key:         "",
			expectedVal: "",
			shouldExist: false,
			message:     "Empty key should not exist",
		},
		{
			key:         "fire",
			expectedVal: "",
			shouldExist: false,
			message:     "Fire should not exist in options",
		},
	}

	for _, tc := range testCases {
		actualValue, exists := wc.Options[tc.key]
		assert.Equal(t, tc.shouldExist, exists, tc.message)
		if exists {
			assert.Equal(t, tc.expectedVal, actualValue, "Expected value for key '%s' should be '%s'", tc.key, tc.expectedVal)
		}
	}
}

func TestIsValidChoice(t *testing.T) {
	// Define the structure for test cases
	type testCase struct {
		name        string
		choice      string
		choices     *model.Choices
		expected    bool
		description string
	}

	// Initialize test cases
	testCases := []testCase{
		{
			name:        "Valid choice - rock",
			choice:      "rock",
			choices:     &model.Choices{Option: []string{"rock", "paper", "scissors"}},
			expected:    true,
			description: "Rock is a valid choice",
		},
		{
			name:        "Valid choice - paper",
			choice:      "paper",
			choices:     &model.Choices{Option: []string{"rock", "paper", "scissors"}},
			expected:    true,
			description: "Paper is a valid choice",
		},
		{
			name:        "Empty choice",
			choice:      "",
			choices:     &model.Choices{Option: []string{"rock", "paper", "scissors"}},
			expected:    false,
			description: "Empty string is not a valid choice",
		},
		{
			name:        "Empty choices",
			choice:      "rock",
			choices:     &model.Choices{Option: []string{}},
			expected:    false,
			description: "Choices option is empty",
		},
		{
			name:        "Duplicate choices - valid",
			choice:      "rock",
			choices:     &model.Choices{Option: []string{"rock", "rock", "paper"}},
			expected:    true,
			description: "Rock is present even with duplicates",
		},
		{
			name:        "Case sensitivity - exact match",
			choice:      "Rock",
			choices:     &model.Choices{Option: []string{"rock", "paper", "scissors"}},
			expected:    false,
			description: "Case-sensitive match: 'Rock' vs 'rock'",
		},
		{
			name:        "Case sensitivity - case-insensitive match",
			choice:      "Rock",
			choices:     &model.Choices{Option: []string{"Rock", "paper", "scissors"}},
			expected:    true, // Adjust based on your function's behavior
			description: "Case-insensitive match: 'Rock' vs 'Rock'",
		},
	}

	// Iterate over each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isValidChoice(tc.choice, tc.choices)
			assert.Equal(t, tc.expected, result, tc.description)
		})
	}
}
