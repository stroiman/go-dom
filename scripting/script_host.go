package scripting

import (
	"errors"
	"runtime"

	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

type ScriptHost struct {
	iso            *v8.Isolate
	windowTemplate *v8.ObjectTemplate
	document       *v8.FunctionTemplate
	node           *v8.FunctionTemplate
	eventTarget    *v8.FunctionTemplate
	contexts       map[*v8.Context]*ScriptContext
}

func (h *ScriptHost) GetContext(v8ctx *v8.Context) (*ScriptContext, bool) {
	ctx, ok := h.contexts[v8ctx]
	return ctx, ok
}

type ScriptContext struct {
	host     *ScriptHost
	v8ctx    *v8.Context
	window   Window
	pinner   runtime.Pinner
	v8nodes  map[ObjectId]*v8.Value
	domNodes map[ObjectId]Node
}

func (c *ScriptContext) GetInstanceForNode(
	prototype *v8.FunctionTemplate,
	node Node,
) (*v8.Value, error) {
	iso := c.host.iso
	if node == nil {
		return v8.Null(iso), nil
	}
	value, err := prototype.GetInstanceTemplate().NewInstance(c.v8ctx)
	if err == nil {
		objectId := node.ObjectId()
		if cached, ok := c.v8nodes[objectId]; ok {
			return cached, nil
		}
		c.v8nodes[objectId] = value.Value
		c.domNodes[objectId] = node
		internal, err := v8.NewValue(iso, objectId)
		if err != nil {
			return nil, err
		}
		value.SetInternalField(0, internal)
		return value.Value, nil
	}
	return nil, err
}

func CreateWindowTemplate(host *ScriptHost) *v8.ObjectTemplate {
	iso := host.iso
	windowTemplate := v8.NewObjectTemplate(iso)
	windowTemplate.SetInternalFieldCount(1)
	windowTemplate.SetAccessorProperty(
		"window",
		v8.AccessProp{
			Get: func(i *v8.FunctionCallbackInfo) *v8.Value {
				return i.This().Value
			},
			Attributes: v8.ReadOnly,
		},
	)
	windowTemplate.SetAccessorPropertyWithError(
		"document",
		v8.AccessPropWithError{
			Get: func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				if ctx, ok := host.GetContext(info.Context()); ok {
					return ctx.GetInstanceForNode(host.document, ctx.window.Document())
				}
				return nil, errors.New("Must have a context")
			},
		})
	windowTemplate.Set("Document", host.document)
	windowTemplate.Set("Node", host.node)
	windowTemplate.Set("EventTarget", host.eventTarget)
	return windowTemplate
}

func NewScriptHost() *ScriptHost {
	host := &ScriptHost{iso: v8.NewIsolate()}
	host.document = CreateDocumentPrototype(host)
	host.node = CreateNode(host.iso)
	host.eventTarget = CreateEventTarget(host)
	host.windowTemplate = CreateWindowTemplate(host)
	host.document.Inherit(host.node)
	host.node.Inherit(host.eventTarget)
	host.contexts = make(map[*v8.Context]*ScriptContext)
	return host
}

func (host *ScriptHost) Dispose() {
	host.iso.Dispose()
}

var global *v8.Object

func (host *ScriptHost) NewContext() *ScriptContext {
	window := NewWindow()
	context := &ScriptContext{
		host:     host,
		v8ctx:    v8.NewContext(host.iso, host.windowTemplate),
		window:   window,
		v8nodes:  make(map[ObjectId]*v8.Value),
		domNodes: make(map[ObjectId]Node),
	}
	global = context.v8ctx.Global()
	host.contexts[context.v8ctx] = context

	return context
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func (host *ScriptHost) createPrototypeChains() {
	host.document.Inherit(host.node)
}

func (ctx *ScriptContext) Dispose() {
	delete(ctx.host.contexts, ctx.v8ctx)
	ctx.v8ctx.Close()
}

func (ctx *ScriptContext) RunScript(script string) (*v8.Value, error) {
	return ctx.v8ctx.RunScript(script, "")
}

func (ctx *ScriptContext) Run(script string) (interface{}, error) {
	return ctx.v8ctx.RunScript(script, "")
}

func (ctx *ScriptContext) Window() Window {
	return ctx.window
}
