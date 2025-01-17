package wrappers

import "io/fs"

func NewGojaWrapperModuleGenerator(idlSources fs.FS) ScriptWrapperModulesGenerator {
	specs := CreateSpecs()
	dom := specs.Module("dom")
	domNode := dom.Type("Node")
	domNode.Method("childNodes").SetNotImplemented()

	return ScriptWrapperModulesGenerator{
		IdlSources:       idlSources,
		Specs:            specs,
		PackagePath:      gojahost,
		TargetGenerators: GojaTargetGenerators{},
	}
}
