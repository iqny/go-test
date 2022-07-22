package sort_string_enq

import "testing"

func TestIsSortString(t *testing.T) {
	if b:=isRegroup("cab","bac");!b{
		t.Error("fail")
	}
}
