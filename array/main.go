package main

import (
	"fmt"

)

func main(){
	array := [5]int{2,4,5,6,7}
	fmt.Printf("%v\n", array)

	var transportations []string
	transportations = append(transportations, "Car")
	transportations = append(transportations, "Bus")
	fmt.Printf("%v\n", transportations)

	fmt.Printf("%v\n", transportations[1])
	
	transportations = append(transportations, "Aeroplane", "Jet")
	transportations = append(transportations, "Rocket")

	landTransportations :=  transportations[:2]
	airTransportations := transportations[2:]

	fmt.Println(transportations)
	fmt.Println(len(transportations), cap(transportations))
	fmt.Println(landTransportations)
	fmt.Println(len(landTransportations), cap(landTransportations))
	fmt.Println(airTransportations)
	fmt.Println(len(airTransportations), cap(airTransportations))
	
	var all []string
	all =  append( all, airTransportations...)
	all =  append( all, landTransportations...)
	fmt.Println(all)

}