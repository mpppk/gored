package etc

import (
	"os"
	"text/template"
)

const (
	tmplDockerComposeFilePath  = "tmpl/docker-compose.tmpl.yml"
	tmplMakefilePath           = "tmpl/Makefile.tmpl"
	tmplCircleCIConfigFilePath = "tmpl/config.tmpl.yml"
)

type fileGeneratorOpt struct {
	DockerImage  string
	UserName     string
	RepoName     string
	BuildPath    string
	VersionPath  string
	MakefilePath string
}

func generateFileFromTemplate(tmplName, tmplFilePath, outputFilePath string, opt *fileGeneratorOpt) error {
	templateContents, err := Asset(tmplFilePath)
	if err != nil {
		return err
	}

	tpl := template.Must(template.New(tmplName).Parse(string(templateContents)))

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}

	err = tpl.Execute(outputFile, opt)
	if err != nil {
		return err
	}
	return nil

}

type FileGenerator struct {
	RepoPath               string
	DockerImageName        string
	DockerComposeFilePath  string
	MakefilePath           string
	CircleCIConfigFilePath string
	BuildPath              string
	VersionPath            string
}

func (f *FileGenerator) GenerateDockerComposeFile() error {
	opt, err := f.NewOpt()
	if err != nil {
		return err
	}

	opt.DockerImage = f.DockerImageName
	return generateFileFromTemplate("docker-compose", tmplDockerComposeFilePath, f.DockerComposeFilePath, opt)
}

func (f *FileGenerator) GenerateMakefile() error {
	opt, err := f.NewOpt()
	if err != nil {
		return err
	}
	return generateFileFromTemplate("makefile", tmplMakefilePath, f.MakefilePath, opt)
}

func (f *FileGenerator) GenerateCircleCIConfigFile() error {
	opt, err := f.NewOpt()
	if err != nil {
		return err
	}
	return generateFileFromTemplate("circleci", tmplCircleCIConfigFilePath, f.CircleCIConfigFilePath, opt)
}

func (f *FileGenerator) NewOpt() (*fileGeneratorOpt, error) {
	remote, err := GetDefaultRemote(f.RepoPath)

	versionPath := f.BuildPath
	if f.VersionPath != "" {
		versionPath = f.VersionPath
	}

	return &fileGeneratorOpt{
		UserName:     remote.Owner,
		RepoName:     remote.RepoName,
		BuildPath:    f.BuildPath,
		VersionPath:  versionPath,
		MakefilePath: f.MakefilePath,
	}, err
}
