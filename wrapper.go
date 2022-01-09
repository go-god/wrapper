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
type Option struct {
	BufCap       int
	RecoveryFunc func()
}

// Options option func
type Options func(o *Option)

// WithBufCap set buf cap
func WithBufCap(c int) Options {
	return func(o *Option) {
		o.BufCap = c
	}
}

// WithRecover set recover func
func WithRecover(recoveryFunc func()) Options {
	return func(o *Option) {
		o.RecoveryFunc = recoveryFunc
	}
}
