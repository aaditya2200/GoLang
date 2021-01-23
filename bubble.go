package temp

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"temp/sockets"
)
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("\n%s took %s", name, elapsed)
}

//Exit Code 1: Scanf Error
func main(){
	//Helper Variables
	i,j := 0,0
	temp := 0
	rand.Seed(time.Now().Unix())
	//var n int
	n := 10000
	//User input
	//fmt.Println("Enter the size of the array: ")
	//_, err := fmt.Scanf("%d", &n)
	//if err != nil {
	//	fmt.Println("Found an error!")
	//	os.Exit(1)
	//}
	var numbers =  make([]int, n)
	numbers = rand.Perm(n)
	//fmt.Println("Enter the array elements: ");
	//for ; i < n; i++ {
	//	_, err = fmt.Scanf("%d", &numbers[i])
	//	if err != nil {
	//		fmt.Println("Found an error!")
	//		os.Exit(1)
	//	}
	//}
	//perform bubble sort
	for i = 0; i < n; i++ {
		for j = 0; j < n -i -1; j++ {
			if numbers[j] > numbers[j+1] {
				temp = numbers[j];
				numbers[j] = numbers[j+1]
				numbers[j+1] = temp
			}
		}
	}
	//done sorting
	//print
	fmt.Println("The sorted array is as follows: ")
	for i = 0; i < n; i++ {
		fmt.Print(numbers[i], " ")
	}
	 defer timeTrack(time.Now(), "Bubble Sort")
}
