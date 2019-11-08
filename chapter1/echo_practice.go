package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	//fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("For Practice！")
	fmt.Println("Q1: ", os.Args[0])
	fmt.Println("Q2: ", strings.Join(os.Args[1:], "\n"))
	t1 := time.Now()
	echo1()
	t2 := time.Now()
	echo3()
	t3 := time.Now()
	fmt.Println("Q3: for_old ->", t2.Sub(t1), ", for_join ->", t3.Sub(t2))
}
