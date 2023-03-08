package operator_test

import "testing"

const (
	Readable  = 1 << iota //0b001
	Writable              //0b010
	Executabe             //0b100
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	//b := [...]int{1, 2, 3, 4, 5}
	c := [...]int{1, 2, 3, 5}
	t.Log(a == b)
	t.Log(a == c)
}

func TestBitClear(t *testing.T) {
	a := 7
	a = a &^ Readable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executabe == Executabe)

}
