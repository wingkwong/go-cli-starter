package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func MakeCommand2() *cobra.Command {
	var command2 = &cobra.Command{
		Use:          "command2",
		Short:        "Vivamus elementum semper nisi",
		Long:         "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor",
		Example:      "cli command2",
		SilenceUsage: true,
	}

	command2.RunE = func(command *cobra.Command, args []string) error {

		// your logic goes here for command 2

		fmt.Println("your logic goes here for command 2")

		return nil
	}

	command2.AddCommand(MakeSubCommand2())

	return command2
}

const Command2InfoMsg = `Lorem ipsum dolor sit amet, consectetuer adipiscing elit.`
