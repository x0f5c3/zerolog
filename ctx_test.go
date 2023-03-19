package zerolog

import (
	"context"
	"io"
	"reflect"
	"testing"
)

func TestCtx(t *testing.T) {
	l := New(io.Discard)
	ctx := l.WithContext(context.Background())
	log2 := Ctx(ctx)
	if !reflect.DeepEqual(l, *log2) {
		t.Error("Ctx did not return the expected logger")
	}

	// update
	l = l.Level(InfoLevel)
	ctx = l.WithContext(ctx)
	log2 = Ctx(ctx)
	if !reflect.DeepEqual(l, *log2) {
		t.Error("Ctx did not return the expected logger")
	}

	log2 = Ctx(context.Background())
	if log2 != disabledLogger {
		t.Error("Ctx did not return the expected logger")
	}

	DefaultContextLogger = l
	t.Cleanup(func() { DefaultContextLogger = nil })
	log2 = Ctx(context.Background())
	if log2 != l {
		t.Error("Ctx did not return the expected logger")
	}
}

func TestCtxDisabled(t *testing.T) {
	dl := New(io.Discard).Level(Disabled)
	ctx := dl.WithContext(context.Background())
	if ctx != context.Background() {
		t.Error("WithContext stored a disabled logger")
	}

	l := New(io.Discard).With().Str("foo", "bar").Logger()
	ctx = l.WithContext(ctx)
	if !reflect.DeepEqual(Ctx(ctx), &l) {
		t.Error("WithContext did not store logger")
	}

	l.UpdateContext(func(c Context) Context {
		return c.Str("bar", "baz")
	})
	ctx = l.WithContext(ctx)
	if !reflect.DeepEqual(Ctx(ctx), &l) {
		t.Error("WithContext did not store updated logger")
	}

	l = l.Level(DebugLevel)
	ctx = l.WithContext(ctx)
	if !reflect.DeepEqual(Ctx(ctx), &l) {
		t.Error("WithContext did not store copied logger")
	}

	ctx = dl.WithContext(ctx)
	if !reflect.DeepEqual(Ctx(ctx), &dl) {
		t.Error("WithContext did not override logger with a disabled logger")
	}
}
