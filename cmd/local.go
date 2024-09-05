/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var local = &cobra.Command{
	Use:   "local",
	Short: "This obtains your local weather using Wttr.in",
	Long: `This command allows you to obtain the local weather using Wttr.in. For example:

Wttr.in is a CLI tool that outputs your local weather in browser or terminal`,
	Run: localRun,
}

func init() {
	rootCmd.AddCommand(local)
}

func localRun(cmd *cobra.Command, args []string) {
	fmt.Println("local command called")
	for _, x := range args {
		cmd := exec.Command("curl", fmt.Sprintf("https://wttr.in/<%s>", x))

		out, err := cmd.Output()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(out))

	}
}
