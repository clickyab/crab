export INVENTORY_ROOT:=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

$(INVENTORY_ROOT)-codegen: tools-codegen
	rm -rf $(INVENTORY_ROOT)/orm/*.gen.go
	rm -rf $(INVENTORY_ROOT)/controllers/*.gen.go
	$(BIN)/codegen -p clickyab.com/crab/modules/inventory/orm
	$(BIN)/codegen -p clickyab.com/crab/modules/inventory/controllers

$(INVENTORY_ROOT)-migration: tools-go-bindata
	rm -rf $(INVENTORY_ROOT)/migrations/*.gen.go
	cd $(INVENTORY_ROOT)/migrations && $(BIN)/go-bindata -nometadata -o $(INVENTORY_ROOT)/migrations/migration.gen.go -nomemcopy=true -pkg=migrations ./db/...

$(INVENTORY_ROOT)-test:
	cd $(INVENTORY_ROOT)/ && $(GO) test ./...

.PHONY: $(INVENTORY_ROOT)-codegen $(USER_ROOT)-migration
