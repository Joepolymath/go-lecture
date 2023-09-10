// main
package main

import (
	"flag"
	"fmt"
	"os"
)

func main()  {
	// defer WordCount("I am learning Go! am am I")
	// defer PlayGround()
	defer DIP()
	defer Ocp()
	defer Interface()
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	fmt.Println(*csvFileName)
	var x, y int = 4, 5
	fmt.Println(x, y)

	file, err := os.Open(*csvFileName)
	if(err != nil) {
		exit(fmt.Sprintf("Failed to open the csv file: %s", *csvFileName))
	} 
	fmt.Println("Opened csv file...")
	_ = file
	
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}