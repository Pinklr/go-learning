package constant_test

import "testing"

const (
	a = 100
	b
	c
	d = "haha"
	e
)

func TestCon(t *testing.T) {
	t.Log(a, b, c, d, e)

}
