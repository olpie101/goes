package command

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/modernice/goes/command/finish"
)

// ContextOption is a Context option.
type ContextOption func(*options)

type options struct {
	whenDone func(context.Context, finish.Config) error
}

type cmdctx[P any] struct {
	context.Context
	Of[P]

	mux sync.Mutex
	options
	finished bool
}

// WhenDone returns an Option that calls the provided function when the Finish()
// method of the context is called.
func WhenDone(fn func(context.Context, finish.Config) error) ContextOption {
	return func(opts *options) {
		opts.whenDone = fn
	}
}

// NewContext returns a context for the given command.
func NewContext[P any](base context.Context, cmd Of[P], opts ...ContextOption) Ctx[P] {
	ctx := cmdctx[P]{
		Context: base,
		Of:      cmd,
	}
	for _, opt := range opts {
		opt(&ctx.options)
	}
	return &ctx
}

func (ctx *cmdctx[P]) AggregateID() uuid.UUID {
	return ctx.Aggregate().ID
}

func (ctx *cmdctx[P]) AggregateName() string {
	return ctx.Aggregate().Name
}

func (ctx *cmdctx[P]) Finish(c context.Context, opts ...finish.Option) error {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	if ctx.finished {
		return nil
	}
	ctx.finished = true
	if ctx.whenDone != nil {
		return ctx.whenDone(c, finish.Configure(opts...))
	}
	return nil
}

// TryCastContext tries to cast the payload of the given context to the given
// `To` type. If the payload is not a `To`, TryCastContext returns false.
func TryCastContext[To, From any](ctx Ctx[From]) (Ctx[To], bool) {
	cmd, ok := TryCast[To, From](ctx)
	if !ok {
		return nil, false
	}

	var opts []ContextOption
	if ctx, ok := ctx.(*cmdctx[From]); ok {
		opts = append(opts, WhenDone(ctx.whenDone))
	}

	return NewContext[To](ctx, cmd, opts...), true
}

// CastContext casts the payload of the given context to the given `To` type.
// If the payload is not a `To`, CastContext panics.
func CastContext[To, From any](ctx Ctx[From]) Ctx[To] {
	cmd := Cast[To, From](ctx)

	var opts []ContextOption
	if ctx, ok := ctx.(*cmdctx[From]); ok {
		opts = append(opts, WhenDone(ctx.whenDone))
	}

	return NewContext[To](ctx, cmd, opts...)
}
