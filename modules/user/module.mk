export USER_ROOT:=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

$(USER_ROOT)-codegen: tools-codegen
	$(BIN)/codegen -p clickyab.com/crab/modules/user/models
	$(BIN)/codegen -p clickyab.com/crab/modules/user/controllers

$(USER_ROOT)-migration: tools-go-bindata
	cd $(USER_ROOT)/migrations && $(BIN)/go-bindata -nometadata -o $(USER_ROOT)/migrations/migration.gen.go -nomemcopy=true -pkg=migrations ./db/...

$(USER_ROOT)-test:
	cd $(USER_ROOT)/ && $(GO) test ./...

.PHONY: $(USER_ROOT)-codegen $(USER_ROOT)-migration