MAKEFLAGS += -j2
codegen: codegen-v8 codegen-goja

codegen-v8:
	rm -f v8host/*_generated.go
	cd v8host && gow -w ../../code-gen -e="" generate

codegen-goja:
	rm -f goja/*_generated.go
	cd goja && gow -w ../../code-gen -e="" generate
