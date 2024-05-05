/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/abibby/dots/config"
	"github.com/abibby/dots/link"
	"github.com/spf13/cobra"
)

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return link.Link(config.HomeDir(), config.DotDir(), replace)
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)
}
