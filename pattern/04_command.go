package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

//это поведенческий паттерн проектирования, который превращает запросы в объекты,
//позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь, логировать их,
//а также поддерживать отмену операций.

//Команда превращает операции в объекты. А объекты можно передавать, хранить и взаимозаменять внутри других объектов.

//отправитель
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

//интерфейс команды
type command interface {
	execute()
}

//конкретная команда
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

//тоже
type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

//интерфейс получателя
type device interface {
	on()
	off()
}

//получатель
type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

//клиентский код
func main() {
	newTv := &tv{}

	newOnCommand := &onCommand{
		device: newTv,
	}

	newOffCommand := &offCommand{
		device: newTv,
	}

	onButton := &button{
		command: newOnCommand,
	}
	onButton.press()

	offButton := &button{
		command: newOffCommand,
	}
	offButton.press()
}
