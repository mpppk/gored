package cmd

import (
	"fmt"

	"github.com/mpppk/gored/etc"
	"github.com/spf13/cobra"
)

const (
	defaultDockerImage = "mpppk/gored"
)

var buildPath string
var versionPath string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate docker-compose.yml and Makefile",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		repoPath := "."

		if len(args) > 0 {
			repoPath = args[0]
		}

		dockerComposeFilePath, err := etc.GetDockerComposeFilePath(repoPath)
		etc.PanicIfError(err)

		makefilePath, err := etc.GetMakeFilePath(repoPath)
		etc.PanicIfError(err)

		circleCIConfigFilePath, err := etc.GetCircleCIConfigFilePath(repoPath)
		etc.PanicIfError(err)

		fg := &etc.FileGenerator{
			RepoPath:               repoPath,
			DockerImageName:        defaultDockerImage,
			DockerComposeFilePath:  dockerComposeFilePath,
			MakefilePath:           makefilePath,
			CircleCIConfigFilePath: circleCIConfigFilePath,
			BuildPath:              buildPath,
			VersionPath:            versionPath,
		}

		err = fg.GenerateDockerComposeFile()
		etc.PanicIfError(err)
		fmt.Printf("Generate dokcer-compose file to %s\n", dockerComposeFilePath)

		err = fg.GenerateMakefile()
		etc.PanicIfError(err)
		fmt.Printf("Generate Makefile to %s\n", makefilePath)

		err = fg.GenerateCircleCIConfigFile()
		etc.PanicIfError(err)
		fmt.Printf("Generate CircleCI config file to %s\n", circleCIConfigFilePath)
	},
}

func init() {
	initCmd.Flags().StringVarP(&buildPath, "build-path", "b", ".", "build path")
	initCmd.Flags().StringVarP(&versionPath, "version-path", "V", "", "version file path")

	RootCmd.AddCommand(initCmd)
}
