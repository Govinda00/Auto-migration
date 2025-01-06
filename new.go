package migrations

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new migration",
	Run: func(cmd *cobra.Command, args []string) {
		message, _ := cmd.Flags().GetString("message")
		if message == "" {
			fmt.Println("Error: migration message is required")
			return
		}

		timestamp := time.Now().Format("20060102150405")
		sanitizedMessage := strings.ReplaceAll(strings.ToLower(message), " ", "_")

		upFile := filepath.Join("internal", "migrations", fmt.Sprintf("%s_%s.up.sql", timestamp, sanitizedMessage))
		downFile := filepath.Join("internal", "migrations", fmt.Sprintf("%s_%s.down.sql", timestamp, sanitizedMessage))

		err := os.WriteFile(upFile, []byte("-- SQL for upgrading the database"), 0644)
		if err != nil {
			fmt.Println("Error creating .up.sql file:", err)
			return
		}
		err = os.WriteFile(downFile, []byte("-- SQL for downgrading the database"), 0644)
		if err != nil {
			fmt.Println("Error creating .down.sql file:", err)
			return
		}

		fmt.Println("New migration created:", message)
	},
}

func init() {
	newCmd.Flags().StringP("message", "m", "", "Migration message")
	MigrationsCmd.AddCommand(newCmd)
}
