package main

import "fmt"

func main() {
	st := stack{}

	for i := 0; i < 10; i++ {
		st.push(i)
	}

	fmt.Println(st)

}

type stack struct {
	values []int
}

// method not function
func (st *stack) push(value int) {
	st.values = append(st.values, value)

}
func (st *stack) pop() int {
	a := st.values [len(st.values) -1]
	st.values = st.values[:len(st.values)-1]

	return a

}
func (st *stack) peek() int {
	a := st.values [len(st.values) -1]

	return a

}


