package etc

import (
	"errors"
	"regexp"
	"strings"

	"os"

	"fmt"

	"path/filepath"

	"gopkg.in/src-d/go-git.v4"
)

type Remote struct {
	URL         string
	ServiceHost string
	Owner       string
	RepoName    string
}

func NewRemote(remoteUrl string) (*Remote, error) {
	var assigned *regexp.Regexp
	if strings.HasPrefix(remoteUrl, "http") {
		assigned = regexp.MustCompile(`https?://[.+]?(.+)/(.+)/(.+)$`)
	} else if strings.HasPrefix(remoteUrl, "git://") {
		assigned = regexp.MustCompile(`git://[.+]?(.+)/(.+)/(.+)$`)
	} else if strings.HasPrefix(remoteUrl, "git@") {
		assigned = regexp.MustCompile(`git@(.+):(.+)/(.+).git`)
	} else {
		return nil, errors.New("unknown remote: " + remoteUrl)
	}

	result := assigned.FindStringSubmatch(remoteUrl)

	if result == nil || len(result) < 4 {
		return nil, errors.New("unknown remoteUrl pattern: " + remoteUrl)
	}
	hostNames := strings.Split(result[1], "@")
	serviceHost := hostNames[len(hostNames)-1]
	return &Remote{
		URL:         remoteUrl,
		ServiceHost: serviceHost,
		Owner:       result[2],
		RepoName:    strings.Replace(result[3], ".git", "", -1),
	}, nil
}

func GetDefaultRemote(path string) (*Remote, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	remote, err := r.Remote(git.DefaultRemoteName)
	if err != nil {
		return nil, err
	}
	return NewRemote(remote.Config().URLs[0]) // FIXME
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func GetDockerComposeFilePath(basePath string) (string, error) {
	filePath := filepath.Clean(basePath + "/docker-compose.yml")
	if !IsExist(filePath) {
		return filePath, nil
	}

	filePath = filepath.Clean(basePath + "/docker-compose.gored.yml")
	if !IsExist(filePath) {
		return filePath, nil
	}

	for i := 1; i < 1000; i++ {
		filePath := filepath.Clean(fmt.Sprintf(basePath+"/docker-compose.gored%d.yml", i))
		if !IsExist(filePath) {
			return filePath, nil
		}
	}
	return "", errors.New("creating docker-compose file name is failed")
}

func GetMakeFilePath(basePath string) (string, error) {
	filePath := filepath.Clean(basePath + "/Makefile")
	if !IsExist(filePath) {
		return filePath, nil
	}

	filePath = filepath.Clean(basePath + "/Makefile.gored")
	if !IsExist(filePath) {
		return filePath, nil
	}

	for i := 1; i < 1000; i++ {
		filePath := filepath.Clean(fmt.Sprintf(basePath+"/Makefile.gored%d", i))
		if !IsExist(filePath) {
			return filePath, nil
		}
	}

	return "", errors.New("creating Makefile name is failed")
}
