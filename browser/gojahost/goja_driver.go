package gojahost

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/stroiman/go-dom/browser/dom"
	"github.com/stroiman/go-dom/browser/html"

	g "github.com/dop251/goja"
)

const INTERNAL_SYMBOL_NAME = "__go_dom_internal_value__"

func New() html.ScriptHost {
	return &gojaScriptHost{}
}

type gojaScriptHost struct{}

type wrapper interface {
	constructor(call g.ConstructorCall, r *g.Runtime) *g.Object
	storeInternal(value any, this *g.Object)
}

type createWrapper func(instance *GojaContext) wrapper

type wrapperPrototypeInitializer interface {
	initializePrototype(prototype *g.Object, r *g.Runtime)
}

type class struct {
	name           string
	superClassName string
	wrapper        createWrapper
}

type classMap map[string]class

var globals classMap = make(classMap)

func installClass(name string, superClassName string, wrapper createWrapper) {
	if _, found := globals[name]; found {
		panic("Class already installed")
	}
	globals[name] = class{name, superClassName, wrapper}
}

func init() {
	installClass("EventTarget", "", newEventTargetWrapper)
	installClass("Window", "Node", newWindowWrapper)
	installClass("Document", "Node", newDocumentWrapper)
	installClass("Event", "", NewEventWrapper)
	installClass("CustomEvent", "Event", NewCustomEventWrapper)
	installClass("Element", "Node", newElementWrapper)

}

type function struct {
	Constructor *g.Object
	Prototype   *g.Object
	Wrapper     wrapper
}

func (d *GojaContext) getObject(obj any, class string) *g.Object {
	result := d.vm.ToValue(obj).(*g.Object)
	g := d.globals[class]
	result.SetPrototype(g.Prototype)
	return result
}

type propertyNameMapper struct{}

func (_ propertyNameMapper) FieldName(t reflect.Type, f reflect.StructField) string {
	return ""
}

func uncapitalize(s string) string {
	return strings.ToLower(s[0:1]) + s[1:]
}

func (_ propertyNameMapper) MethodName(t reflect.Type, m reflect.Method) string {
	var doc dom.Document
	var document = reflect.TypeOf(&doc).Elem()
	if t.Implements(document) && m.Name == "Location" {
		return uncapitalize(m.Name)
	} else {
		return ""
	}
}

func (d *GojaContext) installGlobals(classes classMap) {
	d.globals = make(map[string]function)
	var assertGlobal func(class) function
	assertGlobal = func(class class) function {
		name := class.name
		wrapper := class.wrapper(d)
		if constructor, alreadyInstalled := d.globals[name]; alreadyInstalled {
			return constructor
		}
		constructor := d.vm.ToValue(wrapper.constructor).(*g.Object)
		constructor.DefineDataProperty(
			"name",
			d.vm.ToValue(name),
			g.FLAG_NOT_SET,
			g.FLAG_NOT_SET,
			g.FLAG_NOT_SET,
		)
		prototype := constructor.Get("prototype").(*g.Object)
		result := function{constructor, prototype, wrapper}
		d.vm.Set(name, constructor)
		d.globals[name] = result

		if super := class.superClassName; super != "" {
			if superclass, found := classes[super]; found {
				superPrototype := assertGlobal(superclass).Prototype
				prototype.SetPrototype(superPrototype)
			} else {
				panic(fmt.Sprintf("Superclass not installed for %s. Superclass: %s", name, super))
			}
		}

		if initializer, ok := wrapper.(wrapperPrototypeInitializer); ok {
			initializer.initializePrototype(prototype, d.vm)
		}

		return result
	}
	for _, class := range classes {
		assertGlobal(class)
	}
}

func (d *gojaScriptHost) NewContext(window html.Window) html.ScriptContext {
	vm := g.New()
	vm.SetFieldNameMapper(propertyNameMapper{})
	result := &GojaContext{
		vm:           vm,
		window:       window,
		wrappedGoObj: g.NewSymbol(INTERNAL_SYMBOL_NAME),
		cachedNodes:  make(map[int32]g.Value),
	}
	result.installGlobals(globals)

	globalThis := vm.GlobalObject()
	globalThis.DefineDataPropertySymbol(
		result.wrappedGoObj,
		vm.ToValue(window),
		g.FLAG_FALSE,
		g.FLAG_FALSE,
		g.FLAG_FALSE,
	)
	globalThis.Set("window", globalThis)
	globalThis.DefineAccessorProperty("document", vm.ToValue(func(c *g.FunctionCall) g.Value {
		return result.getObject(window.Document(), "Document")
	}), nil, g.FLAG_FALSE, g.FLAG_TRUE)
	globalThis.SetPrototype(result.globals["Window"].Prototype)

	return result
}

func (d *gojaScriptHost) Close() {}

type GojaContext struct {
	vm           *g.Runtime
	window       html.Window
	globals      map[string]function
	wrappedGoObj *g.Symbol
	cachedNodes  map[int32]g.Value
}

func (i *GojaContext) Close() {}

func (i *GojaContext) Run(str string) error {
	_, err := i.vm.RunString(str)
	return err
}

func (i *GojaContext) Eval(str string) (res any, err error) {
	if gojaVal, err := i.vm.RunString(str); err == nil {
		return gojaVal.Export(), nil
	} else {
		return nil, err
	}
}
