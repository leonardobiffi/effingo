/*
Copyright © 2022 Leonardo Biffi <leonardobiffi@outlook.com>
*/
package cmd

import (
	"effingo/version"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "effingo",
	Short:   "Copy docker images from Dockerhub to ECR Registry",
	Long:    `Copy docker images from Dockerhub to ECR Registry`,
	Version: version.String(),
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Set the version template of the application
	rootCmd.SetVersionTemplate(version.Template())
}
