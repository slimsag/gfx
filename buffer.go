package gfx

type BufferUsage int

const(
	StaticDraw BufferUsage = iota
	DynamicDraw
	StreamDraw
)

// Buffer is a buffer object that contains data such as vertices or colors.
type Buffer interface {
	// DataSize prepares this buffer with size bytes of storage, initialized to
	// zero.
	//
	// The usage hint is only a performance hint, it has no effect on the
	// actual usage of the buffer.
	DataSize(size int, usage UsageHint)

	// Data prepares this buffer with the given data.
	//
	// The usage hint is only a performance hint, it has no effect on the
	// actual usage of the buffer.
	//
	// TODO(slimsag): typeof(data) == ArrayBuffer
	Data(data interface{}, usage UsageHint)
}
