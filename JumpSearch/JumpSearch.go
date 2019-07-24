package JumpSearch

import "math"

func JumpSearch(array []int, number int) int {
	jumpStep := int(math.Floor(math.Sqrt(float64(len(array)))))
	minIndex := 0
	maxIndex := jumpStep
	for array[maxIndex] <= number {
		minIndex += jumpStep
		maxIndex = minIndex + jumpStep
		if maxIndex >= len(array) {
			maxIndex = len(array)
			minIndex = maxIndex - jumpStep
			break
		}
	}

	for i := minIndex; i < maxIndex; i++ {
		if array[i] == number{
			return i
		}
	}
	return -1
}
