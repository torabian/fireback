package workspaces

// This struct represents a general field. It could be a field in an go struct, useful
// for defining golang structs, database entities, DTOs, or database query results.
// an entitiy can have an array of fields, and each field might have their own children based on
// type.

type Module3Field struct {
	IsVirtualObject     bool
	Recommended         bool                 `yaml:"recommended,omitempty" json:"recommended,omitempty"`
	LinkedTo            string               `yaml:"linkedTo,omitempty" json:"linkedTo,omitempty"`
	Description         string               `yaml:"description,omitempty" json:"description,omitempty"`
	Name                string               `yaml:"name,omitempty" json:"name,omitempty"`
	Type                string               `yaml:"type,omitempty" json:"type,omitempty"`
	Primitive           string               `yaml:"primitive,omitempty" json:"primitive,omitempty"`
	Target              string               `yaml:"target,omitempty" json:"target,omitempty"`
	RootClass           string               `yaml:"rootClass,omitempty" json:"rootClass,omitempty"`
	Validate            string               `yaml:"validate,omitempty" json:"validate,omitempty"`
	ExcerptSize         int                  `yaml:"excerptSize,omitempty" json:"excerptSize,omitempty"`
	Default             interface{}          `yaml:"default,omitempty" json:"default,omitempty"`
	Translate           bool                 `yaml:"translate,omitempty" json:"translate,omitempty"`
	Unsafe              bool                 `yaml:"unsafe,omitempty" json:"unsafe,omitempty"`
	AllowCreate         bool                 `yaml:"allowCreate,omitempty" json:"allowCreate,omitempty"`
	Module              string               `yaml:"module,omitempty" json:"module,omitempty"`
	Provider            string               `yaml:"provider,omitempty" json:"provider,omitempty"`
	Json                string               `yaml:"json,omitempty" json:"json,omitempty"`
	OfType              []*Module3FieldOf    `yaml:"of,omitempty" json:"of,omitempty"`
	Yaml                string               `yaml:"yaml,omitempty" json:"yaml,omitempty"`
	IdFieldGorm         string               `yaml:"idFieldGorm,omitempty" json:"idFieldGorm,omitempty"`
	ComputedType        string               `yaml:"computedType,omitempty" json:"computedType,omitempty"`
	ComputedTypeClass   string               `yaml:"computedTypeClass,omitempty" json:"computedTypeClass,omitempty"`
	BelongingEntityName string               `yaml:"-" json:"-"`
	Matches             []*Module3FieldMatch `yaml:"matches,omitempty" json:"matches,omitempty"`
	Gorm                string               `yaml:"gorm,omitempty" json:"json,omitempty"`
	GormMap             GormOverrideMap      `yaml:"gormMap,omitempty" json:"gormMap,omitempty"`

	Sql string `yaml:"sql,omitempty" json:"sql,omitempty"`
	// This is the name of field considering how deep it is
	FullName string          `yaml:"fullName,omitempty" json:"fullName,omitempty"`
	Fields   []*Module3Field `yaml:"fields,omitempty" json:"fields,omitempty"`
}
