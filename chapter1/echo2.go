package main

import (
	"fmt"
	"os"
)
// range 产生一对值：索引即索引处的元素值
// 空标识符'_'可用于任何语法需要变量名但程序逻辑不需要的时候
// 短变量声明 s := ""（只能用于函数内部），等价于var s string/ var s = ""/ var s string = ""
func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
