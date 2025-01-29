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
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/atotto/clipboard"
	"github.com/manifoldco/promptui"
)

func main() {
	if len(os.Args) == 1 {
		snippets, err := getSnippets()
		if err != nil {
			fmt.Printf("Error getting snippets: %v\n", err)
			os.Exit(1)
		}

		if len(snippets) == 0 {
			fmt.Println("No snippets found.")
			os.Exit(0)
		}

		prompt := promptui.Select{
			Label: "Select a snippet",
			Items: snippets,
		}

		_, snippetName, err := prompt.Run()
		if err != nil {
			fmt.Println("Prompt cancelled by user.")
			os.Exit(0)
		}

		promptAction(snippetName)
	} else if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("Usage:")
		fmt.Println("  snipit                      List and manage snippets")
		fmt.Println("  snipit <snippet-name>       Create or manage a snippet")
		fmt.Println("  snipit -h | --help          Show this help message")
		os.Exit(0)
	} else {
		snippetName := os.Args[1]
		dir := filepath.Join(getConfigDir(), "snippets")
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Printf("Error creating snippets directory: %v\n", err)
				os.Exit(1)
			}
		}
		filePath := filepath.Join(dir, snippetName)

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Printf("Creating new snippet: %s\n", snippetName)
			openEditor(filePath)
		} else {
			promptAction(snippetName)
		}
	}
}

func promptAction(snippetName string) {
	dir := filepath.Join(getConfigDir(), "snippets")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("Error creating snippets directory: %v\n", err)
			os.Exit(1)
		}
	}
	filePath := filepath.Join(dir, snippetName)

	prompt := promptui.Select{
		Label: "Choose an action",
		Items: []string{"Run", "Print", "Copy", "Edit", "Delete"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt cancelled by user.")
		return
	}

	switch result {
	case "Run":
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading snippet: %v\n", err)
			os.Exit(1)
		}
		if len(content) == 0 {
			fmt.Printf("Snippet %v is empty!", snippetName)
			os.Exit(1)
		}

		cmd := exec.Command("sh", "-c", string(content))
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error executing %v: %v\n", snippetName, err)
			os.Exit(1)
		}
	case "Copy":
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading snippet: %v\n", err)
			os.Exit(1)
		}
		if len(content) == 0 {
			fmt.Printf("Snippet %v is empty!", snippetName)
			os.Exit(1)
		}

		if err := clipboard.WriteAll(string(content)); err != nil {
			fmt.Printf("Error copying %v to clipboard: %v\n", snippetName, err)
			os.Exit(1)
		}
		fmt.Println("Snippet copied to clipboard.")
	case "Print":
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading snippet: %v\n", err)
			os.Exit(1)
		}
		if len(content) == 0 {
			fmt.Printf("Snippet %v is empty!", snippetName)
			os.Exit(1)
		}
		fmt.Println(string(content))
	case "Edit":
		openEditor(filePath)
	case "Delete":
		if err := os.Remove(filePath); err != nil {
			fmt.Printf("Error deleting snippet: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Snippet deleted successfully.")
	}
}

func getConfigDir() string {
	var configDir string
	if runtime.GOOS == "windows" {
		configDir = os.Getenv("AppData")
	} else {
		configDir = filepath.Join(os.Getenv("HOME"), ".config")
	}
	return filepath.Join(configDir, "snipit")
}

func getSnippets() ([]string, error) {
	dir := filepath.Join(getConfigDir(), "snippets")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("Error creating snippets directory: %v\n", err)
			os.Exit(1)
		}
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var snippets []string
	for _, file := range files {
		if !file.IsDir() {
			snippets = append(snippets, file.Name())
		}
	}
	return snippets, nil
}

func openEditor(filePath string) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}

	cmd := exec.Command(editor, filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error opening editor: %v\n", err)
		os.Exit(1)
	}
}
