APP 			:= UserInfo
BUILD_VERSION   := 1.0.0
BUILD_TIME      := $(shell date "+%F %T")
BUILD_TAG       := $(shell date "+%Y%m%d-%H%M")
GO_VERSION      := $(shell go version)
COMMIT_BRANCH   := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT_SHA1     := $(shell git rev-parse HEAD )

BUILD_DIR = build
TARGET_PATH = ${BUILD_DIR}/bin
TARGET_NAME = user-go

OS=darwin
ARCH=arm64


build:
	GOOS=${OS} GOARCH=${ARCH} go build -ldflags  \
        "                                        \
        -X 'user-go/cmd.App=${APP}'        \
        -X 'user-go/cmd.BuildTag=${BUILD_TAG}'        \
        -X 'user-go/cmd.GoVersion=${GO_VERSION}'        \
        -X 'user-go/cmd.BuildVersion=${BUILD_VERSION}'  \
        -X 'user-go/cmd.BuildTime=${BUILD_TIME}'        \
        -X 'user-go/cmd.CommitBranch=${COMMIT_BRANCH}'  \
        -X 'user-go/cmd.CommitID=${COMMIT_SHA1}'        \
        "                                        \
    	-o ${TARGET_PATH}/${TARGET_NAME}

clean:
	rm -r build/