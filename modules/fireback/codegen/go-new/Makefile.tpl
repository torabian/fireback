default:
	cd cmd/{{ .ctx.Name}}-server && make dev

ui:
	cd front-end && npm start


init:
	go mod tidy && make


{{ if .ctx.CreateReactProject }}
bundle:
	cd front-end && npm run build && cd - && rm -rf cmd/{{ .ctx.Name}}-server/ui && cp -R front-end/build cmd/{{ .ctx.Name}}-server/ui
{{ end }}