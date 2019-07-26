package GnomeSort

func GnomeSort(array []int){
	itemIndex :=0
	for itemIndex < len(array){
		if itemIndex != 0 && array[itemIndex-1] > array[itemIndex]{
			array[itemIndex-1],array[itemIndex] = array[itemIndex],array[itemIndex-1]
			itemIndex--
		}else{
			itemIndex++
		}
	}
}