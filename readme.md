gored is release engineering (on docker) tool for golang project

[![CircleCI](https://circleci.com/gh/mpppk/gored.svg?style=svg)](https://circleci.com/gh/mpppk/gored) [![codebeat badge](https://codebeat.co/badges/5e37d2ec-d6d0-4cfb-9c71-5bbcbe22c63e)](https://codebeat.co/projects/github-com-mpppk-gored-master)

# Usage

```Shell
$ cd gopath/src/github.com/yourname/yourrepo
$ docker run -v $(pwd):/init mpppk/gored     # create files for execute release engineering tasks
$ docker-compose run bump-and-commit         # update version and changelog
$ docker-compose run release                 # upload binaries to GitHub releases
```

Now, changelog and version are updated. 
[![Image](https://gyazo.com/0e5f0377caed3747c61401ab00969089/thumb/1000)](https://gyazo.com/0e5f0377caed3747c61401ab00969089)

and, binaries are uploaded to GitHub releases.
[![Image](https://gyazo.com/777ecc0fcf2dcd6aa91967f0347ee459/thumb/1000)](https://gyazo.com/777ecc0fcf2dcd6aa91967f0347ee459)<br>

## Initialize
The first command `docker run -v $(pwd):/init mpppk/gored` creates below files.
- `docker-compose.yml`
- `Makefile`
- `.circleci/config.yml`

(for more details, check gored's own files)

If go file that including version variable is not located in root directory of project, you can execute below command with `--build-path` instead.

```Shell
$ docker run -v $(pwd):/init mpppk/gored gored init --build-path [build path]
```

### Note
You can use `gored` command to initialize instead.

```Shell
$ go get github.com/mpppk/gored
$ gored init --build-path [build path]
```

## Run tasks on docker container

`docker-compose run [task name]` execute same name target of `Makefile`.  
for example, here is `docker-compose.yml` contents related to `docker-compose bump-and-commit`. 

```yaml
   bump-and-commit:
     image: mpppk/gored
     volumes:
       - $PWD:/go/src/github.com/yourname/yourrepo
     environment:
       GITHUB_TOKEN: $GITHUB_TOKEN
     working_dir: /go/src/github.com/yourname/yourrepo
     command: make bump-and-commit -f Makefile
```

and Makefile is

```Makefile
deps:
	dep ensure

test: deps
	go test ./...

bump-and-commit: deps test
 	gobump patch -w $(BUILD_PATH)
 	ghch -w -N v`gobump show -r $(BUILD_PATH)`
 	git add CHANGELOG.md
 	git commit -am "Checking in changes prior to tagging of version v`gobump show -r $(BUILD_PATH)` [skip ci]"
 	git tag `gobump show -r $(BUILD_PATH)`
 	git push "https://$(GITHUB_TOKEN)@github.com/$(REPO_OWNER)/$(REPO_NAME)" HEAD:master
```

### Frequency used tasks

* bump-and-commit
    * Download dependencies by `deps` target
    * Test by `test` target
    * Update version by [motemen/gobump](https://github.com/motemen/gobump)
    * Update CHANGELOG.md by [Songmu/ghch](https://github.com/Songmu/ghch)
    * Commit & push above files with version tag
* crossbuild
    * Execute `deps` and `test` target
    * Build binaries for windows, mac, linux by [Songmu/goxz](https://github.com/Songmu/goxz)
* release
    * Execute `deps`, `test` and `crossbuild` target
    * Upload binaries to GitHub releases by [tcnksm/ghr](https://github.com/tcnksm/ghr)

## Automatic release via Circle CI
gored generate `.circleci/config.yml`.
Once you commit this file, the project will be released automatically each time you push to master branch.

## Modify task and workflow
You can modify any task and workflow by update Makefile.
for example, if you want to use [Masterminds/glide](https://github.com/Masterminds/glide) instead of [golang/dep](https://github.com/golang/dep), change `deps` target as below.

```Makefile
deps:
	glide install
```

## Using non builtin tools
Default docker image(mpppk/gored) includes below tools
* [golang/dep](https://github.com/golang/dep)
* [Masterminds/glide](https://github.com/Masterminds/glide)
* [motemen/gobump](https://github.com/motemen/gobump)
* [Songmu/ghch](https://github.com/Songmu/ghch)
* [Songmu/goxz](https://github.com/Songmu/goxz)
* [tcnksm/ghr](https://github.com/tcnksm/ghr)

If you want to use other tools, it may be better to create your own Dockerfile.

Makefile

```Makefile
setup:
	go get ${u} github.com/golang/dep/cmd/dep
	go get (other tool)
```

Dockerfile

```Dockerfile
FROM golang:1.9.2 AS builder
COPY Makefile /go/src/github.com/yourname/yourrepo/Makefile
WORKDIR /go/src/github.com/yourname/yourrepo
RUN make setup

FROM circleci/golang:1.9.2
COPY --from=builder /go/bin/* /go/bin/
RUN git config --global user.email "gored-bot@dammy.com"
RUN git config --global user.name  "gored-bot"
```

docker-compose.yml

```yaml
services:
  image: yourname/image
```
