export PUB_ROOT:=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

$(PUB_ROOT)-codegen: tools-codegen
	$(BIN)/codegen -p clickyab.com/crab/modules/publisher/pub
	$(BIN)/codegen -p clickyab.com/crab/modules/publisher/controllers

$(PUB_ROOT)-migration: tools-go-bindata
	cd $(PUB_ROOT)/migrations && $(BIN)/go-bindata -nometadata -o $(PUB_ROOT)/migrations/migration.gen.go -nomemcopy=true -pkg=migrations ./db/...

$(PUB_ROOT)-test:
	cd $(PUB_ROOT)/ && $(GO) test ./...

.PHONY: $(PUB_ROOT)-codegen $(PUB_ROOT)-migration