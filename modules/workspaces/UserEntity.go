package workspaces

/*
* Main entity of the module. Do not define multiple in one module, each module must provide all crud,
* actions, and cli operation independently. Create larger modules by combining smaller modules
 */

func (x *UserEntity) FullName() string {
	if x.Person == nil {
		return ""
	}

	full := ""

	if x.Person.FirstName != nil {
		full += *x.Person.FirstName
	}

	if x.Person.LastName != nil {
		full += " " + *x.Person.LastName
	}

	return full

}

// @meta(include)
// type UserEntity struct {
// 	Model
// 	FirstName string `json:"firstName"`
// 	Lastname  string `json:"lastName"`
// 	Photo     string `json:"photo"`
// 	UniqueId  string `json:"uniqueId" gorm:"primarykey;uniqueId;unique;not null;size:100;"`
// }

// @meta(include)
// type Token struct {
// 	Model
// 	Hash       string     `json:"hash"`
// 	User       UserEntity `gorm:"foreignKey:UserID;references:UniqueId" json:"-"`
// 	UserID     string     `json:"userId"`
// 	ValidUntil time.Time
// 	UniqueId   string `json:"uniqueId"`
// }

// @meta(include)
// type Preference struct {
// 	Model
// 	ItemKey   string `json:"itemKey"`
// 	Value     string `json:"value"`
// 	ValueType string `json:"valueType"`
// 	Scope     string `json:"scope"`

// 	User   UserEntity `gorm:"foreignKey:UserID;references:UniqueId" json:"-"`
// 	UserID string     `json:"userId"`
// }
