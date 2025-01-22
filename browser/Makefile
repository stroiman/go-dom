.PHONY: main
main: codegen-watch

.PHONY: codegen-clean
codegen-clean:
	rm -f scripting/*_generated.go
	rm -f scripting/**/*_generated.go

.PHONY: codegen-watch
codegen-watch: codegen-clean
	gow -w ../code-gen -e="" generate ./...


.PHONY: test-watch
test: 
	gow -s -v -w=./.. test -vet=off ./...

.PHONY: test-dom
test-dom: 
	cd dom && ginkgo watch -vet=off

.PHONY: test-html
test-html: 
	cd html && ginkgo watch -vet=off

.PHONY: test-v8
test-v8: 
	cd scripting/v8host && ginkgo watch -vet=off

.PHONY: test-goja
test-goja:
	cd scripting/gojahost && ginkgo watch -vet=off
