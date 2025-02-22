## This is no longer maintained.


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




Github actions:

vscode-extension:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Use Node.js 18
        uses: actions/setup-node@v3
        with:
          node-version: 18

      - name: Install react-new dependencies since it contains the designer files
        run: cd modules/workspaces/codegen/react-new && npm i --force
        
      - name: Build
        run: cd clients/fireback-tools-vs-code-extension && npm i --force && npm run package
      - uses: actions/upload-artifact@master
        with:
          name: artifacts-vscode
          path: artifacts


deploy_vscode:
    needs:
      - vscode-extension

    if: ${{ inputs.deploy_vscode_market == true }}
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Publish to Visual Studio Marketplace
        uses: HaaLeo/publish-vscode-extension@v1
        with:
          pat: ${{ secrets.FIREBACK_VSCODE_MARKETPLACE }}
          extensionFile: artifacts/fireback-vscode/fireback-tools.vsix
          registryUrl: https://marketplace.visualstudio.com

deploy_vscode_market:
    description: "Publish vscode marketplace"
    required: false
    default: false
    type: boolean

- uses: actions/download-artifact@master
    with:
        name: artifacts-vscode
        path: artifacts-vscode


- name: upload vscode extension
    uses: actions/upload-release-asset@v1
    env:
        GITHUB_TOKEN: ${{ github.token }}
    with:
        upload_url: ${{ steps.create_release.outputs.upload_url }}
        asset_path: artifacts-vscode/fireback-vscode/fireback-tools.vsix
        asset_name: fireback-tools.vsix
        asset_content_type: application/vsix