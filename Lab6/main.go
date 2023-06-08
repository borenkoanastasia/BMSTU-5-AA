package main

/*
#include <pthread.h>
#include <time.h>
#include <stdio.h>

static long long getThreadCpuTime() {
    struct timespec t;
    if (clock_gettime(CLOCK_THREAD_CPUTIME_ID, &t)) {
        perror("clock_gettime");
        return 0;
    }
    return clock();//t.tv_sec * 10+ t.tv_nsec;
}
*/
import "C"
//import "fmt"

func main(){
	menu()
	/*var m2 float_matrix_t
	var min_weight float64


	input_float_matrix(&m2)
	m2.print()

	cputime1 := C.clock()
	rout:= search_all(m2, &min_weight)
	cputime2 := C.clock()

	fmt.Println("Полный перебор: ", min_weight, rout, cputime2 - cputime1)
//	fmt.Println("Время (nsec):", get_time(m2, search_all))

	var colony colony_t;
	var l float64
	colony.generate_colony(m2, 3.0, 7.0, 20.0, 0.6, 10)

	cputime3 := C.clock()
	var ans =ant_alg(colony, &l)
	cputime4 := C.clock()
	fmt.Println("\nМуравьиный алгоритм:", ans, l,  cputime4 - cputime3)
//	fmt.Println("Время (nsec):", get_time(m2, AllAntsAlg))

	//
//	enviroment.alpha = 3.0
//	enviroment.betta = 7.0
//	enviroment.q = 20.0
//	enviroment.p = 0.6

	//print_matrix(m)
	//fmt.Println(generate_routs_arr(3))


	//fmt.Println("get: ", get_factorial(4), "expected: ", 1*2*3*4)
	//fmt.Println("get: ", get_factorial(0), "expected: ", 1)

	//var pn list_t
	//pn.head = nil
	//fmt.Println("Answer: ", all_search(m, &pn), pn)
	//pn.print()*/
}