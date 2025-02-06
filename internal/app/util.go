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
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/viper"
)

func GetSnippetFilePath(snippetName string) string {
	return filepath.Join(viper.GetString("SnippetsDir"), snippetName)
}

func EnsureSnippetsDir() {
	dir := viper.GetString("SnippetsDir")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("Error creating snippets directory: %v\n", err)
			os.Exit(1)
		}
	}
}

func GetConfigDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error retrieving home directory: %v\n", err)
		os.Exit(1)
	}
	return filepath.Join(home, ".snipit")
}

func OpenEditor(filePath string) {
	editor := viper.GetString("Editor")
	if editor == "nano" && os.Getenv("EDITOR") != "" {
		editor = os.Getenv("EDITOR")
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

func GetSnippets() ([]string, error) {
	dir := viper.GetString("SnippetsDir")
	EnsureSnippetsDir()
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	snippets := []string{}
	for _, entry := range entries {
		if !entry.IsDir() {
			snippets = append(snippets, entry.Name())
		}
	}
	return snippets, nil
}
