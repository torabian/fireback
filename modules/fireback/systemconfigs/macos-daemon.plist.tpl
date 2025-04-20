<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>EnvironmentVariables</key>
    <dict>
    {{if .CONFIG_PATH}}
        <key>CONFIG_PATH</key>
        <string>{{ .CONFIG_PATH}}</string>
    {{end}}
    </dict>
    <key>Label</key>
    <string>{{ .Label}}</string>
    <key>ProgramArguments</key>
    <array>
        <string>{{ .Program}}</string>
        <string>start</string>
    </array>
    {{if .StdOut}}
    <key>StandardOutPath</key>
    <string>{{ .StdOut}}</string>
    {{end}}
    {{if .StdErr}}
    <key>StandardErrorPath</key>
    <string>{{ .StdErr}}</string>
    {{end}}
    <key>KeepAlive</key>
    <true/>
    <key>RunAtLoad</key>
    <true/>
</dict>
</plist>
