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
- [x] React Native login using self-service
- [ ] Black overlay on some cases on react.js
- [ ] Auto completion for seeder and mock files, as well as casting tree to array in yaml files
- [ ] WorkspaceType, WorkspaceConfig, Notification, SelfService UIs
- [ ] Demo with integration with Nest.js for authentication
- [ ] Include the docurus, githubactions, and cypress test kits into the new projects created
- [ ] Include the desktop kit.
- [ ] Think if we can archive IOT legacy code into the archive or examples of this app.
- [ ] Complete the Root menu
- [ ] Add production flag, and disable mock, wipe, delete functionality.
- [ ] On react UI ask for signout instead of immediate signout
- [ ] Add google authentication
- [ ] Add internal captcha alongside with recaptcha 3
- [ ] Interface to manage users for Workspace Admin and root.
- [ ] Remove test command internally, and plan for testing product beforehand.
- [ ] Add demo to add extra configuration for child project.
- [ ] Add vite React.js project demo with self service portal, both using react.js components or directly the binary.
- [ ] Add non-pointer dto or entity generation so we use them directly easier.
- [ ] Document the file upload system, allow integration with S3.
- [ ] Self service user can change the phone number or email address
- [ ] Self service can delete account
- [ ] Self service can quite from workspace
- [ ] Allow project init to contain the workspace modules into the new project and be independent of the original fireback project.
- [ ] Add the tabs in the project
- [ ] Add the drawer in the project with promise/resolve/reject
- [ ] Configure the search functionality
- [ ] Consider mongodb query system
- [ ] Explore the option to have custom header, query params in actions
- [ ] Test the custom actions in the entities.
- [ ] Prepare react pre-templates. For lists, single form, single view screen, single tabbed screen, wizard
- [ ] Extensive testing for http json requests, to match the json and more
- [ ] Document better the reactive actions, and make them custom types other than string only.
- [ ] Explore the audio streams and how they work to have endpoints
- [ ] Test the import system, and make all endpoints for importing/exporting also accessible via http
- [ ] Role creation needs to be shortened to the role that user specific permissions, or admin only can create new roles.

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
