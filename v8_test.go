package go_dom_test

import (
	// . "github.com/stroiman/go-dom"
	"fmt"
	"runtime"

	v8 "github.com/tommie/v8go"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Test struct {
	count int
}

var _ = Describe("Parser", func() {
	It("Can run scripts", func() {
		iso := v8.NewIsolate()
		wrapper := v8.NewObjectTemplate(iso)
		wrapper.SetInternalFieldCount(1)
		dummy := &Test{4}
		getter := v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
			var field interface{} = info.This().GetInternalField(0)
			obj, ok := field.(Test)
			fmt.Printf("OK: %v, obj: %v", ok, obj)
			return nil
		})
		setter := v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
			var field = info.This().GetInternalField(0)
			ext, _ := (field.ExternalInterface()).(*Test)
			ext.count = (int)(info.Args()[0].Integer())
			return nil
		})
		wrapper.SetAccessorProperty("count", *getter, *setter)
		ctx := v8.NewContext(iso)
		defer ctx.Close()
		global := ctx.Global()
		object, _ := wrapper.NewInstance(ctx)
		p := new(runtime.Pinner)
		p.Pin(dummy)

		err := object.SetInternalField(0, v8.NewExternalFromInterface(iso, dummy))
		Expect(err).ToNot(HaveOccurred())
		fmt.Printf("OBJECT: %v\n", dummy)
		global.Set("root", object)
		_, err = (ctx.RunScript("root.count = 2", ""))
		Expect(err).ToNot(HaveOccurred())
		Expect(dummy.count).To(Equal(2))
	})
})
