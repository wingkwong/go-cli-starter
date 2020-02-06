package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func MakeCommand1() *cobra.Command {
	var command1 = &cobra.Command{
		Use:          "command1",
		Short:        "Vivamus elementum semper nisi",
		Long:         "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor",
		Example:      "cli command1",
		SilenceUsage: true,
	}

	command1.RunE = func(command *cobra.Command, args []string) error {

		// your logic goes here for command 1

		fmt.Println("your logic goes here for command 1")

		return nil
	}

	command1.AddCommand(MakeSubCommand1())

	return command1
}

const Command1InfoMsg = `2Lorem ipsum dolor sit amet, consectetuer adipiscing elit.`
