package etc

import (
	"os"
	"text/template"
)

const (
	tmplDockerComposeFilePath = "tmpl/docker-compose.tmpl.yml"
	tmplMakefilePath          = "tmpl/Makefile.tmpl"
)

type InitOpt struct {
	DockerImage string
	UserName    string
	RepoName    string
	BuildPath   string
}

func NewInitOpt(repoPath string) (*InitOpt, error) {
	remote, err := GetDefaultRemote(repoPath)
	return &InitOpt{
		UserName:  remote.Owner,
		RepoName:  remote.RepoName,
		BuildPath: "./cmd/",
	}, err
}

func generateFileFromTemplate(tmplName, tmplFilePath, outputFilePath string, opt *InitOpt) error {
	templateContents, err := Asset(tmplFilePath)
	if err != nil {
		return err
	}

	tpl := template.Must(template.New(tmplName).Parse(string(templateContents)))

	outputFile, err := os.OpenFile(outputFilePath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}

	err = tpl.Execute(outputFile, opt)
	if err != nil {
		return err
	}
	return nil

}

func GenerateDockerComposeFile(repoPath, dockerImageName, dockerComposeFilePath string) error {
	opt, err := NewInitOpt(repoPath)
	if err != nil {
		return err
	}

	opt.DockerImage = dockerImageName
	return generateFileFromTemplate("makefile", tmplDockerComposeFilePath, dockerComposeFilePath, opt)
}

func GenerateMakefile(repoPath, makefilePath string) error {
	opt, err := NewInitOpt(repoPath)
	if err != nil {
		return err
	}
	return generateFileFromTemplate("makefile", tmplMakefilePath, makefilePath, opt)
}
