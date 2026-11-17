.PHONY: copy generate

SWAGGER_FILE := api/swagger.yaml
CODEGEN_DIR ?= /tmp/loadtest-codegen

copy:
	go run ./scripts/copy_swagger.go

generate: copy
	rm -rf "$(CODEGEN_DIR)"
	mkdir -p "$(CODEGEN_DIR)"
	swagger-codegen generate \
		-i "$(SWAGGER_FILE)" \
		-l go \
		-o "$(CODEGEN_DIR)" \
		--additional-properties=packageName=loadtest,packageVersion=0.1.0 \
		--http-user-agent="Harness-LoadTest-SDK/0.1.0" \
		--git-user-id="harness" \
		--git-repo-id="harness-go-sdk"
	find "$(CODEGEN_DIR)" -maxdepth 1 -name '*.go' -exec cp {} . \;
	go run ./scripts/patch_generated.go
	gofmt -w *.go
