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

BENCHMARKS:= benchmark0_float_arithmetic benchmark1_float_arithmetic benchmark2_int_arithmetic benchmark3_sqlite \
 benchmark4_sqlite benchmark5_fibonacci benchmark6_slice_operations benchmark7_map_operations benchmark8_http_download

EXECUTABLES:=$(addprefix $(GOPATH)/bin/$(GOOS)_$(GOARCH)/,$(BENCHMARKS))

.PHONY: all

all: $(EXECUTABLES)

$(GOPATH)/bin/$(GOOS)_$(GOARCH)/%: $(GOPATH)/src/github.com/lintek/%/*
	go install -v github.com/lintek/$(notdir $@)

.PHONY: clean

clean:
	rm -rf bin pkg
