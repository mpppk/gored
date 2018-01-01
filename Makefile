CURRENT_REVISION = $(shell git rev-parse --short HEAD)
BUILD_LDFLAGS = "-X github.com/Songmu/goxz.revision=$(CURRENT_REVISION)"
REPO_OWNER = "mpppk"
REPO_NAME = "goreleng-test"
CURRENT_VERSION = $(shell gobump show -r ./cmd/)
ifdef update
  u=-u
endif

deps:
	dep ensure

devel-deps:
	go get ${u} github.com/golang/dep/cmd/dep
	go get ${u} github.com/motemen/gobump/cmd/gobump
	go get ${u} github.com/Songmu/goxz/cmd/goxz
	go get ${u} github.com/Songmu/ghch/cmd/ghch
	go get ${u} github.com/tcnksm/ghr

test: deps
	go test ./...

build: deps
	go build -ldflags=$(BUILD_LDFLAGS) ./cmd/goxz

crossbuild: deps test
	goxz -pv=v${CURRENT_VERSION} -d=./dist/v${CURRENT_VERSION} ./cmd

bump: deps test
	gobump patch -w ./cmd
	git commit -am "Checking in changes prior to tagging of version v$(CURRENT_VERSION)"
	git tag $(CURRENT_VERSION)
	git push "https://${GITHUB_TOKEN}@github.com/${REPO_OWNER}/${REPO_NAME}" HEAD:master

release: bump crossbuild
	ghr --username ${REPO_OWNER} ${CURRENT_VERSION} dist/v${CURRENT_VERSION}

.PHONY: bump test deps devel-deps crossbuild release