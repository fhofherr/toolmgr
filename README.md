# toolmgr

`toolmgr` helps you manage go tools using a
[tools.go](https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module)
file.

## Usage

Create a `tools.go` file with at least the following content:

```go
//+build tools

package tools

import (
	_ "github.com/fhofherr/toolmgr"
)
```

Add imports for the used tools to the above files as required. Then make
sure to install `toolmgr`:

```sh
export GOBIN="$PWD/bin"
go mod tidy
go install github.com/fhofherr/toolmgr
"$PWD"/bin/toolmgr -bin-dir "$PWD/bin"
```

Whenever you add new tools to `tools.go` be sure to re-execute the above
steps.

### Makefile targets

In many cases it is convenient to use a `Makefile` to make the above
easier. Simply add the following targets to your `Makefile`.

```Makefile
GO ?= go

## Change this to your liking:
TOOLS_BIN_DIR ?= bin
TOOLS_GO := tools.go

.PHONY: tools
tools: $(TOOLS_BIN_DIR)

## Dummy target to allow depending on all tools
$(TOOLS_BIN_DIR): $(TOOLS_BIN_DIR)/toolmgr

$(TOOLS_BIN_DIR)/%: $(TOOLS_GO)
	GOBIN=$(abspath ./$(TOOLS_BIN_DIR)) $(GO) install github.com/fhofherr/toolmgr
	$(TOOLS_BIN_DIR)/toolmgr -bin-dir $(TOOLS_BIN_DIR) -tools-go $(TOOLS_GO)
```

Whenever one of your other `Makefile` targets depends on one or more
tools you can either depend on the `$(TOOLS_BIN_DIR)` target or on
`$(TOOTOOLS_BIN_DIR)/<toolname>`.

## License

Copyright © 2021 Ferdinand Hofherr

Distributed under the MIT License.
