package math

const firstCheckOnLCM = 2

// LCMSlice is the LCM of all the element in the list.
// Use with caution. I'm not 100% sure yet if it's correct.
func LCMSlice(list []int64) int64 {
	var (
		lcm  int64 = 1
		keep bool  = true
	)

	for keep {
		keep = false

		for i := int64(firstCheckOnLCM); i < MaximumInt64(list); i++ {
			found := 0

			for _, v := range list {
				if v%i == 0 {
					found++
				}
			}

			if found > 1 {
				for idx, v := range list {
					if v%i == 0 {
						list[idx] = v / i
					}
				}

				lcm *= i
				keep = true

				break
			}
		}
	}

	for _, v := range list {
		lcm *= v
	}

	return lcm
}
