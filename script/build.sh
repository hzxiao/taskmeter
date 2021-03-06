#!/usr/bin/env bash

set -e
PROJ_DIR=github.com/hzxiao
PROJ_NAME=taskmeter
BRANCH=master
PROJ_GIT_REPO=https://github.com/hzxiao/taskmeter.git
PROJ_FULL_DIR=$GOPATH/src/$PROJ_DIR
PROJ_FULL_NAME=$PROJ_FULL_DIR/$PROJ_NAME
OUTPUT_DIR=build

# build with verison infos
versionDir=$PROJ_DIR/$PROJ_NAME/pkg/version
gitTag=$(if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate=$(TZ=Asia/Shanghai date +%FT%T%z)
gitCommit=$(git log --pretty=format:'%H' -n 1)
gitTreeState=$(if git status|grep -q 'clean';then echo clean; else echo dirty; fi)
#
#ldflags="-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"


if [ ! -d ${PROJ_FULL_DIR} ];then
	echo "mkdir project dir: "$PROJ_FULL_DIR;
	mkdir -p $PROJ_FULL_DIR;
fi


if [ ! -d ${PROJ_FULL_NAME} ]; then
	cd $PROJ_FULL_DIR
	git clone $PROJ_GIT_REPO
fi

cd $PROJ_FULL_NAME
git checkout .
git checkout $BRANCH
git pull

swag init

rm -rf $OUTPUT_DIR
mkdir -p $OUTPUT_DIR/$PROJ_NAME
mkdir -p $OUTPUT_DIR/$PROJ_NAME/conf
mkdir -p $OUTPUT_DIR/$PROJ_NAME/script
mkdir -p $OUTPUT_DIR/$PROJ_NAME/docs

go build -v -ldflags "-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}" .
cp -R conf/* $OUTPUT_DIR/$PROJ_NAME/conf
cp -R script/* $OUTPUT_DIR/$PROJ_NAME/script
cp -R docs/* $OUTPUT_DIR/$PROJ_NAME/docs
mv $PROJ_NAME $OUTPUT_DIR/$PROJ_NAME
cp Makefile $OUTPUT_DIR/$PROJ_NAME
cd $OUTPUT_DIR
zip -r -q $PROJ_NAME.zip $PROJ_NAME

if [ -n "${SAVE_PKG_DIR}" ]; then
    cp $PROJ_NAME.zip ${SAVE_PKG_DIR}
fi
