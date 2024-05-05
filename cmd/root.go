/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var replace bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dots",
	Short: "A simple dot file manager",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	cfg, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path.Join(home, ".dots"))
	viper.AddConfigPath(path.Join(home, ".config/dots"))
	viper.AddConfigPath(path.Join(cfg, "dots"))
	viper.AddConfigPath(".")

	rootCmd.PersistentFlags().BoolVarP(&replace, "replace", "r", false, "Replace existing files")
	rootCmd.PersistentFlags().String("home", home, "home directory")
	rootCmd.PersistentFlags().StringP("dot", "d", path.Join(home, "dotfiles"), "dotfile directory")

	err = viper.BindPFlag("home_dir", rootCmd.Flag("home"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("dot_dir", rootCmd.Flag("dot"))
	if err != nil {
		panic(err)
	}

	err = viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		// fallthrough
	} else if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
