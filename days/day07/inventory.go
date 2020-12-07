package day07

// Bag is a bag with color and color tone.
type Bag struct {
	Tone  string
	Color string
	Name  string
}

// ContentStatement is a bag and how many of them.
type ContentStatement struct {
	Bag   Bag
	Count uint64
}

type inventory struct {
	storage map[string][]ContentStatement
}

func (p *inventory) Store(name string, path []ContentStatement) {
	p.storage[name] = path
}

func (p *inventory) Has(name string) bool {
	_, found := p.storage[name]

	return found
}

func (p *inventory) Get(name string) []ContentStatement {
	return p.storage[name]
}

func newInventory() inventory {
	return inventory{
		storage: map[string][]ContentStatement{},
	}
}
