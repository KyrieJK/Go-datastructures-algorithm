package MergeSort

func MergeSort(array []int,leftIndex,rightIndex int){
	if leftIndex < rightIndex{
		midIndex := int((leftIndex+rightIndex)/2)
		MergeSort(array,leftIndex,midIndex)
		MergeSort(array,midIndex+1,rightIndex)
		merge(array.leftIndex,midIndex,rightIndex)
	}
}

func merge(array []int,leftIndex,midIndex,rightIndex int){
	var tempArray1,tempArray2 []int
	for i:=leftIndex;i<=midIndex;i++{
		tempArray1 = append(tempArray1,array[i])
	}
	for i:=midIndex+1;i<=rightIndex;i++{
		tempArray2 = append(tempArray2,array[i])
	}

	arrayIndex := leftIndex
	tempArray1Index:=0
	tempArray2Index:=0
	for tempArray1Index<len(tempArray1) && tempArray2Index<len(tempArray2){
		if tempArray1[tempArray1Index] <= tempArray2[tempArray2Index]{
			array[arrayIndex] = tempArray1[tempArray1Index]
			tempArray1Index++
		}else {
			array[arrayIndex] = tempArray2[tempArray2Index]
			tempArray2Index++
		}
		arrayIndex++
	}
	for tempArray1Index < len(tempArray1){
		array[arrayIndex] = tempArray1[tempArray1Index]
		tempArray1Index++
		arrayIndex++
	}
	for tempArray2Index<len(tempArray2){
		array[arrayIndex] = tempArray2[tempArray2Index]
		tempArray2Index++
		arrayIndex++
	}
}