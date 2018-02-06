export TIME_ROOT:=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

$(TIME_ROOT)-codegen: tools-codegen
	rm -rf $(TIME_ROOT)/tim/*.gen.go
	rm -rf $(TIME_ROOT)/controllers/*.gen.go
	$(BIN)/codegen -p clickyab.com/crab/modules/time/tim
	$(BIN)/codegen -p clickyab.com/crab/modules/time/controllers

$(TIME_ROOT)-migration: tools-go-bindata
	rm -rf $(TIME_ROOT)/migrations/*.gen.go
	cd $(TIME_ROOT)/migrations && $(BIN)/go-bindata -nometadata -o $(TIME_ROOT)/migrations/migration.gen.go -nomemcopy=true -pkg=migrations ./db/...

$(TIME_ROOT)-test:
	cd $(TIME_ROOT)/ && $(GO) test ./...

.PHONY: $(TIME_ROOT)-codegen $(TIME_ROOT)-migration