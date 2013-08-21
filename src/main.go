/**
 * Created with IntelliJ IDEA.
 * Title: ${NAME}
 * Description:
 * User: xieguoqiang
 * @since: 13-8-18 下午10:09
 * @version 1.0
 */
package main

import (
	"com/acorn"
	"fmt"
)

func main() {
	acorn.Hello()
	acorn.PrintMath()
	acorn.InitVar()
	fmt.Println(acorn.E) //调用其他包的常量
	acorn.ArrayProcess()
}
