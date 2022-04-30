package waitgroup

import (
	"sync"

	"github.com/go-god/wrapper"
	"github.com/go-god/wrapper/grecover"
)

var _ wrapper.Wrapper = (*wrapImpl)(nil)

// wrapImpl sync.WaitGroup wrap impl
type wrapImpl struct {
	sync.WaitGroup
	recoveryFunc func()
}

// New create wrapper entity
func New(opts ...wrapper.Option) wrapper.Wrapper {
	w := &wrapImpl{}
	option := &wrapper.Options{}
	for _, o := range opts {
		o(option)
	}

	w.recoveryFunc = option.RecoveryFunc
	if w.recoveryFunc == nil {
		w.recoveryFunc = grecover.DefaultRecovery
	}

	return w
}

// Wrap fn func in goroutine to run
func (w *wrapImpl) Wrap(fn func()) {
	w.Add(1)
	go func() {
		defer w.Done()
		fn()
	}()
}

// WrapWithRecover exec func with recover
func (w *wrapImpl) WrapWithRecover(fn func()) {
	w.Add(1)
	go func() {
		defer w.recoveryFunc()
		defer w.Done()
		fn()
	}()
}
