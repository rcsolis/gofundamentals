package main

import (
	"fmt"
	"github.com/rcsolis/gofundamentals/concurrency"
	"reflect"
)

// Package variables
var Message string

// Declare variables in block
var (
	street     string
	country    string = "Mexico"
	state      string
	postalcode int = 12345
)

func primitives() {
	fmt.Println("*** PRIMITIVES")
	// Local variables
	// Only declaration
	var name string
	name = "Rafael"
	// Full declaration and initialization
	var age int = 0
	age = 35
	// Implicit declaration
	lastname := "Chavez"
	// Shadowing package variable
	var postalcode string
	postalcode = "cp 11230"
	// Assign values to variables
	Message = "Welcome"
	street = "Rio Madeira"
	state = "CDMX"
	fmt.Printf("%v %v %v %T \n", name, lastname, age, age)
	fmt.Printf("Address: %v %v %v %v \n", street, state, country, postalcode)
	// Zero values
	var var1 bool   // false for boolean types
	var var2 int    //0 for numeric types
	var var3 string // "" for string types
	fmt.Printf("%v %T \n", var1, var1)
	fmt.Printf("%v %T \n", var2, var2)
	fmt.Printf("%v %T \n", var3, var3)
	fmt.Printf("%v %T \n", var1, var1)

	// String concatenation
	var str1 string = "Hello"
	var str2 string = " Golang World!\n"
	// Strings are utf8, inmutable, and are collections of bytes
	var byt []byte = []byte(str1 + str2)
	fmt.Printf("%v char: %v %T \n", str1+str2, str1[2], str2[2])
	fmt.Printf("%v , %T \n", byt, byt)
	// Runes - utf32 characters, use single quotes
	var run rune = '\n'
	fmt.Printf("%v %T \n", run, run)
}

func constants() {
	fmt.Println("*** CONSTANTS")
	// Constants are determined at compile time and can be shadowing
	// Typed constants
	const myConstant int = 11230
	fmt.Printf("%v %T \n", myConstant, myConstant)
	// Untyped constants
	const newConstant = "Rafael \n"
	fmt.Printf("%v %T \n", newConstant, newConstant)
	// Enumerated constants
	const enumConstant = iota
	const enumConstant2 = iota
	fmt.Printf("enumerated Constants")
	fmt.Printf("%v %T \n", enumConstant, enumConstant)
	fmt.Printf("%v %T \n", enumConstant2, enumConstant2)
	// Block enumerated constants
	fmt.Println("enumerated constants with blocks")
	const (
		a = iota
		b
		c
	)
	fmt.Println("enum from 0")
	fmt.Printf("%v %T \n", a, a)
	fmt.Printf("%v %T \n", b, b)
	fmt.Printf("%v %T \n", c, c)
	// Enumeration expression, iota reset in each block
	const (
		_ = iota + 2
		d
		e
		f
	)
	fmt.Println("new enum and start from 3 (reset iota per block)")
	fmt.Printf("%v %T \n", d, d)
	fmt.Printf("%v %T \n", e, e)
	fmt.Printf("%v %T \n", f, f)

}

