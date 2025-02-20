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

	"github.com/charmbracelet/huh"
)

func InteractiveMode(action string) {
	snippets, err := GetSnippets()
	if err != nil {
		fmt.Printf("Error getting snippets: %v\n", err)
		os.Exit(1)
	}

	var snippetName string

	if len(snippets) == 0 {
		if action == "" || action == "edit" {
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("Enter your snippet's name.").
						Value(&snippetName),
				),
			)
	
			if err := form.Run(); err != nil {
				log.Fatalf("Form failed: %v\n", err)
			}
	
			EditSnippet(snippetName)
			os.Exit(0)
		}
		fmt.Println("Snippet list is empty!")
		os.Exit(1)
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select a snippet").
				Options(huh.NewOptions(snippets...)...).
				Value(&snippetName),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatalf("Form failed: %v\n", err)
	}

	switch action {
	case "run":
		RunSnippet(snippetName, nil)
	case "copy":
		CopySnippet(snippetName)
	case "print":
		PrintSnippet(snippetName)
	case "edit":
		EditSnippet(snippetName)
	case "delete":
		DeleteSnippet(snippetName)
	}

	if action == "" {
		PromptAction(snippetName)
	}
}

func PromptAction(snippetName string) {
	var action string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select an action").
				Options(huh.NewOptions("Run", "Print", "Copy", "Edit", "Delete")...).
				Value(&action),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatalf("Form failed: %v\n", err)
	}

	switch action {
	case "Run":
		RunSnippet(snippetName, nil)
	case "Copy":
		CopySnippet(snippetName)
	case "Print":
		PrintSnippet(snippetName)
	case "Edit":
		EditSnippet(snippetName)
	case "Delete":
		DeleteSnippet(snippetName)
	}
}
