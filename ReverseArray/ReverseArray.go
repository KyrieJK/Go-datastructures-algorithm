package ReverseArray

/* 数组翻转 */
func ReverseArray(a []int){
	i:=0
	u:=len(a)-1

	for i<u{
		a[i],a[u] = a[u],a[i]
		i++
		u--
	}
}