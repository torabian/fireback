package workspaces

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/stretchr/testify/assert"
)

func RunExecCmd(exePath string, args []string) (*string, *string, error) {
	cmd := exec.Command(exePath, args...)

	return RunExec(cmd, exePath, args)
}

func RunExec(cmd *exec.Cmd, exePath string, args []string) (*string, *string, error) {

	// Capture command output
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Start the command
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		return nil, nil, err
	}

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		ferr := stderr.String()
		fout := stdout.String()
		return &fout, &ferr, err
	}

	ferr := stderr.String()
	fout := stdout.String()
	return &fout, &ferr, nil
}

var TestNewModuleProjectGen = Test{
	Name: "New project generator",
	Function: func(t *TestContext) error {

		// Find the path of the current running binary.
		// assume we are using itself to test itself.
		exePath, err := os.Executable()
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		tempDir := os.TempDir()
		t.Log("Working in:", tempDir)

		t.Log("Found the binary for testing:", exePath)

		{
			_, serr, err := RunExecCmd(exePath, []string{"new"})

			assert.NotNil(t, err, "There should be given error on empty args")
			assert.Contains(t, *serr, "name, module")

			t.Log("Command has error correctly:", *serr)
		}

		{

			args := []string{"new", "--name", "test", "--module", "example.com/torabian/fireback-test"}
			cmd := exec.Command(exePath, args...)
			cmd.Dir = tempDir
			_, serr, err := RunExec(cmd, exePath, args)

			if err != nil {
				t.Error(*serr)
			}
			assert.Nil(t, err, "Should create project boilerplate given name and moduleName")
		}

		{

			p := "go"
			args := []string{"mod", "tidy"}
			cmd := exec.Command(p, args...)
			cmd.Dir = path.Join(tempDir, "test")

			res, serr, err := RunExec(cmd, p, args)

			if err != nil {
				t.Error(*serr)
			}
			assert.Nil(t, err, "Should be able to run 'go mod tidy' and prepare project for build")
			t.Log(*res)
		}

		{

			p := "make"
			args := []string{}
			cmd := exec.Command(p, args...)
			cmd.Dir = path.Join(tempDir, "test")

			res, serr, err := RunExec(cmd, p, args)

			if err != nil {
				t.Error(*serr)
			}
			assert.Nil(t, err, "Should be able to make the project successfully")
			t.Log(*res)
			assert.Contains(t, *res, "go build -ldflags", "Command go build should be visible to the user")
		}

		{
			bpath := path.Join(tempDir, "test", "artifacts", "test-server", "test")
			exists := Exists(bpath)
			assert.True(t, exists, "Binary should be existing in:", bpath)
			t.Log("Binary found correctly:", bpath)
		}

		{

			p := "test"
			args := []string{"doctor"}
			cmd := exec.Command(p, args...)
			cmd.Dir = path.Join(tempDir, "test", "artifacts", "test-server")

			res, serr, err := RunExec(cmd, p, args)

			if err != nil {
				t.Error(*serr)
			}
			assert.Nil(t, err, "Should be able to run doctor command")
			t.Log(*res)

			assert.Contains(t, *res, "currentDirectory", "Current directory needs to be present in doctor")
			assert.Contains(t, *res, "binaryDirectory", "Binary directory needs to be present in doctor")
			assert.Contains(t, *res, "configFileName", "Config filename needs to be present in doctor")
		}

		folder := path.Join(tempDir, "test")

		t.Log("Testing dir: ", folder)
		{
			err := os.RemoveAll(folder)
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		}

		return nil
	},
}
