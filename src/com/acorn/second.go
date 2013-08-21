/**
 * Created with IntelliJ IDEA.
 * Title: ${NAME}
 * Description:
 * User: xieguoqiang
 * @since: 13-8-21 下午3:31
 * @version 1.0
 */
package acorn

import (
	"fmt"
	"math"
)

//打印包 math包方法调用
func PrintMath() {
	fmt.Println("--------------second method start-----------")
	fmt.Printf("%t\n", 1==2)
	fmt.Printf("二进制：%b\n", 255)
	fmt.Printf("八进制：%o\n", 255)
	fmt.Printf("十六进制：%X\n", 255)
	fmt.Printf("十进制：%d\n", 255)
	fmt.Printf("浮点数：%f\n", math.Pi)
	fmt.Printf("字符串：%s\n", "hello world")
	fmt.Println("--------------second method end-----------")

}

