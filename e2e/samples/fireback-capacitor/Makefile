build-ui:
	rm -rf dist;
	cd ../../../modules/fireback/codegen/react-new && npm run build:bundle:demo;
	cp -R ../../../modules/fireback/codegen/react-new/build dist;
	npx cap sync;

apk:
	npx cap sync && \
	cd android && \
	./gradlew assembleRelease \
	-Pandroid.injected.signing.store.file=$(PWD)/release.keystore \
	-Pandroid.injected.signing.store.password=123321 \
	-Pandroid.injected.signing.key.alias=my-key-alias \
	-Pandroid.injected.signing.key.password=123321

rebuild:
	make build-ui apk