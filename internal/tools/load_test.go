package tools_test

import (
	"path/filepath"
	"testing"

	"github.com/fhofherr/toolmgr/internal/tools"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name    string
		imports []string
		errMsg  string
	}{
		{
			name:    "valid",
			imports: []string{"golang.org/x/tools/cmd/stringer"},
		},
		{
			name:   "no imports",
			errMsg: "load: testdata/TestLoad/no_imports.go: no imports",
		},
		{
			name:   "missing",
			errMsg: "load: open testdata/TestLoad/missing.go: no such file or directory",
		},
		{
			name:   "invalid",
			errMsg: "load: testdata/TestLoad/invalid.go:1:1: expected 'package', found invalid",
		},
		{
			name:   "empty",
			errMsg: "load: testdata/TestLoad/empty.go:1:1: expected 'package', found 'EOF'",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			file := filepath.Join("testdata", t.Name()+".go")
			actual, err := tools.Load(file)
			if tt.errMsg != "" {
				assert.EqualError(t, err, tt.errMsg)
				return
			}
			assert.Equal(t, tt.imports, actual)
			assert.NoError(t, err)
		})
	}
}
