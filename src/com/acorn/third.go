/**
 * Created with IntelliJ IDEA.
 * Title: ${NAME}
 * Description:
 * User: xieguoqiang
 * @since: 13-8-21 下午3:39
 * @version 1.0
 */
package acorn

import (
	"fmt"
	"time"
)

const E string = "hello world"  //如果需要提供给外部调用 必须大写

//变量与常量
func InitVar() {
	//声明初始化一个变量 类型放在后面
	var  x int = 100
	var str string = "hello world"
	//声明初始化多个变量
	var  i, j, k int = 1, 2, 3

	//不用指明类型，通过初始化值来推导
	var b = true //bool型

	a := 100
	c := "弱类型"
	d := time.Now()

	fmt.Println("--------------third method start-----------")
	fmt.Println(x)
	fmt.Println(str)
	fmt.Println(i+j+k)
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(E)
	fmt.Println("--------------third method end-----------")
}

