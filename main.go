package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fhofherr/toolmgr/internal/tools"
)

var (
	toolsGo = flag.String("tools-go", "tools.go", "Name of the tool management file.")
	binDir  = flag.String("bin-dir", "bin", "Path to the directory where binaries should be installed.")
)

func fatalf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", a...)
	os.Exit(1)
}

func main() {
	flag.Parse()

	ts, err := tools.Load(*toolsGo)
	if err != nil {
		fatalf("Parse %s: %v", *toolsGo, err)
	}

	gobin := *binDir
	if !filepath.IsAbs(gobin) {
		wd, err := os.Getwd()
		if err != nil {
			fatalf("Get working directory: %v", err)
		}
		gobin = filepath.Join(wd, gobin)
	}
	gobin = filepath.Clean(gobin)

	if err := tools.Install(ts, tools.WithInstallEnv(map[string]string{"GOBIN": gobin})); err != nil {
		fatalf("Install tools: %v", err)
	}
}
