package cmd

import (
	"fmt"

	"github.com/mpppk/gored/etc"
	"github.com/spf13/cobra"
)

const (
	outputDockerComposeFileName = "docker-compose.gored.yml"
	outputMakeFileName          = "Makefile.gored"
	defaultDockerImage          = "mpppk/gored"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate docker-compose.yml and Makefile",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		repoPath := "."
		etc.GenerateDockerComposeFile(repoPath, defaultDockerImage, outputDockerComposeFileName)
		fmt.Printf("Generate dokcer-compose file to %s", repoPath+"/"+outputDockerComposeFileName+"\n")
		etc.GenerateMakefile(repoPath, outputMakeFileName)
		fmt.Printf("Generate Makefile to %s", repoPath+"/"+outputMakeFileName+"\n")
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
