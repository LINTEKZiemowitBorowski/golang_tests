###############################################################################
#
#     Makefile for golang_tests
#
#     Copyright (c) 2015 LINTEK Ziemowit Borowski
#
###############################################################################

# User setup (if exists)
-include Makefile.user

# Set GOPATH, GOOS and GOARCH if not set in the Makefile.user
export GOPATH ?= $(CURDIR)
export GOOS ?= linux
export GOARCH ?= amd64

BENCHMARKS:= benchmark0 benchmark1 benchmark2 benchmark3 benchmark4 benchmark5 benchmark6 benchmark7 benchmark8

EXECUTABLES:=$(addprefix $(GOPATH)/bin/$(GOOS)_$(GOARCH)/,$(BENCHMARKS))

.PHONY: all

all: $(EXECUTABLES)

$(GOPATH)/bin/$(GOOS)_$(GOARCH)/%: $(GOPATH)/src/github.com/lintek/%/*
	go install -v github.com/lintek/$(notdir $@)

.PHONY: clean

clean:
	rm -rf bin pkg
