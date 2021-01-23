package temp

import (
	"fmt"
	"os"
)


func main(){
	i,j := 0,0
	temp := 0
	gap := 0
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

	//performing sort
	for gap = n/2; gap > 0; gap /= 2 {

		for i = gap; i < n; i++ {
			temp = numbers[i]

			for j = i; j >= gap && numbers[j-gap]  > temp; j -= gap {
				numbers[j] = numbers[j - gap]
				}
			numbers[j] = temp

		}

	}
	fmt.Println("The sorted array is as follows: ")
	for i = 0; i < n; i++ {
		fmt.Print(numbers[i], " ")
	}
	//timeTrack(time.Now(), "Shell sort")
}