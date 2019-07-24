package LinearSearch

/* 线性查找 */
func LinearSearch(array []int,number int) int{
	for index,value := range array{
		if value == number{
			return index
		}
	}
	return -1
}