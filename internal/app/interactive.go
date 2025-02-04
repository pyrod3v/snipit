package snipit

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func InteractiveMode() {
	snippets, err := GetSnippets()
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

	PromptAction(snippetName)
}

func PromptAction(snippetName string) {
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
