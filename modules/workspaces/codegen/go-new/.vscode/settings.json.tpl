{
    "emeraldwalk.runonsave": {
      "commands": [
        {
          "match": "\\.go$",
          "cmd": "cd ${workspaceFolder} && make"
        },
        {
          "match": "\\Module3.yml$",
          "cmd": "fireback gen gof --path ${workspaceFolder} --def ${file} --no-cache true --gof-module {{ .ctx.ModuleName }}/modules"
        },
      ]
    }
}
  