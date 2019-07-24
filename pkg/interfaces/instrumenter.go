package interfaces

// Instrumenter could be used for metrics:
type Instrumenter interface {
	Count()
	Guage()
	Timing()
}
