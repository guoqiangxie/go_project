/**
 * Created with IntelliJ IDEA.
 * Title: ${NAME}
 * Description:
 * User: xieguoqiang
 * @since: 13-8-21 下午4:07
 * @version 1.0
 */
package acorn

import "fmt"

func ArrayProcess() {
	fmt.Println("--------------fourth method start-----------")
	var a [5]int
	fmt.Println("array a:", a)

	a[1] = 10
	a[3] = 30
	fmt.Println("assign:", a)

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("init:", b)

	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println("2d: ", c)

	d := b[2:4]
	fmt.Println("d:=b[2:4]:", d)

	e := b[:4]
	fmt.Println("d:=b[:4]:", e)

	f := b[2:]
	fmt.Println("d:=b[2:]:", f)
	fmt.Println("--------------fourth method end-----------")
}

