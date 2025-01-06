package migrations

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize migrations",
	Run: func(cmd *cobra.Command, args []string) {
		// Create the necessary directories
		err := os.MkdirAll("internal/migrations", os.ModePerm)
		if err != nil {
			fmt.Println("Error creating migrations directory:", err)
			return
		}

		// Check if internal/database.go exists
		if _, err := os.Stat("internal/database.go"); os.IsNotExist(err) {
			fmt.Println("Error: internal/database.go does not exist. Please ensure the file is in place.")
			return
		} else {
			fmt.Println("internal/database.go already exists, skipping creation.")
		}

		// Create initial migration files
		timestamp := time.Now().Format("20060102150405")
		upFile := filepath.Join("internal", "migrations", fmt.Sprintf("%s_initial_version.up.sql", timestamp))
		downFile := filepath.Join("internal", "migrations", fmt.Sprintf("%s_initial_version.down.sql", timestamp))

		err = os.WriteFile(upFile, []byte("-- SQL for upgrading the database"), 0644)
		if err != nil {
			fmt.Println("Error creating .up.sql file:", err)
			return
		}
		err = os.WriteFile(downFile, []byte("-- SQL for downgrading the database"), 0644)
		if err != nil {
			fmt.Println("Error creating .down.sql file:", err)
			return
		}

		fmt.Println("Migrations initialized.")
	},
}

func init() {
	MigrationsCmd.AddCommand(initCmd)
}
