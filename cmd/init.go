package cmd

import (
	"os"
	"text/template"

	"fmt"

	"github.com/mpppk/gored/etc"
	"github.com/spf13/cobra"
)

type InitProps struct {
	DockerImage string
	UserName    string
	RepoName    string
}

const (
	outputDockerComposeFileName = "docker-compose.gored.yml"
	tmplDockerComposeFilePath   = "tmpl/docker-compose.tmpl.yml"
	defaultDockerImage          = "mpppk/gored"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate docker-compose.yml and Makefile",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		repoPath := "."
		templateContents, err := etc.Asset(tmplDockerComposeFilePath)
		etc.PanicIfError(err)

		tpl := template.Must(template.New("docker-compose").Parse(string(templateContents)))
		remote, err := etc.GetDefaultRemote(repoPath)
		etc.PanicIfError(err)

		dockerComposeFile, err := os.OpenFile(outputDockerComposeFileName, os.O_WRONLY|os.O_CREATE, 0600)
		etc.PanicIfError(err)

		err = tpl.Execute(dockerComposeFile, &InitProps{
			DockerImage: defaultDockerImage,
			UserName:    remote.Owner,
			RepoName:    remote.RepoName,
		})
		etc.PanicIfError(err)

		fmt.Printf("Generate dokcer-compose file to %s", repoPath+"/"+outputDockerComposeFileName)
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
