package workspaces
import (
	"github.com/gin-gonic/gin"
    "github.com/urfave/cli"
)
var SendEmailSecurityModel *SecurityModel = nil
type SendEmailActionReqDto struct {
    ToAddress   *string `json:"toAddress" yaml:"toAddress"  validate:"required"       `
    // Datenano also has a text representation
    Body   *string `json:"body" yaml:"body"  validate:"required"       `
    // Datenano also has a text representation
}
func ( x * SendEmailActionReqDto) RootObjectName() string {
	return "workspaces"
}
var SendEmailCommonCliFlagsOptional = []cli.Flag{
    &cli.StringFlag{
      Name:     "to-address",
      Required: true,
      Usage:    "toAddress",
    },
    &cli.StringFlag{
      Name:     "body",
      Required: true,
      Usage:    "body",
    },
}
func SendEmailActionReqValidator(dto *SendEmailActionReqDto) *IError {
    err := CommonStructValidatorPointer(dto, false)
    return err
  }
func CastSendEmailFromCli (c *cli.Context) *SendEmailActionReqDto {
	template := &SendEmailActionReqDto{}
      if c.IsSet("to-address") {
        value := c.String("to-address")
        template.ToAddress = &value
      }
      if c.IsSet("body") {
        value := c.String("body")
        template.Body = &value
      }
	return template
}
type SendEmailActionResDto struct {
    QueueId   *string `json:"queueId" yaml:"queueId"       `
    // Datenano also has a text representation
}
func ( x * SendEmailActionResDto) RootObjectName() string {
	return "workspaces"
}
type sendEmailActionImpSig func(
    req *SendEmailActionReqDto, 
    q QueryDSL) (*SendEmailActionResDto,
    *IError,
)
var SendEmailActionImp sendEmailActionImpSig
func SendEmailActionFn(
    req *SendEmailActionReqDto, 
    q QueryDSL,
) (
    *SendEmailActionResDto,
    *IError,
) {
    if SendEmailActionImp == nil {
        return nil,  nil
    }
    return SendEmailActionImp(req,  q)
}
var SendEmailActionCmd cli.Command = cli.Command{
	Name:  "email",
	Usage: "Send a email using default root notification configuration",
	Flags: SendEmailCommonCliFlagsOptional,
	Action: func(c *cli.Context) {
		query := CommonCliQueryDSLBuilderAuthorize(c, SendEmailSecurityModel)
		dto := CastSendEmailFromCli(c)
		result, err := SendEmailActionFn(dto, query)
		HandleActionInCli(c, result, err, map[string]map[string]string{})
	},
}
var SendEmailWithProviderSecurityModel *SecurityModel = nil
type SendEmailWithProviderActionReqDto struct {
    EmailProvider   *  EmailProviderEntity `json:"emailProvider" yaml:"emailProvider"    gorm:"foreignKey:EmailProviderId;references:UniqueId"     `
    // Datenano also has a text representation
        EmailProviderId *string `json:"emailProviderId" yaml:"emailProviderId"`
    ToAddress   *string `json:"toAddress" yaml:"toAddress"  validate:"required"       `
    // Datenano also has a text representation
    Body   *string `json:"body" yaml:"body"  validate:"required"       `
    // Datenano also has a text representation
}
func ( x * SendEmailWithProviderActionReqDto) RootObjectName() string {
	return "workspaces"
}
var SendEmailWithProviderCommonCliFlagsOptional = []cli.Flag{
    &cli.StringFlag{
      Name:     "email-provider-id",
      Required: false,
      Usage:    "emailProvider",
    },
    &cli.StringFlag{
      Name:     "to-address",
      Required: true,
      Usage:    "toAddress",
    },
    &cli.StringFlag{
      Name:     "body",
      Required: true,
      Usage:    "body",
    },
}
func SendEmailWithProviderActionReqValidator(dto *SendEmailWithProviderActionReqDto) *IError {
    err := CommonStructValidatorPointer(dto, false)
    return err
  }
