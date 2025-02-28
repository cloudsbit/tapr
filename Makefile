SHELL:=/bin/sh

# Utilities
RM          ?= $(shell which rm)
CP          ?= $(shell which cp)
MKDIR       ?= $(shell which mkdir)
CHMOD       ?= $(shell which chmod)
ECHO        ?= $(shell which echo)
GO          ?= $(shell which go)
OS          ?= linux
OSBIT       ?= 64
DSTPATH     ?= /usr/local/tapr

export GOARCH=amd64
#export PATH=$PATH:/usr/bin:/usr/local/go/bin
export GO111MODULE=on
export CGO_ENABLED=0
#export GOPROXY=https://goproxy.io,direct
export GOPROXY=https://goproxy.cn,direct

# Path Related
MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MKFILE_DIR  := $(dir $(MKFILE_PATH))
RELEASE_DIR := ${MKFILE_DIR}build/bin
INSTALL_DIR := $(DSTPATH)

# Date time
DATETIME := $(shell date +"%Y%m%d%H%M%S")
# Build Time
BUILD_TIME :=$(shell date +"%Y-%m-%d %H:%M:%S")

# go source files, ignore vendor directory
SOURCE = $(shell find ${MKFILE_DIR} -type f -name "*.go")

all: tapr tapradm taprd

prepare:
	@$(ECHO) "Compile GOOS=$(GOOS), GOARCH=$(GOARCH), CGO_ENABLED=$(CGO_ENABLED)..."
	$(GO) mod tidy
	@$(MKDIR) -p ${RELEASE_DIR}
	@$(MKDIR) -p ${INSTALL_DIR}

tapr: prepare
	$(GO) build -o ${RELEASE_DIR}/$@ -ldflags "-s -w" ${MKFILE_DIR}cmd/tapr

tapradm: prepare
	$(GO) build -o ${RELEASE_DIR}/$@ -ldflags "-s -w" ${MKFILE_DIR}cmd/tapradm

taprd: prepare
	$(GO) build -o ${RELEASE_DIR}/$@ -ldflags "-s -w" ${MKFILE_DIR}cmd/taprd

build: all

install: all
	@$(CP) -rf ${RELEASE_DIR}/conf/* ${INSTALL_DIR}/
	@$(CP) -rf ${RELEASE_DIR}/*      ${INSTALL_DIR}/

clean:
	@$(RM) -rf ${RELEASE_DIR}/*
	@$(ECHO) Tapr Cleaning completed

.PHONY: all build clean
