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

// HandleConsoleInput は標準入力からの入力を待機し、コンソールコマンドの実行を処理します。
func HandleConsoleInput(tokenUseCase usecase.TokenUseCase, canceller context.CancelFunc) {
	commands := initCommandMap(tokenUseCase)

	slog.Info("Type 'help' to show available commands!")

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

		if inputs[0] == "stop" {
			canceller()
			return
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

func initCommandMap(tokenUseCase usecase.TokenUseCase) map[string]command {
	return map[string]command{
		"stop": newCmd("Stop the application", func(string) {
			// do nothing; stop command will be handled in HandleConsoleInput
		}),
		"newtoken":        newCmd("Create a new token", createNewToken(tokenUseCase), "{bytes}"),
		"tokens":          newCmd("Show tokens", getTokens(tokenUseCase)),
		"validatetoken":   newCmd("Validate the token", validateToken(tokenUseCase), "<token>"),
		"invalidatetoken": newCmd("Invalidate the token", invalidateToken(tokenUseCase), "<token>"),
		"cleartokens":     newCmd("Invalidate all tokens", clearTokens(tokenUseCase)),
	}
}

func newCmd(description string, handler func(input string), args ...string) command {
	return command{args, description, handler}
}

func createNewToken(tokenUseCase usecase.TokenUseCase) func(input string) {
	return func(input string) {
		args := strings.Split(input, " ")

		var length int

		if len(args) == 2 {
			if val, err := strconv.Atoi(args[1]); err != nil {
				slog.Error("Invalid number:" + args[1])
				return
			} else if val <= 0 {
				slog.Error("Length must be positive:" + args[1])
			} else {
				length = val
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
			if err = tokenUseCase.InvalidateToken(context.Background(), token.Value); err != nil {
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
