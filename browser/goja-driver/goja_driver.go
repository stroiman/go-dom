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
	InstallClass("Node", "EventTarget", WindowWrapper{})
	InstallClass("Document", "Node", WindowWrapper{})
	InstallClass("EventTarget", "", EventTargetWrapper{})
}

type Function struct {
	Constructor *Object
	Prototype   *Object
}

func (d *GojaInstance) GetObject(obj any, class string) *Object {
	result := d.vm.ToValue(obj).(*Object)
	g := d.globals[class]
	result.SetPrototype(g.Prototype)
	return result
}

func (d *GojaInstance) installGlobals(classes ClassMap) {
	d.globals = make(map[string]Function)
	var assertGlobal func(Class) Function
	assertGlobal = func(class Class) Function {
		name := class.Name
		if constructor, alreadyInstalled := d.globals[name]; alreadyInstalled {
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
		prototype := constructor.Get("prototype").(*Object)
		result := Function{constructor, constructor.Get("prototype").(*Object)}
		d.vm.Set(name, constructor)
		d.globals[name] = result

		if super := class.SuperClassName; super != "" {
			if superclass, found := classes[super]; found {
				superPrototype := assertGlobal(superclass).Prototype
				prototype.SetPrototype(superPrototype)
			} else {
				panic(fmt.Sprintf("Superclass not installed for %s. Superclass: %s", name, super))
			}
		}

		return result
	}
	for _, class := range classes {
		assertGlobal(class)
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
	globalThis.DefineAccessorProperty("document", vm.ToValue(func(c *FunctionCall) Value {
		return result.GetObject(window.Document(), "Document")
	}), nil, FLAG_FALSE, FLAG_TRUE)
	globalThis.SetPrototype(result.globals["Window"].Prototype)

	return result
}

func (d *GojaDriver) Close() {
}

type GojaInstance struct {
	vm      *goja.Runtime
	globals map[string]Function
}

func (i *GojaInstance) Close() {
}

func (i *GojaInstance) Run(string) error {
	return nil
}

func (i *GojaInstance) Eval(str string) (res any, err error) {
	if gojaVal, err := i.vm.RunString(str); err == nil {
		return gojaVal.Export(), nil
	} else {
		return nil, err
	}
}
