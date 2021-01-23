package temp

import (
	"fmt"
	"log"
	"os"
	"time"
)



func swap1(a *int, b *int){
	temp := *a
	*a = *b
	*b = temp
}

func heapify(numbers []int ,n int, i int) []int {
	largest := i
	l := 2*i + 1
	r := 2*i + 2
	if l < n && numbers[l] > numbers[largest] {
		largest = l
	}
	if r < n && numbers[r] > numbers[largest] {
		largest = r
	}
	if largest != i {
		swap1(&numbers[i], &numbers[largest])
		heapify(numbers, n , largest)
	}
	return numbers
}

func HeapSort(numbers [] int, n int) []int {
	for i:=n/2-1; i >= 0; i-- {
		numbers = heapify(numbers,n,i)
	}

	for i := n-1; i > 0; i-- {
		swap1(&numbers[0], &numbers[i])
		numbers = heapify(numbers, i , 0)
	}
	return numbers
}

func main(){
	//Helper Variables
	i := 0
	//temp := 0
	var n int
	//rand.Seed(time.Now().Unix())
	//User input
	fmt.Println("Enter the size of the array: ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Found an error!")
		os.Exit(1)
	}
	var numbers = make([]int, n)
	//numbers = rand.Perm(n)
	fmt.Println("Enter the array elements: ");
	for ; i < n; i++ {
		_, err = fmt.Scanf("%d", &numbers[i])
		if err != nil {
			fmt.Println("Found an error!")
			os.Exit(1)
		}
	}

	//sort
	numbers = HeapSort(numbers,n)
	fmt.Println("The sorted array is as follows: ")
	for i = 0; i < n; i++ {
		fmt.Print(numbers[i], " ")
	}
	//timeTrack(time.Now(), "Heap sort")
}
