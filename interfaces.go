package main

import "fmt"



type Door struct {
		width		float32
		height	float32
		colour	string
		name		string
}
type Window struct {
	width		float32
	height	float32
	colour	string
	name		string
}
type Openings interface {
	Open()
	Close()
	Lock()
}

func (d Door) Open() {
	fmt.Printf("%s is opening \n", d.name)
}

func (d Door) Close() {
	fmt.Printf("%s is closing", d.name)
}
func (d Door) Lock() {
	fmt.Printf("%s is locking", d.name)
}

func (d Window) Open() {
	fmt.Printf("%s is opening \n", d.name)
}

func (d Window) Close() {
	fmt.Printf("%s is closing", d.name)
}
func (d Window) Lock() {
	fmt.Printf("%s is locking", d.name)
}

func Interface() {
	var window Openings = Window{
		width: 4.2,
		height: 21.1,
		name: "window",
	}

	var door Openings = Door{
		name: "door",
	}

	window.Open()
	door.Open()
}