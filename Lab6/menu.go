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
import "fmt"

func line(length int) string{
	var str string = ""
	for i:=0;i<length;i++{
		str = str + "_"
	}
	return str
}
func title(length int) string{
	var str string = ""
	for i:=0;i<(length - 4)/2;i++{
		str = str + " "
	}
	str = str + "MENU"
	for i:=0;i<(length - 4)/2;i++{
		str = str + " "
	}
	return str
}

func manual_testing(){

	var min_weight float64
	var m float_matrix_t
	input_float_matrix(&m)
	m.print()
	if (len(m) < 2){
		print("Нет пути\n")
		return 
	}
	if (len(m) == 2){
		print("Путь", m[0][1], "\n")
		return 
	}
	cputime1 := C.clock()
	rout:= search_all(m, &min_weight)
	cputime2 := C.clock()

	fmt.Println("Полный перебор: ", min_weight, rout, "время:", cputime2 - cputime1)
//	fmt.Println("Время (nsec):", get_time(m2, search_all))

	var colony colony_t;
	var l float64
	colony.generate_colony(m, 3.0, 7.0, 20.0, 0.6, 10)

	cputime3 := C.clock()
	var ans =ant_alg(colony, &l)
	cputime4 := C.clock()
	fmt.Println("\nМуравьиный алгоритм:",l, ans,  "время:", cputime4 - cputime3)
//	fmt.Println("Время (nsec):", get_time(m2, AllAntsAlg))

}

func menu(){
	var length int = 111
	var choise string
    for ;true;{
		fmt.Println(line(length))
		fmt.Println()
		fmt.Println(title(length))
		fmt.Println(line(length))

		fmt.Println("\t1.Протестировать")
		fmt.Println("\t2.Общее тестирование")
		fmt.Println("\tИначе - выход")

		choise = input_string()

		fmt.Println(line(length))

		if	(choise == "1"){
			manual_testing()
		} else if	(choise == "2"){
			autotesting();
		} else if	(choise == "3"){
			graph_testing()
			break
		} else if	(choise == "4"){
			autotesting_parameters()
			break
		} else {
			break
		}
    }
}
