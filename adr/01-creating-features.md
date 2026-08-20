# Creating feature steps.

Feature consisting of backend section, and front-end need to follow:

1- On backend, we do not generate API signature, and entities manually. We are using "Emi compiler", which
definition should be available in `.vscode/emi-module-spec.json`. So we need to modify emi yaml files in
case of adding to existing module, or create a folder and emi definition `modules/[module-name]/[module-name].emi.yml`

2- Using `make defs` the typescript code for the feature needs to be created and updated, not generated manually.

3- In react side, we need to create components for communicating to the sdk generated.
