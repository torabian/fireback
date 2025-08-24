# Fireback core microservice - v1.2.4
Total modules: 11
Modules overview: workspaces, workspaces, workspaces, workspaces, geo, accessibility, widget, commonprofile, currency, licenses, worldtimezone
## Workspaces
Description: This is the fireback core module, which includes everything. In fact you could say workspaces is fireback itself. Maybe in the future that would be changed


### Workspaces Entities
| Name | Usage | Main data |
| --- | --- | --- |
| File | File manager, uploading files and actions related. | Name, DiskPath, Size, VirtualPath, Type, Variations |
| TableViewSizing | Used to store meta data about user tables (in front-end, or apps for example) about the size of the columns | TableName, Sizes |
| AppMenu | Manages the menus in the app, (for example tab views, sidebar items, etc.) | Label, Href, Icon, ActiveMatcher, ApplyType, Capability |
| BackupTableMeta | Keeps information about which tables to be used during backup (mostly internal) | TableNameInDb |
| NotificationConfig | Configuration for the notifications used in the app, such as default gsm number, email senders, and many more | CascadeToSubWorkspaces, ForcedCascadeEmailProvider, GeneralEmailProvider, GeneralGsmProvider, InviteToWorkspaceContent, InviteToWorkspaceContentExcerpt, InviteToWorkspaceContentDefault, InviteToWorkspaceContentDefaultExcerpt, InviteToWorkspaceTitle, InviteToWorkspaceTitleDefault, InviteToWorkspaceSender, AccountCenterEmailSender, ForgetPasswordContent, ForgetPasswordContentExcerpt, ForgetPasswordContentDefault, ForgetPasswordContentDefaultExcerpt, ForgetPasswordTitle, ForgetPasswordTitleDefault, ForgetPasswordSender, AcceptLanguage, ConfirmEmailSender, ConfirmEmailContent, ConfirmEmailContentExcerpt, ConfirmEmailContentDefault, ConfirmEmailContentDefaultExcerpt, ConfirmEmailTitle, ConfirmEmailTitleDefault |
| PassportMethod | Login/Signup methods which are available in the app for different regions (Email, Phone Number, Google, etc) | Name, Type, Region |
| WorkspaceInvite | Active invitations for non-users or already users to join an specific workspace | CoverLetter, TargetUserLocale, Value, Workspace, FirstName, LastName, Used, Role |
| PendingWorkspaceInvite |  | Value, Type, CoverLetter, WorkspaceName, Role |
| Preference |  | Timezone |
| Token |  | User, ValidUntil |
| Person |  | FirstName, LastName, Photo, Gender, Title, BirthDate |
| UserWorkspace | Manage the workspaces that user belongs to (either its himselves or adding by invitation) | User, Workspace, UserPermissions, RolePermission, WorkspacePermissions |
| WorkspaceRole | Manage roles assigned to an specific workspace or created by the workspace itself | UserWorkspace, Role |
| User | Manage the users who are in the current app (root only) | Person, Avatar |
| UserProfile |  | FirstName, LastName |
| Workspace | Fireback general user role, workspaces services. | Description, Name, Type |
| Role | Manage roles within the workspaces, or root configuration | Name, Capabilities |
| Capability | Manage the capabilities inside the application, both builtin to core and custom defined ones | Name, Description |
| WorkspaceConfig |  | DisablePublicWorkspaceCreation, Workspace, ZoomClientId, ZoomClientSecret, AllowPublicToJoinTheWorkspace |
| GsmProvider |  | ApiKey, MainSenderNumber, Type, InvokeUrl, InvokeBody |
| WorkspaceType | Defines a type for workspace, and the role which it can have as a whole. In systems with multiple types of services, e.g. student, teachers, schools this is useful to set those default types and limit the access of the users. | Title, Description, Slug, Role |
| EmailProvider | Thirdparty services which will send email, allows each workspace graphically configure their token without the need of restarting servers | Type, ApiKey |
| EmailSender | All emails going from the system need to have a virtual sender (nick name, email address, etc) | FromName, FromEmailAddress, ReplyTo, NickName |
| PhoneConfirmation |  | User, Status, PhoneNumber, Key, ExpiresAt |
| PublicJoinKey | Joining to different workspaces using a public link directly | Role, Workspace |
| EmailConfirmation |  | User, Status, Email, Key, ExpiresAt |
| Passport | Represent a mean to login in into the system, each user could have multiple passport (email, phone) and authenticate into the system. | Type, User, Value, Password, Confirmed, AccessToken |
| RegionalContent | Email templates, sms templates or other textual content which can be accessed. | Content, Region, Title, LanguageId, KeyGroup |
| ForgetPassword |  | User, Passport, Status, ValidUntil, BlockedUntil, SecondsToUnblock, Otp, RecoveryAbsoluteUrl |





## Geo
Description: Geo location tools, and data set, cities, and provinces2


### Geo Entities
| Name | Usage | Main data |
| --- | --- | --- |
| GeoLocationType |  | Name |
| GeoLocation |  | Name, Code, Type, Status, Flag, OfficialName |
| GeoCountry |  | Status, Flag, CommonName, OfficialName |
| GeoProvince |  | Name, Country |
| GeoState |  | Name |
| GeoCity |  | Name, Province, State, Country |




### Geo actions (0)



## Accessibility


### Accessibility Entities
| Name | Usage | Main data |
| --- | --- | --- |
| KeyboardShortcut | Manage the keyboard shortcuts in web and desktop apps (accessibility) | Os, Host, DefaultCombination, UserCombination, Action, ActionKey |





## Widget


### Widget Entities
| Name | Usage | Main data |
| --- | --- | --- |
| Widget | Widget is an item which can be placed on a widget area, such as weather widget | Name, Family, ProviderKey |
| WidgetArea | Widget areas are groups of widgets, which can be placed on a special place such as dashboard | Name, Layouts, Widgets |





## Commonprofile


### Commonprofile Entities
| Name | Usage | Main data |
| --- | --- | --- |
| CommonProfile | A common profile issues for every user (Set the living address, etc) | FirstName, LastName, PhoneNumber, Email, Company, Street, HouseNumber, ZipCode, City, Gender |





## Currency


### Currency Entities
| Name | Usage | Main data |
| --- | --- | --- |
| Currency | List of all famous currencies, both internal and user defined ones | Symbol, Name, SymbolNative, DecimalDigits, Rounding, Code, NamePlural |
| PriceTag | Price tag is a definition of a price, in different currencies or regions | Variations |





## Licenses


### Licenses Entities
| Name | Usage | Main data |
| --- | --- | --- |
| LicensableProduct |  | Name, PrivateKey, PublicKey |
| ProductPlan |  | Name, Duration, Product, PriceTag, Permissions |
| ActivationKey |  | Series, Used, Plan |
| License | Manage the licenses in the app (either to issue, or to activate current product) | Name, SignedLicense, ValidityStartDate, ValidityEndDate, Permissions |





## Worldtimezone


### Worldtimezone Entities
| Name | Usage | Main data |
| --- | --- | --- |
| TimezoneGroup | World timezone information | Value, Abbr, Offset, Isdst, Text, UtcItems |




