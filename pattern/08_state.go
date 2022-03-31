package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

//это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости от своего
//состояния. Извне создаётся впечатление, что изменился класс объекта.

//type vendingMachine struct {
//	hasItem       state
//	itemRequested state
//	hasMoney      state
//	noItem        state
//
//	currentState state
//
//	itemCount int
//	itemPrice int
//}
//
//func newVendingMachine(itemCount, itemPrice int) *vendingMachine {
//	v := &vendingMachine{
//		itemCount: itemCount,
//		itemPrice: itemPrice,
//	}
//	hasItemStateN := &hasItemState{
//		vendingMachine: v,
//	}
//	itemRequestedStateN := &itemRequestedState{
//		vendingMachine: v,
//	}
//	hasMoneyStateN := &hasMoneyState{
//		vendingMachine: v,
//	}
//	noItemStateN := &noItemState{
//		vendingMachine: v,
//	}
//
//	v.setState(hasItemStateN)
//	v.hasItem = hasItemStateN
//	v.itemRequested = itemRequestedStateN
//	v.hasMoney = hasMoneyStateN
//	v.noItem = noItemStateN
//	return v
//}
//
//func (v *vendingMachine) requestItem() error {
//	return v.currentState.requestItem()
//}
//
//func (v *vendingMachine) addItem(count int) error {
//	return v.currentState.addItem(count)
//}
//
//func (v *vendingMachine) insertMoney(money int) error {
//	return v.currentState.insertMoney(money)
//}
//
//func (v *vendingMachine) dispenseItem() error {
//	return v.currentState.dispenseItem()
//}
//
//func (v *vendingMachine) setState(s state) {
//	v.currentState = s
//}
//
//func (v *vendingMachine) incrementItemCount(count int) {
//	fmt.Printf("Adding %d items\n", count)
//	v.itemCount = v.itemCount + count
//}
//
////интерфейс состояния
//type state interface {
//	addItem(int) error
//	requestItem() error
//	insertMoney(money int) error
//	dispenseItem() error
//}
//
////интерфейс1
//type noItemState struct {
//	vendingMachine *vendingMachine
//}
//
//func (i *noItemState) requestItem() error {
//	return fmt.Errorf("Item out of stock")
//}
//
//func (i *noItemState) addItem(count int) error {
//	i.vendingMachine.incrementItemCount(count)
//	i.vendingMachine.setState(i.vendingMachine.hasItem)
//	return nil
//}
//
//func (i *noItemState) insertMoney(money int) error {
//	return fmt.Errorf("Item out of stock")
//}
//func (i *noItemState) dispenseItem() error {
//	return fmt.Errorf("Item out of stock")
//}
//
////интерфейс2
//type hasItemState struct {
//	vendingMachine *vendingMachine
//}
//
//func (i *hasItemState) requestItem() error {
//	if i.vendingMachine.itemCount == 0 {
//		i.vendingMachine.setState(i.vendingMachine.noItem)
//		return fmt.Errorf("No item present")
//	}
//	fmt.Printf("Item requestd\n")
//	i.vendingMachine.setState(i.vendingMachine.itemRequested)
//	return nil
//}
//
//func (i *hasItemState) addItem(count int) error {
//	fmt.Printf("%d items added\n", count)
//	i.vendingMachine.incrementItemCount(count)
//	return nil
//}
//
//func (i *hasItemState) insertMoney(money int) error {
//	return fmt.Errorf("Please select item first")
//}
//func (i *hasItemState) dispenseItem() error {
//	return fmt.Errorf("Please select item first")
//}
//
////интерфейс3
//type itemRequestedState struct {
//	vendingMachine *vendingMachine
//}
//
//func (i *itemRequestedState) requestItem() error {
//	return fmt.Errorf("Item already requested")
//}
//
//func (i *itemRequestedState) addItem(count int) error {
//	return fmt.Errorf("Item Dispense in progress")
//}
//
//func (i *itemRequestedState) insertMoney(money int) error {
//	if money < i.vendingMachine.itemPrice {
//		fmt.Errorf("Inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
//	}
//	fmt.Println("Money entered is ok")
//	i.vendingMachine.setState(i.vendingMachine.hasMoney)
//	return nil
//}
//func (i *itemRequestedState) dispenseItem() error {
//	return fmt.Errorf("Please insert money first")
//}
//
////интерфейс4
//type hasMoneyState struct {
//	vendingMachine *vendingMachine
//}
//
//func (i *hasMoneyState) requestItem() error {
//	return fmt.Errorf("Item dispense in progress")
//}
//
//func (i *hasMoneyState) addItem(count int) error {
//	return fmt.Errorf("Item dispense in progress")
//}
//
//func (i *hasMoneyState) insertMoney(money int) error {
//	return fmt.Errorf("Item out of stock")
//}
//func (i *hasMoneyState) dispenseItem() error {
//	fmt.Println("Dispensing Item")
//	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
//	if i.vendingMachine.itemCount == 0 {
//		i.vendingMachine.setState(i.vendingMachine.noItem)
//	} else {
//		i.vendingMachine.setState(i.vendingMachine.hasItem)
//	}
//	return nil
//}
//
////пользовательский код
//func main() {
//	VendingMachine := newVendingMachine(1, 10)
//
//	err := VendingMachine.requestItem()
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//
//	err = VendingMachine.insertMoney(10)
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//
//	err = VendingMachine.dispenseItem()
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//
//	fmt.Println()
//
//	err = VendingMachine.addItem(2)
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//
//	fmt.Println()
//
//	err = VendingMachine.requestItem()
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//
//	err = VendingMachine.insertMoney(10)
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//
//	err = VendingMachine.dispenseItem()
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//}

// MobileAlertStater обеспечивает общий интерфейс для различных состояний.
type MobileAlertStater interface {
	Alert() string
}

// MobileAlert реализует предупреждение в зависимости от его состояния.
type MobileAlert struct {
	state MobileAlertStater
}

// Alert возвращает строку предупреждения
func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

// SetState изменяет состояние
func (a *MobileAlert) SetState(state MobileAlertStater) {
	a.state = state
}

// NewMobileAlert конструктор MobileAlert.
func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MobileAlertVibration{}}
}

// MobileAlertVibration реализует вибро аллерт
type MobileAlertVibration struct {
}

// Alert возвращает строку предупреждения
func (a *MobileAlertVibration) Alert() string {
	return "Vrrr... Brrr... Vrrr..."
}

// MobileAlertSong реализует звуковой сигнал
type MobileAlertSong struct {
}

// Alert возвращает строку предупреждения
func (a *MobileAlertSong) Alert() string {
	return "Белые розы, Белые розы. Беззащитны шипы..."
}

// StatePattern выводит пример использования состояния
func StatePattern() {

	mobile := NewMobileAlert()

	result := mobile.Alert()
	result += mobile.Alert()

	mobile.SetState(&MobileAlertSong{})

	result += mobile.Alert()
	fmt.Println(result)
}
