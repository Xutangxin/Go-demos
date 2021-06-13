package main

import (
	"errors"
	"fmt"
)

func main() {
	//有错误
	result, error := getResult(12, 0)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(result)
	}

	//无错误
	result2, error2 := getResult(12, 22)
	if error2 != nil {
		fmt.Println(error2)
	} else {
		fmt.Println(result2)
	}
}

//多返回值，返回值第二个为异常
func getResult(num1 float64, num2 float64) (float64, error) {
	if num2 == 0 {
		//使用errors.New()来创建一个异常
		return 0, errors.New("the divisor can not be 0")
	} else {
		return num1 / num2, nil
	}
}
