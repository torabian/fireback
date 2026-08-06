// Command checkendpointtests makes sure every Emi action - hand-declared or
// entity-synthesized alike - has a dedicated Go test file sitting next to its
// generated wiring file.
//
// Source of truth: each module's own preprocessed.yml (the "preprocessor" compiler
// target's output, produced by `emi compile --path <module>.emi.yml`). Since
// core.ReadEmiFromString runs the full Preprocess() pipeline before that file is
// written, preprocessed.yml's actions: list already includes entity CRUD actions
// synthesized from entities: (see ../../../emi/lib/core/preprocess-entity-actions.go)
// alongside hand-declared actions: entries - exactly what actually gets compiled into
// <Name>Action.go files, without needing to reimplement any of Emi's own entity
// expansion logic here.
//
// Usage:
//
//	go run ./tools/checkendpointtests [--strict] [modules/dir ...]
//
// With no directory arguments, it walks the whole repo for *.emi.yml files. Exits 1
// only when --strict is passed and at least one action is missing its test file;
// otherwise this is a report-only tool (see the Makefile's checkendpointtests target).
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

type preprocessedAction struct {
	Name string `yaml:"name"`
}

type preprocessedFile struct {
	Actions []preprocessedAction `yaml:"actions"`
}

// goActionName mirrors emi's core.ToUpper(name) + "Action" (see
// lib/core/EmiAction.go's ActionName()/lib/core/utilities.go's ToUpper) - capitalizes
// only the first rune, same transform Emi's own Go backend uses to derive
// <Name>Action.go's filename from the action's (possibly lowerCamel, for
// entity-synthesized actions) declared name.
func goActionName(name string) string {
	if name == "" {
		return ""
	}
	return strings.ToUpper(name[:1]) + name[1:] + "Action"
}

var testFuncRe = regexp.MustCompile(`(?m)^func\s+Test`)

type moduleReport struct {
	emiPath string
	dir     string
	missing []string
	total   int
}

func findEmiFiles(roots []string) ([]string, error) {
	var found []string
	for _, root := range roots {
		err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				// node_modules/dotdirs: vendored, irrelevant. project-generator: the
				// front-end project - it ships its own *.emi.yml fixtures (e.g.
				// sdk/sdk/envelopes/.../google-envelop.emi.yml) that describe TS SDK
				// envelope classes, not backend Go modules, so they're not part of what
				// this tool checks. Any "sdk" dir is generated output for the same
				// reason, wherever it's nested.
				switch d.Name() {
				case "node_modules", "project-generator", "sdk":
					return filepath.SkipDir
				}
				if strings.HasPrefix(d.Name(), ".") {
					return filepath.SkipDir
				}
				return nil
			}
			if strings.HasSuffix(strings.ToLower(d.Name()), ".emi.yml") {
				found = append(found, path)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	sort.Strings(found)
	return found, nil
}

func checkModule(emiPath string) (*moduleReport, error) {
	dir := filepath.Dir(emiPath)
	preprocessedPath := filepath.Join(dir, "preprocessed.yml")

	data, err := os.ReadFile(preprocessedPath)
	if err != nil {
		return nil, fmt.Errorf("%s: no preprocessed.yml next to it (run `go run github.com/torabian/emi/cmd/emi compile --path %s` first): %w", emiPath, emiPath, err)
	}

	var pf preprocessedFile
	if err := yaml.Unmarshal(data, &pf); err != nil {
		return nil, fmt.Errorf("%s: %w", preprocessedPath, err)
	}

	report := &moduleReport{emiPath: emiPath, dir: dir, total: len(pf.Actions)}
	for _, a := range pf.Actions {
		if a.Name == "" {
			continue
		}
		goName := goActionName(a.Name)
		testPath := filepath.Join(dir, goName+"_test.go")

		content, err := os.ReadFile(testPath)
		if err != nil || !testFuncRe.Match(content) {
			report.missing = append(report.missing, goName)
		}
	}
	sort.Strings(report.missing)
	return report, nil
}

func main() {
	strict := false
	var roots []string
	for _, arg := range os.Args[1:] {
		if arg == "--strict" {
			strict = true
			continue
		}
		roots = append(roots, arg)
	}
	if len(roots) == 0 {
		roots = []string{"modules"}
	}

	emiFiles, err := findEmiFiles(roots)
	if err != nil {
		fmt.Fprintln(os.Stderr, "checkendpointtests:", err)
		os.Exit(2)
	}
	if len(emiFiles) == 0 {
		fmt.Println("checkendpointtests: no *.emi.yml files found under", roots)
		return
	}

	var reports []*moduleReport
	totalActions, totalMissing := 0, 0
	hadReadError := false

	for _, emiPath := range emiFiles {
		report, err := checkModule(emiPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "checkendpointtests:", err)
			hadReadError = true
			continue
		}
		reports = append(reports, report)
		totalActions += report.total
		totalMissing += len(report.missing)
	}

	for _, report := range reports {
		covered := report.total - len(report.missing)
		fmt.Printf("\n%s (%d/%d actions have a test file)\n", report.emiPath, covered, report.total)
		for _, name := range report.missing {
			fmt.Printf("  MISSING  %s\n", filepath.Join(report.dir, name+"_test.go"))
		}
	}

	fmt.Printf("\ncheckendpointtests: %d/%d actions covered across %d module(s), %d missing\n",
		totalActions-totalMissing, totalActions, len(reports), totalMissing)

	if hadReadError {
		os.Exit(2)
	}
	if strict && totalMissing > 0 {
		fmt.Fprintln(os.Stderr, "checkendpointtests: --strict set, failing due to missing test files above")
		os.Exit(1)
	}
}
