package handler

import (
	"bufio"
	"context"
	"fmt"
	"github.com/Siroshun09/plugin-list/usecase"
	"log/slog"
	"os"
	"sort"
	"strconv"
	"strings"
)

type command struct {
	args        []string
	description string
	handler     func(input string)
}

var noArgs = make([]string, 0)

func args(args ...string) []string {
	return args
}

func HandleConsoleInput(tokenUseCase usecase.TokenUseCase, canceller context.CancelFunc) {
	commands := initCommandMap(tokenUseCase, canceller)

	slog.Info("Type 'help' for show available commands!")

	scanner := bufio.NewScanner(os.Stdin)
	printNewLine()

	for scanner.Scan() {
		raw := scanner.Text()

		if raw == "" {
			printNewLine()
			continue
		}

		inputs := strings.SplitN(scanner.Text(), " ", 2)

		if inputs[0] == "" {
			printNewLine()
			continue
		}

		if inputs[0] == "help" {
			showHelp(commands)
			printNewLine()
			continue
		}

		cmd, exists := commands[inputs[0]]

		if !exists {
			slog.Error("Unknown command: " + inputs[0])
			printNewLine()
			continue
		}

		cmd.handler(raw)
		printNewLine()
	}
}

func printNewLine() {
	fmt.Print("> ")
}

func initCommandMap(tokenUseCase usecase.TokenUseCase, canceller context.CancelFunc) map[string]command {
	return map[string]command{
		"stop": {[]string{}, "Stop the application", func(string) {
			canceller()
		}},
		"newtoken":        {args("{bytes}"), "Create a new token", createNewToken(tokenUseCase)},
		"tokens":          {noArgs, "Show tokens", getTokens(tokenUseCase)},
		"validatetoken":   {args("<token>"), "Validate the token", validateToken(tokenUseCase)},
		"invalidatetoken": {args("<token>"), "Invalidate the token", invalidateToken(tokenUseCase)},
		"cleartokens":     {noArgs, "Invalidate all tokens", clearTokens(tokenUseCase)},
	}
}

func createNewToken(tokenUseCase usecase.TokenUseCase) func(input string) {
	return func(input string) {
		args := strings.Split(input, " ")

		var length int

		if len(args) == 2 {
			var err error
			length, err = strconv.Atoi(args[1])

			if err != nil {
				slog.Error("Invalid number", args[1], err)
				return
			}
		} else {
			length = 16
		}

		token, err := tokenUseCase.CreateNewRandomToken(context.Background(), length)

		if err != nil {
			slog.Error("An error occurred while creating new token", err)
			return
		}

		slog.Info("Created a new token (" + strconv.Itoa(length) + " bytes): " + token.Value)
	}
}

func getTokens(tokenUseCase usecase.TokenUseCase) func(string) {
	return func(string) {
		tokens, err := tokenUseCase.GetAllTokens(context.Background())

		if err != nil {
			slog.Error("An error occurred while creating new token", err)
			return
		}

		if len(tokens) == 0 {
			slog.Info("There are no tokens")
			return
		}

		slog.Info("Current valid tokens:")
		for _, token := range tokens {
			slog.Info(" " + token.Value)
		}
	}
}

func validateToken(tokenUseCase usecase.TokenUseCase) func(input string) {
	return func(input string) {
		args := strings.Split(input, " ")

		if len(args) < 2 {
			slog.Error("No tokens provided")
			return
		}

		for i := 1; i < len(args); i++ {
			token := args[i]

			if len(token) == 0 {
				continue
			}

			valid, err := tokenUseCase.ValidateToken(context.Background(), token)

			if err != nil {
				slog.Error("An error occurred while validating the token", err)
				return
			}

			if valid {
				slog.Info(token + " is valid")
			} else {
				slog.Info(token + " is invalid")
			}
		}
	}
}

func invalidateToken(tokenUseCase usecase.TokenUseCase) func(input string) {
	return func(input string) {
		args := strings.Split(input, " ")

		if len(args) < 2 {
			slog.Error("No tokens provided")
			return
		}

		for i := 1; i < len(args); i++ {
			token := args[i]

			if len(token) == 0 {
				continue
			}

			if err := tokenUseCase.InvalidateToken(context.Background(), token); err != nil {
				slog.Error("An error occurred while invalidating the token", err)
				return
			}

			slog.Info("Invalidated token: " + token)
		}
	}
}

func clearTokens(tokenUseCase usecase.TokenUseCase) func(string) {
	return func(string) {
		tokens, err := tokenUseCase.GetAllTokens(context.Background())

		if err != nil {
			slog.Error("An error occurred while collecting tokens", err)
			return
		}

		if len(tokens) == 0 {
			slog.Info("There are no tokens")
			return
		}

		for _, token := range tokens {
			if err := tokenUseCase.InvalidateToken(context.Background(), token.Value); err != nil {
				slog.Error("An error occurred while invalidating the token", err)
				return
			}

			slog.Info("Invalidated token: " + token.Value)
		}

		slog.Info(strconv.Itoa(len(tokens)) + " tokens are invalidated.")
	}
}

func showHelp(commands map[string]command) {
	slog.Info("Available commands:")
	labels := make([]string, len(commands))

	i := 0
	for label := range commands {
		labels[i] = label
		i++
	}

	sort.Strings(labels)

	for _, label := range labels {
		cmd := commands[label]
		var args string
		if len(cmd.args) == 0 {
			args = ""
		} else {
			args = " " + strings.Join(cmd.args, " ")
		}
		slog.Info(" " + label + args + " - " + cmd.description)
	}
}
