package type_test

import "testing"

type Myint int32

func TestType(t *testing.T) {
	var a int32 = 1
	var b int64 = 2
	a = int32(b)
	var c Myint = Myint(a)
	t.Log(c)

}
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)

}
