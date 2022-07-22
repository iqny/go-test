package append

import (
	"fmt"
	"testing"
)

func TestAppendInt(t *testing.T) {
	z:=appendInt([]int{1,2,3,4,5,6},7,8,9,10,11,12)
	fmt.Printf("%v\n",z)
}
