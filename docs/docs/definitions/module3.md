---
sidebar_position: 2
---

# Fireback Module3 Definition

One of core features of Fireback framework, is that a lot of code is generated based on some yaml files.
Basically, fireback project consists of only these kind of modules, and generated code, slightly depends
on the fireback library version, but in the other hand we have tried to make it in a way that
generated code is usuable without fireback, or in other words make it more explicit.

The thing is, the Fireback itself, is written with Fireback. Fireback original and important codebase is in
`modules/workspaces` folder, which is Generated from `modules/workspaces/WorkspaceModule3.yml`

It took a long time before Fireback reaches to a stability that could be rewritten by itself.

Now in this document we will discuss about the Module3.yml files. When you create a new Fireback project,
the repository is configured to compile `*Module3.yml` file ending. It's a convension in the Fireback ecosystem
only, you can configure this in other code editors. The general idea is something like code below:

```
{
    "match": "\\Module3.yml$",
    "cmd": "${workspaceFolder}/artifacts/fireback/f gen gof --def ${relativeFile} --relative-to ${workspaceFolder} --gof-module github.com/torabian/fireback"
},
```

If you are using VSCode, this already exists but you need to install "Run on Save" plugin from extensions. At the
moment Fireback VSCode extension does not do this.

## Goal of the definition

Fireback Module3 file format has a goal to contain as much as detail possible about a project. Up to version 1.1.27, Entities, Dtos, Actions, Remotes, Config, and few other important services are possible to be defined via these files.

In fact, you could create Fireback modules and use them on a non Fireback project and it would work totally fine,
you might use only the definitions, structs, and helper codes which are generated, although this is not the goal of Fireback.

## Entities in Module3

Entities, is a an array of entity in Fireback, which basically represents a table in database. Fireback entities
offer more feature comparing to common ORMs, when they make entities. When you create an entity, it would
also create very common actions (CRUD, nested operations, array, ...) for you as well and make them a part
of the http router and CLI. Also each entity, can define permissions on fields, or other kind of permissions
and events related to them.

When you define an entity, you often get most of things that you might need to manage them from a administration perspective, you need to create actions on top of them to add your business logic.

## DistinctBy Feature on entities

In Fireback we allow some of entities to be unique, per specific workspace, user or some other conditions.
This is useful, when you want to make sure only one record per that condition exists on the database.
For example, you might want to have settings per user, but only per single user.

In such scenarios, you can define `distinctBy` on the entity:

```yaml
entities:
    - name: config
      distinctBy: workspace
      fields:
      - name: title
        type: string
```

When you make it distinct by workspace, on the entity, WorkspaceId field becomes unique, therefor you cannot have multiple create on the same, and need to use update instead of create.

**Important** Make sure that the `migration apply` has been called if the entity existed before. Migration for unique workspace Id might not occure via gorm migration (which Fireback is using underneath), so you might
need to manually migration add the unique constraint. For performance reasons, In 1.1.27 Fireback doesn't query
and only relies on the constraint.


## Module3 Messages

Messages, is a powerful yet simple feature of Fireback, which aims to organize the error messages across
the app. Often in many backends, we do not provide clear, translated error messages for actions.

We recommend to add every possible error message that an backend (cli or http server) return into the module
definition, and then translate them into different languages on the same place and give explanation.

At the 1.1.27, you cannot extend them from another file, but I think it's useful to be able to keep translation
of the error messages outside of the project, for those projects are having many languages, but maybe will be provided in later versions of Fireback.

### Define messages on the module:

Consider that the messages are for module at the moment, they will become available for everything,
there for you need to put the messages of actions also in the the same object.

```yaml
messages:
  dataTypeDoesNotExistsInFireback:
    en: This data type does not exist in fireback. %name %location
```

When compiled, Fireback will create few objects in Golang which you could use later on (`WorkspacesModule.dyno.go` in this case)

```go
...
const (
	AlreadyConfirmed                   workspacesCode = "AlreadyConfirmed"
	BodyIsMissing                      workspacesCode = "BodyIsMissing"
	DataTypeDoesNotExistsInFireback    workspacesCode = "DataTypeDoesNotExistsInFireback"
...
```

and Also: 

```go

var WorkspacesMessages = newWorkspacesMessageCode()

func newWorkspacesMessageCode() *workspacesMsgs {
	return &workspacesMsgs{
		DataTypeDoesNotExistsInFireback: ErrorItem{
			"$":  "DataTypeDoesNotExistsInFireback",
			"en": "This data type does not exist in fireback. %name %location",
		},
```

As you see, `$` is the key of the error, which needs to be always present and is equal to uppercase of the key.

Later on, in your actions or other go codes you can use when `IError` is required:

```
return Create401Error(&WorkspacesMessages.DataTypeDoesNotExistsInFireback, []string{})
```

