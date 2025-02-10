// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	log "github.com/gost-dom/browser/internal/log"
	v8 "github.com/tommie/v8go"
)

type hTMLAnchorElementV8Wrapper struct {
	nodeV8WrapperBase[html.HTMLAnchorElement]
}

func newHTMLAnchorElementV8Wrapper(scriptHost *V8ScriptHost) *hTMLAnchorElementV8Wrapper {
	return &hTMLAnchorElementV8Wrapper{newNodeV8WrapperBase[html.HTMLAnchorElement](scriptHost)}
}

func init() {
	registerJSClass("HTMLAnchorElement", "HTMLElement", createHTMLAnchorElementPrototype)
}

func createHTMLAnchorElementPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newHTMLAnchorElementV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (e hTMLAnchorElementV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := e.scriptHost.iso

	prototypeTmpl.SetAccessorProperty("target",
		v8.NewFunctionTemplateWithError(iso, e.target),
		v8.NewFunctionTemplateWithError(iso, e.setTarget),
		v8.None)
	prototypeTmpl.SetAccessorProperty("href",
		v8.NewFunctionTemplateWithError(iso, e.href),
		v8.NewFunctionTemplateWithError(iso, e.setHref),
		v8.None)
	prototypeTmpl.SetAccessorProperty("origin",
		v8.NewFunctionTemplateWithError(iso, e.origin),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("protocol",
		v8.NewFunctionTemplateWithError(iso, e.protocol),
		v8.NewFunctionTemplateWithError(iso, e.setProtocol),
		v8.None)
	prototypeTmpl.SetAccessorProperty("username",
		v8.NewFunctionTemplateWithError(iso, e.username),
		v8.NewFunctionTemplateWithError(iso, e.setUsername),
		v8.None)
	prototypeTmpl.SetAccessorProperty("password",
		v8.NewFunctionTemplateWithError(iso, e.password),
		v8.NewFunctionTemplateWithError(iso, e.setPassword),
		v8.None)
	prototypeTmpl.SetAccessorProperty("host",
		v8.NewFunctionTemplateWithError(iso, e.host),
		v8.NewFunctionTemplateWithError(iso, e.setHost),
		v8.None)
	prototypeTmpl.SetAccessorProperty("hostname",
		v8.NewFunctionTemplateWithError(iso, e.hostname),
		v8.NewFunctionTemplateWithError(iso, e.setHostname),
		v8.None)
	prototypeTmpl.SetAccessorProperty("port",
		v8.NewFunctionTemplateWithError(iso, e.port),
		v8.NewFunctionTemplateWithError(iso, e.setPort),
		v8.None)
	prototypeTmpl.SetAccessorProperty("pathname",
		v8.NewFunctionTemplateWithError(iso, e.pathname),
		v8.NewFunctionTemplateWithError(iso, e.setPathname),
		v8.None)
	prototypeTmpl.SetAccessorProperty("search",
		v8.NewFunctionTemplateWithError(iso, e.search),
		v8.NewFunctionTemplateWithError(iso, e.setSearch),
		v8.None)
	prototypeTmpl.SetAccessorProperty("hash",
		v8.NewFunctionTemplateWithError(iso, e.hash),
		v8.NewFunctionTemplateWithError(iso, e.setHash),
		v8.None)
}

func (e hTMLAnchorElementV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(e.scriptHost.iso, "Illegal Constructor")
}

func (e hTMLAnchorElementV8Wrapper) target(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.target")
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Target()
	return e.toDOMString(ctx, result)
}

func (e hTMLAnchorElementV8Wrapper) setTarget(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setTarget")
	args := newArgumentHelper(e.scriptHost, info)
	instance, err0 := e.getInstance(info)
	val, err1 := tryParseArg(args, 0, e.decodeDOMString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetTarget(val)
		return nil, nil
	}
	return nil, errors.New("HTMLAnchorElement.setTarget: Missing arguments")
}

func (e hTMLAnchorElementV8Wrapper) href(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.href")
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Href()
	return e.toUSVString(ctx, result)
}

func (e hTMLAnchorElementV8Wrapper) setHref(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setHref")
	args := newArgumentHelper(e.scriptHost, info)
	instance, err0 := e.getInstance(info)
	val, err1 := tryParseArg(args, 0, e.decodeUSVString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetHref(val)
		return nil, nil
	}
	return nil, errors.New("HTMLAnchorElement.setHref: Missing arguments")
}

func (e hTMLAnchorElementV8Wrapper) origin(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.origin")
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Origin()
	return e.toUSVString(ctx, result)
}

func (e hTMLAnchorElementV8Wrapper) protocol(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.protocol")
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Protocol()
	return e.toUSVString(ctx, result)
}

func (e hTMLAnchorElementV8Wrapper) setProtocol(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setProtocol")
	args := newArgumentHelper(e.scriptHost, info)
	instance, err0 := e.getInstance(info)
	val, err1 := tryParseArg(args, 0, e.decodeUSVString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetProtocol(val)
		return nil, nil
	}
	return nil, errors.New("HTMLAnchorElement.setProtocol: Missing arguments")
}

func (e hTMLAnchorElementV8Wrapper) username(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.username")
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Username()
	return e.toUSVString(ctx, result)
}

func (e hTMLAnchorElementV8Wrapper) setUsername(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setUsername")
	args := newArgumentHelper(e.scriptHost, info)
	instance, err0 := e.getInstance(info)
	val, err1 := tryParseArg(args, 0, e.decodeUSVString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetUsername(val)
		return nil, nil
	}
	return nil, errors.New("HTMLAnchorElement.setUsername: Missing arguments")
}

func (e hTMLAnchorElementV8Wrapper) password(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.password")
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Password()
	return e.toUSVString(ctx, result)
}

func (e hTMLAnchorElementV8Wrapper) setPassword(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setPassword")
	args := newArgumentHelper(e.scriptHost, info)
	instance, err0 := e.getInstance(info)
	val, err1 := tryParseArg(args, 0, e.decodeUSVString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetPassword(val)
		return nil, nil
	}
	return nil, errors.New("HTMLAnchorElement.setPassword: Missing arguments")
}

func (e hTMLAnchorElementV8Wrapper) host(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.host")
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Host()
	return e.toUSVString(ctx, result)
}

func (e hTMLAnchorElementV8Wrapper) setHost(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setHost")
	args := newArgumentHelper(e.scriptHost, info)
	instance, err0 := e.getInstance(info)
	val, err1 := tryParseArg(args, 0, e.decodeUSVString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetHost(val)
		return nil, nil
	}
	return nil, errors.New("HTMLAnchorElement.setHost: Missing arguments")
}

func (e hTMLAnchorElementV8Wrapper) hostname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.hostname")
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Hostname()
	return e.toUSVString(ctx, result)
}

