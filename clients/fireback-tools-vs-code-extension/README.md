# Fireback tools

This plugin is a language server connector for fireback and vscode.
You can find guidelines how to install it on: https://torabi.io/fireback/docs/download-and-install-fireback

**Plugin requires that `fireback` binary to be installed already on your computer.**
If you have installed it in a different place, go to your `.vscode/settings.json` of project
and add the following settings:

```
{
    "fireback.serverPath": "/var/apps/exactplaceoffirebackbinary",
    ... other configs
}
```

and then restart the extension to get access to language server.

## Code completion

This extension offers a code completion feature. For any file which would end on `Module3.yml`, it provides
intelisense to some extension.

## Module3 Designer

You can access module3 visual designer, when you are in a Module3.yml file, and Command+Shift+P and then type:
`Fireback: Module3 Designer` it would popup on the sidebar, and you can edit the modules using forms.

It might remove extra comments you have added on the yaml though. 



