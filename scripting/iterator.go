package scripting

import (
	"errors"
	"iter"

	"github.com/stroiman/go-dom/browser"
	v8 "github.com/tommie/v8go"
)

type Iterator[T any] struct {
	host           *ScriptHost
	ot             *v8.ObjectTemplate
	resultTemplate *v8.ObjectTemplate
	entityLookup   EntityLookup[T]
}

type EntityLookup[T any] func(value T, ctx *ScriptContext) (*v8.Value, error)

func NewIterator[T any](host *ScriptHost, entityLookup EntityLookup[T]) Iterator[T] {
	iso := host.iso
	// TODO, once we have weak handles in v8, we can release the iterator when it
	// goes out of scope.
	iterator := Iterator[T]{
		host,
		v8.NewObjectTemplate(host.iso),
		v8.NewObjectTemplate(host.iso),
		entityLookup,
	}
	iterator.ot.Set("next", v8.NewFunctionTemplateWithError(host.iso, iterator.Next))
	iterator.ot.SetSymbol(
		v8.SymbolIterator(iso),
		v8.NewFunctionTemplateWithError(host.iso, iterator.NewIterator),
	)
	iterator.ot.SetInternalFieldCount(2)
	return iterator
}

type IteratorInstance[T any] struct {
	browser.Entity
	items []T
	next  func() (T, bool)
	stop  func()
}

func SeqOfSlice[T any](items []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range items {
			if !yield(item) {
				return
			}
		}
	}

}

func (i Iterator[T]) NewIteratorInstance(context *ScriptContext, items []T) (*v8.Value, error) {
	seq := SeqOfSlice(items)
	next, stop := iter.Pull(seq)
	iterator := &IteratorInstance[T]{
		browser.NewEntity(),
		items,
		next,
		stop,
	}
	res, err := i.ot.NewInstance(context.v8ctx)
	if err == nil {
		return context.CacheNode(res, iterator)
	}
	return res.Value, err
}

func (i Iterator[T]) iso() *v8.Isolate { return i.host.iso }

func (i Iterator[T]) Next(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := i.host.MustGetContext(info.Context())
	instance, err := getInstanceFromThis[*IteratorInstance[T]](ctx, info.This())
	if err != nil {
		return nil, err
	}
	next := instance.next
	stop := instance.stop
	index := info.This().GetInternalField(1).Int32()
	if item, ok := next(); !ok {
		stop()
		return i.createDoneIteratorResult(ctx.v8ctx)
	} else {
		value, err1 := i.entityLookup(item, ctx)
		result, err2 := i.createNotDoneIteratorResult(ctx.v8ctx, value)
		err3 := info.This().SetInternalField(1, index+1)
		return result, errors.Join(err1, err2, err3)
	}
}

func (i Iterator[T]) NewIterator(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := i.host.MustGetContext(info.Context())
	instance, err := getInstanceFromThis[*IteratorInstance[T]](ctx, info.This())
	if err != nil {
		return nil, err
	}
	return i.NewIteratorInstance(ctx, instance.items)
}

func (i Iterator[T]) createDoneIteratorResult(ctx *v8.Context) (*v8.Value, error) {
	result, err := i.resultTemplate.NewInstance(ctx)
	if err != nil {
		return nil, err
	}
	result.Set("done", true)
	return result.Value, nil
}

func (i Iterator[T]) createNotDoneIteratorResult(
	ctx *v8.Context,
	value interface{},
) (*v8.Value, error) {
	result, err := i.resultTemplate.NewInstance(ctx)
	if err != nil {
		return nil, err
	}
	result.Set("done", false)
	result.Set("value", value)
	return result.Value, nil
}