func arraysSlices() {
	fmt.Println("*** Arrays and Slices")
	// Arrays use copy when assign, need to use pointers
	//Array with zero values
	var grades [3]float32
	grades[1] = 92.3
	fmt.Printf("Grades, %v \n", grades)
	// Array static size with initialization
	var subjects [3]string = [3]string{"Physics", "Math", "Science"}
	fmt.Printf("Subjects, %v \n", subjects)
	// Implicit array with initialization
	teachers := [2]string{"Rafael", "Pamela"}
	students := [...]string{"Sofia", "Emiliano"}
	fmt.Printf("Teachers, %v \n", teachers)
	fmt.Printf("Students, %v \n", students)
	// Number of elements
	fmt.Printf("#students: %v \n", len(students))
	// Slices (automatically use references/pointers)
	// no specify the size of the list
	careers := []string{"ITC", "LAF", "LATI"}
	fmt.Printf("Careers: %v \n", careers)
	fmt.Printf("#careers %v, Capacity:%v \n", len(careers), cap(careers))
	// Slice operations like python (apply to arrays too)
	campus := []string{"CSF", "CCM", "CEM", "MTY", "CGD"}
	cityCampus := campus[:3]
	foreingCampus := campus[3:]
	fmt.Printf("Campus: %v \n", campus)
	fmt.Printf("City Campus: %v \n", cityCampus)
	fmt.Printf("Foreign Campus: %v \n", foreingCampus)
	foreingCampus[1] = "GDL"
	fmt.Printf("Change Foreign Campus: %v \n", foreingCampus)
	fmt.Printf("Campus: %v \n", campus)
	// Using make function
	generations := make([]int, 3, 5)
	generations[0], generations[1], generations[2] = 2010, 2011, 2012
	fmt.Printf("Generations: %v, #%v, C.%v \n",
		generations,
		len(generations),
		cap(generations))
	// adding elements
	generations = append(generations, 2013)
	fmt.Printf("Generations: %v, #%v, C.%v \n",
		generations,
		len(generations),
		cap(generations))
	generations = append(generations, 2014, 2015)
	fmt.Printf("Generations: %v, #%v, C.%v \n",
		generations,
		len(generations),
		cap(generations))
	// Using spread operator after slice
	var newSlice = []int{2016, 2017, 2018}
	generations = append(generations, newSlice...)
	fmt.Printf("Generations: %v, #%v, C.%v \n",
		generations,
		len(generations),
		cap(generations))
	// Create a slice (reference) without some values
	// original repeat last elements to fill the spaces
	newGenerations := append(
		generations[:2],
		generations[5:]...,
	)
	fmt.Printf("New Gen: %v, #%v, C.%v \n",
		newGenerations,
		len(newGenerations),
		cap(newGenerations))
	fmt.Printf("Generations: %v, #%v, C.%v \n",
		generations,
		len(generations),
		cap(generations))
}

func mapsStructs() {
	fmt.Println("*** Maps and Structs")
	// MAPS, are key values, like dictionaries,
	// copy/assignment is by Reference.
	var studentGrades = map[string]int{
		"Rafael":   90,
		"Pamela":   91,
		"Sofia":    98,
		"Emiliano": 98,
	}
	fmt.Println(studentGrades)
	// Using Make
	statesCapital := make(map[string]string)
	statesCapital = map[string]string{
		"CDMX":  "MX",
		"KIEV":  "UC",
		"MOSCU": "RU",
	}
	statesCapital["BOGOTA"] = "CO"
	statesCapital["PARIS"] = "FR"
	fmt.Println(statesCapital)
	// Delete a element
	delete(statesCapital, "CDMX")
	fmt.Println(statesCapital["CDMX"], statesCapital)
	// Check for an element
	_, ok := statesCapital["WASHINGTON"]
	_, ok2 := statesCapital["PARIS"]
	fmt.Println("Has Washington?", ok, "Has Paris?", ok2)
	// Count elements
	fmt.Println("#items,", len(statesCapital))

	// STRUCTS, use by value not reference
	fmt.Println("STRUCTS:")

	type MetaHuman struct {
		alias  string
		name   string
		power  []string
		isHero bool
		age    int
	}

	wolverine := MetaHuman{
		alias: "Wolverine",
		name:  "Logan",
		power: []string{
			"Inmortality",
			"Regeneration",
			"Adamantium",
		},
		isHero: true,
		age:    50,
	}
	magneto := MetaHuman{
		"Magneto",
		"",
		[]string{
			"Magnetism",
			"Fly",
		},
		false,
		60,
	}
	fmt.Println("Heroe:", wolverine)
	fmt.Println("Villian:", magneto)
	fmt.Println("Hero Name:", wolverine.name)
	fmt.Println("First Villian Power:", magneto.power[0])

	// Anonymous structures
	aPet := struct {
		name string
		race string
		age  int
	}{
		name: "Pringris",
		race: "Dalmatian",
		age:  2,
	}
	fmt.Println("Pet:", aPet)
	aPet2 := aPet
	aPet2.name = "Turbo"
	aPet2.race = "Bichon"
	aPet2.age = 8
	pPet := &aPet
	pPet.age = 3
	fmt.Println("Pet 1:", aPet)
	fmt.Println("Pet 2:", aPet2)

	// Inheritance is not supported, Use Composition
	type language struct {
		name   string
		useFor string
	}
	type compiled struct {
		language
		mainFeature string
	}
	type interpreted struct {
		language
		mainFeature string
		isScripting bool
	}
	golang := compiled{
		language: language{
			name:   "Golang",
			useFor: "devops",
		},
		mainFeature: "concurrency",
	}
	python := interpreted{}
	python.useFor = "AI"
	python.name = "Python"
	python.mainFeature = "simple"
	python.isScripting = true

	fmt.Println("GO: ", golang)
	fmt.Println("Py:", python)

	// Tags for structs
	type Company struct {
		name      string `required max:"100"`
		employees int
	}
	// get the tag
	t := reflect.TypeOf(Company{})
	field, _ := t.FieldByName("name")
	fmt.Println("Company, tag for name:", field.Tag)
}

