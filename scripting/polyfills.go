package scripting

import (
	_ "embed"
)

func installPolyfills(context *ScriptContext) error {
	return context.Run(`
		FormData.prototype.forEach = function(cb) {
			return Array.from(this).forEach(cb)
		}
	`)
}
