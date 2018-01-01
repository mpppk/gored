package cmd

import (
	"fmt"

	"os"
	"text/template"

	"github.com/mpppk/gored/etc"
	"github.com/spf13/cobra"
)

type InitProps struct {
	DockerImage string
	UserName    string
	RepoName    string
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate docker-compose.yml and Makefile",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		templateContents, err := etc.Asset("tmpl/docker-compose.tmpl.yml")
		//templateContents, err := ioutil.ReadFile("tmpl/docker-compose.tmpl.yml")
		if err != nil {
			panic(err)
		}
		tpl := template.Must(template.New("docker-compose").Parse(string(templateContents)))
		remote, err := etc.GetDefaultRemote(".")
		if err != nil {
			panic(err)
		}
		err = tpl.Execute(os.Stdout, &InitProps{
			DockerImage: "mpppk/gored",
			UserName:    remote.Owner,
			RepoName:    remote.RepoName,
		})
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
