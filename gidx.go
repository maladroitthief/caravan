package caravan

type GIDX struct {
	generation int
	index      int
}

type allocatorEntry struct {
	isLive     bool
	generation int
}

type GIDXAllocator struct {
	entries []allocatorEntry
	free    []int
}

func NewGIDXAllocator() *GIDXAllocator {
	return &GIDXAllocator{
		entries: []allocatorEntry{},
		free:    []int{},
	}
}

func (g *GIDXAllocator) Allocate() GIDX {
	if len(g.free) <= 0 {
		g.entries = append(
			g.entries,
			allocatorEntry{isLive: true, generation: 1},
		)

		return GIDX{
			index:      len(g.entries) - 1,
			generation: 1,
		}
	}

	n := len(g.free) - 1
	index := g.free[n]
	g.free = g.free[:n]
	g.entries[index].generation++
	g.entries[index].isLive = true

	return GIDX{
		index:      index,
		generation: g.entries[index].generation,
	}
}

func (g *GIDXAllocator) Deallocate(i GIDX) bool {
	if g.IsLive(i) == true {
		g.entries[i.index].isLive = false
		g.free = append(g.free, i.index)

		return true
	}

	return false
}

func (g *GIDXAllocator) IsLive(i GIDX) bool {
	if i.index >= len(g.entries) {
		return false
	}

	if g.entries[i.index].generation != i.generation {
		return false
	}

	if g.entries[i.index].isLive == false {
		return false
	}

	return true
}

type GIDXArrayEntry[T any] struct {
	value      T
	generation int
}

type GIDXArray[T any] []GIDXArrayEntry[T]

func NewGIDXArray[T any]() GIDXArray[T] {
	return GIDXArray[T]{}
}

func (g GIDXArray[T]) Set(index GIDX, value T) GIDXArray[T] {

	for len(g) <= index.index {
		g = append(g, GIDXArrayEntry[T]{generation: -1})
	}

	g[index.index] = GIDXArrayEntry[T]{
		value:      value,
		generation: index.generation,
	}

	return g
}

func (g GIDXArray[T]) Remove(index GIDX) GIDXArray[T] {

	if index.index < len(g) {
		g[index.index].generation = -1
	}

	return g
}

func (g GIDXArray[T]) Get(index GIDX) T {
	var defaultValue T
	if index.index >= len(g) {
		return defaultValue
	}

	entry := g[index.index]
	if entry.generation != index.generation {
		return defaultValue
	}

	return entry.value
}

func (g GIDXArray[T]) GetAllIndices(a *GIDXAllocator) []GIDX {

	result := []GIDX{}

	for i, entry := range g {
		if entry.generation <= 0 {
			continue
		}

		index := GIDX{index: i, generation: entry.generation}
		if a.IsLive(index) {
			result = append(result, index)
		}
	}

	return result
}

func (g GIDXArray[T]) GetAll(a *GIDXAllocator) []T {
	result := []T{}

	for i, entry := range g {
		if entry.generation <= 0 {
			continue
		}

		index := GIDX{index: i, generation: entry.generation}
		if a.IsLive(index) {
			result = append(result, entry.value)
		}
	}

	return result
}

func (g GIDXArray[T]) First(a *GIDXAllocator) (GIDX, T) {

	var defaultValue T
	for i, entry := range g {
		if entry.generation <= 0 {
			continue
		}

		index := GIDX{index: i, generation: entry.generation}
		if a.IsLive(index) {
			return index, entry.value
		}
	}

	return GIDX{}, defaultValue
}
