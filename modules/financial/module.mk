export FINANCIAL_ROOT:=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

$(FINANCIAL_ROOT)-codegen: tools-codegen
	rm -rf $(FINANCIAL_ROOT)/orm/*.gen.go
	rm -rf $(FINANCIAL_ROOT)/controllers/*.gen.go
	$(BIN)/codegen -p clickyab.com/crab/modules/financial/orm
	$(BIN)/codegen -p clickyab.com/crab/modules/financial/controllers

$(FINANCIAL_ROOT)-migration: tools-go-bindata
	rm -rf $(FINANCIAL_ROOT)/migrations/*.gen.go
	cd $(FINANCIAL_ROOT)/migrations && $(BIN)/go-bindata -nometadata -o $(FINANCIAL_ROOT)/migrations/migration.gen.go -nomemcopy=true -pkg=migrations ./db/...

$(FINANCIAL_ROOT)-test:
	cd $(FINANCIAL_ROOT)/ && $(GO) test ./...

.PHONY: $(FINANCIAL_ROOT)-codegen $(USER_ROOT)-migration