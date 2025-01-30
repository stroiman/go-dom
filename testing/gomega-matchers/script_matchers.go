package matchers

import (
	"errors"
	"reflect"

	"github.com/onsi/gomega/gcustom"
	"github.com/onsi/gomega/types"
	"github.com/gost-dom/browser/html"
)

// MatcherScriptContext defines a set of functions a script context must support
// for the matchers to work.
type MatcherScriptContext interface {
	html.ScriptContext
	EvalCore(string) (any, error)
	RunFunction(string, ...any) (any, error)
	Export(any) (any, error)
}

type ScriptMatchers struct{ Ctx MatcherScriptContext }

func (m ScriptMatchers) eval(script string) (any, error) { return m.Ctx.EvalCore(script) }
func (m ScriptMatchers) runScript(script string, args ...any) (any, error) {
	return m.Ctx.RunFunction(script, args...)
}

func (m ScriptMatchers) exportTo(src, dst any) error {
	v, err := m.Ctx.Export(src)
	if err != nil {
		return err
	}
	var (
		srcVal = reflect.ValueOf(v)
		dstVal = reflect.ValueOf(dst)
	)
	if dstVal.Kind() != reflect.Pointer {
		return errors.New("dst is not a pointer")
	}
	dstVal = dstVal.Elem()

	if !dstVal.CanSet() || !srcVal.Type().AssignableTo(dstVal.Type()) {
		return errors.New("Cannot assign value to dst")
	}

	dstVal.Set(srcVal)
	return nil
}

func (m ScriptMatchers) BeInstanceOf(name string) types.GomegaMatcher {
	return gcustom.MakeMatcher(func(script string) (res bool, err error) {
		var v any
		if v, err = m.eval(script); err == nil {
			if v, err = m.runScript("x => x instanceof "+name, v); err == nil {
				err = m.exportTo(v, &res)
			}
		}
		return
	})
}
