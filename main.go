package main

import (
	"project1/note"
	hello "project1/note/lib1" // 这里后面的参数是目录，前面的参数是包名如果路径和包名一致的话可以省略前面一个参数
)

func main() {
	Hello()
	note.Test()
	note.Server()
	hello.Lib()
}