func CastSendEmailWithProviderFromCli (c *cli.Context) *SendEmailWithProviderActionReqDto {
	template := &SendEmailWithProviderActionReqDto{}
      if c.IsSet("email-provider-id") {
        value := c.String("email-provider-id")
        template.EmailProviderId = &value
      }
      if c.IsSet("to-address") {
        value := c.String("to-address")
        template.ToAddress = &value
      }
      if c.IsSet("body") {
        value := c.String("body")
        template.Body = &value
      }
	return template
}
type SendEmailWithProviderActionResDto struct {
    QueueId   *string `json:"queueId" yaml:"queueId"       `
    // Datenano also has a text representation
}
func ( x * SendEmailWithProviderActionResDto) RootObjectName() string {
	return "workspaces"
}
type sendEmailWithProviderActionImpSig func(
    req *SendEmailWithProviderActionReqDto, 
    q QueryDSL) (*SendEmailWithProviderActionResDto,
    *IError,
)
var SendEmailWithProviderActionImp sendEmailWithProviderActionImpSig
func SendEmailWithProviderActionFn(
    req *SendEmailWithProviderActionReqDto, 
    q QueryDSL,
) (
    *SendEmailWithProviderActionResDto,
    *IError,
) {
    if SendEmailWithProviderActionImp == nil {
        return nil,  nil
    }
    return SendEmailWithProviderActionImp(req,  q)
}
var SendEmailWithProviderActionCmd cli.Command = cli.Command{
	Name:  "emailp",
	Usage: "Send a text message using an specific gsm provider",
	Flags: SendEmailWithProviderCommonCliFlagsOptional,
	Action: func(c *cli.Context) {
		query := CommonCliQueryDSLBuilderAuthorize(c, SendEmailWithProviderSecurityModel)
		dto := CastSendEmailWithProviderFromCli(c)
		result, err := SendEmailWithProviderActionFn(dto, query)
		HandleActionInCli(c, result, err, map[string]map[string]string{})
	},
}
var InviteToWorkspaceSecurityModel *SecurityModel = nil
type inviteToWorkspaceActionImpSig func(
    req *WorkspaceInviteEntity, 
    q QueryDSL) (*WorkspaceInviteEntity,
    *IError,
)
var InviteToWorkspaceActionImp inviteToWorkspaceActionImpSig
func InviteToWorkspaceActionFn(
    req *WorkspaceInviteEntity, 
    q QueryDSL,
) (
    *WorkspaceInviteEntity,
    *IError,
) {
    if InviteToWorkspaceActionImp == nil {
        return nil,  nil
    }
    return InviteToWorkspaceActionImp(req,  q)
}
var InviteToWorkspaceActionCmd cli.Command = cli.Command{
	Name:  "invite",
	Usage: "Invite a new person (either a user, with passport or without passport)",
	Flags: WorkspaceInviteCommonCliFlagsOptional,
	Action: func(c *cli.Context) {
		query := CommonCliQueryDSLBuilderAuthorize(c, InviteToWorkspaceSecurityModel)
		dto := CastWorkspaceInviteFromCli(c)
		result, err := InviteToWorkspaceActionFn(dto, query)
		HandleActionInCli(c, result, err, map[string]map[string]string{})
	},
}
var GsmSendSmsSecurityModel *SecurityModel = nil
type GsmSendSmsActionReqDto struct {
    ToNumber   *string `json:"toNumber" yaml:"toNumber"  validate:"required"       `
    // Datenano also has a text representation
    Body   *string `json:"body" yaml:"body"  validate:"required"       `
    // Datenano also has a text representation
}
func ( x * GsmSendSmsActionReqDto) RootObjectName() string {
	return "workspaces"
}
var GsmSendSmsCommonCliFlagsOptional = []cli.Flag{
    &cli.StringFlag{
      Name:     "to-number",
      Required: true,
      Usage:    "toNumber",
    },
    &cli.StringFlag{
      Name:     "body",
      Required: true,
      Usage:    "body",
    },
}
func GsmSendSmsActionReqValidator(dto *GsmSendSmsActionReqDto) *IError {
    err := CommonStructValidatorPointer(dto, false)
    return err
  }
