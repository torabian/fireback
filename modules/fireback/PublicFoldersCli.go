// PublicFoldersCli exposes xapp.PublicFolders (the embedded static folders the
// web server mounts - see gintools.EmbedFoldersForGin/SetupHttpServer) as a debug
// CLI, so an operator can inspect what's actually baked into this binary without
// spinning up the HTTP server at all. Tagged !wasm like every other CLI-only file
// here (CliActions.go, Entrypoint.go, ...) - a wasm build never calls
// CommonHeadlessAppStart (whose GetCommonWebServerCliActions is what registers
// this), so it's never reachable there, but the build tag keeps that true at
// compile time too rather than relying on it going unused.
//go:build !wasm

package fireback

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"sort"
	"strings"

	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/torabian/fireback/modules/fireback/gintools"
	"github.com/urfave/cli/v3"
)

// GetPublicFoldersCli builds the `public-folders` command group: `list` walks
// every xapp.PublicFolders entry and prints each embedded file (as the URL
// path it's actually served at), `get` reads one of those files out by that
// same request path - the same lookup gintools.EmbedFoldersForGin's mounted
// static.Serve middleware does, just without a running HTTP server.
func GetPublicFoldersCli(xapp *application.Application) *cli.Command {
	return &cli.Command{
		Name:  "public-folders",
		Usage: "Inspect xapp.PublicFolders - the embedded static folders the web server mounts (see cmd/fireback/main.go's PublicFolders and gintools.PublicFolderInfo)",
		Commands: []*cli.Command{
			publicFoldersListCli(xapp),
			publicFoldersGetCli(xapp),
		},
	}
}

func publicFoldersListCli(xapp *application.Application) *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "Lists every registered public folder, and every file embedded inside it",
		Action: func(ctx context.Context, c *cli.Command) error {
			if len(xapp.PublicFolders) == 0 {
				fmt.Println("No public folders are registered on this application.")
				return nil
			}

			for _, item := range xapp.PublicFolders {
				prefix := item.Prefix
				if prefix == "" {
					prefix = "/"
				}

				fmt.Printf("%s  (embedded folder: %s)\n", prefix, item.Folder)

				files, err := listEmbeddedFiles(item)
				if err != nil {
					fmt.Println("  error reading folder:", err)
					continue
				}

				if len(files) == 0 {
					fmt.Println("  (empty)")
				}
				for _, f := range files {
					fmt.Println("  " + joinUrlPath(prefix, f))
				}
				fmt.Println()
			}

			return nil
		},
	}
}

func publicFoldersGetCli(xapp *application.Application) *cli.Command {
	return &cli.Command{
		Name:      "get",
		Usage:     "Reads one file out of the embedded public folders by request path (e.g. /manage/fonts/irs/irs.ttf) and writes its raw bytes to stdout, or --out",
		ArgsUsage: "<path>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "path",
				Usage: "Request path of the file, e.g. /manage/fonts/irs/irs.ttf - can also be given as a bare positional argument",
			},
			&cli.StringFlag{
				Name:  "out",
				Usage: "Save the file to this local path instead of writing raw bytes to stdout",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			requestPath := c.String("path")
			if requestPath == "" {
				requestPath = c.Args().First()
			}
			if requestPath == "" {
				return fmt.Errorf("give the file's request path, either as --path or a bare argument (e.g. public-folders get /manage/fonts/irs/irs.ttf)")
			}
			if !strings.HasPrefix(requestPath, "/") {
				requestPath = "/" + requestPath
			}

			data, matchedPrefix, err := readPublicFolderFile(xapp.PublicFolders, requestPath)
			if err != nil {
				return err
			}

			out := c.String("out")
			if out == "" {
				_, err := os.Stdout.Write(data)
				return err
			}

			if err := os.WriteFile(out, data, 0644); err != nil {
				return err
			}
			fmt.Fprintf(os.Stderr, "wrote %d bytes from %s (prefix %s) to %s\n", len(data), requestPath, matchedPrefix, out)
			return nil
		},
	}
}

// listEmbeddedFiles walks item's own embedded sub-filesystem (rooted at
// item.Folder, exactly what gets mounted at item.Prefix) and returns every
// regular file inside it, sorted.
func listEmbeddedFiles(item gintools.PublicFolderInfo) ([]string, error) {
	subFS, err := fs.Sub(*item.Fs, item.Folder)
	if err != nil {
		return nil, err
	}

	var files []string
	walkErr := fs.WalkDir(subFS, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		files = append(files, p)
		return nil
	})
	if walkErr != nil {
		return nil, walkErr
	}

	sort.Strings(files)
	return files, nil
}

// joinUrlPath renders prefix+relativeFile the way it's actually requested over
// HTTP - item.Folder never appears in a served URL, only item.Prefix does.
func joinUrlPath(prefix string, relativeFile string) string {
	if prefix == "/" {
		return "/" + relativeFile
	}
	return strings.TrimSuffix(prefix, "/") + "/" + relativeFile
}

// readPublicFolderFile finds whichever PublicFolders entry actually serves
// requestPath and reads it - the same static.ServeFileSystem.Exists check
// gintools.mountEmbedFolder makes before serving a request for real (see
// EmbedTools.go), just run directly against the embed.FS instead of over
// HTTP. The longest matching prefix wins when more than one folder's Exists
// agrees (mirrors EmbedFoldersForGin's own NoRoute fallback ordering), so a
// path under both a specific folder ("/manage") and a catch-all one ("/")
// resolves to the specific folder, matching what the live server would do.
func readPublicFolderFile(items []gintools.PublicFolderInfo, requestPath string) (data []byte, matchedPrefix string, err error) {
	var best *gintools.PublicFolderInfo
	bestPrefixLen := -1

	for i := range items {
		item := &items[i]
		prefix := item.Prefix
		if prefix == "" {
			prefix = "/"
		}

		fsys := gintools.EmbedFolder(*item.Fs, item.Folder, true)
		if !fsys.Exists(prefix, requestPath) {
			continue
		}
		if len(prefix) > bestPrefixLen {
			best = item
			bestPrefixLen = len(prefix)
		}
	}

	if best == nil {
		return nil, "", fmt.Errorf("no public folder serves path %q (run `public-folders list` to see what's available)", requestPath)
	}

	prefix := best.Prefix
	if prefix == "" {
		prefix = "/"
	}

	fsys := gintools.EmbedFolder(*best.Fs, best.Folder, true)
	relPath := requestPath
	if prefix != "/" {
		relPath = strings.TrimPrefix(requestPath, prefix)
	}
	if relPath == "" {
		relPath = "/"
	}

	f, openErr := fsys.Open(relPath)
	if openErr != nil {
		return nil, "", fmt.Errorf("path %q matched folder %q (prefix %q) but failed to open: %w", requestPath, best.Folder, prefix, openErr)
	}
	defer f.Close()

	if stat, statErr := f.Stat(); statErr == nil && stat.IsDir() {
		return nil, "", fmt.Errorf("path %q is a directory inside folder %q (prefix %q), not a file - see `public-folders list`", requestPath, best.Folder, prefix)
	}

	data, err = io.ReadAll(f)
	if err != nil {
		return nil, "", fmt.Errorf("path %q matched folder %q (prefix %q) but failed to read: %w", requestPath, best.Folder, prefix, err)
	}

	return data, prefix, nil
}
