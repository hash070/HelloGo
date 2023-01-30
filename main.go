// package 表示包名，main是一个特殊的包名，表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包
package main

//import 语句用于导入依赖包
import "fmt" //fmt包实现了格式化IO（输入/输出）的函数

func main() { //main函数是程序开始执行的函数
	fmt.Print("HelloWorld") //fmt包中的Print函数用于打印字符串
}
