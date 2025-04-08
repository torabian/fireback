# Fireback – Microservice-First Fullstack Framework in Go

Fireback is a **microservice-oriented backend framework** written in Go. It's designed for developers who want to build fast, scalable, and modular systems without boilerplate. It can accelerate development **20x-30x** for many common backends.

<img src=".github/logo.svg" alt="Fireback logo" width="200"/>



## Getting started

If you want to build a new go project using fireback, easiest is to download the https://github.com/torabian/fireback/releases/download/latest/fireback-microservice-boilerplate.zip 

and continue from there. Make sure you have installed necessary vscode extensions, for code completion.

Or you can go deeper, and install CLI:

- [Download and install fireback](./download-and-install-fireback)
- [Your first fireback project](./first-fireback-project)


### Microservices Made Easy

At its core, Fireback is built around **definition files** that let you describe your services, entities, and APIs in a structured way. From those definitions, Fireback generates:

- Go microservices with REST endpoints  
- React SDKs  
- Android/iOS SDKs (via Swift/Kotlin)  
- DTOs and boilerplate handlers

Your backend becomes a network of services, each independently deployable and easy to integrate across clients.



### Fullstack Codegen for Microservices

Once your microservices are defined, Fireback’s CLI tools generate frontends and SDKs for various platforms—React, Android, iOS—making it dead-simple for apps to consume your services. Backend and client are always in sync.


### Modular by Design

Fireback services are cleanly separated and **decoupled**, built on top of **Gin** and **Gorm**, with minimal overhead. You can add Fireback to new or existing systems. Each generated microservice can live on its own or be composed in larger systems.

## License

This project is licensed under the [GNU Affero General Public License (AGPL)](LICENSE.md). Contributions and modifications must be shared under the same terms. 

## Documents

The original documents are provided in https://torabi.io/fireback but for quick understanding you
can continue reading this document.


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
- Reactive stream query based on WebSocket on Module3 definition
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
 

## Sample client key OAuth google

1040931581120-prdfdml80dl7uaq3999jkge72dph280l.apps.googleusercontent.com

## Change the fireback version

In order to release a new fireback, follow these steps:

- cmd/fireback/fireback-deb.sh
- OVERVIEW.md
- cmd/fireback/msi/Product.wxs
- .github/workflows/fireback-build.yml
- modules/workspaces/version.go