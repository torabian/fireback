package storage

import "encoding/json"

// ComplexFile represents the value a file field holds: either just an id
// (e.g. a plain tus upload id/url already stored on a record) or a richer
// object carrying whatever metadata was captured at upload time (filesize,
// filename, ...). Mirrors admin/src/modules/resumable-uploader/ComplexFile.ts
// so the same value round-trips between frontend and backend without either
// side having to care which shape it currently holds.
//
// JSON sent as a bare string ("abc123") unmarshals with only Id set. JSON
// sent as an object round-trips its known fields plus anything else under
// Extra. Marshaling always produces the object form, regardless of which
// shape was originally unmarshaled - this is also the shape emi's "complex"
// field type will target once ComplexFile is wired up as one.
type ComplexFile struct {
	Id        string                 `json:"id"`
	Filesize  int64                  `json:"filesize,omitempty"`
	Thumbnail string                 `json:"thumbnail,omitempty"`
	Filename  string                 `json:"filename,omitempty"`
	MimeType  string                 `json:"mimeType,omitempty"`
	Extra     map[string]interface{} `json:"-"`
}

// complexFileFields lists the JSON keys ComplexFile decodes into named
// struct fields - anything else in an object payload is preserved in Extra
// instead of being silently dropped.
var complexFileFields = []string{"id", "filesize", "thumbnail", "filename", "mimeType"}

// NewComplexFile builds a ComplexFile the same way UnmarshalJSON would from
// v: a string becomes an id-only value, anything else is marshaled then
// unmarshaled through ComplexFile's own JSON handling so the two paths can
// never disagree on shape.
func NewComplexFile(v interface{}) (ComplexFile, error) {
	if id, ok := v.(string); ok {
		return ComplexFile{Id: id}, nil
	}

	data, err := json.Marshal(v)
	if err != nil {
		return ComplexFile{}, err
	}

	var f ComplexFile
	if err := json.Unmarshal(data, &f); err != nil {
		return ComplexFile{}, err
	}
	return f, nil
}

// String returns the file's id, so a ComplexFile can be used wherever a
// plain id string is expected (fmt.Sprintf("%s", f), %v formatting, etc.)
func (f ComplexFile) String() string {
	return f.Id
}

// IsEmpty reports whether f holds no file: true both when the field was
// left out of the request entirely (the zero value) and when it was
// explicitly sent as JSON null - either way, there's no Id to act on. A
// bare `{}` object, or a string field with no id, count as empty too.
//
// For distinguishing "field omitted" from "field explicitly null" (rather
// than treating both as the same "no value"), wrap the field as
// emigo.Nullable[ComplexFile] instead - its IsSet()/IsNull() already give
// you that three-way state, and ComplexFile's own UnmarshalJSON/MarshalJSON
// compose with it with no extra work.
func (f ComplexFile) IsEmpty() bool {
	return f.Id == ""
}

// UnmarshalJSON accepts either a bare JSON string (id-only) or a JSON
// object carrying id plus any metadata, so a field of type ComplexFile can
// be decoded from whichever shape the frontend currently sends.
func (f *ComplexFile) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*f = ComplexFile{}
		return nil
	}

	var asString string
	if err := json.Unmarshal(data, &asString); err == nil {
		*f = ComplexFile{Id: asString}
		return nil
	}

	type complexFileAlias ComplexFile
	var alias complexFileAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	for _, known := range complexFileFields {
		delete(raw, known)
	}

	var extra map[string]interface{}
	if len(raw) > 0 {
		extra = make(map[string]interface{}, len(raw))
		for k, v := range raw {
			var value interface{}
			if err := json.Unmarshal(v, &value); err != nil {
				return err
			}
			extra[k] = value
		}
	}

	*f = ComplexFile(alias)
	f.Extra = extra
	return nil
}

// MarshalJSON always produces the full object form (id plus any set
// metadata, plus whatever extra keys were preserved on unmarshal).
func (f ComplexFile) MarshalJSON() ([]byte, error) {

	if f.Id == "" {
		return []byte("null"), nil
	}

	type complexFileAlias ComplexFile

	encoded, err := json.Marshal(complexFileAlias(f))
	if err != nil {
		return nil, err
	}
	if len(f.Extra) == 0 {
		return encoded, nil
	}

	var base map[string]interface{}
	if err := json.Unmarshal(encoded, &base); err != nil {
		return nil, err
	}
	for k, v := range f.Extra {
		if _, exists := base[k]; !exists {
			base[k] = v
		}
	}
	return json.Marshal(base)
}
