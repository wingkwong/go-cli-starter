package main

import (
	"os"

	"github.com/go-cli-starter/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "Lorem ipsum dolor sit amet",
		Long: `Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, 
		sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, l
		orem. Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(cmd.MakeCommand1())
	rootCmd.AddCommand(cmd.MakeCommand2())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
