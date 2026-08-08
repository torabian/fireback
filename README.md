# Fireback – Authentication server

Fireback is an authentication server, which can be used to build commercial products in general category,
or be used as a tool to create authentication server for projects within seconds.

It includes a set of features, which most apps need, and uses famous libraries for that:

- Gin
- Urfave Cli v3
- Tus resumable file upload.
- User role workspace

And also comes with a react project, to handle all those things, in UI.

You can clone fireback, or create a new project using it, or add it's libraries as a dependency
to your project, or install it only for a user authentication solution. Totally up to you.

Codegen features are fully removed. We are using https://github.com/torabian/emi for codegen,
and rest will be written manually.

<img src=".github/logo.svg" alt="Fireback logo" width="200"/>

## How to use the project.

There is no longer a "new" command in fireback. You can clone the repository, and rename some `main.go` and continue.
Also if you already have a project, you can install `github.com/torabian/fireback` into it.

## How to use Fireback for managing?

Fireback binaries are built, with self-service and manage enabled, both as binaries as well as disk image called
fireback. If you only need an authentication service use them instead of source code.
