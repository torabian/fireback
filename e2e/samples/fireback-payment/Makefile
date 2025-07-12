default:
	cd cmd/fireback-payment-server && make dev

ui:
	cd front-end && npm start


init:
	go mod tidy && make

refresh:
	make && ./app gen gof --relative-to ${PWD} --def ${PWD}/modules/product/ProductModule3.yml --no-cache true --gof-module github.com/torabian/fireback/fireback-payment >> ${PWD}/artifacts/module3.log  && make