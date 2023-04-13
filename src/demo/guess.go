// guess a random number
package demo

import (
	"fmt"
	"github.com/Intellectual-lu/headfirstgo"
	"log"
	"math/rand"
	"time"
)

func GuessNum() {

	second := time.Now().Unix()
	rand.Seed(second)
	target := rand.Intn(100) + 1
	//fmt.Println(target)

	fmt.Println("i have give you a value between 1 to 100,can you guess it ?")
	var status string
	success := false
	for i := 0; i < 10; i++ {
		guess, err := keyborad.GetIntInPut()
		if err != nil {
			log.Fatal(err)
		}

		if target == guess {
			status = "bingo , you guess it !"
			fmt.Println(status)
			success = true
			break
		} else if target < guess && 10-i != 1 {
			status = "oh on, you guess high,try again"
			fmt.Println(status)
		} else if target > guess && 10-i != 1 {
			status = "oh on, you guess low,try again"
			fmt.Println(status)
		}
	}

	if !success {
		fmt.Println("you times is empty,the number is: ", target)
	}

}
