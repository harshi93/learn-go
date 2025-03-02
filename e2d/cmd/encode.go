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

		// Check if file is empty
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			fmt.Printf("Error getting file info: %v\n", err)
			return
		} else if fileInfo.Size() == 0 {
			fmt.Printf("Skipping encoding, file is empty: %s\n", filePath)
			return
		} else if fileInfo.IsDir() {
			fmt.Printf("Skipping encoding, file is a directory: %s\n", filePath)
			return
		} else if filepath.Ext(filePath) == ".docx" || filepath.Ext(filePath) == ".pptx" {
			fmt.Printf("Skipping encoding, file is a docx or pptx file: %s\n", filePath)
			return
		} else if filepath.Ext(filePath) == ".raw" || filepath.Ext(filePath) == ".avchd" || filepath.Ext(filePath) == ".prores" || filepath.Ext(filePath) == ".dnxhd" {
			// RAW, AVCHD, ProRes, and DNxHD not to be encoded
			fmt.Printf("Skipping encoding, file is a RAW, AVCHD, ProRes, or DNxHD file: %s\n", filePath)
			return
		} else if filepath.Ext(filePath) == ".gif" || filepath.Ext(filePath) == ".jpeg" || filepath.Ext(filePath) == ".svg" || filepath.Ext(filePath) == ".bmp" || filepath.Ext(filePath) == ".png" || filepath.Ext(filePath) == ".tiff" || filepath.Ext(filePath) == ".webp" || filepath.Ext(filePath) == ".raw" {
			// gif, jpeg, svg, bmp, png, tiff, webp, and raw are not to be encoded
			fmt.Printf("Skipping encoding, file is an image: %s\n", filePath)
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
