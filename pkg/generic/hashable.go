package generic

// Hashable object. The hash can be used for cache.
type Hashable interface {
	Hash() interface{}
}
