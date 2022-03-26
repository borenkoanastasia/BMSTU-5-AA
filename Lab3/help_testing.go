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

var MAX_REPEAT int = 100;
func get_time(v_src vector_t, sort(func(vector_t))) float64{
    var v vector_t
    var res float64 = 0
    for i:=0;i<int(MAX_REPEAT);i++{
		v = vector_copy(v_src)
		cputime1 := C.clock()
                sort(v)
		cputime2 := C.clock()
		res += float64(cputime2-cputime1)
    }
    return ((res)/float64(MAX_REPEAT))
}

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
