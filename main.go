package main

import "fmt"

const secondsInHour = 60 * 60

func main() {
	fmt.Println("Hello Go World!!!")
	distance := 60.8 // distance in kilometers
	fmt.Printf("The distance in miles is %f \n", distance*0.621371)
}
