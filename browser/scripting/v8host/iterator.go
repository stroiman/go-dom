package v8host

import (
	"errors"
	"iter"

	"github.com/stroiman/go-dom/browser/dom"
	v8 "github.com/tommie/v8go"
)

type iterator[T any] struct {
	host           *V8ScriptHost
	ot             *v8.ObjectTemplate
	resultTemplate *v8.ObjectTemplate
	entityLookup   entityLookup[T]
}

type entityLookup[T any] func(value T, ctx *V8ScriptContext) (*v8.Value, error)

func newIterator[T any](host *V8ScriptHost, entityLookup entityLookup[T]) iterator[T] {
	iso := host.iso
	// TODO, once we have weak handles in v8, we can release the iterator when it
	// goes out of scope.
	iterator := iterator[T]{
		host,
		v8.NewObjectTemplate(host.iso),
		v8.NewObjectTemplate(host.iso),
		entityLookup,
	}
	iterator.ot.Set("next", v8.NewFunctionTemplateWithError(host.iso, iterator.next))
	iterator.ot.SetSymbol(
		v8.SymbolIterator(iso),
		v8.NewFunctionTemplateWithError(host.iso, iterator.newIterator),
	)
	iterator.ot.SetInternalFieldCount(2)
	return iterator
}

type iterable[T any] interface {
	All() iter.Seq[T]
}

type iteratorInstance[T any] struct {
	dom.Entity
	items iterable[T]
	next  func() (T, bool)
	stop  func()
}

func seqOfSlice[T any](items []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range items {
			if !yield(item) {
				return
			}
		}
	}
}

type sliceIterable[T any] struct {
	items []T
}

func (i sliceIterable[T]) All() iter.Seq[T] {
	return seqOfSlice(i.items)
}

func (i iterator[T]) newIteratorInstance(context *V8ScriptContext, items []T) (*v8.Value, error) {
	return i.newIteratorInstanceOfIterable(context, sliceIterable[T]{items})
}

func (i iterator[T]) newIteratorInstanceOfIterable(
	context *V8ScriptContext,
	items iterable[T],
) (*v8.Value, error) {
	seq := items.All()
	next, stop := iter.Pull(seq)

	iterator := &iteratorInstance[T]{
		dom.NewEntity(),
		items,
		next,
		stop,
	}
	res, err := i.ot.NewInstance(context.v8ctx)
	if err == nil {
		return context.cacheNode(res, iterator)
	}
	return res.Value, err
}

func (i iterator[T]) iso() *v8.Isolate { return i.host.iso }

func (i iterator[T]) next(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := i.host.mustGetContext(info.Context())
	instance, err := getInstanceFromThis[*iteratorInstance[T]](ctx, info.This())
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

func (i iterator[T]) newIterator(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := i.host.mustGetContext(info.Context())
	instance, err := getInstanceFromThis[*iteratorInstance[T]](ctx, info.This())
	if err != nil {
		return nil, err
	}
	return i.newIteratorInstanceOfIterable(ctx, instance.items)
}

func (i iterator[T]) createDoneIteratorResult(ctx *v8.Context) (*v8.Value, error) {
	result, err := i.resultTemplate.NewInstance(ctx)
	if err != nil {
		return nil, err
	}
	result.Set("done", true)
	return result.Value, nil
}

func (i iterator[T]) createNotDoneIteratorResult(
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
