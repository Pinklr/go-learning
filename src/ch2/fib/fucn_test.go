package my_test

import (
	"fmt"
	"testing"
)

func TestMyFunc(t *testing.T) {

	a := 1
	b := 1
	fmt.Println(a)
	for i := 1; i < 5; i++ {
		fmt.Println(b)
		temp := b
		b = a + b
		a = temp
	}
}
