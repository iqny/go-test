package decorator

import "fmt"

type WeaponAircraft struct {
	Aircraft
}

func (w *WeaponAircraft) fly() {
	w.Aircraft.fly()
	fmt.Println("增加武装功能")
}