func CastGsmSendSmsFromCli (c *cli.Context) *GsmSendSmsActionReqDto {
	template := &GsmSendSmsActionReqDto{}
      if c.IsSet("to-number") {
        value := c.String("to-number")
        template.ToNumber = &value
      }
      if c.IsSet("body") {
        value := c.String("body")
        template.Body = &value
      }
	return template
}
type GsmSendSmsActionResDto struct {
    QueueId   *string `json:"queueId" yaml:"queueId"       `
    // Datenano also has a text representation
}
func ( x * GsmSendSmsActionResDto) RootObjectName() string {
	return "workspaces"
}
type gsmSendSmsActionImpSig func(
    req *GsmSendSmsActionReqDto, 
    q QueryDSL) (*GsmSendSmsActionResDto,
    *IError,
)
var GsmSendSmsActionImp gsmSendSmsActionImpSig
func GsmSendSmsActionFn(
    req *GsmSendSmsActionReqDto, 
    q QueryDSL,
) (
    *GsmSendSmsActionResDto,
    *IError,
) {
    if GsmSendSmsActionImp == nil {
        return nil,  nil
    }
    return GsmSendSmsActionImp(req,  q)
}
var GsmSendSmsActionCmd cli.Command = cli.Command{
	Name:  "sms",
	Usage: "Send a text message using default root notification configuration",
	Flags: GsmSendSmsCommonCliFlagsOptional,
	Action: func(c *cli.Context) {
		query := CommonCliQueryDSLBuilderAuthorize(c, GsmSendSmsSecurityModel)
		dto := CastGsmSendSmsFromCli(c)
		result, err := GsmSendSmsActionFn(dto, query)
		HandleActionInCli(c, result, err, map[string]map[string]string{})
	},
}
var GsmSendSmsWithProviderSecurityModel *SecurityModel = nil
type GsmSendSmsWithProviderActionReqDto struct {
    GsmProvider   *  GsmProviderEntity `json:"gsmProvider" yaml:"gsmProvider"    gorm:"foreignKey:GsmProviderId;references:UniqueId"     `
    // Datenano also has a text representation
        GsmProviderId *string `json:"gsmProviderId" yaml:"gsmProviderId"`
    ToNumber   *string `json:"toNumber" yaml:"toNumber"  validate:"required"       `
    // Datenano also has a text representation
    Body   *string `json:"body" yaml:"body"  validate:"required"       `
    // Datenano also has a text representation
}
func ( x * GsmSendSmsWithProviderActionReqDto) RootObjectName() string {
	return "workspaces"
}
var GsmSendSmsWithProviderCommonCliFlagsOptional = []cli.Flag{
    &cli.StringFlag{
      Name:     "gsm-provider-id",
      Required: false,
      Usage:    "gsmProvider",
    },
    &cli.StringFlag{
      Name:     "to-number",
      Required: true,
      Usage:    "toNumber",
    },
    &cli.StringFlag{
      Name:     "body",
      Required: true,
      Usage:    "body",
    },
}
func GsmSendSmsWithProviderActionReqValidator(dto *GsmSendSmsWithProviderActionReqDto) *IError {
    err := CommonStructValidatorPointer(dto, false)
    return err
  }
func CastGsmSendSmsWithProviderFromCli (c *cli.Context) *GsmSendSmsWithProviderActionReqDto {
	template := &GsmSendSmsWithProviderActionReqDto{}
      if c.IsSet("gsm-provider-id") {
        value := c.String("gsm-provider-id")
        template.GsmProviderId = &value
      }
      if c.IsSet("to-number") {
        value := c.String("to-number")
        template.ToNumber = &value
      }
      if c.IsSet("body") {
        value := c.String("body")
        template.Body = &value
      }
	return template
}
type GsmSendSmsWithProviderActionResDto struct {
    QueueId   *string `json:"queueId" yaml:"queueId"       `
    // Datenano also has a text representation
}
func ( x * GsmSendSmsWithProviderActionResDto) RootObjectName() string {
	return "workspaces"
}
type gsmSendSmsWithProviderActionImpSig func(
    req *GsmSendSmsWithProviderActionReqDto, 
    q QueryDSL) (*GsmSendSmsWithProviderActionResDto,
    *IError,
)
var GsmSendSmsWithProviderActionImp gsmSendSmsWithProviderActionImpSig
func GsmSendSmsWithProviderActionFn(
    req *GsmSendSmsWithProviderActionReqDto, 
    q QueryDSL,
) (
    *GsmSendSmsWithProviderActionResDto,
    *IError,
) {
    if GsmSendSmsWithProviderActionImp == nil {
        return nil,  nil
    }
    return GsmSendSmsWithProviderActionImp(req,  q)
}
var GsmSendSmsWithProviderActionCmd cli.Command = cli.Command{
	Name:  "smsp",
	Usage: "Send a text message using an specific gsm provider",
	Flags: GsmSendSmsWithProviderCommonCliFlagsOptional,
	Action: func(c *cli.Context) {
		query := CommonCliQueryDSLBuilderAuthorize(c, GsmSendSmsWithProviderSecurityModel)
		dto := CastGsmSendSmsWithProviderFromCli(c)
		result, err := GsmSendSmsWithProviderActionFn(dto, query)
		HandleActionInCli(c, result, err, map[string]map[string]string{})
	},
}
var ClassicSigninSecurityModel = &SecurityModel{
    ActionRequires: []PermissionInfo{ 
    },
}
type ClassicSigninActionReqDto struct {
    Value   *string `json:"value" yaml:"value"  validate:"required"       `
    // Datenano also has a text representation
    Password   *string `json:"password" yaml:"password"  validate:"required"       `
    // Datenano also has a text representation
}
func ( x * ClassicSigninActionReqDto) RootObjectName() string {
	return "workspaces"
}
var ClassicSigninCommonCliFlagsOptional = []cli.Flag{
    &cli.StringFlag{
      Name:     "value",
      Required: true,
      Usage:    "value",
    },
    &cli.StringFlag{
      Name:     "password",
      Required: true,
      Usage:    "password",
    },
}
func ClassicSigninActionReqValidator(dto *ClassicSigninActionReqDto) *IError {
    err := CommonStructValidatorPointer(dto, false)
    return err
  }
