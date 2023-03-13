package _1_start

import (
	"fmt"
	"os"
	"strings"
)

func printArgs() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("s1参数:" + s)

}

func printArgs2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("s2参数:" + s)
}

func printArgs3() {
	fmt.Println("s3参数:" + strings.Join(os.Args[1:], " "))
}

func printArgs4() {
	fmt.Println("s4参数:" + strings.Join(os.Args[0:1], " "))
}

func printArgs5() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i)
		fmt.Println(os.Args[i:])
	}
}
