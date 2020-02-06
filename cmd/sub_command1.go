package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func MakeSubCommand1() *cobra.Command {
	var subcommand1 = &cobra.Command{
		Use:          "sub_command1",
		Short:        "Vivamus elementum semper nisi",
		Long:         "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor",
		Example:      "cli sub_command1",
		SilenceUsage: true,
	}

	subcommand1.RunE = func(command *cobra.Command, args []string) error {

		// your logic goes here for sub-command 1

		fmt.Println("your logic goes here for sub-command 1")

		return nil
	}

	return subcommand1
}
