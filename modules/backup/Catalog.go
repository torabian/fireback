package backup

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// Catalog records operational metadata wal-g itself doesn't track: whether
// a backup has been through a restore drill (backup verify) and when it was
// pruned. wal-g's own `backup-list` stays the source of truth for which
// backups actually exist - this file is never used to decide that, only to
// annotate it, so it can never drift into claiming a backup exists when
// wal-g disagrees.
type Catalog struct {
	Entries map[string]*CatalogEntry `json:"entries"`
}

type CatalogEntry struct {
	BackupName string     `json:"backup_name"`
	Type       string     `json:"type"` // "full" or "delta"
	PushedAt   time.Time  `json:"pushed_at"`
	VerifiedAt *time.Time `json:"verified_at,omitempty"`
	VerifyOK   *bool      `json:"verify_ok,omitempty"`
	Notes      string     `json:"notes,omitempty"`
}

func catalogPath(cfg *ModuleConfig) string {
	return filepath.Join(cfg.FilePrefix, ".nima-backup-catalog.json")
}

// LoadCatalog reads the catalog file, returning an empty (not nil) Catalog
// if it doesn't exist yet - the common case on first ever `backup push`.
func LoadCatalog(cfg *ModuleConfig) (*Catalog, error) {
	data, err := os.ReadFile(catalogPath(cfg))
	if os.IsNotExist(err) {
		return &Catalog{Entries: map[string]*CatalogEntry{}}, nil
	}
	if err != nil {
		return nil, err
	}

	var c Catalog
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	if c.Entries == nil {
		c.Entries = map[string]*CatalogEntry{}
	}
	return &c, nil
}

// Save writes the catalog back to disk atomically (write to a temp file,
// then rename) so a crash mid-write can't leave a corrupt/partial catalog
// behind.
func (c *Catalog) Save(cfg *ModuleConfig) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	path := catalogPath(cfg)
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

func (c *Catalog) RecordPush(backupName, backupType string) {
	c.Entries[backupName] = &CatalogEntry{
		BackupName: backupName,
		Type:       backupType,
		PushedAt:   time.Now().UTC(),
	}
}

func (c *Catalog) RecordVerify(backupName string, ok bool, notes string) {
	entry, exists := c.Entries[backupName]
	if !exists {
		entry = &CatalogEntry{BackupName: backupName}
		c.Entries[backupName] = entry
	}
	now := time.Now().UTC()
	entry.VerifiedAt = &now
	entry.VerifyOK = &ok
	entry.Notes = notes
}
