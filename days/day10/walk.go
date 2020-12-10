package day10

func walk(data []int64, pathCache *cache, target int64) int64 {
	if len(data) < 1 {
		return 0
	}

	if len(data) == 1 {
		if data[0] == target {
			return 1
		}

		return 0
	}

	if pathCache.Has(data[0]) {
		return pathCache.Get(data[0])
	}

	paths := int64(0)

	for _, idx := range getPossibleNextAdapterIndexes(data) {
		paths += walk(data[idx+1:], pathCache, target)
	}

	pathCache.Store(data[0], paths)

	return paths
}

func getPossibleNextAdapterIndexes(data []int64) []int {
	possible := []int{}

	for idx, value := range data[1:] {
		if value-data[0] > adapterOffset {
			break
		}

		possible = append(possible, idx)
	}

	return possible
}
