
# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

all: build

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

help:  ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

lint:  ## Lint the files
ifeq (, $(shell which golangci-lint))
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN)
endif
	golangci-lint run -v
ifeq (, $(shell which staticcheck))
	GO111MODULE=on go install honnef.co/go/tools/cmd/staticcheck@latest
endif
	staticcheck ./...

fmt:  ## Run go fmt against code.
	go fmt ./...

tidy:  ## Run go mod tidy
	go mod tidy

vet:  ## Run go vet against code.
	go vet ./...
	
##@ Build

build: fmt vet tidy lint  ## Build manager binary.
	go build -o bin/gn cmd/main.go

run: fmt vet  ## Run a controller from your host.
	go run ./cmd/main.go

clean:  ## delete the bin folder containing binaries
	rm -rf $(PROJECT_DIR)/bin

