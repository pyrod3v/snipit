// Copyright 2025 pyrod3v
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pyrod3v/snipit/internal/app"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

func main() {
	if _, err := os.Stat(snipit.GetConfigDir()); os.IsNotExist(err) {
		if err := os.MkdirAll(snipit.GetConfigDir(), 0755); err != nil {
			fmt.Printf("Error creating config directory: %v\n", err)
			os.Exit(1)
		}
	}

	configFile := filepath.Join(snipit.GetConfigDir(), "config.yaml")
	if _, err := os.Stat(configFile); err != nil && os.IsNotExist(err) {
		_, err := os.Create(configFile)
		if err != nil {
			fmt.Printf("Error creating config file: %v\n", err)
			os.Exit(1)
		}
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./.snipit")
	viper.AddConfigPath(snipit.GetConfigDir())
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error loading config: %w", err))
	}

	viper.SetDefault("SnippetsDir", filepath.Join(snipit.GetConfigDir(), "snippets"))
	viper.SetDefault("Editor", "nano")
	viper.BindEnv("EDITOR")
	viper.WriteConfig()

	snipit.EnsureSnippetsDir()

	app := &cli.App{
		Name:  "snipit",
		Usage: "An easy to use, interactive snippet manager",
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "Run the snippet with optional arguments",
				Action: func(c *cli.Context) error {
					if c.NArg() < 1 {
						return cli.Exit("Missing snippet name for run", 1)
					}
					snippetName := c.Args().Get(0)
					args := c.Args().Slice()[1:]
					filePath := snipit.GetSnippetFilePath(snippetName)
					if _, err := os.Stat(filePath); os.IsNotExist(err) {
						return cli.Exit(fmt.Sprintf("Snippet %s does not exist.", snippetName), 1)
					}
					snipit.RunSnippet(snippetName, args)
					return nil
				},
			},
			{
				Name:    "copy",
				Aliases: []string{"c"},
				Usage:   "Copy the snippet to the clipboard",
				Action: func(c *cli.Context) error {
					if c.NArg() < 1 {
						return cli.Exit("Missing snippet name for copy", 1)
					}
					snippetName := c.Args().Get(0)
					snipit.CopySnippet(snippetName)
					return nil
				},
			},
			{
				Name:    "print",
				Aliases: []string{"p"},
				Usage:   "Print the snippet content",
				Action: func(c *cli.Context) error {
					if c.NArg() < 1 {
						return cli.Exit("Missing snippet name for print", 1)
					}
					snippetName := c.Args().Get(0)
					snipit.PrintSnippet(snippetName)
					return nil
				},
			},
			{
				Name:    "edit",
				Aliases: []string{"e"},
				Usage:   "Edit or create a snippet",
				Action: func(c *cli.Context) error {
					if c.NArg() < 1 {
						return cli.Exit("Missing snippet name for edit", 1)
					}
					snippetName := c.Args().Get(0)
					snipit.EditSnippet(snippetName)
					return nil
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "Delete a snippet",
				Action: func(c *cli.Context) error {
					if c.NArg() < 1 {
						return cli.Exit("Missing snippet name for delete", 1)
					}
					snippetName := c.Args().Get(0)
					snipit.DeleteSnippet(snippetName)
					return nil
				},
			},
			{
				Name:    "config",
				Aliases: []string{"cfg"},
				Usage:   "Update or get a config value",
				Action: func(c *cli.Context) error {
					if c.NArg() == 0 {
						for _, key := range viper.AllKeys() {
							fmt.Printf("%v: %v\n", key, viper.Get(key))
						}
					} else if c.NArg() == 1 {
						fmt.Println(viper.Get(c.Args().Get(0)))
					} else {
						key := c.Args().Get(0)
						value := c.Args().Get(1)
						viper.Set(key, value)
						viper.WriteConfig()
					}
					return nil
				},
			},
		},

		// run in interactive mode if no command is provided
		Action: func(c *cli.Context) error {
			// if a snippet is provided, prompt an action for it
			if len(os.Args) > 1 {
				snippet := os.Args[1]
				snippetPath := snipit.GetSnippetFilePath(snippet)
				if _, err := os.Stat(snippetPath); os.IsNotExist(err) {
					fmt.Printf("Creating new snippet: %s\n", snippet)
					snipit.OpenEditor(snippetPath)
				} else {
					snipit.PromptAction(snippet)
				}
			} else {
				snipit.InteractiveMode()
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
