# v1.2.2

Shift to user user interface provided for the admin and new projects.

- [ ] Auto completion for seeder and mock files, as well as casting tree to array in yaml files
- [ ] Include the desktop kit.
- [ ] Demo with integration with Nest.js for authentication
- [ ] Add internal captcha alongside with recaptcha 3
- [ ] Work on the user invitation accepting and joining to workspace, all the scenarios.
- [ ] Self service user can change the phone number or email address
- [ ] Self service can delete account
- [ ] Self service can quite from workspace
- [ ] Allow project init to contain the workspace modules into the new project and be independent of the original fireback project.
- [ ] Consider mongodb query system
- [ ] Add the tabs in the project
- [ ] Add the drawer in the project with promise/resolve/reject
- [ ] Prepare react pre-templates. For lists, single form, single view screen, single tabbed screen, wizard
- [ ] Explore the audio streams and how they work to have endpoints
- [ ] Extensive testing for http json requests, to match the json and more

# v1.2.1

Goal of this version is to improve the fireback backend mostly, not the UI or usecases.
The plan is the backend be perfect, and can generate major code for front-end technologies
and reduce the time in between.

- [ ] Review the nested modules generated and their behavior on the front-end
- [ ] Make sure the query params, headers, url params are available and accessible via generated sdk
- [ ] Improve the SDK with more functions, specially on react.js one it has some dead code, remove, and comment everything
- [ ] Make sure the documentation (openapi) does contain all of the necessary explanations.
- [ ] Improve the SwiftUI codegen, test it via github actions
- [ ] Improve the Android code gen
- [ ] More detailed errors if the json matching for post requests is not available
- [ ] Streaming content such as audio or video
- [ ] Autoscaling for the project demo to queue every request.
- [ ] Extensive test for the socket connection and changes coming from server to be reflected on UI.
- [ ] Revise the code generated for the typescript, specially fields on nested objects to have the path.
- [ ] Add the replica option for the clickhouse for time series data, investigate the time series formats.

# v1.2.0

The changes planned for 1.2.0 are the following:

- [x] Separate the build flows on github actions, create 2 test flow for app on mysql and sqlite
- [x] Test and fix the issues with exporting
- [x] Use nullable values across the project
- [x] Make sure capabilities use . instead of slash, and add title/description to all
- [x] Complete all functions to use optional override functions for EntityActions instead of direct function
- [x] React.js sidebar bugfix on mobile, also fix the pull to refresh plugin.
- [x] Build the capacitor version as well on the githubactions for android.
- [x] Fix the child project OpenAPI documentation
- [x] Disable the ID numeric primary key in the entities
- [x] Self service user can change the password
- [x] Add duration field type into the fireback
- [x] React Native login using self-service
- [x] Black overlay on some cases on react.js
- [x] WorkspaceType, WorkspaceConfig, SelfService UIs
- [x] Complete the Root menu
- [x] Refactor front-end code into selfservice and manage
- [x] Add google authentication
- [x] Remove test command internally, and plan for testing product beforehand.
- [x] Interface to manage users for Workspace Admin and root.
- [x] Add vite React.js project demo with self service portal
- [x] Role creation needs to be shortened to the role that user specific permissions, or admin only can create new roles.
- [x] Allow customization of the token reading/creation scenario, for microservice projects.
- [x] Add production flag, and disable mock, wipe, delete functionality.
- [ ] Complete the all scenarios for user accepting invitation and joining workspaces.
- [ ] Generate the typings for the queries.
- [ ] Passports list UI for manage
- [ ] Cursor pagination along side with UI changes to the data table.
- [ ] Bring the Arura functionality from a legacy project
- [ ] Include the docurus, githubactions, and cypress test kits into the new projects created
- [ ] On react UI ask for signout instead of immediate signout
- [ ] Add demo to add extra configuration for child project env files
- [ ] Document the file upload system, allow integration with S3.
- [ ] Configure the search functionality
- [ ] AuthroizeOS functionality revise, document.
- [ ] Explore the option to have custom header, query params in actions, entities actions.
- [ ] Test the custom actions in the entities.
- [ ] Document better the reactive actions, and make them custom types other than string only.
- [ ] Unify the exporting mechanism for both cli and http, to get output in csv, json, and yaml formats.
- [ ] Revise the notifications service, use integrate email service the same in otp

# v1.1.28

Add's the React.js and React native project building. For example:

To create a 'front-end' folder inside project with react.js:
```
fireback new --ui
```

Similar, you can add also a react native project inside the new project:
```
fireback new --mobile
```

Both projects will have a similar structure, and include a copy of fireback UI components,
hooks, and you can modify them.
