package etc

import (
	"os"
	"text/template"
)

const (
	tmplDockerComposeFilePath = "tmpl/docker-compose.tmpl.yml"
)

type InitOpt struct {
	DockerImage string
	UserName    string
	RepoName    string
}

func GenerateDockerComposeFile(repoPath, dockerImageName, dockerComposeFilePath string) error {
	templateContents, err := Asset(tmplDockerComposeFilePath)
	if err != nil {
		return err
	}

	tpl := template.Must(template.New("docker-compose").Parse(string(templateContents)))
	remote, err := GetDefaultRemote(repoPath)
	if err != nil {
		return err
	}

	dockerComposeFile, err := os.OpenFile(dockerComposeFilePath, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	err = tpl.Execute(dockerComposeFile, &InitOpt{
		DockerImage: dockerImageName,
		UserName:    remote.Owner,
		RepoName:    remote.RepoName,
	})
	if err != nil {
		return err
	}
	return nil
}
