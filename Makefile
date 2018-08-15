SHELL = /bin/bash
BASEDIR = $(shell pwd)

# env variables
# you should set these follow variables before exec `make build`
# GOPATH: go path
# PROJ_PATH: project path
PROJ_PATH = github.com/hzxiao
PROJ_NAME = taskmeter
#BRANCH: project git branch
BRANCH = master

PROJ_GIT_REPO = https://github.com/hzxiao/taskmeter.git
PROJ_FULL_NAME = $(GOPATH)/src/$(PROJ_PATH)/$(PROJ_NAME)
OUTPUT_DIR=build

# build with verison infos
versionDir = $(PROJ_PATH)/$(PROJ_NAME)/pkg/version
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags="-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"

vendor:
	go get "github.com/spf13/pflag"
test:
	go test ./...

build:
	@if test ! -d $(dir $(PROJ_FULL_NAME));then \
		echo "mkdir project path: "$(dir $(PROJ_FULL_NAME)); \
		mkdir -p $(dir $(PROJ_FULL_NAME)); \
	fi

	@if test ! -d $(PROJ_FULL_NAME); then \
		cd $(dir $(PROJ_FULL_NAME)); \
		git clone $(PROJ_GIT_REPO); \
	fi

	cd $(PROJ_FULL_NAME) 
	# git checkout .
	# git checkout $(BRANCH)
	# git pull 
	-rm -r $(OUTPUT_DIR) && \
	mkdir -p $(OUTPUT_DIR)/$(PROJ_NAME) && \
	mkdir -p $(OUTPUT_DIR)/$(PROJ_NAME)/conf && \
	mkdir -p $(OUTPUT_DIR)/$(PROJ_NAME)/script && \
	go build -v -ldflags ${ldflags} . && \
	cp -R conf/* $(OUTPUT_DIR)/$(PROJ_NAME)/conf && \
	cp -R script/* $(OUTPUT_DIR)/$(PROJ_NAME)/script && \
	cp $(PROJ_NAME) $(OUTPUT_DIR)/$(PROJ_NAME) && \
	cd $(OUTPUT_DIR) && \
	zip -r -q $(PROJ_NAME).zip $(PROJ_NAME)
fmt:
	@gofmt -w .

.PHONY: clean gotool ca help build