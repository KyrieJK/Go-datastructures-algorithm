package BinarySearch


/* 二分查找算法是针对有序数组的 */
func BinarySearch(array []int, number int) int {
	start := 0
	end := len(array) - 1
	for start <= end {
		middle := int((start + end) / 2)
		midItem := array[middle]
		if number == midItem {
			return middle
		}
		if midItem < number {
			start = middle + 1
		} else if midItem > number {
			end = middle - 1
		}
	}
	return -1
}
