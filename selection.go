package temp

import (
	"fmt"
	"os"
	"time"

	//"os"
	//"math/rand"
)
//Exit Code 1: Scanf Error

func main(){
	//Helper Variables
	i,j := 0,0
	index := 0
	temp := 0
	var n int

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
	//min := numbers[0]
	//perform selection sort
	for i = 0; i < n - 1; i++ {
		//min := numbers[i]
		index = i;
		for j = i + 1; j < n; j++ {
			if numbers[index] > numbers[j] {
				index = j
			}
		}
			temp = numbers[index];
			numbers[index] = numbers[i]
			numbers[i] = temp

	}
	//done sorting
	//print
	fmt.Println("The sorted array is as follows: ")
	for i = 0; i < n; i++ {
		fmt.Print(numbers[i], " ")
	}
	timeTrack(time.Now(), "Selection Sort")
}
