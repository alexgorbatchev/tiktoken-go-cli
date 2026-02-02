package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var modelsCmd = &cobra.Command{
	Use:   "models",
	Short: "List available models and their encodings",
	Long:  `Display a list of available OpenAI models and their corresponding tokenization encodings.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available Models and Encodings:")
		fmt.Println()
		fmt.Println("Encoding: o200k_base")
		fmt.Println("  - gpt-4o")
		fmt.Println("  - gpt-4.1")
		fmt.Println("  - gpt-4.5")
		fmt.Println()
		fmt.Println("Encoding: cl100k_base")
		fmt.Println("  - gpt-4")
		fmt.Println("  - gpt-3.5-turbo")
		fmt.Println("  - text-embedding-ada-002")
		fmt.Println("  - text-embedding-3-small")
		fmt.Println("  - text-embedding-3-large")
		fmt.Println()
		fmt.Println("Encoding: p50k_base")
		fmt.Println("  - text-davinci-002")
		fmt.Println("  - text-davinci-003")
		fmt.Println("  - code-davinci-002")
		fmt.Println("  - code-cushman-001")
		fmt.Println()
		fmt.Println("Encoding: r50k_base (gpt2)")
		fmt.Println("  - davinci")
		fmt.Println("  - curie")
		fmt.Println("  - babbage")
		fmt.Println("  - ada")
		fmt.Println()
		fmt.Println("Available Encodings:")
		fmt.Println("  - o200k_base   (newest, used by GPT-4o models)")
		fmt.Println("  - cl100k_base  (used by GPT-4 and GPT-3.5-turbo)")
		fmt.Println("  - p50k_base    (used by Codex models)")
		fmt.Println("  - p50k_edit    (used by edit models)")
		fmt.Println("  - r50k_base    (used by GPT-3 models)")
	},
}

func init() {
	rootCmd.AddCommand(modelsCmd)
}
