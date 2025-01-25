package wrappers

func NewGojaWrapperModuleGenerator() ScriptWrapperModulesGenerator {
	specs := CreateSpecs()
	dom := specs.Module("dom")
	domNode := dom.Type("Node")
	domNode.Method("childNodes").SetNotImplemented()

	return ScriptWrapperModulesGenerator{
		Specs:            specs,
		PackagePath:      gojahost,
		TargetGenerators: GojaTargetGenerators{},
	}
}
