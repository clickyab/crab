export NOTIF_ROOT=$(dir $(realpath $(dir $(lastword $(MAKEFILE_LIST)))))notification

$(NOTIF_ROOT)-codegen: tools-codegen
	rm -rf $(NOTIF_ROOT)/notif/*.gen.go
	$(BIN)/codegen -p clickyab.com/crab/modules/notification/notif

$(NOTIF_ROOT)-migration: tools-go-bindata
	rm -rf $(NOTIF_ROOT)/migrations/*.gen.go
	cd $(NOTIF_ROOT)/migrations && $(BIN)/go-bindata -nometadata -o $(NOTIF_ROOT)/migrations/migration.gen.go -nomemcopy=true -pkg=migrations ./db/...

$(NOTIF_ROOT)-test:

.PHONY: $(NOTIF_ROOT)-codegen $(NOTIF_ROOT)-migration