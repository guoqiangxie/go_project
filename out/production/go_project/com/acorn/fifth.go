/**
 * Created with IntelliJ IDEA.
 * Title: ${NAME}
 * Description:
 * User: xieguoqiang
 * @since: 13-8-21 下午4:27
 * @version 1.0
 */
package acorn

import "fmt"

func IfProcess(x int) {
	if x==1 {
		fmt.Println("a")
	} else if x==2 {
		fmt.Println("b")
	} else {
		fmt.Println("c")
	}

}

func SwithProcess(i int) {
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4,5,6:
		fmt.Println("four or five or six")
	default:
		fmt.Println("invalid value!")
	}
}

