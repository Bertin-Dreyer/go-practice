package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Exercise 14.3:
// Define Level type (underlying string) with Debug and Info constants.
// Create functions to store and extract log level from context.
// Create middleware reading log_level query parameter (valid: debug, info).
// Fill in TODO in Log to extract level from context.

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
)

// TODO: Write ctxKey type for context keys
// TODO: Write SetLogLevel(ctx, Level) context.Context
// TODO: Write GetLogLevel(ctx) Level (returns Info if not set or invalid)
// TODO: Write middleware reading "log_level" query param, validating, setting context
// TODO: Implement Log function: extract level from context, print based on level rules

func Log(ctx context.Context, level Level, message string) {
	var inLevel Level
	// TODO: Get log level from context into inLevel
	// If not set or invalid, print nothing

	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}
	if level == Info && (inLevel == Debug || inLevel == Info) {
		fmt.Println(message)
	}
}

func main() {
	// TODO: Set up server with middleware, test Log with different levels
	fmt.Println("TODO: implement ex3-log-level-context")
	_ = time.Second
	_ = http.DefaultServeMux
}
