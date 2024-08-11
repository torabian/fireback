{
  "version": "2.0.0",
  "tasks": [
    {
      "problemMatcher": [],
      "label": "Generate go module",
      "type": "shell",
      "command": "./artifacts/{{ .ctx.Name}}-server/{{ .ctx.Name}} gen module --name ${input:modulename} --auto-import cmd/{{ .ctx.Name}}-server/main.go",
      "group": "test",
      "presentation": {
        "reveal": "always",
        "panel": "new"
      }
    },
    {
      "problemMatcher": [],
      "label": "Rebuild (go tidy & make)",
      "type": "shell",
      "command": "make init",
      "group": "test",
      "presentation": {
        "reveal": "always",
        "panel": "new"
      }
    },
    {
      "problemMatcher": [],
      "label": "Generate front-end Entity from backend",
      "type": "shell",
      "command": "./artifacts/{{ .ctx.Name}}-server/{{ .ctx.Name}} gen react-ui --entity-path ${input:getmodules} --path ${input:getuipath} --sdk-dir @/modules/sdk/{{ .ctx.Name}}",
      "group": "test",
      "presentation": {
        "reveal": "always",
        "panel": "new"
      }
    },
    {
      "problemMatcher": [],
      "label": "Sync the backend sdk to front-end",
      "type": "shell",
      "command": "./artifacts/{{ .ctx.Name}}-server/{{ .ctx.Name}} gen react --path front-end/src/modules/sdk/{{ .ctx.Name}}",
      "group": "test",
      "presentation": {
        "reveal": "always",
        "panel": "new"
      }
    }
  ],
  "inputs": [
    {
      "id": "modulename",
      "description": "Module name (a-z0-9), and you can nest them",
      "default": "modules/",
      "type": "promptString"
    },
    {
      "id": "getuipath",
      "description": "The address of the generated content",
      "default": "front-end/src/modules/",
      "type": "promptString"
    },
    {
      "id": "getmodules",
      "type": "command",
      "command": "shellCommand.execute",
      "args": {
        "command": "./artifacts/{{ .ctx.Name}}-server/{{ .ctx.Name}} gen entities",
        "cwd": "${workspaceFolder}",
        "description": "The module that you are looking for entity inside",
        "env": {
          "HOME": "~",
          "WORKSPACE": "${workspaceFolder[0]}",
          "FILE": "${file}",
          "PROJECT": "${workspaceFolderBasename}"
        }
      }
    }
  ]
}
