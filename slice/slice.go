package slice

func SortBy[T any](slice []T, comprator func(a, b T) bool) {
}

// quickSortBy quck sort based on comprator
func quickSortBy[T any](slice []T, lowIndex, highIndex int, comprator func(a, b T) bool) {
	if lowIndex < highIndex {
		p := partitionAnySlice(slice, lowIndex, highIndex, comprator)
		quickSortBy(slice, lowIndex, p-1, comprator)
		quickSortBy(slice, p+1, highIndex, comprator)
	}
}

// partitionAnySlice split and slice into two parts for quick sort
func partitionAnySlice[T any](slice []T, lowIndex, highIndex int, comprator func(a, b T) bool) int {
	p := slice[highIndex]
	i := lowIndex
	for j := lowIndex; j < highIndex; j++ {
		if comprator(slice[j], p) {
			swap(slice, i, j)
			i++
		}
	}
	swap(slice, i, highIndex)
	return i
}

// swap swap two slice vale at index i and j
func swap[T any](slice []T, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
