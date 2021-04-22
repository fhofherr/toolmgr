package tools_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/fhofherr/toolmgr/internal/tools"
	"github.com/stretchr/testify/assert"
)

var mockInstallCommand = []string{os.Args[0], "-test.run=TestInstall_HelperProcess", "--"}

const (
	envVarHelperProcess = "TOOLMGR_TEST_INSTALL_HELPER_PROCESS"
	envVarExpectPkg     = "TOOLMGR_TEST_INSTALL_EXPECT_PKG"
	envVarExitCode      = "TOOLMGR_TEST_INSTALL_EXIT_CODE"
)

func TestInstall_Success(t *testing.T) {
	pkg := "golang.org/x/tools/cmd/stringer"
	err := tools.Install(
		[]string{pkg},
		tools.WithInstallCommand(mockInstallCommand),
		tools.WithInstallEnv(map[string]string{
			envVarHelperProcess: "1",
			envVarExpectPkg:     pkg,
		}),
	)
	assert.NoError(t, err)
}

func TestInstall_Failure(t *testing.T) {
	err := tools.Install(
		[]string{"golang.org/x/tools/cmd/stringer"},
		tools.WithInstallCommand(mockInstallCommand),
		tools.WithInstallEnv(map[string]string{
			envVarHelperProcess: "1",
			envVarExitCode:      "1",
		}),
	)
	assert.EqualError(t, err, "install golang.org/x/tools/cmd/stringer: exit status 1")
}

func TestInstall_HelperProcess(t *testing.T) {
	if os.Getenv(envVarHelperProcess) == "" {
		return
	}

	if exp := os.Getenv(envVarExpectPkg); exp != "" {
		act := os.Args[len(os.Args)-1]
		assert.Equal(t, exp, act)
	}

	ecStr := os.Getenv(envVarExitCode)
	if ecStr == "" {
		return
	}
	ec, err := strconv.Atoi(ecStr)
	if !assert.NoError(t, err) {
		return
	}
	os.Exit(ec)
}
