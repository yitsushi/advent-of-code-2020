package slice

// ContainsInt64 returns true if the request key can be found
// in the provided list.
func ContainsInt64(list []int64, key int64) (int, bool) {
	for idx, value := range list {
		if value == key {
			return idx, true
		}
	}

	return -1, false
}

// ContainsInt returns true if the request key can be found
// in the provided list.
func ContainsInt(list []int, key int) (int, bool) {
	for idx, value := range list {
		if value == key {
			return idx, true
		}
	}

	return -1, false
}

// ContainsString returns true if the request key can be found
// in the provided list.
func ContainsString(list []string, key string) (int, bool) {
	for idx, value := range list {
		if value == key {
			return idx, true
		}
	}

	return -1, false
}
