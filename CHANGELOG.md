# v1.2.0

The changes planned for 1.2.0 are the following:

- [ ] React.js sidebar bugfix on mobile, also fix the pull to refresh plugin.
- [x] Separate the build flows on github actions, create 2 test flow for app on mysql and sqlite
- [ ] Build the capacitor version as well on the githubactions for android.
- [ ] Complete all functions to use optional override functions for EntityActions instead of direct function
- [x] Test and fix the issues with exporting
- [ ] Use nullable values across the project
- [ ] Fix the child project OpenAPI documentation
- [ ] React Native login using self-service
- [ ] Auto completion for seeder and mock files, as well as casting tree to array in yaml files
- [ ] WorkspaceType, WorkspaceConfig, Notification, SelfService UIs
- [ ] Demo with integration with Nest.js for authentication
- [ ] Complete the Root menu
- [ ] Make sure capabilities use . instead of slash, and add title/description to all
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
- [ ] Self service user can change the password
- [ ] Self service user can change the phone number or email address
- [ ] Self service can delete account
- [ ] Self service can quite from workspace
- [ ] Allow project init to contain the workspace modules into the new project and be independent of the original fireback project.



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