func CastClassicSigninFromCli (c *cli.Context) *ClassicSigninActionReqDto {
	template := &ClassicSigninActionReqDto{}
      if c.IsSet("value") {
        value := c.String("value")
        template.Value = &value
      }
      if c.IsSet("password") {
        value := c.String("password")
        template.Password = &value
      }
	return template
}
type classicSigninActionImpSig func(
    req *ClassicSigninActionReqDto, 
    q QueryDSL) (*UserSessionDto,
    *IError,
)
var ClassicSigninActionImp classicSigninActionImpSig
func ClassicSigninActionFn(
    req *ClassicSigninActionReqDto, 
    q QueryDSL,
) (
    *UserSessionDto,
    *IError,
) {
    if ClassicSigninActionImp == nil {
        return nil,  nil
    }
    return ClassicSigninActionImp(req,  q)
}
var ClassicSigninActionCmd cli.Command = cli.Command{
	Name:  "in",
	Usage: "Signin publicly to and account using class passports (email, password)",
	Flags: ClassicSigninCommonCliFlagsOptional,
	Action: func(c *cli.Context) {
		query := CommonCliQueryDSLBuilderAuthorize(c, ClassicSigninSecurityModel)
		dto := CastClassicSigninFromCli(c)
		result, err := ClassicSigninActionFn(dto, query)
		HandleActionInCli(c, result, err, map[string]map[string]string{})
	},
}
var ClassicSignupSecurityModel = &SecurityModel{
    ActionRequires: []PermissionInfo{ 
    },
}
type ClassicSignupActionReqDto struct {
    Value   *string `json:"value" yaml:"value"  validate:"required"       `
    // Datenano also has a text representation
    Type   *string `json:"type" yaml:"type"  validate:"required"       `
    // Datenano also has a text representation
    Password   *string `json:"password" yaml:"password"  validate:"required"       `
    // Datenano also has a text representation
    FirstName   *string `json:"firstName" yaml:"firstName"  validate:"required"       `
    // Datenano also has a text representation
    LastName   *string `json:"lastName" yaml:"lastName"  validate:"required"       `
    // Datenano also has a text representation
    InviteId   *string `json:"inviteId" yaml:"inviteId"       `
    // Datenano also has a text representation
    PublicJoinKeyId   *string `json:"publicJoinKeyId" yaml:"publicJoinKeyId"       `
    // Datenano also has a text representation
    WorkspaceTypeId   *string `json:"workspaceTypeId" yaml:"workspaceTypeId"  validate:"required"       `
    // Datenano also has a text representation
}
func ( x * ClassicSignupActionReqDto) RootObjectName() string {
	return "workspaces"
}
var ClassicSignupCommonCliFlagsOptional = []cli.Flag{
    &cli.StringFlag{
      Name:     "value",
      Required: true,
      Usage:    "value",
    },
    &cli.StringFlag{
      Name:     "type",
      Required: true,
      Usage:    "One of: 'phonenumber', 'email'",
    },
    &cli.StringFlag{
      Name:     "password",
      Required: true,
      Usage:    "password",
    },
    &cli.StringFlag{
      Name:     "first-name",
      Required: true,
      Usage:    "firstName",
    },
    &cli.StringFlag{
      Name:     "last-name",
      Required: true,
      Usage:    "lastName",
    },
    &cli.StringFlag{
      Name:     "invite-id",
      Required: false,
      Usage:    "inviteId",
    },
    &cli.StringFlag{
      Name:     "public-join-key-id",
      Required: false,
      Usage:    "publicJoinKeyId",
    },
    &cli.StringFlag{
      Name:     "workspace-type-id",
      Required: true,
      Usage:    "workspaceTypeId",
    },
}
func ClassicSignupActionReqValidator(dto *ClassicSignupActionReqDto) *IError {
    err := CommonStructValidatorPointer(dto, false)
    return err
  }
