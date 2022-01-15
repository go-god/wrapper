package chanwrap

import (
	"github.com/go-god/wrapper"
	"github.com/go-god/wrapper/grecover"
)

var est = struct{}{}

// WrapImpl wrapper impl
type WrapImpl struct {
	bufCap       int
	bufCh        chan struct{}
	recoveryFunc func()
}

// New create wrapImpl entity
// If the wrapper using the chan method needs to specify the number of
// goroutines to be executed,the wrapper.WithBufCap method needs to be called.
// Otherwise, after the Wait method is executed, some goroutines
// will exit without execution.
func New(opts ...wrapper.Option) wrapper.Wrapper {
	w := &WrapImpl{}

	var option = &wrapper.Options{}
	for _, o := range opts {
		o(option)
	}

	w.recoveryFunc = option.RecoveryFunc
	if w.recoveryFunc == nil {
		w.recoveryFunc = grecover.DefaultRecovery
	}

	w.bufCap = option.BufCap
	w.bufCh = make(chan struct{}, w.bufCap)

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
	for i := 0; i < c.bufCap; i++ {
		<-c.bufCh
	}
}

func (c *WrapImpl) done() {
	c.bufCh <- est
}
