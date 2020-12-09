package day09

// PairNotFound occurs when we find an invalid value.
type PairNotFound struct{}

func (e PairNotFound) Error() string {
	return "pair not found"
}

// EndOfStream occurs when we reach the end of the stream
// without finding an invalid value.
type EndOfStream struct{}

func (e EndOfStream) Error() string {
	return "end of stream"
}
