/**
 * Created with IntelliJ IDEA.
 * Title: ${NAME}
 * Description:
 * User: xieguoqiang
 * @since: 13-8-21 下午1:21
 * @version 1.0
 */
	//这里的包都只是上一级而不是所有
	//方法名必须大写才能提供给外部调用 相当于public
	//由于调用时都是 包名.方法名 与文件名无关 所以 同一个包内方法名不能相同
package acorn

import "fmt"

// hello world
func Hello() {
	fmt.Println("--------------first method start-----------")
	fmt.Println("Hello world!")
	fmt.Println("--------------first method start-----------")
}

