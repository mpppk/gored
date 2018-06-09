REPO_OWNER = mpppk
REPO_NAME = gored
BUILD_PATH = .
VERSION_PATH = cmd/
SHELL = /bin/bash

ifdef update
  u=-u
endif

bindata:
	go-bindata -pkg etc -o etc/bindata.go tmpl/

deps:
	dep ensure

setup:
	go get ${u} github.com/golang/dep/cmd/dep
	go get ${u} github.com/Masterminds/glide
	go get ${u} github.com/motemen/gobump/cmd/gobump
	go get ${u} github.com/Songmu/goxz/cmd/goxz
	go get ${u} github.com/Songmu/ghch/cmd/ghch
	go get ${u} github.com/tcnksm/ghr

lint: deps
	gometalinter

test: deps
	go test ./...

coverage: deps
	go test -race -coverprofile=coverage.txt -covermode=atomic ./etc

codecov: deps coverage
	bash <(curl -s https://codecov.io/bash)

build: deps
	go build $(BUILD_PATH)

install: deps
	go install $(BUILD_PATH)

crossbuild: deps test
	goxz -pv=v`gobump show -r $(VERSION_PATH)` -d=./dist/v`gobump show -r $(VERSION_PATH)` $(BUILD_PATH)

bump-and-commit: deps test
	gobump patch -w $(VERSION_PATH)
	ghch -w -N v`gobump show -r $(VERSION_PATH)`
	git add CHANGELOG.md
	git commit -am "Checking in changes prior to tagging of version v`gobump show -r $(VERSION_PATH)` [skip ci]"
	git tag `gobump show -r $(VERSION_PATH)`
	git push "https://$(GITHUB_TOKEN)@github.com/$(REPO_OWNER)/$(REPO_NAME)" HEAD:master

release: crossbuild
	ghr --username $(REPO_OWNER) `gobump show -r $(VERSION_PATH)` dist/v`gobump show -r $(VERSION_PATH)`

circleci:
	circleci build -e GITHUB_TOKEN=$GITHUB_TOKEN

.PHONY: bindata deps setup lint test coverage codecov build install crossbuild bump-and-commit release circleci