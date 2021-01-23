package temp

import (
	"fmt"
	"math/rand"
	"time"
)


func DeepCopy(vals []int) []int {
	tmp := make([]int, len(vals))
	copy(tmp, vals)
	return tmp
}


func MergeSort(numbers []int)  {
	if len(numbers) > 1{
		mid := len(numbers) / 2
		left := DeepCopy(numbers[0:mid])
		right := DeepCopy(numbers[mid:])
		MergeSort(left)
		MergeSort(right)
		l,r := 0,0
		for i := 0; i < len(numbers); i++ {
			lval := -1
			rval := -1
			if l < len(left) {
				lval = left[l]
			}

			if r < len(right) {
				rval = right[r]
			}

			if (lval != -1 && rval != -1 && lval < rval) || rval == -1 {
				numbers[i] = lval
				l += 1
			} else if (lval != -1 && rval != -1 && lval >= rval) || lval == -1 {
				numbers[i] = rval
				r += 1
			}
		}
	}
}

func main(){
	i := 0
	//temp := 0
	n := 10000
	rand.Seed(time.Now().Unix())
	//User input
	//fmt.Println("Enter the size of the array: ")
	//_, err := fmt.Scanf("%d", &n)
	//if err != nil {
	//	fmt.Println("Found an error!")
	//	os.Exit(1)
	//}
	var numbers = make([]int, n)
	numbers = rand.Perm(n)
	//fmt.Println("Enter the array elements: ");
	//for ; i < n; i++ {
	//	_, err = fmt.Scanf("%d", &numbers[i])
	//	if err != nil {
	//		fmt.Println("Found an error!")
	//		os.Exit(1)
	//	}
	//}
	MergeSort(numbers)
	fmt.Println("The sorted array is as follows: ")
	for i = 0; i < n; i++ {
		fmt.Print(numbers[i], " ")
	}
	timeTrack(time.Now(), "MergeSort")
}
