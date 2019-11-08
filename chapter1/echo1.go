// print command-line arguments.
package main

import (
	"fmt"
	"os"
)
// 注：i++是语句，而非表达式，因此j = i++非法，且++和--只能放在变量后
// go只有for循环一种循环语句
/*
for initialization; condition; post {
    // zero or more statements
}
for的三个部分都可以省略，如果省略initialization和post，分号也可以省略
*/
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}