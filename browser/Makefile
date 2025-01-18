.PHONY: main
main: codegen-watch

.PHONY: codegen-clean
codegen-clean:
	rm -f scripting/*_generated.go
	rm -f scripting/**/*_generated.go

.PHONY: codegen-watch
codegen-watch: codegen-clean
	gow -w ../code-gen -e="" generate ./...
