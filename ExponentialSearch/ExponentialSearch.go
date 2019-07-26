package ExponentialSearch

func ExponentialSearch(array []int,number int) int{
	if array == nil || len(array) ==0{
		return -1
	}

	bound := 1
	for bound < len(array) && array[bound] < number{
		bound *= 2
	}
	return BinarySearch(array,bound/2,Min(bound+1,len(array)),number)
}

func Min(x,y int) int{
	if x<y{
		return x
	}
	return y
}

func BinarySearch(array []int,low,high int,number int) int{
	start := low
	end := high-1
	for start <= end{
		middle := int((start+end)/2)
		midItem := array[middle]
		if number == midItem{
			return middle
		}
		if midItem < number{
			start = middle+1
		}else if midItem > number{
			end = middle - 1
		}
	}
	return -1
}