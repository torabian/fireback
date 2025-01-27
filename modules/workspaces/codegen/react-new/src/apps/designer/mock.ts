export const mockYaml = `name: workspaces
meta-workspace: true
config:
  - name: name
    description: Project name, which could be used in couple of places. better lower case - only.
  - name: dbName
    default: ":memory:"
    description: Database name for vendors which provide database names, such as mysql. Filename on disk for sqlite.
  - name: dbPort
    description: Database port for those which are having a port, 3306 on mysql for example
    type: int64
  - name: driveEnabled
    type: bool
    description: Drive is a mechanism to have file upload and download, inlining integrated into the fireback
    default: true
  - name: dbDsn
    description: Connection dsn to database. Some databases allow connection using a string with all credentials and configs. This has hight priority, if set other details will be ignored.
  - name: dbHost
    description: Database host, such as localhost, or 127.0.0.1
  - name: dbUsername
    description: Database username for connection, such as root.
  - name: dbPassword
    description: Database password for connection. "Can be" empty if there is no password
  - name: ginMode
    description: Gin framework mode, which could be 'test', 'debug', 'release'
    enum:
      - value: test
      - value: debug
      - value: release
  - name: storage
    description: This is the storage url which files will be uploaded to
  - name: dbVendor
    default: sqlite
    description: Database vendor name, such as sqlite, mysql, or any other supported database.
  - name: stdOut
    description: Writes the logs instead of std out into these log files.
  - name: workerAddress
    description: This is the url (host and port) of a queue service. If not set, we use the internal queue system
    default: 127.0.0.1:6379
  - name: workerConcurrency
    description: How many tasks worker can take concurrently
    default: 10
    type: int
  - name: stdErr
    description: Writes the errors instead of std err into these log files.
  - name: tusPort
    description: Resumable file upload server port.
  - name: cliToken
    description: Authorization token for cli apps, to access resoruces similar on http api
  - name: cliRegion
    description: Region, for example us or pl
    default: us
  - name: cliLanguage
    description: Language of the cli operations, for example en or pl
    default: en
  - name: cliWorkspace
    description: Selected workspace in the cli context.
  - name: port
    description: The port which application would be lifted
    default: 4500
    type: int64
  - name: host
    default: localhost
    description: Application host which http server will be lifted
  - name: macIdentifier
    default: fireback
    description: Used name for installing app as system service on macos installers
  - name: debianIdentifier
    description: Used name for installing app as system service on ubuntu installers
    default: fireback
  - name: windowsIdentifier
    description: Used name for installing app as system service on windows installers
    default: fireback

actions:
  - name: importUser
    url: /user/import
    cliName: userImport
    method: post
    description: Imports users, and creates their passports, and all details
    in:
      fields:
      - name: path
        type: string
    out:
      dto: OkayResponseDto
  - name: sendEmail
    url: /email/send
    cliName: email
    method: post
    description: Send a email using default root notification configuration
    in:
      
      fields:
        - name: toAddress
          validate: required
          type: string

        - name: body
          validate: required
          type: string
        
        
    out:
      fields:
      - name: queueId
        type: string

  - name: sendEmailWithProvider

    url: /emailProvider/send
    cliName: emailp
    method: post
    description: Send a text message using an specific gsm provider
    in:
      fields:
        - name: emailProvider
          type: one
          target: EmailProviderEntity
        - name: toAddress
          validate: required
          type: string
        - name: body
          validate: required
          type: string
    out:
      fields:
      - name: queueId
        type: string
  - name: inviteToWorkspace
    url: /workspace/invite
    cliName: invite
    method: post
    description: Invite a new person (either a user, with passport or without passport)
    in:
      entity: WorkspaceInviteEntity
    out:
      entity: WorkspaceInviteEntity
  
  - name: gsmSendSms
    url: /gsm/send/sms
    cliName: sms
    method: post
    description: Send a text message using default root notification configuration
    in:
      fields:
        - name: toNumber
          validate: required
          type: string
        - name: body
          validate: required
          type: string
    out:
      fields:
      - name: queueId
        type: string
  - name: gsmSendSmsWithProvider
    url: /gsmProvider/send/sms
    cliName: smsp
    method: post
    description: Send a text message using an specific gsm provider
    in:
      fields:
        - name: gsmProvider
          type: one
          target: GsmProviderEntity
        - name: toNumber
          validate: required
          type: string
        - name: body
          validate: required
          type: string
    out:
      fields:
      - name: queueId
        type: string
  - name: classicSignin
    url: /passports/signin/classic
    cliName: in
    description: Signin publicly to and account using class passports (email, password)
    method: post
    in:
      fields:
      - name: value
        type: string
        validate: required
      - name: password
        type: string
        validate: required
    out:
      dto: UserSessionDto

  - name: classicSignup
    url: /passports/signup/classic
    cliName: up
    description: Signup a user into system via public access (aka website visitors) using either email or phone number
    method: post
    in:
      fields:
      - name: value
        validate: required
        type: string
      - name: type
        type: enum
        of:
          - k: phonenumber
          - k: email
        validate: required
      - name: password
        type: string
        validate: required
      - name: firstName
        type: string
        validate: required
      - name: lastName
        type: string
        validate: required
      - name: inviteId
        type: string
      - name: publicJoinKeyId
        type: string
      - name: workspaceTypeId
        type: string
        validate: required
    out:
      dto: UserSessionDto
      
  - name: createWorkspace
    method: post
    url: /workspaces/create
    in:
      fields:
      - name: name
        type: string
      - name: workspace
        type: one
        target: WorkspaceEntity
      - name: workspaceId
        type: string
    out:
      entity: WorkspaceEntity
    fn: CreateWorkspaceAction

  # Checks if the passport or email address exists
  - name: checkClassicPassport
    method: post
    cliName: ccp
    description: Checks if a classic passport (email, phone) exists or not, used in multi step authentication
    url: /workspace/passport/check
    in:
      fields:
      - name: value
        type: string
        validate: required
    out:
      fields:
      - name: exists
        type: bool

  - name: classicPassportOtp
    description: Authenticate the user publicly for classic methods using communication service, such as sms, call, or email
    method: post
    cliName: otp
    url: /workspace/passport/otp
    in:
      fields:
      - name: value
        type: string
        validate: required
      - name: otp
        type: string
    out:
      fields:
      - name: suspendUntil
        type: int64
      - name: session
        type: one
        target: UserSessionDto
      - name: validUntil
        type: int64
      - name: blockedUntil
        type: int64
      - name: secondsToUnblock
        type: int64
dtos:
  - name: userImport
    fields:
    - name: avatar
      type: string
    - name: passports
      type: array
      fields:
      - name: value
        type: string
      - name: password
        type: string
    - name: person
      type: one
      target: PersonEntity
    - name: address
      type: object
      fields:
      - name: street
        type: string
      - name: zipCode
        type: string
      - name: city
        type: string
      - name: country
        type: string
  - name: userRoleWorkspacePermission
    fields:
    - name: workspaceId
      type: string
    - name: userId
      type: string
    - name: roleId
      type: string
    - name: capabilityId
      type: string
    - name: type
      type: string

  - name: permissionInfo
    fields:
    - name: name
      type: string
    - name: description
      type: string
    - name: completeKey
      type: string
    
  - name: userRoleWorkspace
    fields:
    - name: roleId
      type: string  
    - name: capabilities
      type: arrayP
      primitive: string

  - name: importRequest
    fields:
    - name: file
      type: string
  - name: okayResponse
  - name: testMail
    fields:
    - name: senderId
      type: string
    - name: toName
      type: string
    - name: toEmail
      type: string
    - name: subject
      type: string
    - name: content
      type: string

  - name: assignRole
    fields:
    - name: roleId
      type: string
    - name: userId
      type: string
    - name: visibility
      type: string
    - name: updated
      type: int64
    - name: created
      type: int64
   

  - name: exchangeKeyInformation
    fields:
    - name: key
      type: string
    - name: visibility
      type: string

  - name: authResult
    fields:
    - name: workspaceId
      type: string
    - name: userRoleWorkspacePermissions
      type: many2many
      target: UserRoleWorkspacePermissionDto
    - name: internalSql
      type: string
    - name: userId
      type: string
    - name: userHas
      type: arrayP
      primitive: string
    - name: workspaceHas
      type: arrayP
      primitive: string
    - name: user
      type: one
      target: UserEntity
    - name: accessLevel
      type: one
      target: UserAccessLevelDto
  - name: authContext
    fields:
    - name: skipWorkspaceId
      type: bool
    - name: workspaceId
      type: string
    - name: token
      type: string
    - name: capabilities
      primitive: PermissionInfo
      type: arrayP

  - name: reactiveSearchResult
    fields:
      - type: string
        name: uniqueId
      - type: string
        name: phrase
      - type: string
        name: icon
      - type: string
        name: description
      - type: string
        name: group
      - type: string
        name: uiLocation
      - type: string
        name: actionFn

  - name: userAccessLevel
    fields:
    - name: capabilities
      type: arrayP
      primitive: string
    - name: userRoleWorkspacePermissions
      type: many2many
      target: UserRoleWorkspacePermissionDto
    - name: workspaces
      type: arrayP
      primitive: string
    - name: SQL
      type: string

  - name: acceptInvite
    fields:
    - name: inviteUniqueId
      type: string
    - name: visibility
      type: string
    - name: updated
      type: int64
    - name: created
      type: int64
     
  - name: classicAuth
    fields:
    - name: value
      type: string
      validate: required
    - name: password
      type: string
      validate: required
    - name: firstName
      type: string
      validate: required
    - name: lastName
      type: string
      validate: required
    - name: inviteId
      type: string
    - name: publicJoinKeyId
      type: string
    - name: workspaceTypeId
      type: string
  - name: emailAccountSignin
    fields:
    - name: email
      type: string
      validate: required
    - name: password
      validate: required
      type: string
  - name: phoneNumberAccountCreation
    fields:
    - name: phoneNumber
      type: string
  - name: userSession
    fields:
    - name: passport
      type: one
      target: PassportEntity
    - name: token
      type: string
    - name: exchangeKey
      type: string
    - name: userWorkspaces
      type: many2many
      target: UserWorkspaceEntity
    - name: user
      type: one
      target: UserEntity
    - name: userId
      type: string
  - name: otpAuthenticate
    fields:
    - name: value
      type: string
      validate: required  
    - name: otp
      type: string
    - name: type
      type: string
      validate: required
    - name: password
      type: string
      validate: required
  - name: emailOtpResponse
    fields:
    - name: request
      type: one
      target: ForgetPasswordEntity 
    - name: userSession
      type: one
      target: UserSessionDto
  - name: resetEmail
    fields:
    - name: password
      type: string
entities:
  - name: file
    description: File manager, uploading files and actions related.
    capabilities:
      - name: "upload"
      - name: "replace"
      - name: "rename"
      - name: "share"
    fields:
      - name: name
        type: string
      - name: diskPath
        type: string
      - name: size
        type: int64
      - name: virtualPath
        type: string
      - name: type
        type: string  
      - name: variations
        type: array
        fields:
        - name: name
          type: string

  - name: tableViewSizing
    cliShort: tvs
      
    description: Used to store meta data about user tables (in front-end, or apps for example) about the size of the columns
    fields:
    - name: tableName
      type: string
      validate: required
    - name: sizes
      type: string
  - name: appMenu
    cte: true
    description: Manages the menus in the app, (for example tab views, sidebar items, etc.)
    fields:
    - name: label
      type: string
      translate: true
      recommended: true
    - name: href
      type: string
      recommended: true
    - name: icon
      type: string
      recommended: true
    - name: activeMatcher
      type: string
    - name: applyType
      type: string
    - name: capability
      type: one
      target: CapabilityEntity

  - name: backupTableMeta
    cliName: backup
    description: Keeps information about which tables to be used during backup (mostly internal)
    fields:
    - name: tableNameInDb
      type: string
    
  - name: notificationConfig
    distinctBy: workspace
    description: Configuration for the notifications used in the app, such as default gsm number, email senders, and many more
    cliShort: config
    fields:
    - name: cascadeToSubWorkspaces
      type: bool
    - name: forcedCascadeEmailProvider
      type: bool
    - name: generalEmailProvider
      type: one
      target: EmailProviderEntity
      allowCreate: false
    - name: generalGsmProvider
      type: one
      target: GsmProviderEntity
      allowCreate: false
    - name: inviteToWorkspaceContent
      type: string
      gorm: text
    - name: inviteToWorkspaceContentExcerpt
      gorm: text
      type: string
    - name: inviteToWorkspaceContentDefault
      gorm: text
      sql: false
      type: string
    - name: inviteToWorkspaceContentDefaultExcerpt
      type: string
      gorm: text
      sql: false
    - name: inviteToWorkspaceTitle
      type: string
    - name: inviteToWorkspaceTitleDefault
      sql: false
      type: string
    - name: inviteToWorkspaceSender
      type: one
      target: EmailSenderEntity
    - name: accountCenterEmailSender
      type: one
      target: EmailSenderEntity
    - name: forgetPasswordContent
      gorm: text
      type: string
    - name: forgetPasswordContentExcerpt
      gorm: text
      type: string
    - name: forgetPasswordContentDefault
      gorm: text
      type: string
      sql: false
    - name: forgetPasswordContentDefaultExcerpt
      gorm: text
      sql: false
      type: string
    - name: forgetPasswordTitle
      gorm: text
      type: string
    - name: forgetPasswordTitleDefault
      gorm: text
      sql: false
      type: string
    - name: forgetPasswordSender
      type: one
      target: EmailSenderEntity
    - name: acceptLanguage
      type: text
    - name: confirmEmailSender
      type: one
      target: EmailSenderEntity


    - name: confirmEmailContent
      gorm: text
      type: string
    - name: confirmEmailContentExcerpt
      gorm: text
      type: string
    - name: confirmEmailContentDefault
      type: string
      gorm: text
      sql: false
    - name: confirmEmailContentDefaultExcerpt
      gorm: text
      sql: false
      type: string
    - name: confirmEmailTitle
      type: string
    - name: confirmEmailTitleDefault
      type: string
      sql: false
    
  - name: passportMethod
    cliShort: method
    description: Login/Signup methods which are available in the app for different regions (Email, Phone Number, Google, etc)
    queryScope: public # Means there is no authentication while querying this
    fields:
      - name: name
        translate: true
        type: string
        validate: required
      - name: type
        type: string
        validate: required
      - name: region
        type: string
        validate: required

  - name: workspaceInvite
    description: Active invitations for non-users or already users to join an specific workspace
    fields:
    - name: coverLetter
      type: string
    - name: targetUserLocale
      type: string
    - name: value
      type: string
      validate: required
    - name: workspace
      type: one
      target: WorkspaceEntity
      validate: required
      allowCreate: false
    - name: firstName
      type: string
      validate: required
    - name: lastName
      validate: required
      type: string
    - name: used
      type: bool
    - name: role
      type: one
      validate: required
      target: RoleEntity
      allowCreate: false

  - name: pendingWorkspaceInvite
    fields:
    - name: value
      type: string
    - name: type
      type: string
    - name: coverLetter
      type: string
    - name: workspaceName
      type: string
    - name: role
      type: one
      target: RoleEntity
      allowCreate: false
  - name: preference
    fields:
    - name: timezone
      type: string
  - name: token
    fields:
    - name: user
      type: one
      allowCreate: false
      target: UserEntity
    - name: validUntil
      type: string
  - name: person
    fields: 
    - name: firstName
      type: string
      validate: required
    - name: lastName
      type: string
      validate: required
    - name: photo
      type: string
    - name: gender
      type: string
    - name: title
      type: string
    - name: birthDate
      type: date

  - name: userWorkspace
    cliShort: user
    description: Manage the workspaces that user belongs to (either its himselves or adding by invitation)
    postFormatter: UserWorkspacePostFormatter
    security:
      resolveStrategy: user
    gormMap:
      workspaceId: index:userworkspace_idx,unique
      userId: index:userworkspace_idx,unique
    fields:
    - name: user
      type: one
      target: UserEntity
    - name: workspace
      type: one
      target: WorkspaceEntity
    - name: userPermissions
      gorm: "-"
      sql: "-"
      type: arrayP
      primitive: string
    - name: rolePermission
      gorm: "-"
      sql: "-"
      type: arrayP
      primitive: UserRoleWorkspaceDto
    - name: workspacePermissions
      gorm: "-"
      sql: "-"
      type: arrayP
      primitive: string

  - name: workspaceRole
    cliShort: role
    description: Manage roles assigned to an specific workspace or created by the workspace itself
    fields:
    - name: userWorkspace
      type: one
      target: UserWorkspaceEntity
      idFieldGorm: index:workspacerole_idx,unique
    - name: role
      type: one
      target: RoleEntity
      idFieldGorm: index:workspacerole_idx,unique


  - name: user
    # noQuery: true
    description: Manage the users who are in the current app (root only)
    fields:
    - name: person
      type: one
      target: PersonEntity
      allowCreate: true
    - name: avatar
      type: string
      
    # - name: passports
    #   type: many2many
    #   target: PassportEntity
  - name: userProfile
    fields:
    - name: firstName
      type: string
    - name: lastName
      type: string
  - name: workspace
    cte: true
    # noQuery: true
    cliName: ws
    description: Fireback general user role, workspaces services.
    fields:
    - name: description
      type: string
    - name: name
      validate: required
      type: string
    - name: type
      type: one
      target: WorkspaceTypeEntity
      validate: required
    
  - name: role
    description: Manage roles within the workspaces, or root configuration
    fields:
    - name: name
      type: string
      validate: required,omitempty,min=1,max=200
    - name: capabilities
      type: many2many
      target: CapabilityEntity
      allowCreate: false

  - name: capability
    cliShort: cap
    description: Manage the capabilities inside the application, both builtin to core and custom defined ones
    fields:
    - name: name
      type: string
    - name: description
      type: string
      translate: true
  - name: workspaceConfig
    cliName: config
    distinctBy: workspace
    fields:
      # in some apps, we do not want to allow the server to allow anonymouse people create their own
      # workspace
      - name: disablePublicWorkspaceCreation
        type: int64
        default: 1
      - name: workspace
        type: one
        target: WorkspaceEntity
      - name: zoomClientId
        type: string
      - name: zoomClientSecret
        type: string
      - name: allowPublicToJoinTheWorkspace
        type: bool
  - name: gsmProvider
    fields:
      - name: apiKey
        type: string
      - name: mainSenderNumber
        type: string
        validate: required
      - name: type
        validate: required
        type: enum
        of:
          - k: url
          - k: terminal
          - k: mediana
      - name: invokeUrl
        type: string
      - name: invokeBody
        type: string

  # Predefined workspace types inside the project. Usually managed by root,
  # these types are only way to allow people create accounts
  - name: workspaceType
    cliName: type
    # distinctBy: workspace
    fields:
      - name: title
        translate: true
        type: string
        validate: required,omitempty,min=1,max=250
      - name: description
        translate: true
        type: string
        # validate: required,omitempty,min=1,max=1500
      - name: slug
        type: string
        # gorm: unique;not null;size:100;
        validate: required,omitempty,min=2,max=50
      - name: role
        type: one
        target: RoleEntity
  - name: emailProvider
    fields:
    - name: type
      validate: required
      type: enum
      of:
      - k: terminal
      - k: sendgrid
    - name: apiKey
      type: string
  - name: emailSender
    description: All emails going from the system need to have a virtual sender (nick name, email address, etc)
    fields:
    - name: fromName
      type: string
      validate: required
    - name: fromEmailAddress
      type: string
      unique: true
      gorm: unique
      validate: required
    - name: replyTo
      type: string
      validate: required
    - name: nickName
      validate: required
      type: string

  - name: phoneConfirmation
    fields:
    - name: user
      target: UserEntity
      allowCreate: false
      type: one
    - name: status
      type: string
    - name: phoneNumber
      type: string
    - name: key
      type: string
    - name: expiresAt
      type: string

  - name: publicJoinKey
    description: Joining to different workspaces using a public link directly
    fields:
    - name: role
      type: one
      allowCreate: false
      target: RoleEntity
    - name: workspace
      type: one
      allowCreate: false
      target: WorkspaceEntity
 
  - name: emailConfirmation
    fields:
    - name: user
      target: UserEntity
      allowCreate: false
      type: one
    - name: status
      type: string
    - name: email
      type: string
    - name: key
      type: string
    - name: expiresAt
      type: string


  - name: passport
    fields:
    - name: type
      type: string
      validate: required
    - name: user
      type: one
      target: UserEntity
    - name: value
      type: string
      validate: required
      unique: true
      gorm: unique
    - name: password
      type: string
      json: "-"
      yaml: "-"
    - name: confirmed
      type: bool
    - name: accessToken
      type:  string
  - name: regionalContent
    cliShort: rc
    description: Email templates, sms templates or other textual content which can be accessed.
    fields:
    - name: content
      type: html
      validate: required
    - name: region
      validate: required
      type: string
    - name: title
      type: string
    - name: languageId
      validate: required
      gorm: index:regional_content_index,unique
      type: string
    - name: keyGroup
      validate: required
      type: enum
      of:
      - k: SMS_OTP
        description: Used when an email would be sent with one time password
      - k: EMAIL_OTP
        description: Used when an sms would be sent with one time password
      gorm: index:regional_content_index,unique

  - name: forgetPassword
    fields:
    - name: user
      target: UserEntity
      allowCreate: false
      type: one
      json: false
    - name: passport
      target: PassportEntity
      allowCreate: false
      type: one
      json: false
    - name: status
      type: string
      json: false
    - name: validUntil
      type: datenano
    - name: blockedUntil
      type: datenano
    - name: secondsToUnblock
      type: int64
    - name: otp
      type: string
      json: false
    - name: recoveryAbsoluteUrl
      type: string
      json: false
      sql: false

messages:
  dataTypeDoesNotExistsInFireback:
    en: This data type does not exist in fireback. %name %location
  inviteToWorkspaceMailSenderMissing:
    en: We cannot send the invitation via email address, because sender email is not available, or not configurated.
    fa: امکان ارسال دعوت نامه از طریق ایمیل وجود ندارد، چون مدیریت تنظیمات لازم برای ایمیل را انجام نداده یا آن را برای شما محدود کرده است.
  userWhichHasThisTokenDoesNotExist:
    en: User which has this token does not exists
    fa: کاربری که با این دسترسی وارد شده بود وجود ندارد. لطفا دوباره به سیستم وارد شوید
  provideTokenInAuthorization:
    en: Request requires authroization, please make sure you are logged in, and have enough access level
    fa: شما باید توکن دسترسی را در بخش هدر و قسمت authorization وارد کنید
  userNotFoundOrDeleted:
    en: User not found, your account might be deleted, or access level has been reduced.
    fa: کاربر پیدا نشد ممکن است اکانت حذف شده باشد یا سطح دسترسی آن کاهش پیدا کرده باشد
  selectWorkspaceId:
    en: You need to select a correct workspace-id in header section
    fa: شما باید تیم یا ورک اسپیس را در بخش هدر با فیلد workspace-id تعیین کنید
  emailConfigurationMissing:
    en: Email configuration is not available
    fa:
  gsmConfigurationIsNotAvailable:
    en: GSM Services configuration is not available
    fa:
  emailConfigurationIsNotAvailable:
    en: Email configuration is not available
    fa:
  passportUserNotAvailable:
    en: User with this passport is not available at this moment
    fa:

  userDoesNotExist:
    en: User is not available.
    fa:
  alreadyConfirmed:
    en: Already confirmed
    fa:
  emailNotFound:
    en: Email is not found
    fa:
  invitationExpired:
    en: Invitation has been expired.
    fa:
  passwordRequired:
    en: Password is required
    fa:
  passportNotAvailable:
    en: This passport is not available. Please check credentials and try again
    fa:
  resetNotFound:
    en: Reset not found
    fa:
  passportNotFound:
    en: This passport is not available. Please check credentials and try again
    fa:
  otaRequestBlockedUntil:
    en: Request is blocked until.
    fa:
  emailIsNotConfigured:
    en: Email server is not configured
    fa:
  otpCodeInvalid:
    en: Otp code is invalid
    fa:
  invalidContent:
    en: Body content is not correct. You need a valid json.
    fa:
  bodyIsMissing:
    en: Body content is not correct. You need a valid json.
    fa:
  notEnoughPermission:
    en: You do not have enough permission for this section
  invalidExchangeKey:
    en:
  smsNotSent:
    en: Sending text message has failed.
  invokeUrlMissing:
    en: Invoking url is missing
  fileNotFound:
    en: File not found
  validationFailedOnSomeFields:
    en: Validation has failed on some fields
  `;
