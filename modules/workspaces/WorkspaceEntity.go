package workspaces

func init() {
	WorkspaceCliCommands = append(WorkspaceCliCommands, GetWorkspacesActionsCli()...)
}

// @meta(include)
// type WorkspaceEntity struct {
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// 	UniqueId    string `json:"uniqueId" gorm:"primarykey;uniqueId;unique;not null;size:100;"`
// }

// // @meta(include)
// type PendingWorkspaceInvite struct {
// 	ID    uint
// 	Email string
// }

// // @meta(include)
// type WorkspaceInvite struct {
// 	CoverLetter      string `json:"description"`
// 	TargetUserLocale string `json:"locale"`
// 	Email            string `json:"email"`
// 	PhoneNumber      string `json:"phoneNumber"`

// 	Workspace   WorkspaceEntity `gorm:"foreignKey:WorkspaceID;references:UniqueId" json:"-"`
// 	WorkspaceID string          `json:"workspaceId" gorm:"size:100;"`

// 	Role   RoleEntity `gorm:"foreignKey:RoleID;references:UniqueId" json:"-"`
// 	RoleID string     `json:"roleID" gorm:"size:100;"`

// 	UniqueId string `json:"uniqueId" gorm:"primarykey;uniqueId;unique;not null;size:100;"`
// }

// // @meta(include)
// type UserRoleWorkspace struct {
// 	Workspace   *WorkspaceEntity `gorm:"foreignKey:WorkspaceID;references:UniqueId" json:"workspace"`
// 	WorkspaceID *string          `json:"workspaceId" gorm:"size:100;"`

// 	Role   *RoleEntity `gorm:"foreignKey:RoleID;references:UniqueId" json:"role"`
// 	RoleID *string     `json:"roleID" gorm:"size:100;"`

// 	User   *UserEntity `gorm:"foreignKey:UserID;references:UniqueId" json:"user"`
// 	UserID *string           `json:"userId" gorm:"size:100;"`

// 	UniqueId string `json:"uniqueId" gorm:"primarykey;uniqueId;unique;not null;size:100;"`
// }