func CastClassicSignupFromCli (c *cli.Context) *ClassicSignupActionReqDto {
	template := &ClassicSignupActionReqDto{}
      if c.IsSet("value") {
        value := c.String("value")
        template.Value = &value
      }
      if c.IsSet("type") {
        value := c.String("type")
        template.Type = &value
      }
      if c.IsSet("password") {
        value := c.String("password")
        template.Password = &value
      }
      if c.IsSet("first-name") {
        value := c.String("first-name")
        template.FirstName = &value
      }
      if c.IsSet("last-name") {
        value := c.String("last-name")
        template.LastName = &value
      }
      if c.IsSet("invite-id") {
        value := c.String("invite-id")
        template.InviteId = &value
      }
      if c.IsSet("public-join-key-id") {
        value := c.String("public-join-key-id")
        template.PublicJoinKeyId = &value
      }
      if c.IsSet("workspace-type-id") {
        value := c.String("workspace-type-id")
        template.WorkspaceTypeId = &value
      }
	return template
}
type classicSignupActionImpSig func(
    req *ClassicSignupActionReqDto, 
    q QueryDSL) (*UserSessionDto,
    *IError,
)
var ClassicSignupActionImp classicSignupActionImpSig
func ClassicSignupActionFn(
    req *ClassicSignupActionReqDto, 
    q QueryDSL,
) (
    *UserSessionDto,
    *IError,
) {
    if ClassicSignupActionImp == nil {
        return nil,  nil
    }
    return ClassicSignupActionImp(req,  q)
}
var ClassicSignupActionCmd cli.Command = cli.Command{
	Name:  "up",
	Usage: "Signup a user into system via public access (aka website visitors) using either email or phone number",
	Flags: ClassicSignupCommonCliFlagsOptional,
	Action: func(c *cli.Context) {
		query := CommonCliQueryDSLBuilderAuthorize(c, ClassicSignupSecurityModel)
		dto := CastClassicSignupFromCli(c)
		result, err := ClassicSignupActionFn(dto, query)
		HandleActionInCli(c, result, err, map[string]map[string]string{})
	},
}
var CreateWorkspaceSecurityModel *SecurityModel = nil
type CreateWorkspaceActionReqDto struct {
    Name   *string `json:"name" yaml:"name"       `
    // Datenano also has a text representation
    Workspace   *  WorkspaceEntity `json:"workspace" yaml:"workspace"    gorm:"foreignKey:WorkspaceId;references:UniqueId"     `
    // Datenano also has a text representation
    WorkspaceId   *string `json:"workspaceId" yaml:"workspaceId"       `
    // Datenano also has a text representation
}
func ( x * CreateWorkspaceActionReqDto) RootObjectName() string {
	return "workspaces"
}
var CreateWorkspaceCommonCliFlagsOptional = []cli.Flag{
    &cli.StringFlag{
      Name:     "name",
      Required: false,
      Usage:    "name",
    },
    &cli.StringFlag{
      Name:     "workspace-id",
      Required: false,
      Usage:    "workspace",
    },
    &cli.StringFlag{
      Name:     "workspace-id",
      Required: false,
      Usage:    "workspaceId",
    },
}
func CreateWorkspaceActionReqValidator(dto *CreateWorkspaceActionReqDto) *IError {
    err := CommonStructValidatorPointer(dto, false)
    return err
  }
