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

BENCHMARKS:= benchmark00_float_arithmetic benchmark01_float_arithmetic benchmark02_int_arithmetic  \
 benchmark03_string_operations benchmark04_slice_operations benchmark05_map_operations \
 benchmark06_fibonacci benchmark07_crc benchmark08_sqlite benchmark09_sqlite benchmark10_http_download

EXECUTABLES:=$(addprefix $(GOPATH)/bin/$(GOOS)_$(GOARCH)/,$(BENCHMARKS))

.PHONY: all

all: $(EXECUTABLES)

$(GOPATH)/bin/$(GOOS)_$(GOARCH)/%: $(GOPATH)/src/github.com/lintek/%/*
	go install -v github.com/lintek/$(notdir $@)

.PHONY: clean

clean:
	rm -rf bin pkg
