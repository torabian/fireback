package workspaces

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ssh/terminal"
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

func RunExec2(exePath string, args []string) (*string, *string, error) {
	// Construct the command
	cmd := exec.Command(exePath, args...)

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

func ExecInteractive(dir string, exePath string, args []string, responses []PipeAction) error {
	// Construct the command
	cmd := exec.Command(exePath, args...)
	cmd.Dir = dir

	// Capture command output
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("error creating stdout pipe: %v", err)
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("error creating stderr pipe: %v", err)
	}

	// Create a pipe for stdin
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("error creating stdin pipe: %v", err)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error starting command: %v", err)
	}

	// Goroutine to read and print stdout in real-time
	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			fmt.Printf("stdout: %s\n", scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("error reading stdout: %v\n", err)
		}
	}()

	// Goroutine to read and print stderr in real-time
	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			fmt.Printf("stderr: %s\n", scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("error reading stderr: %v\n", err)
		}
	}()

	// Check if the input is coming from a terminal
	if terminal.IsTerminal(int(os.Stdin.Fd())) {
		oldState, err := terminal.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			return fmt.Errorf("error setting terminal to raw mode: %v", err)
		}
		defer terminal.Restore(int(os.Stdin.Fd()), oldState)
	} else {
		fmt.Println("Not running in a terminal, skipping raw mode.")
	}

	// Handle interactive prompts
	go func() {
		for _, response := range responses {
			_, err := stdinPipe.Write([]byte(response.Write))
			if err != nil {
				fmt.Printf("error writing to stdin: %v\n", err)
				return
			}
			time.Sleep(time.Duration(response.Wait) * time.Millisecond) // Adjust sleep time as needed
		}
		stdinPipe.Close()
	}()

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("command finished with error: %v", err)
	}

	return nil
}

type PipeAction struct {
	Write string
	Wait  int
}

func testInitializeProject(dir string, exePath string, args []string) error {
	shortDelay := 400
	initProjectSequence := []PipeAction{
		{
			Write: "testproject\n",
			Wait:  shortDelay,
		},
		{
			Write: "\x1b[B",
			Wait:  shortDelay,
		},
		{
			Write: "\n",
			Wait:  shortDelay,
		},
		{
			Write: "testdb.db",
			Wait:  shortDelay,
		},
		{
			Write: "\n",
			Wait:  shortDelay,
		},
		{
			Write: "\x1b[B",
			Wait:  shortDelay,
		},
		{
			Write: "\n",
			Wait:  shortDelay,
		},
		{
			Write: "4600",
			Wait:  shortDelay,
		},
		{
			Write: "\n",
			Wait:  shortDelay,
		},
		{
			Write: "fireback-file-storage2",
			Wait:  shortDelay,
		},
		{
			Write: "\n",
			Wait:  shortDelay,
		},
		{
			Write: "4508",
			Wait:  shortDelay,
		},
		{
			Write: "\n",
			Wait:  shortDelay,
		},
		{
			Write: "\n",
			Wait:  shortDelay,
		},
		{
			Write: "\n",
			Wait:  shortDelay,
		},
		{
			Write: "\n",
			Wait:  shortDelay,
		},
		{
			Write: "\n",
			Wait:  2000,
		},
		{
			Write: "\n",
			Wait:  2000,
		},
		{
			Write: "\n",
			Wait:  2000,
		},
		{
			Write: "\n",
			Wait:  2000,
		},
		{
			Write: "\n",
			Wait:  2000,
		},
		{
			Write: "\n",
			Wait:  2000,
		},
		{
			Write: "\n",
			Wait:  2000,
		},
		{
			Write: "\n",
			Wait:  2000,
		},
		{
			Write: "\n",
			Wait:  2000,
		},
		{
			Write: "\n",
			Wait:  2000,
		},
	}
	if err := ExecInteractive(dir, exePath, args, initProjectSequence); err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	} else {
		fmt.Println("Completed successfully.")
		return nil
	}
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

		folder := path.Join(tempDir, "test")

		t.Log("Lets remove test directory first: ", folder)
		{
			err := os.RemoveAll(folder)
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		}

		t.Log("Found the binary for testing:", exePath)

		{

			args := []string{"new", "--name", "test", "--module", "example.com/torabian/fireback-test"}

			// For testing, if we specifiy the sdk location, it would read project from there
			// instead, useful for testing with different versios, ejecting, or any other
			// testing before publish to pkg
			fb := os.Getenv("FIREBACK_SDK_LOCATION")
			if fb != "" {
				args = append(args, "--replace-fb", fb)
			}
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

			bpath := path.Join(tempDir, "test", "artifacts", "test-server", "test")
			stdout, stderr, err := RunExec2(bpath, []string{"doctor"})

			if err != nil {
				t.ErrorLn("Command error:", err)
			}

			assert.Nil(t, err, "Should be able to run doctor command")
			t.Log(*stdout)

			assert.Contains(t, *stdout, "currentDirectory", "Current directory needs to be present in doctor")
			assert.Contains(t, *stdout, "binaryDirectory", "Binary directory needs to be present in doctor")
			assert.Contains(t, *stdout, "configFileName", "Config filename needs to be present in doctor")
			t.Log(*stderr)
		}

		{

			bpath := path.Join(tempDir, "test", "artifacts", "test-server", "test")
			err := testInitializeProject(path.Join(tempDir, "test"), bpath, []string{"init"})

			assert.Nil(t, err, "Should be able to init the project")
			if err != nil {
				t.Log("Error %v:", err)
			}
		}

		{

			bpath := path.Join(tempDir, "test", "artifacts", "test-server", "test")
			args := []string{"migration", "apply"}
			cmd := exec.Command(bpath, args...)
			cmd.Dir = path.Join(tempDir, "test")
			res, serr, err := RunExec(cmd, bpath, args)

			if err != nil {
				t.ErrorLn("Command error:", err)
			}

			t.Log("Migration function did not throw any comment or issues")
			assert.Nil(t, err, "Should be able to apply the migration without any error")
			t.Log(*serr)
			t.Log(*res)
		}

		{

			bpath := path.Join(tempDir, "test", "artifacts", "test-server", "test")
			args := []string{"tz", "sync"}
			cmd := exec.Command(bpath, args...)
			cmd.Dir = path.Join(tempDir, "test")
			res, serr, err := RunExec(cmd, bpath, args)

			if err != nil {
				t.ErrorLn("Command error:", err)
			}

			t.Log("Timezone sync ran correctly")
			assert.Nil(t, err, "Should be able to apply the migration without any error")
			t.Log(*serr)
			t.Log(*res)

		}

		return nil
	},
}
