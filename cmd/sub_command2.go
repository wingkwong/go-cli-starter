package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func MakeSubCommand2() *cobra.Command {
	var sub_command2 = &cobra.Command{
		Use:          "sub_command2",
		Short:        "Vivamus elementum semper nisi",
		Long:         "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor",
		Example:      "cli sub_command2",
		SilenceUsage: true,
	}

	sub_command2.RunE = func(command *cobra.Command, args []string) error {

		// your logic goes here for sub-command 2

		fmt.Println("your logic goes here for sub-command 2")

		return nil
	}

	return sub_command2
}
