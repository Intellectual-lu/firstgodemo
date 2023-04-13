package demo

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"
)

func Hello() {
	fmt.Println("hello go!", "hello world")
	fmt.Println(floor())
	fmt.Println(title())

	//trans test
	c := 'A'
	var v int = int(c) //类型转换
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(c)) //获取类型
	fmt.Println(reflect.TypeOf(floor()))
	fmt.Println(reflect.TypeOf(title()))
	fmt.Println(reflect.TypeOf(v))

	//time test
	now := time.Now()
	fmt.Println(now.Year())

	// replace test
	hello := "bcncnc"
	replacer := strings.NewReplacer("c", "a")
	fmt.Println(replacer.Replace(hello))

	//format test
	x := 1.0 / 3.0
	fmt.Printf("%%6.2f: %6.2f\n", x)

	format := fmt.Sprintf("%%6.2f: %6.3f", x)
	fmt.Println(format)

	fmt.Printf("this is a test: %v %v %v", 13, "\t", true)
	fmt.Println()
	fmt.Printf("this is a test: %#v %#v %#v", 13, "\t", true)
	fmt.Println()
}

func floor() float64 {
	return math.Floor(0.618)
}

func title() string {
	return strings.Title("hello go")
}
