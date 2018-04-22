export DATAWORKER_ROOT:=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

$(DATAWORKER_ROOT)-codegen: tools-codegen
	rm -rf $(DATAWORKER_ROOT)/orm/*.gen.go
	rm -rf $(DATAWORKER_ROOT)/workers/*.gen.go
	$(BIN)/codegen -p clickyab.com/crab/modules/dataworker/orm
	$(BIN)/codegen -p clickyab.com/crab/modules/dataworker/workers

$(DATAWORKER_ROOT)-migration: tools-go-bindata
	rm -rf $(DATAWORKER_ROOT)/migrations/*.gen.go
	cd $(DATAWORKER_ROOT)/migrations && $(BIN)/go-bindata -nometadata -o $(DATAWORKER_ROOT)/migrations/migration.gen.go -nomemcopy=true -pkg=migrations ./db/...

$(DATAWORKER_ROOT)-test:
	cd $(DATAWORKER_ROOT)/ && $(GO) test ./...

.PHONY: $(DATAWORKER_ROOT)-codegen $(DATAWORKER_ROOT)-migration