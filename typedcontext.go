package typedcontext

import (
	"context"
	"reflect"
)

type key struct{ t reflect.Type }

func New[T any](ctx context.Context, val T) context.Context {
	return context.WithValue(ctx, key{reflect.TypeOf(val)}, val)
}

func Get[T any](ctx context.Context) (T, bool) {
	v, ok := ctx.Value(key{reflect.TypeOf((*T)(nil)).Elem()}).(T)
	return v, ok
}

func MustGet[T any](ctx context.Context) T {
	v, ok := Get[T](ctx)
	if !ok {
		panic("context doesn't have the value of that type")
	}
	return v
}
