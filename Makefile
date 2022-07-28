BINARY_NAME=team-divider
EXE =
ifeq ($(GOOS),windows)
EXE = .exe
endif

build: ## Build your project and put the output binary in out/bin/
	mkdir -p out/bin
	GO111MODULE=on go build  -o out/bin/$(BINARY_NAME)$(EXE) ./cmd/team-divider

clean: ## Remove build related file
	rm -fr ./out