func controlFlow() {
	// If must have curly braces
	//Simple if
	b := true
	if b {
		fmt.Println("This is a simple if")
	}
	// Initializer into if
	movies := map[string]int{
		"Sample movie": 2000,
		"Second movie": 1997,
		"Third movie":  2019,
	}
	if pop, ok := movies["Third movie"]; ok {
		fmt.Println("Year of third movie:", pop)
	}
	number := 50
	guess := 50
	if guess < 1 || guess > 100 {
		fmt.Println("Guess number is not valid.")
	} else {
		fmt.Println("Guess number is valid.")
		if guess < number {
			fmt.Println("Too low")
		} else if guess > number {
			fmt.Println("Too high")
		} else {
			fmt.Println("You got it!")
		}
	}
	//Switch
	fmt.Println("SWITCH")
	var var1 int = 9
	switch var1 {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Default")
	}
	// Switch with multiple tests (case value must be unique)
	// and can use initializers
	switch i := var1 + 2; i {
	case 1, 3, 5:
		fmt.Println("Odd")
	case 2, 4, 6:
		fmt.Println("Even")
	default:
		fmt.Println("Out of range")
	}
	// Tagless sintax, could have multiple cases
	// that are evaluated to true,
	// but execute only the first case if there is not
	// an explicit fallthrough
	// this always executes the following test case code
	// even if the test condition could be false
	switch {
	case var1 <= 10:
		fmt.Println("Less than or equal to 10:", var1)
		fallthrough
	case var1 >= 20:
		fmt.Println("More than or equal to 20:", var1)
	default:
		fmt.Println("Out of range 2")
	}
	// Typed Switch
	var in interface{} = "use Break to exit the case block"
	switch in.(type) {
	case int:
		fmt.Println("Variable is int")
	case float32:
		fmt.Println("Variable is float32")
	case string:
		fmt.Println(in)
		break
		fmt.Println("Variable is string")
	default:
		fmt.Printf("Variable is unknown: %T \n", in)
	}
}

func looping() {
	// Basic fors
	for i := 0; i < 5; i++ {
		fmt.Println("Into a simple for:", i)
	}
	for i, j := 0, 1; i < 5; i, j = i+1, j+1 {
		fmt.Println("Using two variables: ", i, j)
	}
	// Omit statements
	var j int = 0
	for j < 5 {
		fmt.Println("No initializer ", j)
		j++
	}
	fmt.Println("Value after for: ", j)
	// Loop over collections
	arr := []string{"Uno", "Dos", "Tres"}
	for k, v := range arr {
		fmt.Println("Element: ", k, v)
	}

}
func panrec() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Func Error:", err)
			panic("Func: Rethrow Panic")
		}
	}()
	defer fmt.Println("Func: defer")
	fmt.Println("Func: Start")
	panic("Func: Panic!!")
	fmt.Println("Func: End")
}

