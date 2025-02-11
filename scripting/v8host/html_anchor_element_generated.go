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
func (w hTMLAnchorElementV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := w.scriptHost.iso

	prototypeTmpl.SetAccessorProperty("target",
		v8.NewFunctionTemplateWithError(iso, w.target),
		v8.NewFunctionTemplateWithError(iso, w.setTarget),
		v8.None)
	prototypeTmpl.SetAccessorProperty("href",
		v8.NewFunctionTemplateWithError(iso, w.href),
		v8.NewFunctionTemplateWithError(iso, w.setHref),
		v8.None)
	prototypeTmpl.SetAccessorProperty("origin",
		v8.NewFunctionTemplateWithError(iso, w.origin),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("protocol",
		v8.NewFunctionTemplateWithError(iso, w.protocol),
		v8.NewFunctionTemplateWithError(iso, w.setProtocol),
		v8.None)
	prototypeTmpl.SetAccessorProperty("username",
		v8.NewFunctionTemplateWithError(iso, w.username),
		v8.NewFunctionTemplateWithError(iso, w.setUsername),
		v8.None)
	prototypeTmpl.SetAccessorProperty("password",
		v8.NewFunctionTemplateWithError(iso, w.password),
		v8.NewFunctionTemplateWithError(iso, w.setPassword),
		v8.None)
	prototypeTmpl.SetAccessorProperty("host",
		v8.NewFunctionTemplateWithError(iso, w.host),
		v8.NewFunctionTemplateWithError(iso, w.setHost),
		v8.None)
	prototypeTmpl.SetAccessorProperty("hostname",
		v8.NewFunctionTemplateWithError(iso, w.hostname),
		v8.NewFunctionTemplateWithError(iso, w.setHostname),
		v8.None)
	prototypeTmpl.SetAccessorProperty("port",
		v8.NewFunctionTemplateWithError(iso, w.port),
		v8.NewFunctionTemplateWithError(iso, w.setPort),
		v8.None)
	prototypeTmpl.SetAccessorProperty("pathname",
		v8.NewFunctionTemplateWithError(iso, w.pathname),
		v8.NewFunctionTemplateWithError(iso, w.setPathname),
		v8.None)
	prototypeTmpl.SetAccessorProperty("search",
		v8.NewFunctionTemplateWithError(iso, w.search),
		v8.NewFunctionTemplateWithError(iso, w.setSearch),
		v8.None)
	prototypeTmpl.SetAccessorProperty("hash",
		v8.NewFunctionTemplateWithError(iso, w.hash),
		v8.NewFunctionTemplateWithError(iso, w.setHash),
		v8.None)
}

func (w hTMLAnchorElementV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(w.scriptHost.iso, "Illegal Constructor")
}

func (w hTMLAnchorElementV8Wrapper) target(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.target")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Target()
	return w.toDOMString(ctx, result)
}

func (w hTMLAnchorElementV8Wrapper) setTarget(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setTarget")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	val, err1 := tryParseArg(args, 0, w.decodeDOMString)
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

func (w hTMLAnchorElementV8Wrapper) href(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.href")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Href()
	return w.toUSVString(ctx, result)
}

func (w hTMLAnchorElementV8Wrapper) setHref(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setHref")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	val, err1 := tryParseArg(args, 0, w.decodeUSVString)
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

func (w hTMLAnchorElementV8Wrapper) origin(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.origin")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Origin()
	return w.toUSVString(ctx, result)
}

func (w hTMLAnchorElementV8Wrapper) protocol(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.protocol")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Protocol()
	return w.toUSVString(ctx, result)
}

func (w hTMLAnchorElementV8Wrapper) setProtocol(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setProtocol")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	val, err1 := tryParseArg(args, 0, w.decodeUSVString)
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

func (w hTMLAnchorElementV8Wrapper) username(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.username")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Username()
	return w.toUSVString(ctx, result)
}

func (w hTMLAnchorElementV8Wrapper) setUsername(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setUsername")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	val, err1 := tryParseArg(args, 0, w.decodeUSVString)
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

func (w hTMLAnchorElementV8Wrapper) password(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.password")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Password()
	return w.toUSVString(ctx, result)
}

func (w hTMLAnchorElementV8Wrapper) setPassword(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setPassword")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	val, err1 := tryParseArg(args, 0, w.decodeUSVString)
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

func (w hTMLAnchorElementV8Wrapper) host(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.host")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Host()
	return w.toUSVString(ctx, result)
}

func (w hTMLAnchorElementV8Wrapper) setHost(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setHost")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	val, err1 := tryParseArg(args, 0, w.decodeUSVString)
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

func (w hTMLAnchorElementV8Wrapper) hostname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.hostname")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Hostname()
	return w.toUSVString(ctx, result)
}

func (w hTMLAnchorElementV8Wrapper) setHostname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setHostname")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	val, err1 := tryParseArg(args, 0, w.decodeUSVString)
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

func (w hTMLAnchorElementV8Wrapper) port(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.port")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Port()
	return w.toUSVString(ctx, result)
}

func (w hTMLAnchorElementV8Wrapper) setPort(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setPort")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	val, err1 := tryParseArg(args, 0, w.decodeUSVString)
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

func (w hTMLAnchorElementV8Wrapper) pathname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.pathname")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Pathname()
	return w.toUSVString(ctx, result)
}

func (w hTMLAnchorElementV8Wrapper) setPathname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setPathname")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	val, err1 := tryParseArg(args, 0, w.decodeUSVString)
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

func (w hTMLAnchorElementV8Wrapper) search(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.search")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Search()
	return w.toUSVString(ctx, result)
}

func (w hTMLAnchorElementV8Wrapper) setSearch(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setSearch")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	val, err1 := tryParseArg(args, 0, w.decodeUSVString)
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

func (w hTMLAnchorElementV8Wrapper) hash(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: HTMLAnchorElement.hash")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Hash()
	return w.toUSVString(ctx, result)
}

func (w hTMLAnchorElementV8Wrapper) setHash(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: HTMLAnchorElement.setHash")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	val, err1 := tryParseArg(args, 0, w.decodeUSVString)
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
