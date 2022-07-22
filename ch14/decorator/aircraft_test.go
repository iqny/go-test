package decorator

import "testing"

func TestWeaponAircraft(t *testing.T) {
	 w:=WeaponAircraft{Aircraft(new(Helicopter))}
	 w.fly()
}

func TestRescueAircraft(t *testing.T) {
	r:=RescueAircraft{Aircraft(new(Helicopter))}
	r.fly()
}
