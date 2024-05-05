/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/abibby/dots/config"
	"github.com/abibby/dots/link"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <path>",
	Short: "",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		homeFile := args[0]

		relFile, err := filepath.Rel(config.HomeDir(), homeFile)
		if err != nil {
			return err
		}
		if strings.HasPrefix(relFile, "../") {
			return fmt.Errorf("file must be in the home directory")
		}
		dotFile := path.Join(config.DotDir(), relFile)

		return link.Move(homeFile, dotFile, replace)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
