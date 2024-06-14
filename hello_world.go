package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hola, Go!")

	//Variables

	var myString string = "Esto es una cadena de texto"

	fmt.Println(myString)

	myString = "aqui cambio el valor de la cadena de texto"

	fmt.Println(myString)

	var myString2 = "Esto es una cadena de texto sin declarar el tipo de variable"

	fmt.Println(myString2)

	var myInt int64 = 7
	fmt.Println(myInt)

	fmt.Println(myString, myInt)
	fmt.Printf("%s %d", myString, myInt)
	fmt.Println()

	fmt.Println(reflect.TypeOf(myInt))

	var myFloat = 6.5
	fmt.Println(myFloat)
	fmt.Println(reflect.TypeOf(myFloat))

	fmt.Println(myInt + int64(myFloat))

	var myBool bool = true
	fmt.Println(myBool)

	myBool = false
	fmt.Println(myBool)

	myString3 := "Esto es una cadena de texto" // := sirve para declarar e inicializar var myString3 String = "Esto es una cadena de texto"
	fmt.Println(myString3)

	const myConst = "Esto es una constante"
	fmt.Println(myConst)

	if myInt == 10 { // && <-> ||
		fmt.Println("el valor es 10")
	} else if myInt == 7 {
		fmt.Println("El valor es 7")
	} else {
		fmt.Println("El valor no es 10")
	}

	//ESTRUCTURAS DE DATOS

	//ARRAY

	var myArray [3]int
	myArray[0] = 1
	fmt.Println(myArray)
	fmt.Println(myArray[1])

	var myArray2 [3]string
	fmt.Println(myArray2)

	// MAP

	myMap := make(map[string]int)
	myMap["Dayana"] = 25
	myMap["Juan"] = 12
	fmt.Println(myMap)
	fmt.Println(myMap["Dayana"])

	myMap2 := map[string]int{"Dayana": 25, "Juan": 12}
	fmt.Println(myMap2)

	//LIST

	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	fmt.Println(myList.Back().Value)

	//Bucles

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	//Funcion

	myFunction()
	fmt.Println(myFunction2())

	//Estructura

	type MyStruct struct {
		name string
		age  int
	}

	myStuct := MyStruct{"Dayana", 25}

	fmt.Println(myStuct)
	fmt.Println(myStuct.name)
}

func myFunction() {
	fmt.Println("mi función")
}

func myFunction2() string {
	return "mi funcion 2"
}
