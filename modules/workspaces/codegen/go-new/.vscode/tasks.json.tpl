{
  "version": "2.0.0",
  "tasks": [
    {
      "problemMatcher": [],
      "label": "Generate new module",
      "type": "shell",
      "command": "fireback gen module --name ${input:modulename} --auto-import cmd/{{ .ctx.Name}}-server/main.go",
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
      "description": "Module name (prefer lower string) it will be created in modules folder",
      "default": "",
      "type": "promptString"
    }
  ]
}
