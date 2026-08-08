//go:build !wasm

// Package clitools holds every fireback CLI feature that needs a real
// interactive terminal or OS process control (promptui/bubbletea prompts,
// os/exec service management, asynq worker processes, graceful HTTP
// shutdown via OS signals). None of this can run under wasm/js, so it lives
// in its own package, tagged !wasm, and registers itself into the core
// fireback package's function-pointer hooks via init() - the same pattern
// already used by modules/abac for optional action overrides. Importing
// this package (even blank-imported) is what makes these features
// available; omitting the import (as cmd/fireback-wasm does) leaves the
// hooks nil and the core package free of any of these dependencies.
package clitools

import (
	"errors"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	fireback.AskForSelect = askForSelect
	fireback.AskBoolean = askBoolean
	fireback.AskForInputOptional = askForInputOptional
	fireback.AskForInput = askForInput
	fireback.AskForPassword = askForPassword
}

func askForSelect(label string, items []string) string {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		if err.Error() == "^C" {
			os.Exit(35)
			return ""
		}
		return ""
	}

	index := strings.Index(result, ">>>")
	if index <= 0 {
		return result
	}
	return strings.Trim(result[0:index], " ")

}

func askBoolean(label string) bool {
	if r := askForSelect(label, []string{"true", "false"}); r == "true" {
		return true
	}

	return false
}

func askForInputOptional(label string, defaultV string) (string, bool, error) {

	promptVariable := promptui.Prompt{
		Label:   label,
		Default: defaultV,
	}

	value, err := promptVariable.Run()
	if err != nil {
		if err.Error() == "^C" {
			os.Exit(35)
			return "", true, err
		}
		return "", false, err
	}

	return value, false, nil
}

func askForInput(label string, defaultV string) string {
	validate := func(input string) error {
		if input == "" {
			return errors.New("this is necessary")
		}
		return nil
	}

	promptVariable := promptui.Prompt{
		Label:    label,
		Validate: validate,
		Default:  defaultV,
	}

	value, err := promptVariable.Run()
	if err != nil {
		if err.Error() == "^C" {
			os.Exit(35)
			return ""
		}
		return ""
	}

	return value
}

func askForPassword(label string, defaultV string) string {
	validate := func(input string) error {
		if input == "" {
			return errors.New("this is necessary")
		}
		return nil
	}

	promptVariable := promptui.Prompt{
		Label:    label,
		Mask:     '*',
		Validate: validate,
		Default:  defaultV,
	}

	value, err := promptVariable.Run()
	if err != nil {
		if err.Error() == "^C" {
			os.Exit(35)
			return ""
		}
		return ""
	}
	return value
}
