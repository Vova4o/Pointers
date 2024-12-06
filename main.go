package main

import "fmt"

// simple example of pointers 5-22
func main() {
	textMain := "Hello"
	var pointer *string = &textMain

	fmt.Println("Value:", textMain)
	fmt.Println("Pointer:", pointer)

	fmt.Println("*Pointer:", *pointer)

	addWorld(pointer)
	fmt.Println("Modified Value:", textMain)
}

// Указатель нужен для того чтобы менять данные в получаемой переменной.
func addWorld(text *string) {
	*text = *text + " world!"
}

// // pointers with types 24-42
// type Creature struct {
// 	Species string
// }

// func main() {
// 	var creature Creature = Creature{
// 		Species: "shark",
// 	}

// 	fmt.Printf("1) %+v\n", creature)
// 	changeCreature(creature)
// 	fmt.Printf("3) %+v\n", creature)
// }

// func changeCreature(c Creature) {
// 	c.Species = "jellyfish"
// 	fmt.Printf("2) %+v\n", c)
// }

// // pointers with types 44-62
// type Creature struct {
// 	Species string
// }

// func main() {
// 	var creature Creature = Creature{
// 		Species: "shark",
// 	}

// 	fmt.Printf("1) %+v\n", creature)
// 	changeCreature(&creature)
// 	fmt.Printf("3) %+v\n", creature)
// }

// func changeCreature(c *Creature) {
// 	c.Species = "jellyfish"
// 	fmt.Printf("2) %+v\n", c)
// }

// // pointers 64-80, we will get error, wonder why?
// type Creature struct {
// 	Species string
// }

// func main() {
// 	var creature *Creature

// 	fmt.Printf("1) %+v\n", creature)
// 	changeCreature(creature)
// 	fmt.Printf("3) %+v\n", creature)
// }

// func changeCreature(c *Creature) {
// 	c.Species = "jellyfish"
// 	fmt.Printf("2) %+v\n", c)
// }

// // modifying above program 82-103, error is gone...
// type Creature struct {
// 	Species string
// }

// func main() {
// 	var creature *Creature

// 	fmt.Printf("1) %+v\n", creature)
// 	changeCreature(creature)
// 	fmt.Printf("3) %+v\n", creature)
// }

// func changeCreature(c *Creature) {
// 	if c == nil {
// 		fmt.Println("creature is nil")
// 		return
// 	}

// 	c.Species = "jellyfish"
// 	fmt.Printf("2) %+v\n", c)
// }

// // lets use structure that we created 105-127
// type Creature struct {
// 	Species string
// }

// func main() {
// 	var creature *Creature
// 	creature = &Creature{Species: "shark"}

// 	fmt.Printf("1) %+v\n", creature)
// 	changeCreature(creature)
// 	fmt.Printf("3) %+v\n", creature)
// }

// func changeCreature(c *Creature) {
// 	if c == nil {
// 		fmt.Println("creature is nil")
// 		return
// 	}

// 	c.Species = "jellyfish"
// 	fmt.Printf("2) %+v\n", c)
// }

// // no pointer, no point... 129-144
// type Creature struct {
// 	Species string
// }

// func (c Creature) Reset() {
// 	c.Species = ""
// }

// func main() {
// 	var creature Creature = Creature{Species: "shark"}

// 	fmt.Printf("1) %+v\n", creature)
// 	creature.Reset()
// 	fmt.Printf("2) %+v\n", creature)
// }

// // here we add the pointer, and get the point 146-161
// type Creature struct {
// 	Species string
// }

// func (c *Creature) Reset() {
// 	c.Species = ""
// }

// func main() {
// 	var creature Creature = Creature{Species: "shark"}

// 	fmt.Printf("1) %+v\n", creature)
// 	creature.Reset()
// 	fmt.Printf("2) %+v\n", creature)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	answer := 42

// 	fmt.Println("Pointer to answer", &answer)

// 	address := &answer
// 	fmt.Println("Value of address pointer", *address)

// 	fmt.Printf(
// 		"Value equal %d, address type equal %T, address it self %d\n",
// 		answer,
// 		address,
// 		&answer,
// 	)

// 	canada := "Canada"

// 	var home *string
// 	fmt.Printf("home is a %T\n", home)

// 	home = &canada
// 	fmt.Println(*home)

// 	var administrator *string

// 	scolese := "Christofer J. Scolese"
// 	administrator = &scolese
// 	fmt.Println(*administrator)

// 	bolden := "Charles F. Bolden"
// 	administrator = &bolden
// 	fmt.Println(*administrator)

// 	countBolden := 0
// 	countAdministrator := 0

// 	startBolden := time.Now()
// 	for i := 0; i < 100; i++ {
// 		fmt.Println(bolden)
// 		countBolden++
// 	}
// 	durationBolden := time.Since(startBolden)

// 	startAdministrator := time.Now()
// 	for i := 0; i < 100; i++ {
// 		fmt.Println(*administrator)
// 		countAdministrator++
// 	}
// 	durationAdministrator := time.Since(startAdministrator)

// 	fmt.Printf("Bolden printed %d times in %v\n", countBolden, durationBolden)
// 	fmt.Printf("Administrator printed %d times in %v\n", countAdministrator, durationAdministrator)
// }

// package main

// import "fmt"

// func modifySlice(arr *[]int) {
// 	for i := 0; i < len(*arr); i++ {
// 		(*arr)[i] = (*arr)[i] * 2
// 	}

// }

// func modSliceNoPointer(arr []int) []int {
// 	res := make([]int, len(arr))
// 	for i := 0; i < len(arr); i++ {
// 		res[i] = arr[i] * 2
// 	}

// 	return res
// }

// func modifyArray(arr *[5]int) {
// 	for i := 0; i < len(arr); i++ {
// 		arr[i] = arr[i] * 2
// 	}
// }

// func main() {
// 	a := []int{2, 3, 5, 7, 9}
//     fmt.Println("Original array:", a)

// 	modifySlice(&a)
// 	fmt.Println("Modified array:", a)

// 	res := modSliceNoPointer(a)
// 	fmt.Println("Mod second time", res)

// 	aa := [5]int{}
// 	copy(aa[:], a)
// 	modifyArray(&aa)
// 	fmt.Println("Mod third time", aa)
// }
