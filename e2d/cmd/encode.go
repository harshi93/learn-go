/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode files in current directory",
	Long:  `Encode files from the current working directory based on input filename`,
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("input")
		if filename == "" {
			fmt.Println("Please provide a filename using --input or -i flag")
			return
		}

		// Get current working directory
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			return
		}

		// Create full file path
		filePath := filepath.Join(currentDir, filename)

		// Check if file exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Printf("File not found: %s\n", filePath)
			return
		} else {
			fmt.Printf("Found file: %s\n", filePath)
		}

		// Read file content
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}

		// Convert to base64
		base64String := base64.StdEncoding.EncodeToString(fileContent)

		fmt.Printf("Base64 encoded content: %s\n", base64String)
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)
	encodeCmd.Flags().StringP("input", "i", "", "input file to encode")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
