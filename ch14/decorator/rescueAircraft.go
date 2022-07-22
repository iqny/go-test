package decorator

import "fmt"

type RescueAircraft struct {
	Aircraft
}
func (r *RescueAircraft) fly() {
	r.Aircraft.fly()
	fmt.Println("增加救援功能")
}
