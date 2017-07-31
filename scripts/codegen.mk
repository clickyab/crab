tools-codegen:
	$(BUILD) clickyab.com/crab/cmd/codegen

tools-go-bindata:
	GOBIN=$(BIN) $(GO) get -v github.com/jteeuwen/go-bindata/go-bindata
	GOBIN=$(BIN) $(GO) install -v github.com/jteeuwen/go-bindata/go-bindata