package main

import "fmt"

// No pointer, no point... 3-19

func main() {
	text := "Hello"

	fmt.Println("1) Value:", text)

	addWorld(text)
	fmt.Println("3) Modified Value:", text)
}

func addWorld(text string) {
	text = text + " world!"
	fmt.Println("2) Text in function:", text)
}

// // simple example of pointers 20-37
// func main() {
// 	textMain := "Hello"
// 	var pointer *string = &textMain

// 	fmt.Println("Value:", textMain) // Hello
// 	fmt.Println("Pointer:", pointer) // 0xc000010200

// 	fmt.Println("*Pointer:", *pointer) // 0xc000010200 -> Hello

// 	addWorld(pointer)
// 	fmt.Println("Modified Value:", textMain) // Hello
// 	// Modified Value: Hello world!
// }

// // Указатель нужен для того чтобы менять данные в получаемой переменной.
// func addWorld(text *string) {
// 	*text = *text + " world!"
// }

// pointers with types 40-58

// // Creature a new type called Creature with a single field called Species of type string.
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

// pointers with types 44-62

// // Creature a new type called Creature with a single field called Species of type string.
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

// pointers 64-80, we will get error, wonder why?

// // Creature a new type called Creature with a single field called Species of type string.
// type Creature struct {
// 	Species string
// }

// func main() {
// 	var creature *Creature

// 	fmt.Printf("1) %+v\n", creature)
// 	changeCreature(creature)
// 	// panic: runtime error: invalid memory address or nil pointer dereference
// 	fmt.Printf("3) %+v\n", creature)
// }

// func changeCreature(c *Creature) {
// 	c.Species = "jellyfish"
// 	fmt.Printf("2) %+v\n", c)
// }

// // modifying above program 107-131, error is gone...

// // Creature a new type called Creature with a single field called Species of type string.
// type Creature struct {
// 	Species string
// }

// func main() {
// 	var creature *Creature
// 	// = &Creature{
// 	// 	Species: "",
// 	// }

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

// // Creature a new type called Creature with a single field called Species of type string.
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

// // Creature a new type called Creature with a single field called Species of type string.
// type Creature struct {
// 	Species string
// }

// // Reset method that sets the Species field to an empty string.
// func (c Creature) Reset() {
// 	fmt.Printf("2) %+v\n", c) // shark
// 	c.Species = ""
// 	fmt.Printf("2a) %+v\n", c) // ??
// }

// func main() {
// 	var creature Creature = Creature{Species: "shark"}

// 	fmt.Printf("1) %+v\n", creature) // shark
// 	creature.Reset()
// 	fmt.Printf("3) %+v\n", creature) // shark
// }

// here we add the pointer, and get the point 146-161

// // Creature - struct
// type Creature struct {
// 	Species string
// }

// // Reset - method that resets value of the stucture.
// func (c *Creature) Reset() {
// 	fmt.Printf("*Pointer in func: %p\n", &c.Species)
// 	c.Species = ""
// }

// func main() {
// 	var creature Creature = Creature{Species: "shark"}

// 	fmt.Printf("*Pointer in main: %p\n", &creature.Species)

// 	fmt.Printf("1) %+v\n", creature) // shark
// 	creature.Reset()

// 	fmt.Printf("*Pointer in main: %p\n", &creature.Species)
// 	fmt.Printf("2) %+v\n", creature) // ??
// }

// TODO:check why pointer is different all the time for a single variable???

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
