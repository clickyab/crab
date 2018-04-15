export LINTER=$(BIN)/gometalinter.v2
# TODO : Ignoring services/codegen is a bad thing. try to get it back to lint
export LINTERCMD=$(LINTER) -e ".*.gen.go" -e ".*_test.go" -e "clickyab.com/crab/vendor/.*" --cyclo-over=15 --line-length=120 --deadline=170s --disable-all --enable=structcheck --enable=deadcode --enable=gocyclo --enable=ineffassign --enable=golint --enable=goimports --enable=errcheck --enable=varcheck --enable=goconst --enable=gosimple --enable=staticcheck --enable=unused --enable=misspell

lint: $(LINTER)
	$(LINTERCMD) $(ROOT)/cmd/...
	$(LINTERCMD) $(ROOT)/modules/...

metalinter:
	$(GO) get -v gopkg.in/alecthomas/gometalinter.v2
	$(GO) install -v gopkg.in/alecthomas/gometalinter.v2
	$(LINTER) --install

$(LINTER):
	@[ -f $(LINTER) ] || make -f $(ROOT)/Makefile metalinter