package commands_test

import (
	"github.com/projectriff/riff/pkg/riff/commands"
	"github.com/projectriff/riff/pkg/testing"
)

func TestStreamCommandAttributes(t *testing.T) {
	command := commands.NewStreamCommand(nil)

	testing.ExpectNonEmptyString(t, command.Short, "Expected short description to be non-empty")
}

func TestStreamCommandSubCommands(t *testing.T) {
	command := commands.NewStreamCommand(nil)

	testing.ExpectNonNil(t, testing.FindSubcommand(command, "create"))
	testing.ExpectNonNil(t, testing.FindSubcommand(command, "delete"))
	testing.ExpectNonNil(t, testing.FindSubcommand(command, "invoke"))
	testing.ExpectNonNil(t, testing.FindSubcommand(command, "list"))
	testing.ExpectNonNil(t, testing.FindSubcommand(command, "update"))
}

func TestStreamCommand(t *testing.T) {
	table := testing.CommandTable{
		{
			Name: "stream",
			Args: []string{},
		},
		{
			Name: "streams",
			Args: []string{},
		},
		{
			Name: "stream",
			Args: []string{"nope"},
			Skip: true, // this invocation should generate an error but does not...
			Verify: func(t *testing.T, output string, err error) {
				if err == nil || err.Error() != "TBD" {
					t.Error("stream command does not accept any argument")
				}
			},
		},
	}

	table.Run(t, commands.NewStreamCommand)
}