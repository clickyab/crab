MISC_ROOT:=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

$(MISC_ROOT)-codegen:
	@echo "Nothing to do here"

$(MISC_ROOT)-migration:
	@echo "Nothing to do here"

$(MISC_ROOT)-test:
	cd $(MISC_ROOT)/ && $(GO) test ./...