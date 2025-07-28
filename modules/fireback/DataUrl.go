package fireback

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
)

type DataURL struct {
	Mime     string
	Charset  string
	Params   map[string]string
	Base64   bool
	Content  []byte
	RawInput string
}

// parses a full data URL and extracts all metadata + content
func ParseDataURL(input string) (*DataURL, error) {
	if !strings.HasPrefix(input, "data:") {
		return nil, nil
	}

	raw := strings.TrimPrefix(input, "data:")
	parts := strings.SplitN(raw, ",", 2)
	if len(parts) != 2 {
		return nil, nil
	}

	meta := parts[0]
	content := parts[1]

	metaParts := strings.Split(meta, ";")

	result := &DataURL{
		Mime:     "text/plain", // default
		Charset:  "US-ASCII",   // default
		Params:   make(map[string]string),
		Base64:   false,
		RawInput: input,
	}

	if metaParts[0] != "" && !strings.Contains(metaParts[0], "=") {
		result.Mime = metaParts[0]
		metaParts = metaParts[1:] // skip MIME
	}

	for _, p := range metaParts {
		if p == "base64" {
			result.Base64 = true
		} else if strings.Contains(p, "=") {
			kv := strings.SplitN(p, "=", 2)
			result.Params[kv[0]] = kv[1]
			if kv[0] == "charset" {
				result.Charset = kv[1]
			}
		}
	}

	var decoded []byte
	var err error

	if result.Base64 {
		decoded, err = base64.StdEncoding.DecodeString(content)
		if err != nil {
			return nil, fmt.Errorf("invalid base64 content: %w", err)
		}
	} else {
		decoded2, err := url.QueryUnescape(content)
		if err != nil {
			return nil, fmt.Errorf("invalid url-encoded content: %w", err)
		}

		decoded = []byte(decoded2)
	}

	result.Content = decoded
	return result, nil
}
