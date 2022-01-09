package waitgroup

import (
	"sync"

	"github.com/go-god/wrapper"
	"github.com/go-god/wrapper/grecover"
)

// WrapImpl sync.WaitGroup wrap impl
type WrapImpl struct {
	sync.WaitGroup
	recoveryFunc func()
}

// New create wrapper entity
func New() wrapper.Wrapper {
	w := &WrapImpl{
		recoveryFunc: grecover.DefaultRecovery,
	}

	return w
}

// Wrap fn func in goroutine to run
func (w *WrapImpl) Wrap(fn func()) {
	w.Add(1)
	go func() {
		defer w.Done()
		fn()
	}()
}

// WrapWithRecover exec func with recover
func (w *WrapImpl) WrapWithRecover(fn func()) {
	w.Add(1)
	go func() {
		defer w.recoveryFunc()
		defer w.Done()
		fn()
	}()
}

func (w *WrapImpl) WithRecover(recoveryFunc func()) {
	w.recoveryFunc = recoveryFunc
}
