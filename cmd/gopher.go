/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/gosuri/uilive"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// gopherCmd represents the gopher command
var gopherCmd = &cobra.Command{
	Use:   "gopher",
	Short: "Print the ASCII art of a gopher",
	Run: func(cmd *cobra.Command, args []string) {
		b := []string{
			"         ,_---~~~~~----._         ",
			"  _,,_,*^____      _____``*g*\"*, ",
			" / __/ /'     ^.  /      \\ ^@q   f ",
			"[  @f | @))    |  | @))   l  0 _/  ",
			" \\`/   \\~____ / __ \\_____/    \\   ",
			"  |           _l__l_           I   ",
			"  }          [______]           I  ",
			"  ]            | | |            |  ",
			"  ]             ~ ~             |  ",
			"  |                            |   ",
			"   |                           | ",
		}

		bw := len(b[0])

		tw, _, err := terminal.GetSize(syscall.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		for i := range b {
			b[i] = strings.Repeat(" ", tw) + b[i]
		}
	
		writer := uilive.New()
		
		writer.Start()
		for i := 0; i <= tw+bw; i++ {
			var out string
			for _, v := range b {
				if i < bw {
					out += v[i:tw+i] + "\n"
					} else if i < len(v) {
						out += v[i:] + "\n"
						} else {
							out += "\n"
						}
					}
					fmt.Fprint(writer, out)
					time.Sleep(10 * time.Millisecond)
				}
			writer.Stop()
		// fmt.Println(strings.Join(b, "\n"))
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
