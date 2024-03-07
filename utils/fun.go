package utils

type Matrix struct {
	Array []int
	Real  bool
}

func SortList(sort []int) []int {
	size := len(sort)
	for i := 0; i < size; i++ {

		for j := i + 1; j < size; j++ {

			if sort[j] < sort[i] {
				temp := sort[j]
				sort[j] = sort[i]
				sort[i] = temp
			}
		}
	}
	return sort
}
