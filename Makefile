.PHONY: main
main: codegen-watch

.PHONY: codegen-clean
codegen-clean:
	rm -f scripting/*_generated.go
	rm -f scripting/**/*_generated.go
	rm -f dom/*_generated.go
	rm -f html/*_generated.go

.PHONY: codegen-watch
codegen-watch: codegen-clean
	gow -w ./internal/code-gen -e="" generate ./...

.PHONY: codegen codegen-build
codegen-build:
	$(MAKE) -C internal/code-gen build

codegen: codegen-clean codegen-build
	go generate ./...

.PHONY: test test-watch test-browser test-v8 test-goja
test: 
	go test -v -vet=all ./...

test-watch: 
	gow -c -e=go -e=js -e=html -v -w=./.. test -vet=off ./...

.PHONY: test-dom
test-browser: 
	gow -s -w=./dom -w=./html -w=. test -vet=off . ./dom ./html

.PHONY: test-html
test-html: 
	cd html && ginkgo watch -vet=off

.PHONY: test-v8
test-v8: 
	gow -s -e=go -e=js -e=html -w ./.. test -vet=off ./scripting/v8host

.PHONY: test-goja
test-goja:
	gow -c -e=go -e=js -e=html -w ./.. test -vet=off ./scripting/gojahost

.PHONY: ci ci-build
ci-build:
	go build -v ./...

ci: codegen ci-build test
	git diff --quiet HEAD
