package fireback

import (
	"database/sql/driver"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (x XFileMeta) Json() string {
	str, _ := json.MarshalIndent(x, "", "  ")
	return (string(str))
}

func (j *XFileMeta) Scan(value interface{}) error {
	res, err := ScanDbField(value)
	if err != nil {
		return err
	}

	return json.Unmarshal(res, j)
}

func (j XFileMeta) Value() (driver.Value, error) {
	return j.Json(), nil
}

func (XFileMeta) GormDataType() string {
	return "json"
}

type XFile struct {
	Meta XFileMeta `json:"meta" yaml:"meta" gorm:"-" sql:"-"`

	// Store the file size in a separate column
	Filesize uint64 `json:"fileSize" yaml:"fileSize"`

	Blob []byte `gorm:"type:blob" json:"-" yaml:"-"`
}

type XFileMeta struct {
	FileID   string `json:"fileId" yaml:"fileId"`
	FileName string `json:"fileName" yaml:"fileName"`
	URL      string `json:"url" yaml:"url"`
	Mime     string `json:"mime" yaml:"mime"`
	IsBase64 bool   `json:"isBase64,omitempty" yaml:"isBase64,omitempty"`
}

func (f *XFile) UnmarshalJSON(data []byte) error {
	var obj interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	return f.parseXFileInput(obj)
}

func (f *XFile) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var obj interface{}
	if err := unmarshal(&obj); err != nil {
		return err
	}
	return f.parseXFileInput(obj)
}

func (f *XFile) parseXFileInput(obj interface{}) error {

	switch v := obj.(type) {
	case string:
		parsed, err := ParseDataURL(v)
		if err != nil {
			return err
		}

		// This means it has been by passed.
		if err == nil && parsed == nil {

			// Let's handle the http files
			if strings.HasPrefix(v, "http") {
				file, errDl := FetchFileFromURL(v)
				if errDl == nil {
					f.Blob = file.Content
					f.Meta.Mime = file.Mime
					f.Filesize = uint64(len(file.Content))

					return nil
				}
			}

		} else {

			f.Blob = parsed.Content
			f.Meta.Mime = parsed.Mime
			f.Filesize = uint64(len(parsed.Content))
			return nil
		}

		// In the end, the string is just []byte as well
		f.Blob = []byte(obj.(string))
		f.Filesize = uint64(len(f.Blob))

		return nil

	case nil:
		return nil

	case map[string]interface{}:
		// Extract known fields
		if b64, ok := v["blob"].(string); ok && b64 != "" {
			blob, err := base64.StdEncoding.DecodeString(b64)
			if err != nil {
				return fmt.Errorf("XFile: invalid base64 blob: %w", err)
			}
			f.Blob = blob
		}
		if name, ok := v["fileName"].(string); ok {
			f.Meta.FileName = name
		}
		if mime, ok := v["mime"].(string); ok {
			f.Meta.Mime = mime
		}
		if size, ok := v["size"].(float64); ok {
			f.Filesize = uint64(size)
		}
		if url, ok := v["url"].(string); ok {
			f.Meta.URL = url
		}
		if id, ok := v["fileId"].(string); ok {
			f.Meta.FileID = id
		}
		return nil

	default:
		return fmt.Errorf("XFile: unsupported format in YAML input: %T", obj)
	}
}

func encodeBlobDataURL(blob []byte, filename string) string {
	if len(blob) == 0 {
		return ""
	}

	mime := detectMimeType(blob, filename)
	encoded := base64.StdEncoding.EncodeToString(blob)
	return fmt.Sprintf("data:%s;base64,%s", mime, encoded)
}

// MarshalJSON - ensure blob is encoded as base64 string in field "blob"
func (f XFile) marshalStructWithEncodedBlob() any {
	const maxInlineSize = 25 * 1024 // 25KB

	var blobData string
	meta := f.Meta
	meta.IsBase64 = false

	// If inlining allowed, inline it.
	if len(f.Blob) > 0 && len(f.Blob) <= maxInlineSize {
		blobData = encodeBlobDataURL(f.Blob, meta.FileName)
		meta.IsBase64 = true
	} else if len(f.Blob) == 0 {
		return nil
	} else {
		// Now we need to give a url, so the user can read the file
		blobData = "large"
	}

	return &struct {
		Blob     string    `json:"blob,omitempty" yaml:"blob,omitempty"`
		Meta     XFileMeta `json:"meta" yaml:"meta"`
		FileSize uint64    `json:"fileSize" yaml:"fileSize"`
	}{
		Blob:     blobData,
		Meta:     meta,
		FileSize: f.Filesize,
	}
}
func (f XFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.marshalStructWithEncodedBlob())
}

func (f XFile) MarshalYAML() (interface{}, error) {
	return f.marshalStructWithEncodedBlob(), nil
}
func NewXFileAutoNull(value string) XFile {
	if Exists(value) {
		fmt.Println("File exists", value)

		info, err := os.Stat(value)
		if err != nil {
			fmt.Println("Failed to stat file:", err)
			return XFile{}
		}

		data, err := os.ReadFile(value)
		if err != nil {
			fmt.Println("Failed to read file:", err)
			return XFile{}
		}

		mime := http.DetectContentType(data)

		return XFile{
			Blob:     data,
			Filesize: uint64(info.Size()),
			Meta: XFileMeta{
				FileName: filepath.Base(value),
				Mime:     mime,
			},
		}
	}
	return XFile{}
}
