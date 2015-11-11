###############################################################################
#
#     Makefile for goexamples
#
#     Copyright (c) 2015 LINTEK Ziemowit Borowski
#
###############################################################################

# User setup (if exists)
-include Makefile.user

# export GOOS:=linux
# export GOARCH:=arm
# export GOARM=7
# export GOROOT:=/usr/lib/go
# export GOPATH:=/media/devel_disk/projects/Sanbox/goexample

BENCHMARK0_SRCS:=src/github.com/lintek/benchmark0/benchmark0.go
BENCHMARK0_BIN:=pkg/$(GOOS)_$(GOARCH)/github.com/lintek/benchmark0

BENCHMARK1_SRCS:=src/github.com/lintek/benchmark1/benchmark1.go
BENCHMARK1_BIN:=pkg/$(GOOS)_$(GOARCH)/github.com/lintek/benchmark1

BENCHMARK2_SRCS:=src/github.com/lintek/benchmark2/benchmark2.go
BENCHMARK2_BIN:=pkg/$(GOOS)_$(GOARCH)/github.com/lintek/benchmark2

BENCHMARK3_SRCS:=src/github.com/lintek/benchmark3/benchmark3.go
BENCHMARK3_BIN:=pkg/$(GOOS)_$(GOARCH)/github.com/lintek/benchmark3

BENCHMARK4_SRCS:=src/github.com/lintek/benchmark4/benchmark4.go
BENCHMARK4_BIN:=pkg/$(GOOS)_$(GOARCH)/github.com/lintek/benchmark4

all: $(BENCHMARK0_BIN) $(BENCHMARK1_BIN) $(BENCHMARK2_BIN) $(BENCHMARK3_BIN) $(BENCHMARK4_BIN)

$(BENCHMARK0_BIN): $(BENCHMARK0_SRCS)
	go install -v github.com/lintek/benchmark0

$(BENCHMARK1_BIN): $(BENCHMARK1_SRCS)
	go install -v github.com/lintek/benchmark1

$(BENCHMARK2_BIN): $(BENCHMARK2_SRCS)
	go install -v github.com/lintek/benchmark2

$(BENCHMARK3_BIN): $(BENCHMARK3_SRCS)
	go install -v github.com/lintek/benchmark3

$(BENCHMARK4_BIN): $(BENCHMARK4_SRCS)
	go install -v github.com/lintek/benchmark4

.PHONY: clean

clean:
	rm -rf bin pkg

