export DOMAIN_ROOT=$(dir $(realpath $(dir $(lastword $(MAKEFILE_LIST)))))domain

$(DOMAIN_ROOT)-codegen: tools-codegen
	$(BIN)/codegen -p clickyab.com/crab/modules/domain/dmn
	$(BIN)/codegen -p clickyab.com/crab/modules/domain/controllers

$(DOMAIN_ROOT)-migration: tools-go-bindata
	cd $(DOMAIN_ROOT)/migrations && $(BIN)/go-bindata -nometadata -o $(DOMAIN_ROOT)/migrations/migration.gen.go -nomemcopy=true -pkg=migrations ./db/...

	echo $(DOMAIN_ROOT)
.PHONY: $(DOMAIN_ROOT)-codegen $(DOMAIN_ROOT)-migration

