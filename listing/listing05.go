//Что выведет программа? Объяснить вывод программы.
package main

import "fmt"

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test1() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test1()
	fmt.Printf("%T", err) //*main.customError
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

//Ответ: error
