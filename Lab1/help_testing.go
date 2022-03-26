package main
/*
#include <pthread.h>
#include <time.h>
#include <stdio.h>

static long long getThreadCpuTimeNs() {
    struct timespec t;
    if (clock_gettime(CLOCK_THREAD_CPUTIME_ID, &t)) {
        perror("clock_gettime");
        return 0;
    }
    return t.tv_sec * 1000000000LL + t.tv_nsec;
}
*/
import "C"
import "math/rand"

// Общая функция замера времени

var MAX_REPEAT float64 = 200;
func get_time(s1 string, s2 string, rast(func (*string, *string, int, int)int)) float64{
	var l1, l2 int 
	l1 = len(s1) - 1
	l2 = len(s2) - 1
    cputime1 := C.clock()
    for i:=0;i<int(MAX_REPEAT);i++{
        rast(&s1, &s2, l1, l2)
    }
    cputime2 := C.clock()
    return (float64)(cputime2 - cputime1)/MAX_REPEAT
}

// Генерирование случайной строки нужной длины
const LETTERS_BYTES = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = LETTERS_BYTES[rand.Intn(len(LETTERS_BYTES))]
    }
    return string(b)
}

// Генерирование строки типа 1 нужной длины
func StringType1Bytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = 'a'
    }
    return string(b)
}
// Генерирование строки типа 2 нужной длины
func StringType2Bytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = 'b'
    }
    return string(b)
}

// Генерирование строки типа 3 нужной длины
func StringType3Bytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = 'a'
    }
	b[n-1] = 'b'
    return string(b)
}
// Генерирование строки типа 4 нужной длины
func StringType4Bytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = 'a'
    }
	b[n-2] = 'b'
    return string(b)
}

/*
func LM_manual_time(s1 string, s2 string) float64{
	var l1, l2 int 
	l1 = len(s1) - 1
	l2 = len(s2) - 1
    cputime1 := C.getThreadCpuTimeNs()
    for i:=0;i<int(MAX_REPEAT);i++{
        LMatrix(&s1, &s2, l1, l2)
    }
    cputime2 := C.getThreadCpuTimeNs()
    return (float64)(cputime2 - cputime1)
}
func LR_manual_time(s1 string, s2 string) float64{
	var l1, l2 int 
	l1 = len(s1) - 1
	l2 = len(s2) - 1
    cputime1 := C.getThreadCpuTimeNs()
    for i:=0;i<int(MAX_REPEAT);i++{
        LRecursion(&s1, &s2, l1, l2)
    }
    cputime2 := C.getThreadCpuTimeNs()
    return (float64)(cputime2 - cputime1)
}
func DLR_manual_time(s1 string, s2 string) float64{
	var l1, l2 int 
	l1 = len(s1) - 1
	l2 = len(s2) - 1
    cputime1 := C.getThreadCpuTimeNs()
    for i:=0;i<int(MAX_REPEAT);i++{
        DLRecursion(&s1, &s2, l1, l2)
    }
    cputime2 := C.getThreadCpuTimeNs()
    //fmt.Printf("CPU time = %f ns\n", float64((cputime2 - cputime1))/MAX_REPEAT)
	return (float64)(cputime2 - cputime1)
}
func LRK_manual_time(s1 string, s2 string) float64{
	var l1, l2 int 
	l1 = len(s1) - 1
	l2 = len(s2) - 1
    cputime1 := C.getThreadCpuTimeNs()

    for i:=0;i<int(MAX_REPEAT);i++{
	    var LRecursionKesh = getLRecursionKesh()
        LRecursionKesh(&s1, &s2, l1, l2)
    }
    cputime2 := C.getThreadCpuTimeNs()
    //fmt.Printf("CPU time = %f ns\n", float64((cputime2 - cputime1))/MAX_REPEAT)
	return (float64)(cputime2 - cputime1)
}
*/