func deferPanicRecover() {
	// Defer puts the execution at the end, use a LIFO
	// Evaluates the expression when its call not when its
	// executed
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Main Func Catch:", err)
		}
	}()
	defer fmt.Println("First Defer executed!")
	defer fmt.Println("This is after first defer")
	fmt.Println("*** DEFER; PANIC; RECOVER")
	a := "Message deferred"
	defer fmt.Println(a)
	a = "Other Message"
	fmt.Println("Other is print?", a)

	// Panic and Recover
	// Panic exectues after defer
	fmt.Println("Panic and Recover")
	panrec()
	fmt.Println("Flow is recovered!!!")

}

func pointers() {
	a := 35
	var b *int = &a
	// Pointers stores a memory location for the element that points
	// Pointers Arithmetics are not allowed, you need to use the unsafe package
	fmt.Println("*** POINTERS")
	fmt.Println("Pointers: a=", a, " b=", b)
	fmt.Println("Pontier address: a=", &a, " b=", b)
	// For get the value, use dereference operator *
	fmt.Println("Pontier values: a=", a, " b=", *b)
	a = 27
	fmt.Println("Pontier values: a=", a, " b=", *b)
	*b = 32
	fmt.Println("Pontier values: a=", a, " b=", *b)

	arr := [3]int{100, 101, 102}
	parr := &arr[0]
	parr2 := &arr[2]

	fmt.Println("%v %p %p \n", arr, parr, parr2)

	// Pointers to Structs
	type sample struct {
		item int
		name string
	}
	var varStruct sample
	varStruct = sample{
		item: 0,
		name: "sample 0",
	}
	var ptrVar1 *sample = &varStruct
	// Initialize to zero values
	var ptrVar2 *sample = new(sample)
	// to print the value you can use or omit deference *
	fmt.Println("Pointers to Structs:", ptrVar1, *ptrVar2)

	fmt.Println("Access to values:", ptrVar1.item, "-", ptrVar1.name)
	fmt.Println("Access to values using *:", (*ptrVar1).item, "-", (*ptrVar1).name)

}

func funcSimpleParam(msg string) {
	fmt.Println("Message", msg)
}
func funcMultiParam(msg string, idx int) {
	fmt.Println("Message ", msg, " Idx ", idx)
}

func funcSameType(msg, msg2 string, idx, idx2 int) {
	fmt.Println("Message ", msg, " Message2 ", msg2, " Idx ", idx, " Idx2 ", idx2)
}

func funcByValue(msg string) {
	fmt.Println("By value Received: ", msg)
	msg = "Changed"
	fmt.Println("Change received: ", msg)
}

func funcByRef(msg *string) {
	fmt.Println("By reference Received: ", *msg)
	*msg = "GOLANG Pointer"
	fmt.Println("Change received: ", *msg)
}

func funcUseStructs(s struct{ idx int }) {
	fmt.Println("Struct Inside func:", s)
	s.idx = 99
	fmt.Println("Struct changed inside func:", s)
}

func funcUseSlices(s []int) {
	fmt.Println("Slice inside func:", s)
	s[2] = 101
	fmt.Println("Slice changed inside func:", s)
}

func funcVariadicParam(numbers ...int) {
	fmt.Println("Variadic parameters:", numbers)
	result := 0
	for _, v := range numbers {
		result += v
	}
	fmt.Println("Sum of parameters:", result)
}

func funcReturnSimple(numbers ...int) int {
	var result int
	for _, v := range numbers {
		result += v
	}
	return result
}

func funcReturnMultiple(numbers ...int) (int, int) {
	var (
		result  int
		result2 int = 1
	)
	for _, v := range numbers {
		result += v
		result2 *= v
	}
	return result, result2
}

func funcReturnPtr(numbers ...int) *int {
	var result int = 0
	for _, v := range numbers {
		result += v
	}
	return &result
}

func funcReturnNamed(numbers ...int) (result int, result2 int) {
	result2 = 1
	for _, v := range numbers {
		result += v
		result2 *= v
	}
	return
}

func funcReturnError(a, b float32) (float32, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Error throwed from function")
	}

	return a / b, nil
}

type object struct {
	content string
}

// Value receiver
func (obj object) valueMethod(message string) {
	obj.content = message
	fmt.Println("Method execution: ", obj.content)
}

