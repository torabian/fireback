package fireback

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type DownloadedFile struct {
	Content      []byte
	Mime         string
	FileName     string
	Size         int64
	LastModified string
	ETag         string
	Server       string
	Encoding     string
	RawHeaders   http.Header
}

// FetchFileFromURL downloads the file and extracts useful metadata.
func FetchFileFromURL(fileURL string) (*DownloadedFile, error) {
	resp, err := http.Get(fileURL)
	if err != nil {
		return nil, fmt.Errorf("failed to download: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("bad response: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	// Try to guess filename
	fileName := getFileName(resp, fileURL)

	// Get content type
	mimeType := resp.Header.Get("Content-Type")
	if mimeType == "" {
		mimeType = http.DetectContentType(body)
	}

	size := resp.ContentLength
	if size < 0 {
		size = int64(len(body))
	}

	return &DownloadedFile{
		Content:      body,
		Mime:         mimeType,
		FileName:     fileName,
		Size:         size,
		LastModified: resp.Header.Get("Last-Modified"),
		ETag:         resp.Header.Get("ETag"),
		Server:       resp.Header.Get("Server"),
		Encoding:     resp.Header.Get("Content-Encoding"),
		RawHeaders:   resp.Header,
	}, nil
}

func getFileName(resp *http.Response, fileURL string) string {
	cd := resp.Header.Get("Content-Disposition")
	if cd != "" {
		if _, params, err := mime.ParseMediaType(cd); err == nil {
			if filename, ok := params["filename"]; ok {
				return filename
			}
		}
	}
	// fallback to URL path
	u, err := url.Parse(fileURL)
	if err != nil {
		return "unknown"
	}
	name := path.Base(u.Path)
	if name == "" || strings.Contains(name, "?") {
		return "file"
	}
	return name
}
