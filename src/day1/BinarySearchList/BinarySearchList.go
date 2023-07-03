package binarysearchlist

func BinarySearch(haystack []int, needle int) bool {
	low := 0
	high := len(haystack)

	for high > low {
		pivot := low + (high-low)/2
		if haystack[pivot] == needle {
			return true
		} else if haystack[pivot] > needle {
			high = pivot
		} else {
			low = pivot + 1
		}
	}

	return false
}
