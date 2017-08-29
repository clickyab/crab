export ASSET_ROOT:=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

$(ASSET_ROOT)-codegen: tools-codegen
	$(BIN)/codegen -p clickyab.com/crab/modules/asset/asst
	$(BIN)/codegen -p clickyab.com/crab/modules/asset/controllers

$(ASSET_ROOT)-migration: tools-go-bindata
	cd $(ASSET_ROOT)/migrations && $(BIN)/go-bindata -nometadata -o $(ASSET_ROOT)/migrations/migration.gen.go -nomemcopy=true -pkg=migrations ./db/...

$(ASSET_ROOT)-test:
	cd $(ASSET_ROOT)/ && $(GO) test ./...

.PHONY: $(ASSET_ROOT)-codegen $(ASSET_ROOT)-migration