# Fireback ‐ Ultimate backend framework

Backend in Golang - SQL Database - Native Android, Native IOS, Lightweight React web and desktop apps. **Idea is**, when you want build a software, there is a ready to use backend, with everything that you might need, and all that backends features, will be converted into SWIFT, Android, React RPC codes. It's coding, but 10-20 times faster.

## Getting started

### Bare minimum

1. You need to clone the repository.
2. `cd cmd/fireback-server`
3. `go run . init` and follow the instruction for database setup
4. `go run . start` will start the server.

### Practical setup

Using a run on save extension for VSCode ide, and moving the binary into executable path on mac, actually reduces a lot of time

1. Install run on save extension for VSCode.
2. Use Cmd+Shift+P and search for `Run task`
3. Run the `Reconfig` task, and follow instruction
4. It will build the binary, and copy the executable (at this moment only on mac) into path, so you can call it directly.

**Project is only Golang app, you can run it how you might want**, but also check the `Makefile` and `cmd/fireback-server/Makefile`, those are convinient
scripts to build (Only on mac at the time of writing this)

## Watch on youtube (How to build a backend in 15 minutes)

https://www.youtube.com/watch?v=G2Wjeq7ZmS0


## About Fireback

Fireback is a backend/system software programming framework, which intends to generate backend in minutes instead of months,
give an opinionated structure of software and deliver process, written in Golang and C++.

Fireback products using Module2 yaml language, as the foundation of the business logic, and products created by that will
benefit from auto generated direct SDK for Android Java, SwiftUI, React.js, with a complete dto, actions, rpc exposure.

Fireback starts with zero configuration, will be always a compiled software, uses 5MB ram and weights less than 10MB for server, 30MB for desktop, and C++ version can be compiled for devices with less 1MB of on chip memory.


## Core principles

The reason I have started this project was to change the course of action for entire software development I am involved for next
30 years, they include but not limited to:

- Compile the fastest (within 10 seconds or less), use the least memory (less than 5MB in idle mode), and ensure type safety.
- Products should have a minimum lifespan of 20 years after their initial build.
- Never rewrite the code, due to back structure or skipping important details (remove the "Nice to have" line)
- Achieve machine independence without relying on Docker or other similar tools.
- Compile for various platforms including web, native mobile (iOS, Android), desktop, and embedded devices.
- Do not address ambiguity regarding roles, workspaces, permissions, or email/SMS/OAuth authentication.
- Utilize SQL with an organized data structure.
- Prioritize code management over code generation.
- Significantly reduce the build times for clients in Android, React, Angular, and SwiftUI.
- Zero setup backend usage for front-end or app developers
- Always provide more than any one's need, always before they ask

## Quick overview of fireback project(s)

Fireback itself is a backend project, based on Golang, but there are few other projects inside `clients` folder. The mindset behind it is, to build backend on Go, front-end on React, Mobile apps on Native Android and IOS. Fireback well supports other frameworks, such as react native, angular, ... but we are not going to give you boilerplates. Build backend on Fireback, and use it's restful api, or sdk generator, if available for your framework.

### fireback-golang [Production Ready]

This is the backend, with few modules coming, such as user management, permissions, codegen based on Gin and Gorm. It's production ready, you get it by cloning this repository.

### fireback-react [Production Ready]

