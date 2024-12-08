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