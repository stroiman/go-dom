package scripting

import (
	_ "embed"
)

//go:embed polyfills/url-polyfill.js
var urlPolyfill []byte

func installPolyfills(context *ScriptContext) error {
	// _, err := context.RunScript(string(urlPolyfill))
	return nil
}
