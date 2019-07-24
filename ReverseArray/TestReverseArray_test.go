package ReverseArray

import(
	"math/rand"
	"testing"
	"time"
)

func TestReverseArray(t *testing.T){
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	array1:=make([]int,random.Intn(100-10)+10)
	for i:=range array1{
		array1[i] = random.Intn(100)
	}
	array2 := make([]int,len(array1))
	copy(array2,array1)
	ReverseArray(array1)
	j:=len(array2)-1
	for i:=0;i<j;i++{
		if array1[i] != array2[j]{
			t.Fail()
		}
		j--
	}
}