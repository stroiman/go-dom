package wrappers

import "github.com/gost-dom/code-gen/script-wrappers/configuration"

func NewGojaWrapperModuleGenerator() ScriptWrapperModulesGenerator {
	specs := configuration.CreateSpecs()
	dom := specs.Module("dom")
	domNode := dom.Type("Node")
	domNode.Method("childNodes").SetNotImplemented()

	return ScriptWrapperModulesGenerator{
		Specs:            specs,
		PackagePath:      gojahost,
		TargetGenerators: GojaTargetGenerators{},
	}
}
