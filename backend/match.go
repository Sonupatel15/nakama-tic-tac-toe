package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

type MatchState struct {
	Board     []int                       // Length 9: [0,0,0, 0,0,0, 0,0,0]
	Marks     map[string]int              // Map userID to Mark (1 or 2)
	Turn      int                         // Which mark's turn it is (1 or 2)
	Presences map[string]runtime.Presence // Track connected players
}

type Match struct{}

// MatchInit initializes the game state
func (m *Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	state := &MatchState{
		Board:     make([]int, 9),
		Marks:     make(map[string]int),
		Presences: make(map[string]runtime.Presence),
		Turn:      1, // Player 1 (X) starts
	}

	tickRate := 1 // How many times MatchLoop runs per second
	label := "tic-tac-toe-match"

	return state, tickRate, label
}

// We will implement MatchJoin, MatchLoop, etc., in the next step.
