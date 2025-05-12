package context

import "context"

// ProgramContext is the context of the program.
type ProgramContext struct {
	appName    string
	appVersion string
	ctx        context.Context
}

// New returns a new context.
func New(ctx context.Context) *ProgramContext {
	return &ProgramContext{
		ctx: ctx,
	}
}

// SetAppName sets the app name.
func (c *ProgramContext) SetAppName(name string) {
	c.appName = name
}

// SetAppVersion sets the app version.
func (c *ProgramContext) SetAppVersion(version string) {
	c.appVersion = version
}

// Context returns the context.
func (c *ProgramContext) Context() context.Context {
	return c.ctx
}

// GetAppName returns the app name.
func (c *ProgramContext) GetAppName() string {
	return c.appName
}

// GetAppVersion returns the app version.
func (c *ProgramContext) GetAppVersion() string {
	return c.appVersion
}
