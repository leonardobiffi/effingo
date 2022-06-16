/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"effingo/template"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	Source string
	Dest   string
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		fmt.Println("Source:", Source)
		fmt.Println("Dest:", Dest)

		file, err := template.File(template.Dockerfile{
			Image: Source,
		})
		if err != nil {
			return
		}
		fmt.Println(file)

		if err = os.MkdirAll("/tmp/effingo", 0755); err != nil {
			return
		}

		f, err := os.Create("/tmp/effingo/Dockerfile")
		if err != nil {
			return
		}
		_, err = f.WriteString(file)
		if err != nil {
			return
		}
		f.Close()

		fmt.Println("Dockerfile created")

		buildCmd := exec.Command("docker", "build", "-t", Dest, "/tmp/effingo")
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr
		if err = buildCmd.Run(); err != nil {
			return
		}

		pushCmd := exec.Command("docker", "push", Dest)
		pushCmd.Stdout = os.Stdout
		pushCmd.Stderr = os.Stderr
		if err = pushCmd.Run(); err != nil {
			return
		}

		return
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// copyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// copyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	copyCmd.Flags().StringVarP(&Source, "source", "s", "", "Help message for flag")
	copyCmd.Flags().StringVarP(&Dest, "dest", "d", "", "Help message for flag")
}
