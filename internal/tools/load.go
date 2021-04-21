package tools

import (
	"fmt"
	"go/parser"
	"go/token"
	"strings"
)

// Load analyzes the package file belongs to and returns all import
// paths of this package.
func Load(file string) ([]string, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, parser.ImportsOnly)

	if err != nil {
		return nil, fmt.Errorf("load: %v", err)
	}
	if len(f.Imports) == 0 {
		return nil, fmt.Errorf("load: %s: no imports", file)
	}

	imports := make([]string, len(f.Imports))
	for i, v := range f.Imports {
		imports[i] = strings.ReplaceAll(v.Path.Value, "\"", "")
	}

	return imports, nil
}
