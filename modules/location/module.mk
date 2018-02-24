export LOCATION_ROOT:=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

$(LOCATION_ROOT)-codegen: tools-codegen
	rm -rf $(LOCATION_ROOT)/location/*.gen.go
	rm -rf $(LOCATION_ROOT)/controllers/*.gen.go
	$(BIN)/codegen -p clickyab.com/crab/modules/location/location
	$(BIN)/codegen -p clickyab.com/crab/modules/location/controllers

$(LOCATION_ROOT)-migration: tools-go-bindata
	cd $(LOCATION_ROOT)/migrations && $(BIN)/go-bindata -nometadata -o $(LOCATION_ROOT)/migrations/migration.gen.go -nomemcopy=true -pkg=migrations ./db/...

$(LOCATION_ROOT)-test:
	cd $(LOCATION_ROOT)/ && $(GO) test ./...

.PHONY: $(LOCATION_ROOT)-codegen $(USER_ROOT)-migration