package tmpl

func IterateRange[T any](slice []T, steps int) (subSlice [][]T) {
	totalSteps := len(slice) / steps
	for i := 0; i < totalSteps; i++ {
		var tmpSlice []T
		for j := 0; j < steps; j++ {
			tmpSlice = append(tmpSlice, slice[i*steps+j])
		}
		subSlice = append(subSlice, tmpSlice)
	}
	return subSlice
}
