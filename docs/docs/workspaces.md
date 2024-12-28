---
sidebar_position: 53
---

# Fireback Workspaces tables and features

As you might noticed so far, Fireback is only `modules/workspaces` folder, and there is not much code
out side of this module. This is intentional, to have everything needed in one module and do not focus
on managing modules and their compatibility.

In this document, we are going to discuss what are the features of the project in a managemental perspective.

## User role permission

Fireback comes with a built in user, role permission, to manage the user access in the projects. Fireback
aims to reduce the need for large user management systems on smaller scale projects, and also give the developers
and management an option to store the user information under the same database as their project. 
This approach would make everything simple, specially for those projects with less than a million users.

### User, Person, Workspace entities

As the name suggests, Fireback manages data under workspaces. Workspaces essentially are a group of users, which
form a workspace, such as a team, or company. Over decades, I have seen a lot of projects which store user data
individually but when they want to extend logic for teams, invitations, their entire database model changes
and tons of rewrites needed. Fireback at the zero day solves this problem by defining workspace and user system.

All workspaces are stored in workspaces entity (table) on the database, and each workspace can have a name,
and other basic columns such as created at, ... All workspaces belong to root workspace by definition,
and when a new workspace has been created, it's workspace id and parent id is `root`. This allows your project,
to have the possbility of nested workspaces, which the parent workspaces have access to content of children but
not wise versa. 

Role entity, defines a set of capabilities which a role can have. You can define the permissions as strings,
which will be stored in Capability entity on the database, and Fireback apps sync those permissions from module
definitions into database upon migration.

By grouping a set of capabilities (array of strings), a role is being created. Roles belong to workspaces, therefor, every single workspace in the product can have multiple roles, which is different from other workspace
and defined by the owner of the workspace within the same app. Some roles, are also enabled by default

You can check the capabilities in a system via `cap` or `capabilities` in the CLI, which is also available
on the http server.

```bash
ali@x fireback % ./app cap tree
root
├ *
├ modules
│ ├ widget
│ │ ├ widget-area
│ │ │ ├ create
│ │ │ ├ update
│ │ │ ├ query
│ │ │ ├ *
│ │ │ └ delete
│ │ └ widget
│ │   ├ *
│ │   ├ delete
│ │   ├ create
```

The * sign, means all of the neighbour permissions are allowed. It's useful because it would automatically include
later added permissions without the need of revising the roles.

## Workspace types

In many different applications, you might need different types of workspaces, with different signup flow.
