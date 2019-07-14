package main

import "fmt"

func sliceTest( slice *[]int){

	*slice=append(*slice,2)
}

func main(){

	a := []int{1}

	a = append(a,3)
	sliceTest(&a)

	fmt.Println(a)
}
