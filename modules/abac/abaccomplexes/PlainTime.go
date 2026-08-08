package abaccomplexes

import "time"

// PlainTime is a genuine alias for time.Time (not a distinct named type), used as an
// Emi "complex" field for the createdAt/updatedAt bookkeeping columns that Module3 used
// to auto-inject on every entity. Emi's field type system has no bare "time" type - the
// only way to declare one is via a complex - but a wrapper *type* (rather than an
// alias) would stop gorm's reflection-based CreatedAt/UpdatedAt auto-population from
// recognizing the field, since gorm checks for literal time.Time by name convention.
// Being a plain alias keeps it indistinguishable from time.Time to gorm and to Go
// itself, so that auto-population keeps working with zero extra plumbing.
type PlainTime = time.Time
