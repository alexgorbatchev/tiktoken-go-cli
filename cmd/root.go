package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "tiktoken",
	Short: "A CLI tool for counting tokens using OpenAI's tiktoken",
	Long: `tiktoken-go-cli is a command line interface for the tiktoken-go library.
It allows you to count tokens in text using various OpenAI tokenization encodings.

Examples:
  # Count tokens using default encoding (cl100k_base)
  echo "Hello, world!" | tiktoken count

  # Count tokens for a specific model
  tiktoken count -m gpt-4o "Hello, world!"

  # Count tokens using a specific encoding
  tiktoken count -e o200k_base "Hello, world!"

  # Encode text to token IDs
  tiktoken encode "Hello, world!"

  # Decode token IDs back to text
  tiktoken decode 15339 1917 0`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("tiktoken version %s\n", version)
		fmt.Printf("  commit: %s\n", commit)
		fmt.Printf("  built:  %s\n", date)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
