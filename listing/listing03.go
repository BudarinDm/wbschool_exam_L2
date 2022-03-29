//Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	//PathError - интерфейс, присваиваем значение nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)        //ппредположу что значение интерфейса равно нилу
	fmt.Println(err == nil) //но эррор не равен нил, пустой интерфейс не равен нилу
	fmt.Printf("%T", err)
	fmt.Println()
	fmt.Printf("%T", nil)
}

//Ответ:
//nil
//false