func (e hTMLAnchorElementV8Wrapper) setHostname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setHostname")
	args := newArgumentHelper(e.scriptHost, info)
	instance, err0 := e.getInstance(info)
	val, err1 := tryParseArg(args, 0, e.decodeUSVString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetHostname(val)
		return nil, nil
	}
	return nil, errors.New("HTMLAnchorElement.setHostname: Missing arguments")
}

func (e hTMLAnchorElementV8Wrapper) port(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.port")
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Port()
	return e.toUSVString(ctx, result)
}

func (e hTMLAnchorElementV8Wrapper) setPort(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setPort")
	args := newArgumentHelper(e.scriptHost, info)
	instance, err0 := e.getInstance(info)
	val, err1 := tryParseArg(args, 0, e.decodeUSVString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetPort(val)
		return nil, nil
	}
	return nil, errors.New("HTMLAnchorElement.setPort: Missing arguments")
}

func (e hTMLAnchorElementV8Wrapper) pathname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.pathname")
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Pathname()
	return e.toUSVString(ctx, result)
}

func (e hTMLAnchorElementV8Wrapper) setPathname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setPathname")
	args := newArgumentHelper(e.scriptHost, info)
	instance, err0 := e.getInstance(info)
	val, err1 := tryParseArg(args, 0, e.decodeUSVString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetPathname(val)
		return nil, nil
	}
	return nil, errors.New("HTMLAnchorElement.setPathname: Missing arguments")
}

func (e hTMLAnchorElementV8Wrapper) search(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.search")
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Search()
	return e.toUSVString(ctx, result)
}

func (e hTMLAnchorElementV8Wrapper) setSearch(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setSearch")
	args := newArgumentHelper(e.scriptHost, info)
	instance, err0 := e.getInstance(info)
	val, err1 := tryParseArg(args, 0, e.decodeUSVString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetSearch(val)
		return nil, nil
	}
	return nil, errors.New("HTMLAnchorElement.setSearch: Missing arguments")
}

func (e hTMLAnchorElementV8Wrapper) hash(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.hash")
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Hash()
	return e.toUSVString(ctx, result)
}

func (e hTMLAnchorElementV8Wrapper) setHash(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setHash")
	args := newArgumentHelper(e.scriptHost, info)
	instance, err0 := e.getInstance(info)
	val, err1 := tryParseArg(args, 0, e.decodeUSVString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetHash(val)
		return nil, nil
	}
	return nil, errors.New("HTMLAnchorElement.setHash: Missing arguments")
}
