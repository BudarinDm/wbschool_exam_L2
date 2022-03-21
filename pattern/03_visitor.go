package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

//это поведенческий паттерн, который позволяет добавить новую операцию для целой иерархии классов,
//не изменяя код этих классов(ну почти)

//общий для фигур интерфейс
type shape interface {
	getType() string
	accept(visitor)
}

//фигуры
//1
type square struct {
	side int
}

func (s *square) accept(v visitor) { //один метод все же придется добавить
	v.visitForSquare(s)
}

func (s *square) getType() string {
	return "Square"
}

//2
type circle struct {
	radius int
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

func (c *circle) getType() string {
	return "Circle"
}

//3
type rectangle struct {
	l int
	b int
}

func (t *rectangle) accept(v visitor) {
	v.visitForrectangle(t)
}

func (t *rectangle) getType() string {
	return "rectangle"
}

//интерфейс визитера
type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForrectangle(*rectangle)
}

//сам визитер
type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	fmt.Println("Calculating area for square")
}

func (a *areaCalculator) visitForCircle(s *circle) {
	fmt.Println("Calculating area for circle")
}
func (a *areaCalculator) visitForrectangle(s *rectangle) {
	fmt.Println("Calculating area for rectangle")
}

//другой визитер
type middleCoordinates struct {
	x int
	y int
}

func (a *middleCoordinates) visitForSquare(s *square) {
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *middleCoordinates) visitForCircle(c *circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}
func (a *middleCoordinates) visitForrectangle(t *rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}

//у пользователя
func main() {
	newSquare := &square{side: 2}
	newCircle := &circle{radius: 3}
	newRectangle := &rectangle{l: 2, b: 3}

	newAreaCalculator := &areaCalculator{}

	newSquare.accept(newAreaCalculator)
	newCircle.accept(newAreaCalculator)
	newRectangle.accept(newAreaCalculator)

	fmt.Println()
	newMiddleCoordinates := &middleCoordinates{}
	newSquare.accept(newMiddleCoordinates)
	newCircle.accept(newMiddleCoordinates)
	newRectangle.accept(newMiddleCoordinates)
}
