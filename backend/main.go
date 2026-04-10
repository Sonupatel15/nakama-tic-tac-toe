package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

// InitModule is the entry point for the Nakama Go runtime.
func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	logger.Info("Go module loaded successfully!")

	// This is where we will register our Tic-Tac-Toe match logic later
	// Example:
	// if err := initializer.RegisterMatch("tic-tac-toe", func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error) {
	// 	return &Match{}, nil
	// }); err != nil {
	// 	return err
	// }

	return nil
}
