package bubblesort

func Sort(a []int) []int {
	for {
		didSwap := false
		for i := 0; i < len(a)-1; i++ {
			if compare(a[i], a[i+1]) == 1 {
				// swap
				swap := a[i+1]
				a[i+1] = a[i]
				a[i] = swap
				didSwap = true
			}
		}
		if !didSwap {
			break
		}
	}
	return a
}

func compare(el1, el2 int) int {
	if el1 > el2 {
		return 1
	} else if el1 < el2 {
		return -1
	}

	return 0
}
