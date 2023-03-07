package main // 包 ， 表明所在的模块

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) > 1 {
		fmt.Println("hello world " + os.Args[1])
	}

}
