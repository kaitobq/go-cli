/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// gopherCmd represents the gopher command
var gopherCmd = &cobra.Command{
	Use:   "gopher",
	Short: "Print the ASCII art of a gopher",
	Run: func(cmd *cobra.Command, args []string) {
		b, err := os.ReadFile("aa/gopher.txt")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Println(string(b))
	},
}

func init() {
	rootCmd.AddCommand(gopherCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gopherCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gopherCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
