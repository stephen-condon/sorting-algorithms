package selectionsort

func Sort(a []int) []int {
	var minimumIndex int
	var sortedSlice = make([]int, len(a))
	for i := range a {
		for j := 0; j < len(a); j++ {
			if j == 0 || a[minimumIndex] > a[j] {
				minimumIndex = j
			}
		}
		sortedSlice[i] = a[minimumIndex]
		if minimumIndex == 0 {
			a = a[1:]
		} else if minimumIndex == len(a)-1 {
			a = a[:len(a)-1]
		} else {
			a = append(a[:minimumIndex], a[minimumIndex+1:]...)
		}
	}

	return sortedSlice
}
