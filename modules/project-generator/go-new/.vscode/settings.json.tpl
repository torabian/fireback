{
  "emeraldwalk.runonsave": {
    "commands": [
      {
        "match": "\\.emi.yml$",
        "cmd": "${workspaceFolder}/artifacts/{{ .ctx.Name}}-server/{{ .ctx.Name}} emi --path ${file}"
      },
    ]
  }
}
  