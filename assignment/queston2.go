package main

import "fmt"

func main() {
	ah := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	kk := []int{}

	sum := 0

	for i := range ah { // this for loop would go in and out of th slice, by taking the values indiivually
		ah[i] *= 2 // and then muliply the values by two

	}
	fmt.Println("Doubled: ", ah)
	

	for _, num := range ah {

		if num%2 == 0 {
			kk = append(kk, num)
			fmt.Printf("Even numbers: %v\n", kk)
		}
	}
	fmt.Println(" ")

	for _, value := range kk {
		sum += value
	}

	fmt.Println("Total sum: ", sum)

}
