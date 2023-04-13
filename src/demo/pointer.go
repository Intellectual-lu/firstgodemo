// pointer rest
package demo

import (
	"fmt"
	"reflect"
)

func Pointer() {
	amount := 6
	fmt.Println(amount)
	fmt.Println(&amount)
	fmt.Println(reflect.TypeOf(&amount))
	var intPointer *int = &amount
	fmt.Println(intPointer)
	fmt.Println(*intPointer)
	*intPointer = 9
	fmt.Println(amount)

	fmt.Println("----------------------------------")

	pointer := createPointer()
	fmt.Println(pointer)
	fmt.Println(*pointer)

	fmt.Println("----------------------------------")

	value := false
	fmt.Println(value)
	changeValue(value)
	fmt.Println("changeValue: ", value)
	updateValue(&value)
	fmt.Println("updateValue: ", value)

	fmt.Println("----------------------------------")

	i := 0.2
	doubleValue(&i)
	fmt.Println(i)

	fmt.Println("----------------------------------")

	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)
}

func createPointer() *float64 {
	myFloat := 3.14
	return &myFloat
}

func changeValue(value bool) bool {
	value = true
	return value
}

func updateValue(value *bool) *bool {
	*value = true
	return value
}

func doubleValue(value *float64) *float64 {
	*value = 2 * *value
	return value
}

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}
