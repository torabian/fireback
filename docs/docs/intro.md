---
sidebar_position: 1
---
# Fireback ‐ Ultimate Go FullStack Framework

Fireback is a backend framework written in Go, for Go developers. Fireback is opinated,
it could be making developing of certain type of projects 20x-30x faster, and easier.
Fireback depends on definition files to create entities, actions, dto and uses those
definitions to create go code, react code, android and ios.


Although fireback is for creating backend, it has a look on Android, Swift, React/Native apps
and generates sdk for them, to make it easier to use your backend or microservice directly
in the apps. When you complete your backend in fireback, cli tools can be used to generate
react sdk, and other platform tools.

Fireback is pure golang, you can organize your code as you want after all, you can use
it in new projects or add it to your existing projects. Code tries to be independent as
much as possible, but it is on top of Gin and Gorm libraries, and has few other things
for cli.

Fireback comes with `workspaces` module, which includes many day to day functionality,
such as user signin, signup, forget password, ABAC, role permission system, and many
more you might need to build a mid to large scale backend.

## License

This project is licensed under the [GNU Affero General Public License (AGPL)](https://github.com/torabian/fireback/blob/main/LICENSE.md). Contributions and modifications must be shared under the same terms. 

## Documents

The original documents are provided in https://torabi.io/fireback but for quick understanding you
can continue reading this document.

## Initial Web UI

When generating a new fireback project, now it's possible to use `--ui` flag, and a complete
react.js dashboard compatible with fireback will be created. It already has a set of fireback components, screens, hooks, which you can modify, delete or create your own set instead.

 
## Getting started

You can create a new project by installing fireback, and using `fireback new` command,
create some actions and entities and build your first backend with tons of functionalities in 10 minutes.
It would be an interactive question set, which initializes everything.

- [Download and install fireback](https://pixelplux.com/en/fireback/download-and-install-fireback)
- [Your first fireback project](https://pixelplux.com/en/fireback/your-first-fireback-project)


## Watch on youtube (How to build a backend in 15 minutes)

[![Fireback in action](https://img.youtube.com/vi/G2Wjeq7ZmS0/0.jpg)](https://www.youtube.com/watch?v=G2Wjeq7ZmS0)

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

## React/Native generator features

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
 
