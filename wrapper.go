package wrapper

// Wrapper wrap goroutine to run
type Wrapper interface {
	// Wrap exec func in goroutine without recover catch
	Wrap(fn func())

	// WrapWithRecover safely execute func in goroutine
	WrapWithRecover(fn func())

	// Wait wait all goroutine finish
	Wait()
}

// Option wrapper option
type Options struct {
	BufCap       int
	RecoveryFunc func()
}

// Options option func
type Option func(o *Options)

// WithBufCap set buf cap
func WithBufCap(c int) Option {
	return func(o *Options) {
		o.BufCap = c
	}
}

// WithRecover set recover func
func WithRecover(recoveryFunc func()) Option {
	return func(o *Options) {
		o.RecoveryFunc = recoveryFunc
	}
}
