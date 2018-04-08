export AD_ROOT:=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

$(AD_ROOT)-codegen: tools-codegen
	rm -rf $(AD_ROOT)/orm/*.gen.go
	rm -rf $(AD_ROOT)/controllers/*.gen.go
	$(BIN)/codegen -p clickyab.com/crab/modules/ad/orm
	$(BIN)/codegen -p clickyab.com/crab/modules/ad/controllers

$(AD_ROOT)-migration: tools-go-bindata
	rm -rf $(AD_ROOT)/migrations/*.gen.go
	cd $(AD_ROOT)/migrations && $(BIN)/go-bindata -nometadata -o $(AD_ROOT)/migrations/migration.gen.go -nomemcopy=true -pkg=migrations ./db/...

$(AD_ROOT)-test:
	cd $(AD_ROOT)/ && $(GO) test ./...

.PHONY: $(AD_ROOT)-codegen $(AD_ROOT)-migration