// Pointer receiver
func (obj *object) method(message string) {
	obj.content = message
	fmt.Println("Method execution: ", obj.content)
}

func functions() {
	fmt.Println("*** Functions")
	// Used to encapsulate/isolated logic
	funcSimpleParam("Simple Param")
	funcMultiParam("Multi Param", 2)
	funcSameType("Hello", "World", 35, 37)
	// Default parameters are passing by value
	var message = "Hello World"
	funcByValue(message)
	fmt.Println("Message after by value:", message)
	// Passing by reference use pointers, are more efficent
	funcByRef(&message)
	fmt.Println("Message after by reference:", message)
	sample := struct {
		idx int
	}{
		0,
	}
	funcUseStructs(sample)
	fmt.Println("Struct after func:", sample)
	// Slices and maps pass reference by default
	slice := []int{99, 100, 99}
	funcUseSlices(slice)
	fmt.Println("Slice after func:", slice)
	// Variadic parameters
	// Wrapping the parameters into a slice automatically
	// We can only have one and this need to be at the end of the parameters
	funcVariadicParam(1, 2, 3, 4, 5)

	// Returns values
	sumOf, powOf := 0, 0
	fmt.Println("Function with simple returns: ", sumOf)
	sumOf = funcReturnSimple(1, 2, 3, 4, 5)
	fmt.Println("Function with simple returns get: ", sumOf)
	sumOf, powOf = funcReturnMultiple(1, 2, 3, 4, 5)
	fmt.Println("Function with multiple returns: ", sumOf, powOf)
	// Return a pointer, this automatically promotes the variable to
	// the heap memory /shared to prevent to be destroyed when the local stacked
	// memory be free
	ptrRes := funcReturnPtr(1, 2, 3)
	fmt.Println("Pointer returned: ", ptrRes, " - ", *ptrRes)

	// Named return values
	sumOf, powOf = funcReturnNamed(1, 2, 3, 4, 5)
	fmt.Println("Function with named retrun values: ", sumOf, powOf)

	// Return errors
	div, e := funcReturnError(9.0, 0.0)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("Division = ", div)
	}

	// Anonymous functions
	func(param int) {
		fmt.Println("Hello from anonymous! ", param)
	}(sumOf)
	// Assign functions to variables
	var function func() = func() {
		fmt.Println("Printing from function assigned to variable")
	}
	function()

	// Methods, are functions assigned to types
	var instance = object{}
	fmt.Println("Create object :", instance)
	instance.valueMethod("Hello Value Methods!")
	fmt.Println("Create object after value method:", instance)
	instance.method("Hello Methods!")
	fmt.Println("Create object after reference method:", instance)
}

// Interfaces describe behaviors
// declares methods
type Writer interface {
	Write([]byte) (int, error)
}

// Use a type to implements the interface
type ConsoleWriter struct{}

// Implements methods implicity for the type
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// Define the behavior into a interface
type Programmer interface {
	Sleep() (int, error)
	Eat([]byte) (int, error)
	Code([]byte) (int, error)
	Debug([]byte) (int, error)
}

// Define the type of data
type GoDeveloper struct {
	language string
}

// Implement the interface on the type
func (pr GoDeveloper) Sleep() (n int, err error) {
	n, err = fmt.Println("I'm sleep ZZZ")
	return
}
func (pr GoDeveloper) Eat(food []byte) (n int, err error) {
	n, err = fmt.Println("I'm eating ", string(food))
	return
}
func (pr GoDeveloper) Code(thing []byte) (n int, err error) {
	n, err = fmt.Println("I'm coding in ", pr.language, " this ", string(thing))
	return
}
func (pr GoDeveloper) Debug(thing []byte) (n int, err error) {
	err = fmt.Errorf("Error throwed while debugging:  %v", string(thing))
	return
}

// Another Type
type RustDeveloper struct {
	language     string
	favoriteFood string
}

