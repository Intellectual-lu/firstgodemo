package demo

import (
	"errors"
	"fmt"
	"log"
)

func CalLiters() {
	total := 0.0
	need, err := calNeed(4.2, 3.0)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	total += need
	need, err = calNeed(5.2, 3.5)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	total += need
	fmt.Printf("%.2f liters,we needed", total)
	fmt.Println()
}

func calNeed(length float64, width float64) (areas float64, err error) {
	if length <= 0 || width <= 0 {
		err = errors.New("length or width can not less 0")
	}
	areas = length * width
	return areas / 10.0, err
}