This is the front-end boilerplate (it's different from react codegen) gives you general screens of an app, signin, signup. It will be updated over time, but it's reliable enough to use it. 

### fireback-android [Developing]

This is a set of components for building an Android app. It does not really depend on fireback backend, but it's a quick start to build project on Android. It's purely an Android Java project without any modification to details of Android structure. Few libraries are added, such as RxJava, ...
Signup, Signin, ... and few screens and activities are there, and you can modify them.

### fireback-ios [Developing]

Similar to fireback-android, this is also a boilerplate to build IOS apps using SWIFT and SWIFTUI. There is cocopod, and fireback generated sdk inside of it, `Promise` library has been installed, and few onboarding screens, etc will be there.

### fireback-ios and fireback-android

You can clone those projects independently as your boilerplate without fireback also. Plus, you can create android or ios app with fireback, without these projects, it's just gonna
move you forward faster by giving few screens. Also these are boilerplate, they are not "library". Once you clone them, any update to it, from us, will be manually, except the fireback gen android/ios which gives sdk from your backend.



## Cli overview

Here is the default plugins which comes free or with every backend products we build:

```
NAME:
   Fireback core microservice

USAGE:
    [global options] command [command options] [arguments...]

COMMANDS:
   capability, cap            Manage the capabilities inside the application, both builtin to core and custom defined ones
   role                       Manage roles within the workspaces, or root configuration
   user                       Manage the users who are in the current app (root only)
   ws                         
   workspaceInvite            Active invitations for non-users or already users to join an specific workspace
   backup                     Keeps information about which tables to be used during backup (mostly internal)
   tableViewSizing, tvs       Used to store meta data about user tables (in front-end, or apps for example) about the size of the columns
   appMenu                    Manages the menus in the app, (for example tab views, sidebar items, etc.)
   regionalContent, rc        Email templates, sms templates or other textual content which can be accessed.
   keyboardShortcut, kbshort  Manage the keyboard shortcuts in web and desktop apps (accessibility)
   notification, nt           Manage email accounts, templates, email providers and so on
   passport                   Manage the methods of authentication in the app, as well as users passports (root only)
   publicJoinKey              Joining to different workspaces using a public link directly
   widgetArea                 Widget areas are groups of widgets, which can be placed on a special place such as dashboard
   widget                     Widget is an item which can be placed on a widget area, such as weather widget
   commonProfile              A common profile issues for every user (Set the living address, etc)
   currency, curr             List of all famous currencies, both internal and user defined ones
   priceTag                   Price tag is a definition of a price, in different currencies or regions
   license                    Manage the licenses in the app (either to issue, or to activate current product)
   init                       Initialize the project, adds yaml configuration in the folder.
   about                      About Fireback, the author of software, support and contact :)
   gen                        Code generation tools, both for internal codes and sdk remote files
   doctor                     Gives some information about the app, operating system, for remote debugging
   service                    Manages the system service on operating system
   migration                  Migration of the data (import or export)
   start, s                   Starts http server only
   mock                       Generates or export mocks based on all available content inside the database
   seeders                    Imports all necessarys eeders
   reports                    Views all the reports available in the system
   help, h                    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```


## Fireback terms

Fireback is a general name, but refers to few products in itself. Fireback microservice is a way to authenticate other backend project, which could be written by any other languages, like node.js, php, etc. Fireback Module2 is a language, a declarative way of defining entities (both models, and business logic), Fireback Core, which is the framework source, that I internally use for client projects,
and not all the toolset is exposed at all. Fireback for Android, is the android tools written in Java to connect to Fireback core products, the same goes for Fireback React, and Fireback Swift. React-Iron and React-Native-Iron are two boilerplates for react and react-native which are open-source, and kinda a way to start react(native?) projects but influenced by the Fireback anyways.

## Core features on backend

- Integrated desktop app support using Wails
- Binaries ready to be embedded on Android and IOS native, as well as cordova for offline backend usage (same code base)
- Build on top of Golang. Fast build time, blazing fast runtime, least memory usage.
- Terminal (CLI) for everything, even user signup, forget password
- User/Role/Workspace built into everything.
- User permission management
- Relational Database support (Sqlite/Mysql/Oracle/Postgres)
- Nested, Object in Object, Array in Object, Object in Array post form handler
- Polyglot (Multilingual) built into the definition
- File upload system, virtual directories based on resumable file upload (TUS)
- Form validation, even for nested objects
- Reactive stream query based on WebSocket on Module2 definition
- Bulk Patch (Update) on every entity.
- Publicly available versus workspace/user protected actions
- Custom action definition, auto cast from HTTP and CLI
- Generate Post, Patch, Query, Delete, Bulk Patch, for every entity
- Html and rich text content storage
- Dynamic data using json support as well as ERD data storage
- QueryDSL for every entity, to query, filter them, sort, without extra dependencies
- Built in search mechanism
- Multilingual menu system, sidebar, tabbar design.
- Generate API and actions to work with children (nested elements) inside another entity (add/remove/etc)
- Auto handle for unique ids and relations between non-nested elements
- Formatting dates, currency, time, based on the client region
- Backend error translated based on the accept language
- Complete and unified google json styleguide response, without even single exception
- Manual and automatic mocking system built into all entities
- Seeder system to include the initial or necessary content for the entities to built to binary (yaml/json/csv)
- Advanced background operation for import/export all entities
- Default and custom defined permission definition
- CSV/yaml data template generation to fill in and import
- Auto sanitize the content
- Advanced and clean form validation, excerpt content creation
- Casting dto/entity from cli parameters
- Pdf export for the queries (beta) without any kind of thirdparty dependency
- JsonDSL (for complex conditions on a query) and QueryDSL (textual query)
- Predefined queries on SQL (Pivot and Recursive CTE operation for all entities)
- Event system for changes on the entity, and broadcasting to rooms using WebSocket
- Wipe entities on CLI (development) and multi row delete 
- Automatic forcing the user privilegs on the content they can create or modify
- Multiple workspaces, and multiple roles inside each workspace
- Multiple workspace type (such as School, Student, etc)
- Direct signup to a team, via public join key
- Distinct by user, workspace operation flag
- Interactive cli tools for creating entities, as well as traditional
- Support of complex enum, which casts to all major programming languages

## Typescript compiler features

- Action, errors, dto and entites definitions
- Fields information
- Autogenerate necessary form elements
- Backend post/get/... methods with typesafety
- Automatic event mapped to socket and update ui
- Caching, filters, multiple update handlers
- Typesafe translations
- Error handling for each input parser

## Android/IOS compiler features

- Full ABAC screens and functionality using native elements
- Entire dto, actions, entities on Java and Swift.
- Form generation helpers

## Passport module

Fireback comes with a passport module, which helps users to signup using email, phone numbers, do forget password, otp, and all the operations in between. Here is an overview of the functionalities:

```
NAME:
   Fireback core microservice passport - Manage the methods of authentication in the app, as well as users passports (root only)

USAGE:
   Fireback core microservice passport command [command options] [arguments...]

COMMANDS:
   ic                      Creates a new user in the system, using an interactive question builder, and adding a passport to it
   append-email            Appends a new passport to an specific user, given by userid in the system
   update, u               Updates a template by passing the parameters
   passportMethod, method  Login/Signup methods which are available in the app for different regions (Email, Phone Number, Google, etc)
   wipe                    Wipes entire passports 
   update, u               Updates a template by passing the parameters
   query, q                Queries all of the entities in database based on the standard query format
   email                   Send a email using default root notification configuration
   emailp                  Send a text message using an specific gsm provider
   invite                  Invite a new person (either a user, with passport or without passport)
   sms                     Send a text message using default root notification configuration
   smsp                    Send a text message using an specific gsm provider
   in                      Signin publicly to and account using class passports (email, password)
   up                      Signup a user into system via public access (aka website visitors) using either email or phone number
   create-workspace        
   ccp                     Checks if a classic passport (email, phone) exists or not, used in multi step authentication
   otp                     Authenticate the user publicly for classic methods using communication service, such as sms, call, or email

OPTIONS:
   --help, -h  show help
```

## Workspace module cli

Workspace module is also accessible.

```
NAME:
   Fireback core microservice ws - Workspaces module actions (sample module to handle complex entities)

USAGE:
   Fireback core microservice ws command [command options] [arguments...]

COMMANDS:
   query, q                Queries all of the entities in database based on the standard query format
   table, t                Table formatted queries all of the entities in database based on the standard query format
   create, c               Create a new template
   update, u               Updates a template by passing the parameters
   ic                      Creates a new template, using requied fields in an interactive name
   wipe                    Wipes entire workspaces 
   remove, r, del, delete  Deletes an entity with given id (uniqueid)
   query-cte, cte          Same as query, but in recursive manner
   query-pivot, pivot      Pivots the the entire table based on conditions
   scope                   Returns the access level, roles, and scopes that an specific user has access to
   cli                     Set some configuration for cli, such as language, region, etc
   config                  Sets the configuration for an specific workspace
   meets                   By given a user id, to will check if user has the capabilities asked for
   as                      Set the workspace in terminal
   tests                   Tests related to the workspace cli
   type                    
   config                  
   workspaceRole, role     Manage roles assigned to an specific workspace or created by the workspace itself
   userWorkspace, user     Manage the workspaces that user belongs to (either its himselves or adding by invitation)
   mock                    Generates mock records based on the entity definition
   init, i                 Creates a basic seeder file for you, based on the definition module we have. You can populate this file as an example
   validate, v             Reads a yaml file containing an array of workspaces, you can run this to validate if your import file is correct, and how it would look like after import
   import                  imports csv/yaml/json file and place it and its children into database

OPTIONS:
   --language value  (default: "en")
   --help, -h        show help
```

## Module2 Definition

This is a language that we define our modules, even fireback workspace module (which everything depends on) is written using this language.
Language itself is valid yaml, or even could be presented in json format.

Document can be created in `modules/mymodule` folder, and by convention needs to end with `Module3.yml` for VSCode auto compile.
Define `path` and  `name`, which represents the folder name (and the generated folder as well) and name of the module.

**Always, camelCaseVariables in entire of the definitions**

```
path: mymodule
name: mymodule
```


### Module2 entities

Defining entities is super straight-forward, it a flat array in the module.yml file. Entity name *Must be singular* by definition.

```
path: mymodule
name: mymodule

entities: 
  - name: customer
    fields:
    - name: firstName
      type: string
    - name: age
      type: int64
```

As visible, each entity has few sections, the `name` and `fields` are the most important one. Each field, needs to at least
define name and it's type. Fireback types are all golang primitives `int64`, `string`, `bool` ... and few custom defined by fireback:

#### Fireback `object` and `array` types

Each field can be object or array, when defined as such, it can have children properties called `fields`, similar to the original entity,
and this relation can be nested indefinitely.

For example, customers can be extended like this:


```
path: mymodule
name: mymodule

entities: 
  - name: customer
    fields:
    - name: firstName
      type: string
    - name: age
      type: int64
    - name: address
      type: object
      fields: 
      - name: postalCode
        type: string
      - name: streetAddress
        type: string
    - name: visitToOffices
      type: array
      fields:
      - name: visitDate
        type: datenano
      - name: visitDurationInSeconds
        type: int64

```

As we see, array and object field types have their own fields, there will be separate tables in the database for them, and they are linked
to the parent table automatically using foreign key and unique id by any fireback product.

**Note1: If the field type is array, its name must be plural**
**Note2: object/array ARE DIFFERENT from one/many2many types**

#### Fireback `one` and `many2many` types.
While defining these types are similar to object/array types, there are significant differences.

Unlike array/object, fireback does not automatically create them. Instead, they are treated as completely separate features/business model, you just add them by giving their uniqueId.

For `one`, there will be a string field added called (name+Id) and for the many2many, will be string array (name+ListId).
When updating the entity with such relation, you need to specify those keys rather than the object itself.
When querying, fireback will use the original fields to give you back your selection.

It's better to use `many2many` and `one` if applicable, to separate the structure of the code. From data base perspective, all 4 of those
fields are gonna have their own tables.

#### Fireback `json` type

Although fireback emphasizes on data normalization, there is a json type supported, which will be JSON field in the database, or string if the target database does not support it. It is useful when the data structure is unknown, does not need strict validation (such as config for an specific widget). Never over use, if data might be queried later. [ In some databases, such as postgres, there is significant improvement on querying or store json fields, but using this field would limit fireback and unforeseeable database operation in future. ]

```
entities: 
  - name: customer
    fields:
    - name: config
      type: json
      matches:
      - dto: CustomerConfigDto
```

If you specify a list of `matches` entities or dto, fireback might use validation from them, and generated Swift/Ts/Java code for front-end 
will be giving some helper functions to detect the actual type of the column there. If you have limited and known dto for that, it's good practise to mention them in matches section.

#### Fireback `html` field
This type is similar to the string, but allows for html tags. Good to use when generating rich text content, etc. `excerptSize` and `unsafe` are also an option to use, since 'text', and 'html' fields in fireback also create an excerpt mechanism automatically. If `unsafe` is true, it would be allowing for `<script>` tags without striping them down.

```
...
      - name: title
        type: html
        unsafe: true
        excerptSize: 10
      - name: body
        type: html
        excerptSize: 10
...
```


## Designed with love by Ali Torabi

Fireback is not fully opensource. It's created to design the reliable software faster. If you want to move away from basic day-to-day issues and step forward, contact me.
+48 783 538 796 - ali-torabian@outlook.com

