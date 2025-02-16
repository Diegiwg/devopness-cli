package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/diegiwg/devopness-cli/core"
	service "github.com/diegiwg/devopness-cli/generated"
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

		response := user.LoginUser(ctx, email, password)

		// Check for empty response
		if response == "" {
			return errors.New("empty response from login API")
		}

		// Define structure to parse JSON response
		var respData struct {
			TokenType    string `json:"token_type"`
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
			ExpiresIn    int    `json:"expires_in"`
		}

		// Decode JSON response
		if err := json.Unmarshal([]byte(response), &respData); err != nil {
			return fmt.Errorf("failed to parse login response: %w", err)
		}

		// Ensure required fields are present
		if respData.AccessToken == "" || respData.RefreshToken == "" {
			return errors.New("invalid response: missing authentication tokens")
		}

		// Update authentication context
		ctx.Client.Auth.TokenType = respData.TokenType
		ctx.Client.Auth.AccessToken = respData.AccessToken
		ctx.Client.Auth.RefreshToken = respData.RefreshToken
		ctx.Client.Auth.ExpiresIn = respData.ExpiresIn

		ctx.Authenticated = true

		if err := ctx.SaveToFile(); err != nil {
			return fmt.Errorf("failed to save authentication context: %w", err)
		}

		fmt.Println("Login successful!")
		return nil
	},
}
