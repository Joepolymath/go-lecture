package main

import (
	"fmt"
	"strings"
	"time"
)

// PlayGround for
func PlayGround() {
	var x, y int = 4, 5
	fmt.Println(x, y)
	today := time.Now().Weekday()
	fmt.Println(today)
	fmt.Printf("The type is %T \n", today)
	fmt.Println(time.Saturday)

	// arrays
	 myArray := []int{2, 3, 4, 5, 6}
	myArray[2] = 12
	fmt.Println(myArray)

	mySlice := make([]int, 5)
	mySlice[4] = 24
	fmt.Println(mySlice)
	mySlice = append(mySlice, 52)
	mySlice = append(mySlice, 52)
	mySlice = append(mySlice, 52)
	mySlice = append(mySlice, 52)
	mySlice = append(mySlice, 52)
	mySlice = append(mySlice, 52)

	for i, element := range mySlice {
		fmt.Println(i, element)
	}

	// maps. Use make() when you are not initializing the map. Use direct map when you wish to initialize the map.
	myMap := make(map[string]string)
	myMap["name"] = "Joshua"
	myMap["occupation"] = "Software Engineer"
	myMap["gender"] = "Male"
	fmt.Println(myMap)
	for key, value := range myMap {
		fmt.Println(key)
		fmt.Println(value)
	}

	person1 := Person{
		name: "Joshua",
		age: 27,
		gender: "Male",
	}
	person1.SayName()
	x = 456
	person1.age++
}

func WordCount(s string) (result map[string]int) {
	result = make(map[string]int)
	var wordsArray = strings.Split(s, " ")
	for _, element := range wordsArray {
		result[element] += 1
	}
	return
}

func fibonacci() func() int {
	prev1 := 0
	prev2 := 0
	return func() int {
		sum := prev1 + prev2
		prev1 = prev2	
		if prev2 == 0 {
			prev2++
		} else {
			prev2 = sum
		}
		return sum
	}
}

type Person struct {
	name	string
	age	int
	gender	string
}

func (p Person) SayName() {
	fmt.Println("My name is", p.name)
}