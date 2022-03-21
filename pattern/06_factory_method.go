package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

//это порождающий паттерн проектирования, который определяет общий интерфейс для создания объектов в суперклассе,
//позволяя подклассам изменять тип создаваемых объектов.

//В Go невозможно реализовать классический вариант паттерна Фабричный метод,
//поскольу в языке отсутствуют возможности ООП, в том числе классы и наследственность.
//Несмотря на это, мы все же можем реализовать базовую версию этого паттерна — Простая фабрика.

//интерфейс продукта
type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

//сам продукт
type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getPower() int {
	return g.power
}

//продукт
type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

//продукт
type musket struct {
	gun
}

func newMusket() iGun {
	return &musket{
		gun: gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

//фабрика
func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

//пользовательский код
func main() {
	Ak47, _ := getGun("ak47")
	Musket, _ := getGun("musket")

	printDetails(Ak47)
	printDetails(Musket)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
