package temp

import (
	"fmt"
	"math"
	"os"
)

func max(numbers []int) int {
	m := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if m < numbers[i] {
			m = numbers[i]
		}
	}
	return m
}

func counting(a []int, radix int, digit int) []int{
	var b = make([]int, len(a))
	var c = make([]int, radix)
	pow:= math.Pow(float64(radix),float64(digit))
	for i := 0;i < len(a); i++ {
		digits := (a[i]/int(pow))%radix
		c[digits]++
	}
	for j := 1; j < radix; j++ {
		c[j] = c[j] + c[j-1]
	}

	for m := len(a)-1; m > -1; m--{
		digits := (a[m]/int(pow))%radix
		c[digits]--
		b[c[digits]] = a[m]
	}

	return b
}

func RadixSort(a []int, radix int) []int{
	k := max(a)
	output := a
	digits := int(math.Floor(math.Log10(float64(k)) + 1))

	for i :=0; i < digits; i++ {
		output = counting(output, radix, i)
	}
	return output
}


func main(){
	//Helper Variables
	i:= 0
	//temp := 0
	var n int
	//User input
	fmt.Println("Enter the size of the array: ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Found an error!")
		os.Exit(1)
	}
	var numbers = make([]int , n)
	fmt.Println("Enter the array elements: ");
	for ; i < n; i++ {
		_, err = fmt.Scanf("%d", &numbers[i])
		if err != nil {
			fmt.Println("Found an error!")
			os.Exit(1)
		}
	}

	numbers = RadixSort(numbers,10)
	fmt.Println("The sorted array is as follows: ")
	for i = 0; i < n; i++ {
		fmt.Print(numbers[i], " ")
	}
}
