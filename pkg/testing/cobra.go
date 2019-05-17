package testing

import "github.com/spf13/cobra"

func FindSubcommand(command *cobra.Command, names ...string) *cobra.Command {
	cmd, unmatchedArgs, err := command.Find(names)
	if err != nil || len(unmatchedArgs) > 0 {
		return nil
	}
	return cmd
}

