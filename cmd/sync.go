/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Execute sync between ECR registries",
	Long:  `Execute sync between ECR registries`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		// TODO: docker login tendo que ser executado antes de executar este comando
		// Ex.: aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 555407960619.dkr.ecr.us-east-1.amazonaws.com

		ctx := context.Background()

		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			log.Fatalf("unable to load SDK config, %v", err)
		}

		svc := ecr.NewFromConfig(cfg)

		result, err := svc.DescribeImages(ctx, &ecr.DescribeImagesInput{
			RepositoryName: aws.String(Source),
			MaxResults:     aws.Int32(100),
		})
		if err != nil {
			log.Fatalf("unable to describe images, %v", err)
		}

		repoSource, err := svc.DescribeRepositories(ctx, &ecr.DescribeRepositoriesInput{
			RepositoryNames: []string{Source},
		})
		if err != nil {
			log.Fatalf("unable to describe repositories, %v", err)
		}

		repoDest, err := svc.DescribeRepositories(ctx, &ecr.DescribeRepositoriesInput{
			RepositoryNames: []string{Destination},
		})
		if err != nil {
			log.Fatalf("unable to describe repositories, %v", err)
		}

		for _, image := range result.ImageDetails {
			imageSource := fmt.Sprintf("%s:%s", *repoSource.Repositories[0].RepositoryUri, image.ImageTags[0])
			imageDest := fmt.Sprintf("%s:%s", *repoDest.Repositories[0].RepositoryUri, image.ImageTags[0])

			fmt.Println("Syncing image...")
			fmt.Println("From:", imageSource)
			fmt.Println("To:", imageDest)

			type Command struct {
				Cmd []string
			}

			commands := []Command{
				{
					Cmd: []string{"docker", "pull", imageSource},
				},
				{
					Cmd: []string{"docker", "tag", imageSource, imageDest},
				},
				{
					Cmd: []string{"docker", "push", imageDest},
				},
			}

			for _, command := range commands {
				cmd := exec.Command(command.Cmd[0], command.Cmd[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					log.Fatalf("unable to execute command, %v", err)
				}
			}

			fmt.Println("")
			fmt.Println("Sync complete!")
		}

		return
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	// Here you will define your flags and configuration settings.
	syncCmd.Flags().StringVarP(&Source, "src", "s", "", "Source ECR Registry")
	syncCmd.Flags().StringVarP(&Destination, "dest", "d", "", "Destination ECR Registry")
}
