package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "goweathercli",
	Short: "A Go based weather application that uses Wttr.in via Curl commands to obtain the user's current weather information",
	Long:  "GoWeatherCLI is designed to obtain the user's weather information via making specific curl commands to Wttr.in for the given user's location",
}
