package main

import "fmt"

var (
	i int //

	i8   int8
	i32  int32
	i64  int64
	ui32 uint32 //unsigned integer

	f float32 //float64
	s string
	r rune

	ii interface{}
)

func main() {
	//i32 = 1<<31 - 1
	//ui32 = 1<<32 - 1
	//
	//var array = []int{5,3,3,4}
	//
	//fmt.Printf("this is int: %v", array, 123 ,123)
	//
	//
	//fmt.Println("ui",ui32)

	//slice
	//var runes []rune
	//
	//for i :=0; i<100 ;i++  {
	//	runes= append(runes, 's')
	//}
	//s :="Almas"
	//fmt.Println([]rune(s))

	var arr1 = []int{1,6,5,3,-31}


	i :=0
	cap0 :=cap(arr1)
	for i<1000000 {
		i++
		arr1 = append(arr1,5)

		if cap(arr1) != cap0{
			fmt.Printf("initial capacity: %1v\n", cap(arr1))
			fmt.Printf("initial capacity: %4.2f\n", float32(cap(arr1))/float32((cap0)))
		}
		cap0 = cap(arr1)
		}





}
