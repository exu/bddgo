package bddgo

func Collect(arr []int) map[string]int {
	out := make(map[string]int)
	min, max, sum, count := arr[0], arr[0], 0, len(arr)

	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}

		sum += v
	}

	out["min"] = min
	out["len"] = count
	out["avg"] = sum / count
	out["max"] = max
	out["sum"] = sum

	return out
}
