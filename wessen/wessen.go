// You can edit this code!
// Click here and start typing.
package wessen

import "fmt"

func DeleteAtIdx(sl []int, idx int) []int {

	res := make([]int, len(sl)-1)
	for I := 0; I < len(sl)-1; I++ {
		if I >= idx {
			res[I] = sl[I+1]
		} else {
			res[I] = sl[I]
		}

	}
	return res
}
func Wessenmain() {
	// slice: a := {1, 2, 3, 4, 5, 6}
	// Delete element at index 3
	// a := []int{1, 2, 3, 4, 5, 6}
	a1 := make([]int, 10)
	b1 := make([]int, 0, 10)
	c1 := make([]int, 10, 10)

	fmt.Println("a1: ", a1)
	fmt.Println("b1: ", b1)
	fmt.Println("c1: ", c1)

	// res := DeleteAtIdx(a, 3)
	// fmt.Println(res)

}

type student struct {
	rollNumber int
}

func Studentmain() {
	m := map[student]string{
		student{1}: "a",
		student{2}: "b",
		student{3}: "c",
	}

	for k, v := range m {
		fmt.Printf("%v - %v\n", k.rollNumber, v)
	}
}
