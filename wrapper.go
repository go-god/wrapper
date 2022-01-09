package wrapper

// Wrapper wrap goroutine to run
type Wrapper interface {
	// Wrap exec func in goroutine without recover catch
	Wrap(fn func())

	// WrapWithRecover safely execute func in goroutine
	WrapWithRecover(fn func())

	// WithRecover set recover func
	WithRecover(recoveryFunc func())

	// Wait wait all goroutine finish
	Wait()
}
