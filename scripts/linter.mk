export LINTER=$(BIN)/gometalinter.v1
# TODO : Ignoring services/codegen is a bad thing. try to get it back to lint
export LINTERCMD=$(LINTER) -e ".*.gen.go" -e "clickyab.com/crab/.*"  -e ".*_test.go" -e "clickyab.com/crab/vendor/.*" --cyclo-over=19 --line-length=120 --deadline=100s --disable-all --enable=structcheck --enable=deadcode --enable=gocyclo --enable=ineffassign --enable=golint --enable=goimports --enable=errcheck --enable=varcheck --enable=goconst --enable=gosimple --enable=staticcheck --enable=unused --enable=misspell

lint: $(LINTER)
	$(LINTERCMD) $(ROOT)/cmd/...
	$(LINTERCMD) $(ROOT)/modules/...

metalinter:
	$(GO) get -v gopkg.in/alecthomas/gometalinter.v1
	$(GO) install -v gopkg.in/alecthomas/gometalinter.v1
	$(LINTER) --install

$(LINTER):
	@[ -f $(LINTER) ] || make -f $(ROOT)/Makefile metalinter