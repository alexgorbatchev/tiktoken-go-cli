package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/pkoukk/tiktoken-go"
	"github.com/spf13/cobra"
)

var (
	decodeModel    string
	decodeEncoding string
)

var decodeCmd = &cobra.Command{
	Use:   "decode [token_ids...]",
	Short: "Decode token IDs back to text",
	Long: `Decode token IDs back to text using the specified model or encoding.

Token IDs can be provided as arguments or piped through stdin (space or newline separated).

Examples:
  # Decode token IDs
  tiktoken decode 15339 1917 0

  # Decode using a specific model
  tiktoken decode -m gpt-4o 15339 1917 0

  # Decode from stdin
  echo "15339 1917 0" | tiktoken decode

  # Chain encode and decode
  tiktoken encode "Hello, world!" | tiktoken decode`,
	RunE: runDecode,
}

func init() {
	rootCmd.AddCommand(decodeCmd)
	decodeCmd.Flags().StringVarP(&decodeModel, "model", "m", "", "OpenAI model name (e.g., gpt-4o, gpt-4, gpt-3.5-turbo)")
	decodeCmd.Flags().StringVarP(&decodeEncoding, "encoding", "e", "cl100k_base", "Encoding name (o200k_base, cl100k_base, p50k_base, r50k_base)")
}

func runDecode(cmd *cobra.Command, args []string) error {
	tokens, err := getTokens(args)
	if err != nil {
		return err
	}

	var enc *tiktoken.Tiktoken

	if decodeModel != "" {
		enc, err = tiktoken.EncodingForModel(decodeModel)
		if err != nil {
			return fmt.Errorf("failed to get encoding for model %s: %w", decodeModel, err)
		}
	} else {
		enc, err = tiktoken.GetEncoding(decodeEncoding)
		if err != nil {
			return fmt.Errorf("failed to get encoding %s: %w", decodeEncoding, err)
		}
	}

	text := enc.Decode(tokens)
	fmt.Println(text)

	return nil
}

func getTokens(args []string) ([]int, error) {
	var tokenStrs []string

	if len(args) > 0 {
		tokenStrs = args
	} else {
		// Check if stdin has data
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			return nil, fmt.Errorf("no token IDs provided. Either pass token IDs as arguments or pipe through stdin")
		}

		reader := bufio.NewReader(os.Stdin)
		var builder strings.Builder

		for {
			line, err := reader.ReadString('\n')
			builder.WriteString(line)
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, fmt.Errorf("error reading stdin: %w", err)
			}
		}

		// Split by whitespace (spaces, tabs, newlines)
		tokenStrs = strings.Fields(builder.String())
	}

	tokens := make([]int, 0, len(tokenStrs))
	for _, s := range tokenStrs {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		t, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("invalid token ID %q: %w", s, err)
		}
		tokens = append(tokens, t)
	}

	if len(tokens) == 0 {
		return nil, fmt.Errorf("no valid token IDs provided")
	}

	return tokens, nil
}
