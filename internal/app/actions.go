package snipit

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/atotto/clipboard"
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
	if err := os.Remove(filePath); err != nil {
		fmt.Printf("Error deleting snippet: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Snippet deleted successfully.")
}