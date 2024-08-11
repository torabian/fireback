default:
	cd cmd/{{ .ctx.Name}}-server && make dev

ui:
	cd front-end && npm start


init:
	go mod tidy && make