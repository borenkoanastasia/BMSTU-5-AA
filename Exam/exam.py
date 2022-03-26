from random import *

def get_arr(n):
    a = []
    for i in range(n):
        a.append(randint(0, 100))
    return a

def sort(a, start, end):
    if (end - start < 2):
        return a
    index_op = (end-start)//2 + start 
    
    for i in range(index_op, start - 1, -1):
        if (i == index_op-1 and a[index_op] < a[i]):
            a[index_op], a[i] = a[i], a[index_op]
            index_op = i
        elif (a[index_op] < a[i]):
            a[index_op], a[index_op - 1] = a[index_op - 1], a[index_op]
            a[index_op], a[i] = a[i], a[index_op]
            index_op = index_op - 1
            
    for i in range(index_op, end+1):
        if (i == index_op+1 and a[index_op] > a[i]):
            a[index_op], a[i] = a[i], a[index_op]
            index_op = i
        elif (a[index_op] > a[i]):
            a[index_op], a[index_op + 1] = a[index_op + 1], a[index_op]
            a[index_op], a[i] = a[i], a[index_op]
            index_op = index_op + 1
    #print(a, start, index_op, end)
    #input()
    a = sort(a, start, index_op)
    a = sort(a, index_op, end)
    return a
            
TEST_COUNT = 10
ARR_SIZE = 10
for i in range(TEST_COUNT):
    ARR = get_arr(ARR_SIZE)
    print("test", ARR, end = "\t")
    RES = sort(ARR, 0, len(ARR) - 1)
    GOOD_RES = sorted(ARR)
    if (RES == GOOD_RES):
        print("pass", RES)
    else:
        print("badd", RES, GOOD_RES)

