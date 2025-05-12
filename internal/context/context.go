package context

import "context"

// ProgramContext is the context of the program.
type ProgramContext struct {
	ctx context.Context
}

// New returns a new context.
func New(ctx context.Context) *ProgramContext {
	return &ProgramContext{
		ctx: ctx,
	}
}

// Context returns the context.
func (c *ProgramContext) Context() context.Context {
	return c.ctx
}
