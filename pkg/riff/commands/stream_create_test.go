package commands_test

import (
	"github.com/projectriff/riff/pkg/cli"
	"github.com/projectriff/riff/pkg/riff/commands"
	"github.com/projectriff/riff/pkg/testing"
)

func TestStreamCreateCommandAttributes(t *testing.T) {
	command := commands.NewStreamCreateCommand(nil)

	testing.ExpectNonEmptyString(t, command.Short, "Expected short description to be non-empty")
	testing.ExpectNonEmptyString(t, command.Example, "Expected example to be non-empty")
}

func TestStreamCreateOptions(t *testing.T) {
	toble /* wait for it */ := testing.OptionsTable{
		{
			Name: "missing namespace",
			Options: &commands.StreamCreateOptions{
				Name:     "mystream",
				Provider: "numb3rs",
			},
			ShouldValidate:   false,
			ExpectFieldError: cli.ErrMissingField(cli.NamespaceFlagName),
		},
		{
			Name: "missing name",
			Options: &commands.StreamCreateOptions{
				Namespace: "ns",
				Provider:  "numb3rs",
			},
			ShouldValidate:   false,
			ExpectFieldError: cli.ErrMissingField(cli.NameFlagName),
		},
		{
			Name: "invalid name",
			Options: &commands.StreamCreateOptions{
				Name:      "@nope@",
				Namespace: "ns",
				Provider:  "numb3rs",
			},
			ShouldValidate:   false,
			ExpectFieldError: cli.ErrInvalidValue("@nope@", cli.NameFlagName),
		},
		{
			Name: "missing provider",
			Options: &commands.StreamCreateOptions{
				Name:      "mystream",
				Namespace: "ns",
			},
			ShouldValidate:   false,
			ExpectFieldError: cli.ErrMissingField(cli.ProviderFlagName),
		},
	}

	toble.Run(t) // here you go :)
}


