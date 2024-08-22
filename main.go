package main

import "fmt"

func anonymousReturnValues() int {
	var res int
	defer func() {
		res++
		fmt.Println("defer")
	}()
	return res
}

func namedReturnValues() (res int) {
	//var s int = 3
	defer func() {
		res++
		fmt.Println("defer")
	}()
	//fmt.Printf("%d\n", s)
	return
}

func c() (i int) {
	defer func() { fmt.Println("third") }()
	defer func() { fmt.Println("second") }()
	defer func() { fmt.Println("first") }()

	return 1
}

/*
Порядок выполнения:

i = 1 (возвращает 1)
"first"
"second"
"third"
 */

func main() {
	fmt.Println(c())
	fmt.Println(anonymousReturnValues())
	fmt.Println(namedReturnValues())
}

/*
func incrementNumber(num int) int {
	//num := 0
	defer fmt.Println(num) // 0
	num++
	return num
}

func main() {
	num := 0
	num = incrementNumber(num) // 1
	fmt.Println(num)           // 1
}
*/
