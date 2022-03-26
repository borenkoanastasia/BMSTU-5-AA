package main

import "fmt"

func main(){
	var m matrix_t
	input_matrix(&m);
	print_matrix(m)
	fmt.Println()

	fmt.Println("get: ", get_factorial(4), "expected: ", 1*2*3*4)
	fmt.Println("get: ", get_factorial(0), "expected: ", 1)

	//var pn list_t
	//pn.head = nil
	//fmt.Println("Answer: ", all_search(m, &pn), pn)
	//pn.print()
}