package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/diegiwg/devopness-cli/core"
	"github.com/diegiwg/devopness-cli/generated/models"
	"github.com/diegiwg/devopness-cli/generated/services"
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
		user := service.Users{}

		response, _ := user.LoginUser(ctx, model.UserLogin{Email: email, Password: password})

		// Check for empty response
		if response == nil {
			return errors.New("empty response from login API")
		}

		// Ensure required fields are present
		if response.AccessToken == "" || response.RefreshToken == "" {
			return errors.New("invalid response: missing authentication tokens")
		}

		// Update authentication context
		ctx.Client.Auth.TokenType = response.TokenType
		ctx.Client.Auth.AccessToken = response.AccessToken
		ctx.Client.Auth.RefreshToken = response.RefreshToken
		ctx.Client.Auth.ExpiresIn = response.ExpiresIn

		ctx.Authenticated = true

		if err := ctx.SaveToFile(); err != nil {
			return fmt.Errorf("failed to save authentication context: %w", err)
		}

		fmt.Println("Login successful!")
		return nil
	},
}
