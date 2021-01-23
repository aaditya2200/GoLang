package temp

import (
	"fmt"
	"log"
	"os"
	"time"
)


func swap(a *int, b *int){
	temp := *a
	*a = *b
	*b = temp
}

func partition(numbers []int, l int, h int) int{
	pivot := numbers[h]
	i := l-1
	for j := l; j <= h-1; j++ {
		if numbers[j] < pivot {
			i++
			swap(&numbers[i], &numbers[j])
		}
	}
	swap(&numbers[i+1], &numbers[h])
	return i+1
}

func QuickSort(numbers []int, l int, h int) {
	if l < h {
		pi := partition(numbers,l,h)

		QuickSort(numbers,l, pi -1)
		QuickSort(numbers, pi+1, h)
	}
}

func main(){
	//Helper Variables
	i:= 0
	//temp := 0
	var n int
	//n := 10000
	//rand.Seed(time.Now().Unix())
	//User input
	fmt.Println("Enter the size of the array: ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Found an error!")
		os.Exit(1)
	}
	var numbers = make([]int , n)
	//numbers = rand.Perm(n)
	fmt.Println("Enter the array elements: ");
	for ; i < n; i++ {
		_, err = fmt.Scanf("%d", &numbers[i])
		if err != nil {
			fmt.Println("Found an error!")
			os.Exit(1)
		}
	}
	QuickSort(numbers,0,n-1)
	fmt.Println("The sorted array is as follows: ")
	for i = 0; i < n; i++ {
		fmt.Print(numbers[i], " ")
	}
	//timeTrack(time.Now(), "Quick Sort")
}
