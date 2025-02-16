package core

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Context struct {
	Client        *API `json:"api"`
	Authenticated bool `json:"authenticated"`
}

func NewContext() *Context {
	ctx := &Context{
		Client: NewAPI(),
	}

	return ctx
}

func (ctx *Context) SaveToFile() error {
	user_home, _ := os.UserHomeDir()
	path := filepath.Join(user_home, "devopness.ctx")

	data, err := json.Marshal(ctx)
	if err != nil {
		panic(err)
	}

	os.WriteFile(path, []byte(data), 0644)

	return nil
}

func (ctx *Context) LoadFromFile() error {
	user_home, _ := os.UserHomeDir()
	path := filepath.Join(user_home, "devopness.ctx")

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, ctx)
	if err != nil {
		return err
	}

	return nil
}
