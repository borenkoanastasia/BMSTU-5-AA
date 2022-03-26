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

// Общая функция замера времени

var MAX_REPEAT int = 10;

func get_time(d float_matrix_t, search(func(float_matrix_t, *float64)[]int)) float64{
    var res float64 = 0
    var l float64
    for i:=0;i<int(MAX_REPEAT);i++{
		cputime1 := C.clock()
        search(d, &l)
		cputime2 := C.clock()
		res += float64(cputime2-cputime1)
    }
    return ((res)/float64(MAX_REPEAT))
}
// Ввод строки
func input_string()string{
	var str1 string
    fmt.Scanf("%s\n", &str1)
	return str1
}
/*
func RandVector(size int) vector_t {
    var b vector_t = make_empty_vector(size)
    for i := 0; i < b.size;i++ {
        b.elem[i] = float32(rand.Intn(100))
    }
    return b
}
func SortVector(size int) vector_t {
    var b vector_t = make_empty_vector(size)
    for i := 0; i < b.size;i++ {
        b.elem[i] = float32(i)
    }
    return b
}
func ReverseVector(size int) vector_t {
    var b vector_t = make_empty_vector(size)
    for i := 0; i < b.size;i++ {
        b.elem[i] = float32(b.size - i)
    }
    return b
}
*/













