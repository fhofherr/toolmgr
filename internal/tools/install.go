package tools

import (
	"fmt"
	"os"
	"os/exec"
)

// InstallOption allows to configure the way Install installs tools.
type InstallOption func(*installOpts)

// Install installs passed tools to the caller's system.
//
// The installation process can be configured by passing various
// installation options.
func Install(tools []string, opts ...InstallOption) error {
	var instOpts installOpts

	for _, opt := range opts {
		opt(&instOpts)
	}

	for _, tool := range tools {
		cmd := instOpts.cmdForTool(tool)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("install %s: %v", tool, err)
		}
	}

	return nil
}

// WithInstallCommand allows to configure the command used to install the go tools.
//
// This is useful for testing purposes. For all other cases the default "go
// install <pkg-name>" should be used.
func WithInstallCommand(cmd []string) InstallOption {
	return func(opts *installOpts) {
		opts.cmd = append([]string(nil), cmd...)
	}
}

// WithInstallEnv allows to specify environment variables to set during
// the installation process.
func WithInstallEnv(vars map[string]string) InstallOption {
	return func(opts *installOpts) {
		if opts.env == nil {
			opts.env = make(map[string]string, len(vars))
		}
		for k, v := range vars {
			opts.env[k] = v
		}
	}
}

type installOpts struct {
	cmd []string
	env map[string]string
}

func (opts installOpts) cmdForTool(tool string) *exec.Cmd {
	var args []string

	args = append(args, opts.cmd...)
	if len(args) == 0 {
		args = append(args, "go", "install")
	}
	args = append(args, tool)
	cmd := exec.Command(args[0], args[1:]...) // nolint: gosec

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = os.Environ()
	for k, v := range opts.env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}

	return cmd
}
