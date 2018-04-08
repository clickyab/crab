export DOMAIN_ROOT:=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

$(DOMAIN_ROOT)-codegen: tools-codegen
	rm -rf $(DOMAIN_ROOT)/orm/*.gen.go
	rm -rf $(DOMAIN_ROOT)/controllers/*.gen.go
	$(BIN)/codegen -p clickyab.com/crab/modules/domain/orm
	$(BIN)/codegen -p clickyab.com/crab/modules/domain/controllers

$(DOMAIN_ROOT)-migration: tools-go-bindata
	rm -rf $(DOMAIN_ROOT)/migrations/*.gen.go
	cd $(DOMAIN_ROOT)/migrations && $(BIN)/go-bindata -nometadata -o $(DOMAIN_ROOT)/migrations/migration.gen.go -nomemcopy=true -pkg=migrations ./db/...

$(DOMAIN_ROOT)-test:
	cd $(DOMAIN_ROOT)/ && $(GO) test ./...

.PHONY: $(DOMAIN_ROOT)-codegen $(DOMAIN_ROOT)-migration

