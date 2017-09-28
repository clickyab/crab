export CAMPAIGN_ROOT:=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

$(CAMPAIGN_ROOT)-codegen: tools-codegen
	$(BIN)/codegen -p clickyab.com/crab/modules/campaign/models
	$(BIN)/codegen -p clickyab.com/crab/modules/campaign/controllers

$(CAMPAIGN_ROOT)-migration: tools-go-bindata
	cd $(CAMPAIGN_ROOT)/migrations && $(BIN)/go-bindata -nometadata -o $(CAMPAIGN_ROOT)/migrations/migration.gen.go -nomemcopy=true -pkg=migrations ./db/...

$(CAMPAIGN_ROOT)-test:
	cd $(CAMPAIGN_ROOT)/ && $(GO) test ./...

.PHONY: $(CAMPAIGN_ROOT)-codegen $(USER_ROOT)-migration