package main

import (
	"github.com/bugbounce/totp-cli/command"
	"github.com/yitsushi/go-commander"
)

func registerCommands(registry *commander.CommandRegistry) {
	// Register available commands
	registry.Register(command.NewAddToken)
	registry.Register(command.NewChangePassword)
	registry.Register(command.NewDelete)
	registry.Register(command.NewGenerate)
	registry.Register(command.NewInstant)
	registry.Register(command.NewList)
	registry.Register(command.NewUpdate)
	registry.Register(command.NewVersion)
	registry.Register(command.NewSearchList)
}

func main() {
	registry := commander.NewCommandRegistry()

	registerCommands(registry)

	registry.Execute()
}
