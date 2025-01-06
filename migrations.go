package migrations

import (
	"github.com/spf13/cobra"
)

var MigrationsCmd = &cobra.Command{
	Use:   "migrations",
	Short: "Manage database migrations using CLI ./devtool migrations init or ./devtool migrations new -m MESSAGE",
}

func init() {
	// Add subcommands to MigrationsCmd
	MigrationsCmd.AddCommand(initCmd)
	MigrationsCmd.AddCommand(newCmd)
}
