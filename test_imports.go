package main

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

func main() {
    // Use the imported packages
    _ = cobra.Command{}
    _ = viper.New()

    fmt.Println("Imports are successful")
}
