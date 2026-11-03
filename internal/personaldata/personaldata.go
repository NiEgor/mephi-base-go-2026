package personaldata

import "fmt"

// Personal хранит персональные данные пользователя.
type Personal struct {
	Name   string  // имя пользователя
	Weight float64 // вес пользователя, кг
	Height float64 // рост пользователя, м
}

// Print выводит персональные данные на экран.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
