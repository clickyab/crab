export ROOT=$(realpath $(dir $(firstword $(MAKEFILE_LIST))))
export APP_NAME=crab
export DEFAULT_PASS=bita123
export GO=$(shell which go)
export GIT=$(shell which git)
export BIN=$(ROOT)/bin
export GOPATH=$(abspath $(ROOT)/../../..)
export GOBIN?=$(BIN)
export DB_PASS?=$(DEFAULT_PASS)
export DB_USER?=$(APP_NAME)
export DB_NAME?=$(APP_NAME)
export R_USER?=$(APP_NAME)
export R_PASS?=$(DEFAULT_PASS)
export WORK_DIR=$(ROOT)/tmp
export LONG_HASH?=$(shell git log -n1 --pretty="format:%H" | cat)
export SHORT_HASH?=$(shell git log -n1 --pretty="format:%h"| cat)
export COMMIT_DATE?=$(shell git log -n1 --date="format:%D-%H-%I-%S" --pretty="format:%cd"| sed -e "s/\//-/g")
export COMMIT_COUNT?=$(shell git rev-list HEAD --count| cat)
export BUILD_DATE=$(shell date "+%D/%H/%I/%S"| sed -e "s/\//-/g")
export FLAGS="-X version.hash=$(LONG_HASH) -X version.short=$(SHORT_HASH) -X version.date=$(COMMIT_DATE) -X version.count=$(COMMIT_COUNT) -X version.build=$(BUILD_DATE)"
export LD_ARGS=-ldflags $(FLAGS)
export BUILD=cd $(ROOT) && $(GO) install -v $(LD_ARGS)
export DATABASE_URL=postgres://$(DB_USER):$(DB_PASS)@localhost/$(DB_NAME)
export SWAGGER_TITLE=The clickyab.com/crab API
export SWAGGER_DESCRIPTION=The crab api, auto generated
export SWAGGER_VERSION=$(COMMIT_COUNT)

all:
	GOPATH=$(GOPATH) $(BUILD) clickyab.com/crab/cmd/...

all-gen: codegen migration all

# Include all files in the scripts folder
include $(wildcard $(ROOT)/scripts/*.mk)

# Include all modules builder
include $(wildcard $(ROOT)/modules/*/module.mk)

# codegen target as wildcard target
# TODO : remove "| true" when the rotes are added
codegen: $(addsuffix -codegen,$(wildcard $(ROOT)/modules/*))
	cp $(ROOT)/tmp/swagger/out.json $(ROOT)/swagger-ui/index.json | true
	cp $(ROOT)/tmp/swagger/out.yaml $(ROOT)/swagger-ui/index.yaml | true

# Migration as a wildcard target
migration: $(addsuffix -migration,$(wildcard $(ROOT)/modules/*))

build-webserver:
	$(BUILD) clickyab.com/crab/cmd/webserver

run-webserver: build-webserver
	$(BIN)/webserver

test: $(addsuffix -test,$(wildcard $(ROOT)/modules/*))
	echo "Done"
