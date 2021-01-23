package temp

import (
	"fmt"
	"os"
)


func main(){

	i,j := 0,0
	temp := 0
	var  n int
	//n := 2500
	////User input
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

	//performing insertion sort
	for i = 1;i < n; i++ {
		temp = numbers[i]
		j = i-1
		for j >= 0 && numbers[j] > temp {
			numbers[j+1] = numbers[j]
			j--
		}
		numbers[j+1] = temp
	}

	fmt.Println("The sorted array is as follows: ")
	for i = 0; i < n; i++ {
		fmt.Print(numbers[i], " ")
	}
	//defer timeTrack(time.Now(), "Insertion Sort")
}