func CastCreateWorkspaceFromCli (c *cli.Context) *CreateWorkspaceActionReqDto {
	template := &CreateWorkspaceActionReqDto{}
      if c.IsSet("name") {
        value := c.String("name")
        template.Name = &value
      }
      if c.IsSet("workspace-id") {
        value := c.String("workspace-id")
        template.WorkspaceId = &value
      }
      if c.IsSet("workspace-id") {
        value := c.String("workspace-id")
        template.WorkspaceId = &value
      }
	return template
}
type createWorkspaceActionImpSig func(
    req *CreateWorkspaceActionReqDto, 
    q QueryDSL) (*WorkspaceEntity,
    *IError,
)
var CreateWorkspaceActionImp createWorkspaceActionImpSig
func CreateWorkspaceActionFn(
    req *CreateWorkspaceActionReqDto, 
    q QueryDSL,
) (
    *WorkspaceEntity,
    *IError,
) {
    if CreateWorkspaceActionImp == nil {
        return nil,  nil
    }
    return CreateWorkspaceActionImp(req,  q)
}
var CreateWorkspaceActionCmd cli.Command = cli.Command{
	Name:  "create-workspace",
	Usage: "",
	Flags: CreateWorkspaceCommonCliFlagsOptional,
	Action: func(c *cli.Context) {
		query := CommonCliQueryDSLBuilderAuthorize(c, CreateWorkspaceSecurityModel)
		dto := CastCreateWorkspaceFromCli(c)
		result, err := CreateWorkspaceActionFn(dto, query)
		HandleActionInCli(c, result, err, map[string]map[string]string{})
	},
}
var CheckClassicPassportSecurityModel = &SecurityModel{
    ActionRequires: []PermissionInfo{ 
    },
}
type CheckClassicPassportActionReqDto struct {
    Value   *string `json:"value" yaml:"value"  validate:"required"       `
    // Datenano also has a text representation
}
func ( x * CheckClassicPassportActionReqDto) RootObjectName() string {
	return "workspaces"
}
var CheckClassicPassportCommonCliFlagsOptional = []cli.Flag{
    &cli.StringFlag{
      Name:     "value",
      Required: true,
      Usage:    "value",
    },
}
func CheckClassicPassportActionReqValidator(dto *CheckClassicPassportActionReqDto) *IError {
    err := CommonStructValidatorPointer(dto, false)
    return err
  }
func CastCheckClassicPassportFromCli (c *cli.Context) *CheckClassicPassportActionReqDto {
	template := &CheckClassicPassportActionReqDto{}
      if c.IsSet("value") {
        value := c.String("value")
        template.Value = &value
      }
	return template
}
type CheckClassicPassportActionResDto struct {
    Exists   *bool `json:"exists" yaml:"exists"       `
    // Datenano also has a text representation
}
func ( x * CheckClassicPassportActionResDto) RootObjectName() string {
	return "workspaces"
}
type checkClassicPassportActionImpSig func(
    req *CheckClassicPassportActionReqDto, 
    q QueryDSL) (*CheckClassicPassportActionResDto,
    *IError,
)
var CheckClassicPassportActionImp checkClassicPassportActionImpSig
func CheckClassicPassportActionFn(
    req *CheckClassicPassportActionReqDto, 
    q QueryDSL,
) (
    *CheckClassicPassportActionResDto,
    *IError,
) {
    if CheckClassicPassportActionImp == nil {
        return nil,  nil
    }
    return CheckClassicPassportActionImp(req,  q)
}
var CheckClassicPassportActionCmd cli.Command = cli.Command{
	Name:  "ccp",
	Usage: "Checks if a classic passport (email, phone) exists or not, used in multi step authentication",
	Flags: CheckClassicPassportCommonCliFlagsOptional,
	Action: func(c *cli.Context) {
		query := CommonCliQueryDSLBuilderAuthorize(c, CheckClassicPassportSecurityModel)
		dto := CastCheckClassicPassportFromCli(c)
		result, err := CheckClassicPassportActionFn(dto, query)
		HandleActionInCli(c, result, err, map[string]map[string]string{})
	},
}
var ClassicPassportOtpSecurityModel = &SecurityModel{
    ActionRequires: []PermissionInfo{ 
    },
}
type ClassicPassportOtpActionReqDto struct {
    Value   *string `json:"value" yaml:"value"  validate:"required"       `
    // Datenano also has a text representation
    Otp   *string `json:"otp" yaml:"otp"       `
    // Datenano also has a text representation
}
func ( x * ClassicPassportOtpActionReqDto) RootObjectName() string {
	return "workspaces"
}
var ClassicPassportOtpCommonCliFlagsOptional = []cli.Flag{
    &cli.StringFlag{
      Name:     "value",
      Required: true,
      Usage:    "value",
    },
    &cli.StringFlag{
      Name:     "otp",
      Required: false,
      Usage:    "otp",
    },
}
func ClassicPassportOtpActionReqValidator(dto *ClassicPassportOtpActionReqDto) *IError {
    err := CommonStructValidatorPointer(dto, false)
    return err
  }
