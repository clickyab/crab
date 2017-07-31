tools-migration:
	$(BUILD) clickyab.com/crab/cmd/migration

mig-up: tools-migration
	$(BIN)/migration --action=up

mig-down: tools-migration
	$(BIN)/migration --action=down