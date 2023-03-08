package condition

import "testing"

func TestCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("是4以内的偶数")
		case 1, 3:
			t.Log("是4以内的奇数")
		default:
			t.Log("This should be four")

		}
	}

}
