{
    "emeraldwalk.runonsave": {
      "commands": [
        {
          "match": "\\.go$",
          "cmd": "cd ${workspaceFolder} && make"
        },
        {
          "match": "\\Module3.yml$",
          "cmd": "./artifacts/{{ .ctx.Name}}-server/{{ .ctx.Name}} gen gof --relative-to ${workspaceFolder} --def ${file} --no-cache true --gof-module {{ .ctx.ModuleName }}"
        },
      ]
    }
}
  