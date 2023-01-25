package tmpl

func IterateRange[T any](slice []T, steps int) (subSlice [][]T) {
	totalSteps := len(slice) / steps
	remaining := len(slice) % steps
	for i := 0; i < totalSteps; i++ {
		var tmpSlice []T
		for j := 0; j < steps; j++ {
			tmpSlice = append(tmpSlice, slice[i*steps+j])
		}
		subSlice = append(subSlice, tmpSlice)
	}

	var remainingSlice []T
	for i := 0; i < remaining; i++ {
		remainingSlice = append(remainingSlice, slice[totalSteps+i])
	}

	if remainingSlice != nil {
		subSlice = append(subSlice, remainingSlice)
	}
	return subSlice
}
