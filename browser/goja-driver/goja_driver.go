package goja_driver

import (
	"fmt"

	"github.com/stroiman/go-dom/browser/html"

	"github.com/dop251/goja"
	. "github.com/dop251/goja"
)

func NewGojaScriptEngine() html.ScriptHost {
	return &GojaDriver{}
}

type GojaDriver struct {
}

func WindowConstructor(call goja.ConstructorCall, r *goja.Runtime) *goja.Object {
	panic(r.NewTypeError("Illegal Constructor"))
}

type Wrapper interface {
	Constructor(call goja.ConstructorCall, r *goja.Runtime) *goja.Object
}

type Class struct {
	Name           string
	SuperClassName string
	Wrapper        Wrapper
}

type WindowWrapper struct{}

func (w WindowWrapper) Constructor(call goja.ConstructorCall, r *goja.Runtime) *goja.Object {
	panic(r.NewTypeError("Illegal Constructor"))
}

type EventTargetWrapper struct{}

func (w EventTargetWrapper) Constructor(call goja.ConstructorCall, r *goja.Runtime) *goja.Object {
	panic(r.NewTypeError("Illegal Constructor"))
}

type ClassMap map[string]Class

var Globals ClassMap = make(ClassMap)

func InstallClass(name string, superClassName string, wrapper Wrapper) {
	if _, found := Globals[name]; found {
		panic("Class already installed")
	}
	Globals[name] = Class{name, superClassName, wrapper}
}

func init() {
	InstallClass("Window", "EventTarget", WindowWrapper{})
	InstallClass("EventTarget", "", EventTargetWrapper{})
}

type Function *Object

func (d *GojaInstance) installGlobals(classes ClassMap) {
	globals := make(map[string]Function)
	var installGlobal func(Class) *Object
	installGlobal = func(class Class) *Object {
		name := class.Name
		if constructor, alreadyInstalled := globals[name]; alreadyInstalled {
			return constructor
		}
		constructor := d.vm.ToValue(class.Wrapper.Constructor).(*goja.Object)
		constructor.DefineDataProperty(
			"name",
			d.vm.ToValue(name),
			FLAG_NOT_SET,
			FLAG_NOT_SET,
			FLAG_NOT_SET,
		)

		if super := class.SuperClassName; super != "" {
			if superclass, found := classes[super]; found {
				superClassConstructor := installGlobal(superclass)
				superPrototype := superClassConstructor.Get("prototype").(*Object)
				prototype := constructor.Get("prototype").(*Object)
				prototype.SetPrototype(superPrototype)
			} else {
				panic(fmt.Sprintf("Superclass not installed for %s. Superclass: %s", name, super))
			}
		}

		d.vm.Set(name, constructor)
		globals[name] = constructor
		return constructor
	}
	for _, class := range classes {
		installGlobal(class)
	}
}

func (d *GojaDriver) NewContext(window html.Window) html.ScriptContext {
	vm := goja.New()
	result := &GojaInstance{
		vm: vm,
	}
	result.installGlobals(Globals)
	globalThis := vm.GlobalObject()
	globalThis.Set("window", globalThis)
	windowPrototype := globalThis.Get("Window").(*Object).Get("prototype").(*goja.Object)
	globalThis.SetPrototype(windowPrototype)
	return result
}

func (d *GojaDriver) Close() {
}

type GojaInstance struct {
	vm *goja.Runtime
}

func (i *GojaInstance) Close() {
}

func (i *GojaInstance) Run(string) error {
	return nil
}

func (i *GojaInstance) Eval(str string) (any, error) {
	if gojaVal, err := i.vm.RunString(str); err == nil {
		return gojaVal.Export(), nil
	} else {
		return nil, err
	}
}
