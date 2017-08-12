export UPLOAD_ROOT:=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

$(UPLOAD_ROOT)-codegen: tools-codegen
	$(BIN)/codegen -p clickyab.com/crab/modules/upload/model
	$(BIN)/codegen -p clickyab.com/crab/modules/upload/controllers

$(UPLOAD_ROOT)-migration: tools-go-bindata
	cd $(UPLOAD_ROOT)/migrations && $(BIN)/go-bindata -nometadata -o $(UPLOAD_ROOT)/migrations/migration.gen.go -nomemcopy=true -pkg=migrations ./db/...

$(UPLOAD_ROOT)-test:
	cd $(UPLOAD_ROOT)/ && $(GO) test ./...

.PHONY: $(UPLOAD_ROOT)-codegen $(USER_ROOT)-migration