# v1.2.3

Shift to user user interface provided for the admin and new projects.

- [x] Extensive testing for http json requests, to match the json and more
- [ ] Include the docurus, githubactions, and cypress test kits into the new projects created
- [x] Revise the event system completely and make sure that the scenario of multiple instances works fine
    with socket connection, test it in docker with multiple instances and remember each user
    with correct session id. 
- [x] VAPID integration with Web Push notification
- [ ] Improve the Android (Kotlin) SDK - generate specifically reactive method due to usage in dashboard
- [ ] Explore the audio streams and how they work to have endpoints
- [ ] Make sure the eventBus sockets is generated via fireback reactive method.
- [ ] Webrtc audio streaming options, and video options
- [ ] Add replica option for the ClickHouse
- [ ] Test with SQL Server, and check if there is easy way to integrate it.
- [ ] Casting tree in yaml files to read the children and automatically assign the parent id.
- [ ] Include the desktop kit.
- [ ] Demo with integration with Nest.js for authentication
- [ ] Add internal captcha alongside with recaptcha 3
- [ ] Work on the user invitation accepting and joining to workspace, all the scenarios.
- [ ] Self service user can change the phone number or email address
- [ ] Self service can delete account
- [ ] Create notification rules mechanism, to define when a notification needs to be created based on some logic
- [ ] Self service can quite from workspace
- [ ] Allow project init to contain the workspace modules into the new project and be independent of the original fireback project.
- [ ] Add the tabs in the project
- [ ] Add the drawer in the project with promise/resolve/reject
- [ ] Prepare react pre-templates. For lists, single form, single view screen, single tabbed screen, wizard
- [ ] Review the nested modules generated and their behavior on the front-end
- [ ] Make sure the query params, headers, url params are available and accessible via generated sdk
- [ ] Improve the SDK with more functions, specially on react.js one it has some dead code, remove, and comment everything
- [ ] Make sure the documentation (openapi) does contain all of the necessary explanations.
- [ ] Improve the SwiftUI codegen, test it via github actions
- [ ] Streaming content such as audio or video
- [ ] Autoscaling for the project demo to queue every request.
- [ ] Extensive test for the socket connection and changes coming from server to be reflected on UI.
- [ ] Revise the code generated for the typescript, specially fields on nested objects to have the path.
- [ ] Add the replica option for the clickhouse for time series data, investigate the time series formats.
- [ ] Comments added on typescript missing @description flag.
- [ ] Complete the all scenarios for user accepting invitation and joining workspaces, with tests.
- [ ] Generate the typings for the queries.
- [ ] Passports list UI for manage
- [ ] Bring the Arura functionality from a legacy project, make it default available on child projects with a sample.
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
