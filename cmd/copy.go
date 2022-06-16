/*
Copyright © 2022 Leonardo Biffi <leonardobiffi@outlook.com>
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
	Source      string
	Destination string
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy docker images",
	Long:  `Copy docker images`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
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

		buildCmd := exec.Command("docker", "build", "-t", Destination, "/tmp/effingo")
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr
		if err = buildCmd.Run(); err != nil {
			return
		}

		pushCmd := exec.Command("docker", "push", Destination)
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
	copyCmd.Flags().StringVarP(&Source, "src", "s", "", "Source image from Dockerhub")
	copyCmd.Flags().StringVarP(&Destination, "dest", "d", "", "Destination image to ECR Registry")
}
