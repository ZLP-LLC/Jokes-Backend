package bootstrap

import (
	"github.com/spf13/cobra"

	"jokes/commands"
)

var rootCmd = &cobra.Command{
	Use:              "jokes",
	Short:            "Jokes REST API server",
	TraverseChildren: true,
}

// App root of application
type App struct {
	*cobra.Command
}

func NewApp() App {
	cmd := App{
		Command: rootCmd,
	}
	cmd.AddCommand(commands.GetSubCommands(CommonModules)...)
	return cmd
}

var RootApp = NewApp()
