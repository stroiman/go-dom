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
	gow -c -e=go -e=js -e=html -v -w=./.. test -vet=off ./...

.PHONY: test-dom
test-dom: 
	cd dom && ginkgo watch -vet=off

.PHONY: test-html
test-html: 
	cd html && ginkgo watch -vet=off

.PHONY: test-v8
test-v8: 
	gow -c -e=go -e=js -e=html -w ./.. test -vet=off ./scripting/v8host

.PHONY: test-goja
test-goja:
	gow -c -e=go -e=js -e=html -w ./.. test -vet=off ./scripting/gojahost