func CastClassicPassportOtpFromCli (c *cli.Context) *ClassicPassportOtpActionReqDto {
	template := &ClassicPassportOtpActionReqDto{}
      if c.IsSet("value") {
        value := c.String("value")
        template.Value = &value
      }
      if c.IsSet("otp") {
        value := c.String("otp")
        template.Otp = &value
      }
	return template
}
type ClassicPassportOtpActionResDto struct {
    SuspendUntil   *int64 `json:"suspendUntil" yaml:"suspendUntil"       `
    // Datenano also has a text representation
    Session   *  UserSessionDto `json:"session" yaml:"session"    gorm:"foreignKey:SessionId;references:UniqueId"     `
    // Datenano also has a text representation
        SessionId *string `json:"sessionId" yaml:"sessionId"`
    ValidUntil   *int64 `json:"validUntil" yaml:"validUntil"       `
    // Datenano also has a text representation
    BlockedUntil   *int64 `json:"blockedUntil" yaml:"blockedUntil"       `
    // Datenano also has a text representation
    SecondsToUnblock   *int64 `json:"secondsToUnblock" yaml:"secondsToUnblock"       `
    // Datenano also has a text representation
}
func ( x * ClassicPassportOtpActionResDto) RootObjectName() string {
	return "workspaces"
}
type classicPassportOtpActionImpSig func(
    req *ClassicPassportOtpActionReqDto, 
    q QueryDSL) (*ClassicPassportOtpActionResDto,
    *IError,
)
var ClassicPassportOtpActionImp classicPassportOtpActionImpSig
func ClassicPassportOtpActionFn(
    req *ClassicPassportOtpActionReqDto, 
    q QueryDSL,
) (
    *ClassicPassportOtpActionResDto,
    *IError,
) {
    if ClassicPassportOtpActionImp == nil {
        return nil,  nil
    }
    return ClassicPassportOtpActionImp(req,  q)
}
var ClassicPassportOtpActionCmd cli.Command = cli.Command{
	Name:  "otp",
	Usage: "Authenticate the user publicly for classic methods using communication service, such as sms, call, or email",
	Flags: ClassicPassportOtpCommonCliFlagsOptional,
	Action: func(c *cli.Context) {
		query := CommonCliQueryDSLBuilderAuthorize(c, ClassicPassportOtpSecurityModel)
		dto := CastClassicPassportOtpFromCli(c)
		result, err := ClassicPassportOtpActionFn(dto, query)
		HandleActionInCli(c, result, err, map[string]map[string]string{})
	},
}
func WorkspacesCustomActions() []Module2Action {
	routes := []Module2Action{
		{
			Method: "POST",
			Url:    "/email/send",
            SecurityModel: SendEmailSecurityModel,
            Group: "WorkspacesCustom",
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
                    // POST_ONE - post
                        HttpPostEntity(c, SendEmailActionFn)
                },
			},
			Format:         "POST_ONE",
			Action:         SendEmailActionFn,
			ResponseEntity: &SendEmailActionResDto{},
			RequestEntity: &SendEmailActionReqDto{},
		},
		{
			Method: "POST",
			Url:    "/emailProvider/send",
            SecurityModel: SendEmailWithProviderSecurityModel,
            Group: "WorkspacesCustom",
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
                    // POST_ONE - post
                        HttpPostEntity(c, SendEmailWithProviderActionFn)
                },
			},
			Format:         "POST_ONE",
			Action:         SendEmailWithProviderActionFn,
			ResponseEntity: &SendEmailWithProviderActionResDto{},
			RequestEntity: &SendEmailWithProviderActionReqDto{},
		},
		{
			Method: "POST",
			Url:    "/workspace/invite",
            SecurityModel: InviteToWorkspaceSecurityModel,
            Group: "WorkspacesCustom",
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
                    // POST_ONE - post
                        HttpPostEntity(c, InviteToWorkspaceActionFn)
                },
			},
			Format:         "POST_ONE",
			Action:         InviteToWorkspaceActionFn,
			ResponseEntity: &WorkspaceInviteEntity{},
			RequestEntity: &WorkspaceInviteEntity{},
		},
		{
			Method: "POST",
			Url:    "/gsm/send/sms",
            SecurityModel: GsmSendSmsSecurityModel,
            Group: "WorkspacesCustom",
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
                    // POST_ONE - post
                        HttpPostEntity(c, GsmSendSmsActionFn)
                },
			},
			Format:         "POST_ONE",
			Action:         GsmSendSmsActionFn,
			ResponseEntity: &GsmSendSmsActionResDto{},
			RequestEntity: &GsmSendSmsActionReqDto{},
		},
		{
			Method: "POST",
			Url:    "/gsmProvider/send/sms",
            SecurityModel: GsmSendSmsWithProviderSecurityModel,
            Group: "WorkspacesCustom",
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
                    // POST_ONE - post
                        HttpPostEntity(c, GsmSendSmsWithProviderActionFn)
                },
			},
			Format:         "POST_ONE",
			Action:         GsmSendSmsWithProviderActionFn,
			ResponseEntity: &GsmSendSmsWithProviderActionResDto{},
			RequestEntity: &GsmSendSmsWithProviderActionReqDto{},
		},
		{
			Method: "POST",
			Url:    "/passports/signin/classic",
            SecurityModel: ClassicSigninSecurityModel,
            Group: "WorkspacesCustom",
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
                    // POST_ONE - post
                        HttpPostEntity(c, ClassicSigninActionFn)
                },
			},
			Format:         "POST_ONE",
			Action:         ClassicSigninActionFn,
			ResponseEntity: &UserSessionDto{},
			RequestEntity: &ClassicSigninActionReqDto{},
		},
		{
			Method: "POST",
			Url:    "/passports/signup/classic",
            SecurityModel: ClassicSignupSecurityModel,
            Group: "WorkspacesCustom",
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
                    // POST_ONE - post
                        HttpPostEntity(c, ClassicSignupActionFn)
                },
			},
			Format:         "POST_ONE",
			Action:         ClassicSignupActionFn,
			ResponseEntity: &UserSessionDto{},
			RequestEntity: &ClassicSignupActionReqDto{},
		},
		{
			Method: "POST",
			Url:    "/workspaces/create",
            SecurityModel: CreateWorkspaceSecurityModel,
            Group: "WorkspacesCustom",
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
                    // POST_ONE - post
                        HttpPostEntity(c, CreateWorkspaceActionFn)
                },
			},
			Format:         "POST_ONE",
			Action:         CreateWorkspaceActionFn,
			ResponseEntity: &WorkspaceEntity{},
			RequestEntity: &CreateWorkspaceActionReqDto{},
		},
		{
			Method: "POST",
			Url:    "/workspace/passport/check",
            SecurityModel: CheckClassicPassportSecurityModel,
            Group: "WorkspacesCustom",
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
                    // POST_ONE - post
                        HttpPostEntity(c, CheckClassicPassportActionFn)
                },
			},
			Format:         "POST_ONE",
			Action:         CheckClassicPassportActionFn,
			ResponseEntity: &CheckClassicPassportActionResDto{},
			RequestEntity: &CheckClassicPassportActionReqDto{},
		},
		{
			Method: "POST",
			Url:    "/workspace/passport/otp",
            SecurityModel: ClassicPassportOtpSecurityModel,
            Group: "WorkspacesCustom",
			Handlers: []gin.HandlerFunc{
				func(c *gin.Context) {
                    // POST_ONE - post
                        HttpPostEntity(c, ClassicPassportOtpActionFn)
                },
			},
			Format:         "POST_ONE",
			Action:         ClassicPassportOtpActionFn,
			ResponseEntity: &ClassicPassportOtpActionResDto{},
			RequestEntity: &ClassicPassportOtpActionReqDto{},
		},
	}
	return routes
}
var WorkspacesCustomActionsCli = []cli.Command {
    SendEmailActionCmd,
    SendEmailWithProviderActionCmd,
    InviteToWorkspaceActionCmd,
    GsmSendSmsActionCmd,
    GsmSendSmsWithProviderActionCmd,
    ClassicSigninActionCmd,
    ClassicSignupActionCmd,
    CreateWorkspaceActionCmd,
    CheckClassicPassportActionCmd,
    ClassicPassportOtpActionCmd,
}