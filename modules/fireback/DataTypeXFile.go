package fireback

import (
	"database/sql/driver"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type XFileMeta struct {
	FileID   string `json:"fileId" yaml:"fileId"`
	FileName string `json:"fileName" yaml:"fileName"`
	URL      string `json:"url" yaml:"url"`
	Mime     string `json:"mime" yaml:"mime"`
	Size     int64  `json:"size" yaml:"size"`
}

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
	Meta XFileMeta

	Blob []byte `gorm:"type:blob" json:"-" yaml:"-"`
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
				}
			}

		} else {

			f.Blob = parsed.Content
			f.Meta.Mime = parsed.Mime
		}

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
			f.Meta.Size = int64(size)
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

// MarshalJSON - ensure blob is encoded as base64 string in field "blob"
func (f XFile) MarshalJSON() ([]byte, error) {
	type Alias XFile

	aux := &struct {
		Blob string `json:"blob,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(&f),
	}
	if len(f.Blob) > 0 {
		aux.Blob = base64.StdEncoding.EncodeToString(f.Blob)
	}
	return json.Marshal(aux)
}

func NewXFileAutoNull(value string) XFile {
	return XFile{}
}
