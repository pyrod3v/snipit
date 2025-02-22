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

package snipit

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
)

func RunSnippet(snippetName string, extraArgs []string) {
	filePath := GetSnippetFilePath(snippetName)
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading snippet: %v\n", err)
		os.Exit(1)
	}
	if len(content) == 0 {
		fmt.Printf("Snippet %s is empty.\n", snippetName)
		os.Exit(1)
	}

	args := []string{"-c", string(content), "script"}
	if len(extraArgs) > 0 {
		args = append(args, extraArgs...)
	}
	cmd := exec.Command("sh", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing %s: %v\n", snippetName, err)
		os.Exit(1)
	}
}

func CopySnippet(snippetName string) {
	filePath := GetSnippetFilePath(snippetName)
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading snippet: %v\n", err)
		os.Exit(1)
	}
	if len(content) == 0 {
		fmt.Printf("Snippet %s is empty!\n", snippetName)
		os.Exit(1)
	}

	if err := clipboard.WriteAll(string(content)); err != nil {
		fmt.Printf("Error copying %s to clipboard: %v\n", snippetName, err)
		os.Exit(1)
	}
	fmt.Println("Snippet copied to clipboard.")
}

func PrintSnippet(snippetName string) {
	filePath := GetSnippetFilePath(snippetName)
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading snippet: %v\n", err)
		os.Exit(1)
	}
	if len(content) == 0 {
		fmt.Printf("Snippet %s is empty.\n", snippetName)
		os.Exit(1)
	}
	fmt.Println(string(content))
}

func EditSnippet(snippetName string) {
	filePath := GetSnippetFilePath(snippetName)
	OpenEditor(filePath)
}

func DeleteSnippet(snippetName string) {
	filePath := GetSnippetFilePath(snippetName)
	var delete bool

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(fmt.Sprintf("Are you sure you want to delete %v?", snippetName)).
				Value(&delete),
		),
	).WithKeyMap(func(k *huh.KeyMap) *huh.KeyMap {
		k.Quit = key.NewBinding(key.WithKeys("q", "ctrl+c"))
		return k
	}(huh.NewDefaultKeyMap()))

	if err := form.Run(); err != nil {
		log.Fatalf("Form failed: %v\n", err)
	}

	if !delete {
		fmt.Println("Snippet deletion cancelled.")
		os.Exit(0)
	}

	if err := os.Remove(filePath); err != nil {
		fmt.Printf("Error deleting snippet: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Snippet deleted successfully.")
}
