default:
	cd cmd/fireback-data-types-server && make dev

ui:
	cd front-end && npm start


init:
	go mod tidy && make

refresh:
	make && ./app gen gof --relative-to ${PWD} --def ${PWD}/modules/product/ProductModule3.yml --no-cache true --gof-module github.com/torabian/fireback/fireback-data-types >> ${PWD}/artifacts/module3.log  && make