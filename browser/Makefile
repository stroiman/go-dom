MAKEFLAGS += -j2
codegen: codegen-scripting codegen-v8 codegen-goja

codegen-scripting:
	rm -f scripting/*_generated.go
	cd scripting && gow -w ../../code-gen -e="" generate

codegen-v8:
	rm -f scripting/v8host/*_generated.go
	cd scripting/v8host && gow -w ../../../code-gen -e="" generate

codegen-goja:
	rm -f scripting/gojahost/*_generated.go
	cd scripting/gojahost && gow -w ../../../code-gen -e="" generate
