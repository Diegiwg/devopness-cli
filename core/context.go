package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Context struct {
	Client        *API `json:"api"`
	Authenticated bool `json:"authenticated"`
}

// NewContext creates and returns a new Context instance.
func NewContext() *Context {
	return &Context{
		Client: NewAPI(),
	}
}

// SaveToFile saves the context to a file in the user's home directory.
func (ctx *Context) SaveToFile() error {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}

	path := filepath.Join(userHome, "devopness.ctx")

	data, err := json.Marshal(ctx)
	if err != nil {
		return fmt.Errorf("failed to marshal context: %w", err)
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write context to file: %w", err)
	}

	return nil
}

// LoadFromFile loads the context from a file in the user's home directory.
func (ctx *Context) LoadFromFile() error {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}

	path := filepath.Join(userHome, "devopness.ctx")

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read context file: %w", err)
	}

	err = json.Unmarshal(data, ctx)
	if err != nil {
		return fmt.Errorf("failed to unmarshal context: %w", err)
	}

	return nil
}
