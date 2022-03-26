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
import "math/rand"

// Общая функция замера времени

var MAX_REPEAT float64 = 500;
func get_time(m1 matrix_t, m2 matrix_t, mulmatrix(func(matrix_t, matrix_t, *bool)matrix_t)) float64{
    var rc bool
	cputime1 := C.clock()
    for i:=0;i<int(MAX_REPEAT);i++{
        mulmatrix(m1, m2, &rc)
    }
    cputime2 := C.clock()
    return (((float64)(cputime2 - cputime1))/MAX_REPEAT)
}

func RandMatrix(rows int, columns int) matrix_t {
    var b matrix_t = make_empty_matrix(rows, columns)
    for i := 0; i < b.rows;i++ {
    	for j := 0; j < b.columns;j++ {
        	b.elem[i][j] = float32(rand.Intn(100))
		}
    }
    return b
}
