// Command checkendpointtests makes sure every Emi action - hand-declared or
// entity-synthesized alike - has a dedicated Go test covering it.
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
// A module is considered to cover an action if either:
//   - <ModuleDir>/<Name>Action_test.go exists with a func Test in it (the plain,
//     in-process convention), or
//   - <ModuleDir>/*_test.go or <ModuleDir>/tests/*_test.go (this repo's established
//     black-box HTTP/CLI test package, see modules/abac/tests/testconfig.go) contains a
//     func Test whose name has the action's bare name as a substring, e.g.
//     TestCapabilitiesTree_HTTP covers the "CapabilitiesTree" action.
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

var testFuncNameRe = regexp.MustCompile(`(?m)^func\s+(Test\w+)\s*\(`)

type moduleReport struct {
	emiPath string
	dir     string
	missing []string
	total   int
}

// collectTestFuncNames gathers every top-level `func TestXxx(` name declared in *_test.go
// files directly inside dir, plus (if present) dir/tests/*_test.go - this repo's
// established location for black-box tests (see modules/abac/tests). It deliberately
// doesn't recurse further than that one "tests" subdir, so an unrelated test package
// elsewhere in the tree can't accidentally satisfy coverage for a module it has nothing
// to do with.
func collectTestFuncNames(dir string) ([]string, error) {
	var files []string
	for _, d := range []string{dir, filepath.Join(dir, "tests")} {
		entries, err := os.ReadDir(d)
		if err != nil {
			continue // fine if dir (or its tests/ subdir) doesn't exist or has none
		}
		for _, e := range entries {
			if e.IsDir() || !strings.HasSuffix(e.Name(), "_test.go") {
				continue
			}
			files = append(files, filepath.Join(d, e.Name()))
		}
	}

	var names []string
	for _, f := range files {
		content, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}
		for _, m := range testFuncNameRe.FindAllStringSubmatch(string(content), -1) {
			names = append(names, m[1])
		}
	}
	return names, nil
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

	testFuncNames, err := collectTestFuncNames(dir)
	if err != nil {
		return nil, fmt.Errorf("%s: reading test files: %w", dir, err)
	}

	report := &moduleReport{emiPath: emiPath, dir: dir, total: len(pf.Actions)}
	for _, a := range pf.Actions {
		if a.Name == "" {
			continue
		}
		goName := goActionName(a.Name)
		bareName := strings.TrimSuffix(goName, "Action")

		covered := false
		for _, fn := range testFuncNames {
			if strings.Contains(fn, bareName) {
				covered = true
				break
			}
		}
		if !covered {
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
			bareName := strings.TrimSuffix(name, "Action")
			fmt.Printf("  MISSING  no func Test*%s* in %s or %s\n",
				bareName, report.dir, filepath.Join(report.dir, "tests"))
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
