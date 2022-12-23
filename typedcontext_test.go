package typedcontext

import (
	"context"
	"reflect"
	"testing"
)

func TestNormalOperation(t *testing.T) {
	ctx := context.Background()
	ctx = New(ctx, 10)
	if MustGet[int](ctx) != 10 {
		t.FailNow()
	}
	if _, ok := Get[float64](ctx); ok {
		t.FailNow()
	}
}

func TestIsolatedFromExplicitTypeReflection(t *testing.T) {
	ctx := context.Background()
	ctx = New(ctx, 10)
	ctx = context.WithValue(ctx, reflect.TypeOf(20), 20)
	if MustGet[int](ctx) != 10 {
		t.FailNow()
	}
}

func TestPanicIfNoValue(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.FailNow()
		}
	}()
	MustGet[int](context.Background())
}