func (pr RustDeveloper) Sleep() (n int, err error) {
	n, err = fmt.Println("Sleeping time!")
	return
}
func (pr RustDeveloper) Eat(beverage []byte) (n int, err error) {
	n, err = fmt.Println("Time to eat ", pr.favoriteFood, "with ", string(beverage))
	return
}
func (pr RustDeveloper) Code(thing []byte) (n int, err error) {
	n, err = fmt.Println("I'm coding with ", pr.language, " this ", string(thing))
	return
}
func (pr RustDeveloper) Debug(thing []byte) (n int, err error) {
	err = fmt.Errorf("Error while debugging:  %v", string(thing))
	return
}

func interfaces() {
	fmt.Println("*** Interfaces")
	// We can use the interface to declare the type and
	// use any type that implements it  for polymorphism
	// If use value all method implementation must be using only value receivers
	// If use a pointer variable, methods could implement both value or pointer receivers
	var w Writer = ConsoleWriter{}
	n, err := w.Write([]byte("Hello Interfaces"))
	if err != nil {
		fmt.Println("Error in processing interface")
	}
	fmt.Println("Interface operation success:", n)

	fmt.Println("New Programmer interface \n +++GO DEVELOPER+++")
	var man Programmer = GoDeveloper{
		language: "Golang",
	}

	n, err = man.Eat([]byte("Tacos"))
	if err != nil {
		fmt.Println("Error eating", err)
	} else {
		fmt.Println("Eat time finish.", n)
	}
	n, err = man.Code([]byte("Microservice"))
	if err != nil {
		fmt.Println("Error code", err)
	} else {
		fmt.Println("Code time finish.", n)
	}
	n, err = man.Debug([]byte("REST API"))
	if err != nil {
		fmt.Println("Error debug", err)
	} else {
		fmt.Println("Debug time finish.", n)
	}
	n, err = man.Sleep()
	if err != nil {
		fmt.Println("Error sleep", err)
	} else {
		fmt.Println("Sleep time finish.", n)
	}

	fmt.Println("Using same Interface with polymorphism \n +++RUST DEVELOPER+++")

	var man2 Programmer = RustDeveloper{
		language:     "RUST",
		favoriteFood: "Tacos",
	}

	n, err = man2.Eat([]byte("Beer"))
	if err != nil {
		fmt.Println("Error eating", err)
	} else {
		fmt.Println("Eat time finish.", n)
	}
	n, err = man2.Code([]byte("Web Assembly application"))
	if err != nil {
		fmt.Println("Error code", err)
	} else {
		fmt.Println("Code time finish.", n)
	}
	n, err = man2.Debug([]byte("Smart contract"))
	if err != nil {
		fmt.Println("Error debug", err)
	} else {
		fmt.Println("Debug time finish.", n)
	}
	n, err = man2.Sleep()
	if err != nil {
		fmt.Println("Error sleep", err)
	} else {
		fmt.Println("Sleep time finish.", n)
	}

	// Everything can be casted to an empty interface
	// We use this to cast to multiple types, we need to convert to the type first
	fmt.Println("Empty interfaces ")
	var objEmpty interface{}
	objEmpty = RustDeveloper{}

	goDev, cerr := objEmpty.(GoDeveloper)
	fmt.Printf(" ConvesionVar: %v %T , %v %T \n", goDev, goDev, cerr, cerr)
	if cerr {
		fmt.Println("Error Conversion ")
	}

	code, err := goDev.Code([]byte("Tests"))
	if err != nil {
		fmt.Println("Error code", err)
	} else {
		fmt.Println("Code time finish.", code)
	}

	// Typed swietches
	fmt.Println("Use switch to get the type")
	var param interface{} = 0
	typeOfVar(param)
	param = true
	typeOfVar(param)
	param = "Hello"
	typeOfVar(param)

}

func typeOfVar(param interface{}) {
	switch param.(type) {
	case int:
		fmt.Println("Its an int")
	case string:
		fmt.Println("Its a string")
	default:
		fmt.Println("Its unknown")

	}
}

func main() {
	Message = "Hello Golang!"
	fmt.Println(Message)
	primitives()
	constants()
	arraysSlices()
	mapsStructs()
	controlFlow()
	looping()
	deferPanicRecover()
	pointers()
	functions()
	interfaces()
	concurrency.Init()
}
