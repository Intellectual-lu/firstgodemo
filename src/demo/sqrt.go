// sqrt test
package demo

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func Sqrt() {
	value, err := calSqrt(4.2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the number sqrt value is : %.2f", value)
	fmt.Println()
}

func calSqrt(number float64) (value float64, err error) {
	if number < 0 {
		err = errors.New("number can not less zero")
	}
	value = math.Sqrt(number)
	return value, err
}
