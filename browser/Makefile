MAKEFLAGS += -j2
codegen: codegen-v8 codegen-goja

codegen-v8:
	rm -f scripting/*_generated.go
	cd scripting && gow -w ../../code-gen -e="" generate

codegen-goja:
	rm -f goja-driver/*_generated.go
	cd goja-driver && gow -w ../../code-gen -e="" generate
