package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// Command definition
var outputFile string

var dnswordlistCmd = &cobra.Command{
	Use:   "dnswordlist",
	Short: "Process subdomains and generate a word count list",
	Long:  `Processes subdomains from stdin, counts unique parts, and outputs the counts in a JSON file or prints to stdout if no output file is specified.

Examples:
  cat subs.txt | wordcount dnswordlist -o best-dns-wordlist.json
`,
	Run: func(cmd *cobra.Command, args []string) {
		processInput()
	},
}

func processSubdomains(parts []string) []string {
	return parts[:len(parts)-2] // Exclude the last two parts (main domain)
}

func processInput() {
	wordCounts := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Split(line, ".")

		if len(parts) > 2 {
			processedParts := processSubdomains(parts)
			for _, part := range processedParts {
				wordCounts[part]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		return
	}

	// Convert word counts to JSON
	jsonOutput, err := json.MarshalIndent(wordCounts, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling JSON: %v\n", err)
		return
	}

	fmt.Println(string(jsonOutput))

	// Write to file if outputFile is specified, otherwise print to stdout
	if outputFile != "" {
		file, err := os.Create(outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			return
		}
		defer file.Close()

		if _, err := file.Write(jsonOutput); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to output file: %v\n", err)
			return
		}
	}
}

func init() {
	rootCmd.AddCommand(dnswordlistCmd)
	dnswordlistCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file (optional)")
}
