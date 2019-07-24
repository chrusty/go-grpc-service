package interfaces

// Tracer could be used for tracing (though it should probably either wrap or just BE openTracing.Tracer):
type Tracer interface {
	Trace()
}
