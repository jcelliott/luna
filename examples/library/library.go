package main

import (
	"fmt"
	"../.."
)

func Empty() {
	fmt.Println("Empty called")
}

func BasicParams(a int, b float64, c string, d bool) {
	fmt.Println("BasicParams called:", a, b, c, d)
}

func BasicRet() (int, float64, string, bool) {
	fmt.Println("BasicRet called")
	return 3, 4.2, "hello", false
}

func main() {
	l := luna.New(luna.AllLibs)

	libMembers := []luna.TableKeyValue{
		{"Empty", Empty},
		{"BasicParams", BasicParams},
		{"BasicRet", BasicRet},
	}
	l.CreateLibrary("testlib", libMembers...)

	l.LoadFile("library.lua")
}
