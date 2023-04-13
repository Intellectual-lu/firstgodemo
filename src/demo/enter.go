// 记录用户键盘输入的值
package demo

import (
	"fmt"
	"github.com/Intellectual-lu/headfirstgo"
	"log"
)

func Grade() {
	grade, err := keyborad.GetFloatInPut()
	if err != nil {
		log.Fatal(err)
	}
	var status string

	if grade >= 60 && grade <= 100 {
		status = "you grade is: pass"
	} else if grade >= 0 && grade < 60 {
		status = "you grade is: fail"
	} else {
		status = "you grade is wrong,please check you enter"
	}
	fmt.Println(status)
}
