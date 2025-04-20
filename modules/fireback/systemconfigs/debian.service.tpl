[Unit]
Description={{ .Label}}

[Service]
Environment="CONFIG_PATH={{ .CONFIG_PATH}}"
Type=simple
Restart=always
RestartSec=5s
ExecStart={{ .Program}} start

[Install]
WantedBy=multi-user.target
