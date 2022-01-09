package chanwrap

import (
	"github.com/go-god/wrapper"
	"github.com/go-god/wrapper/grecover"
)

var est = struct{}{}

// WrapImpl wrapper impl
type WrapImpl struct {
	bufNum       int
	bufCh        chan struct{}
	recoveryFunc func()
}

// New create wrapImpl entity
func New(c int) wrapper.Wrapper {
	w := &WrapImpl{
		bufNum:       c,
		bufCh:        make(chan struct{}, c),
		recoveryFunc: grecover.DefaultRecovery,
	}

	return w
}

// Wrap exec func in goroutine without recover catch
func (c *WrapImpl) Wrap(fn func()) {
	go func() {
		defer c.done()
		fn()
	}()
}

// WrapWithRecover safely execute func in goroutine
func (c *WrapImpl) WrapWithRecover(fn func()) {
	go func() {
		defer c.recoveryFunc()
		defer c.done()
		fn()
	}()
}

// Wait wait all goroutine finish
func (c *WrapImpl) Wait() {
	for i := 0; i < c.bufNum; i++ {
		<-c.bufCh
	}
}

// WithRecover set recover func
func (c *WrapImpl) WithRecover(recoveryFunc func()) {
	c.recoveryFunc = recoveryFunc
}

func (c *WrapImpl) done() {
	c.bufCh <- est
}
