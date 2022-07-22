package singleton

import "testing"

func TestGetSingleInstance(t *testing.T) {
	s:=GetSingleInstance()
	s.Show()
}
