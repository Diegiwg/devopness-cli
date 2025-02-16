package cmd

import (
	"fmt"
	"log"

	"github.com/diegiwg/devopness-cli/core"
	"github.com/diegiwg/devopness-cli/service"
	"github.com/spf13/cobra"
)

func init() {
	loginCmd.Flags().StringP("email", "e", "", "Email (required)")
	loginCmd.Flags().StringP("password", "p", "", "Password (required)")

	if err := loginCmd.MarkFlagRequired("email"); err != nil {
		log.Fatalf("Error marking email as required: %v", err)
	}

	if err := loginCmd.MarkFlagRequired("password"); err != nil {
		log.Fatalf("Error marking password as required: %v", err)
	}

	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with Devopness",
	Long:  `Login to Devopness using your email and password.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		email, err := cmd.Flags().GetString("email")
		if err != nil {
			return fmt.Errorf("failed to retrieve email flag: %w", err)
		}

		password, err := cmd.Flags().GetString("password")
		if err != nil {
			return fmt.Errorf("failed to retrieve password flag: %w", err)
		}

		ctx := core.NewContext()
		user := service.User{}

		if err := user.Login(ctx, email, password); err != nil {
			return fmt.Errorf("login failed: %w", err)
		}

		if err := ctx.SaveToFile(); err != nil {
			return fmt.Errorf("failed to save authentication context: %w", err)
		}

		fmt.Println("Login successful!")
		return nil
	},
}
