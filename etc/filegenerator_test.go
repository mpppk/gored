package etc

import (
	"os"
	"path/filepath"
	"testing"
)

const testDirPath = "../testdata"

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func TestFileGenerator(t *testing.T) {
	cases := []struct {
		RepoPath               string
		DockerImageName        string
		DockerComposeFilePath  string
		MakefilePath           string
		CircleCIConfigFilePath string
		BuildPath              string
		VersionPath            string
	}{
		{
			RepoPath:               "..",
			DockerImageName:        filepath.Join(testDirPath, "test_docker_image_name"),
			DockerComposeFilePath:  filepath.Join(testDirPath, "docker-compose.yml"),
			MakefilePath:           filepath.Join(testDirPath, "Makefile"),
			CircleCIConfigFilePath: filepath.Join(testDirPath, ".circleci/config.yml"),
			BuildPath:              ".",
			VersionPath:            ".",
		},
	}

	for _, c := range cases {
		fg := &FileGenerator{
			RepoPath:               c.RepoPath,
			DockerImageName:        c.DockerImageName,
			DockerComposeFilePath:  c.DockerComposeFilePath,
			MakefilePath:           c.MakefilePath,
			CircleCIConfigFilePath: c.CircleCIConfigFilePath,
			BuildPath:              c.BuildPath,
			VersionPath:            c.VersionPath,
		}

		if err := fg.GenerateMakefile(); err != nil {
			t.Fatalf("Unexpected error occured when generate Makefile: %s", err)
		} else if !isExist(c.MakefilePath) {
			t.Fatalf("generated Makefile does not exist: %s", c.MakefilePath)
		}

		if err := fg.GenerateDockerComposeFile(); err != nil {
			t.Fatalf("Unexpected error occured when generate docker-compose file: %s", err)
		} else if !isExist(c.DockerComposeFilePath) {
			t.Fatalf("generated docker-compose does not exist: %s", c.DockerComposeFilePath)
		}

		if err := fg.GenerateCircleCIConfigFile(); err != nil {
			t.Fatalf("Unexpected error occured when generate circle ci file: %s", err)
		} else if !isExist(c.CircleCIConfigFilePath) {
			t.Fatalf("generated circle ci does not exist: %s", c.CircleCIConfigFilePath)
		}
	}
}

func removeFileAndDirs(rootDir string, fileAndDirs []string) error {
	for _, fad := range fileAndDirs {
		if err := os.RemoveAll(filepath.Join(rootDir, fad)); err != nil {
			return err
		}
	}
	return nil
}

func setup() {
	if _, err := os.Stat(testDirPath); err != nil {
		if err := os.Mkdir(testDirPath, 0755); err != nil {
			panic(err)
		}
	}
}

func tearDown() {
	files := []string{"Makefile", "docker-compose.yml", ".circleci"} // FIXME
	removeFileAndDirs(testDirPath, files)
}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	if ret == 0 {
		tearDown()
	}
	os.Exit(ret)
}
