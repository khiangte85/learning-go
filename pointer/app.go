package main

import "fmt"

func main() {
	age := 39
	agePointer := &age

	fmt.Println("Age: ", *agePointer)
	fmt.Println("Age Pointer: ", agePointer)

	fmt.Println("Adult Years: ", getAdultYears(age))
	fmt.Println("Adult Years: ", getAdultYearsPtr(&age))

	fmt.Println("Adult Years: ", getAdultYearsPtr(agePointer))
	getAdultYearsPtrMutate(agePointer)
	fmt.Println("Adult Years: ", age)

}

func getAdultYears(age int) int {
	return age - 18
}

func getAdultYearsPtr(age *int) int {
	return *age - 18
}

func getAdultYearsPtrMutate(age *int) {
	*age = *age - 18
}
