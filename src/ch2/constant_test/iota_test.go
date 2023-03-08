package iota_test

import "testing"

const (
	Readable  = 1 << iota //0b001
	Writable              //0b010
	Executabe             //0b100
)

func TestIota(t *testing.T) {
	t.Log(Readable | Writable | Executabe) //0b111

}
