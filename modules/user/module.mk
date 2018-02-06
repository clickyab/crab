export USER_ROOT:=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

$(USER_ROOT)-codegen: tools-codegen
	rm -rf $(USER_ROOT)/aaa/*.gen.go
	rm -rf $(USER_ROOT)/controllers/*.gen.go
	$(BIN)/codegen -p clickyab.com/crab/modules/user/aaa
	$(BIN)/codegen -p clickyab.com/crab/modules/user/controllers

$(USER_ROOT)-migration: tools-go-bindata
	rm -rf $(USER_ROOT)/migrations/*.gen.go
	cd $(USER_ROOT)/migrations && $(BIN)/go-bindata -nometadata -o $(USER_ROOT)/migrations/migration.gen.go -nomemcopy=true -pkg=migrations ./db/...

$(USER_ROOT)-test:
	cd $(USER_ROOT)/ && $(GO) test ./...

.PHONY: $(USER_ROOT)-codegen $(USER_ROOT)-migration