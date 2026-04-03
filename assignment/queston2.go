package assignment

import "fmt"

func PrintEvenNumbers() {
	countingNumbers := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	evenNumbers := []int{}

	sum := 0

	for i := range countingNumbers { // this for loop would go in and out of the slice, by taking the values indiivually
		countingNumbers[i] *= 2 // and then muliply the values by two

	}
	fmt.Println("Doubled: ", countingNumbers)

	for _, num := range countingNumbers {

		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
			fmt.Printf("Even numbers: %v\n", evenNumbers)
		}
	}
	fmt.Println(" ")

	for _, value := range evenNumbers {
		sum += value
	}

	fmt.Println("Sum: ", sum)

}
