# Fireback module milestones

Content here are the next tasks related to the Fireback core module. Fireback core includes
only the codegeneration, entities, http, socket, connections.

It has no authentication layer in itself, everything is 'root' or null when it coems to database
requirement. "ABAC" module needs to be included for user/multitenant access level, or implemented customly.
Assume the "module/fireback" is the main framework, and "modules/abac" is a project written using that.
Fireback itself doesn't reference the ABAC module - at all.

Fireback module goal is to provide everything other projects might need, in terms of code and solving complex
problems, and the scope is backend + react.js sdk (rpc) + kotlin sdk (rpc), documentation.
It doesn't provide a front-end (which is available though, but it's not meant for subsequent projects)


- [ ] Include the docurus, githubactions, and cypress test kits into the new projects created
- [ ] Improve the Android (Kotlin) SDK - generate specifically reactive method due to usage in dashboard
- [ ] Add replica option for the ClickHouse
- [ ] Test with SQL Server, and check if there is easy way to integrate it.
- [ ] Casting tree in yaml files to read the children and automatically assign the parent id.
- [ ] Add internal captcha alongside with recaptcha 3
- [ ] Allow project init to contain the workspace modules into the new project and be independent of the original fireback project.
- [ ] Review the nested modules generated and their behavior on the front-end
- [ ] Make sure the query params, headers, url params are available and accessible via generated sdk
- [ ] Improve the SDK with more functions, specially on react.js one it has some dead code, remove, and comment everything
- [ ] Make sure the documentation (openapi) does contain all of the necessary explanations.
- [ ] Improve the SwiftUI codegen, test it via github actions
- [ ] Autoscaling for the project demo to queue every request.
- [ ] Extensive test for the socket connection and changes coming from server to be reflected on UI.
- [ ] Revise the code generated for the typescript, specially fields on nested objects to have the path.
- [ ] Add the replica option for the clickhouse for time series data, investigate the time series formats.
- [ ] Comments added on typescript missing @description flag.
- [ ] Complete the all scenarios for user accepting invitation and joining workspaces, with tests.
- [ ] Generate the typings for the queries.
- [ ] Bring the Arura functionality from a legacy project, make it default available on child projects with a sample.
- [ ] Add demo to add extra configuration for child project env files
- [ ] Document the file upload system, allow integration with S3.
- [ ] Explore the option to have custom header, query params in actions, entities actions.
- [ ] Test the custom actions in the entities.
- [ ] Document better the reactive actions, and make them custom types other than string only.
- [ ] Unify the exporting mechanism for both cli and http, to get output in csv, json, and yaml formats.
