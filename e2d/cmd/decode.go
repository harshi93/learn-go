/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		contents, _ := cmd.Flags().GetString("decode")

		output, err := base64.StdEncoding.DecodeString(contents)
		if err != nil {
			fmt.Println("Error decoding string:", err)
			return
		}
		fmt.Printf("Decoded content: %s\n", output)
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)
	decodeCmd.Flags().StringP("decode", "o", "", "string to decode")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	decodeCmd.Flags().StringP("cipher", "c", "b64", "cipher to use")
	decodeCmd.MarkFlagRequired("cipher")
	decodeCmd.MarkFlagRequired("decode")

}
