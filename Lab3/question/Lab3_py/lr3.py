import random
import time

def sort_bubble(a, n):
    for i in range(1, n):
        for j in range(n-i):
            if (a[j] > a[j+1]):
                a[j], a[j+1] = a[j+1], a[j]
    return a
    
def sort_bubble_test(a, n):
    change = 0
    for i in range(1, n):
        for j in range(n-i):
            if (a[j] > a[j+1]):
                a[j], a[j+1] = a[j+1], a[j]
                change+=1
    return a, change

def test():
    rep = 500
    size = 500
    for i in range(1, 20):
        arr1 = gen_sortarr(10)
        get_time(sort_bubble, arr1, rep)
        get_time(sort_bubble, arr1, rep)
    
    arr1 = gen_sortarr(size)
    arr2 = gen_reversearr(size)
    arr3 = gen_randomarr(size)
    print(arr1, arr2, arr3)
    
    tb1 = get_time(sort_bubble, arr1, rep)
    tb2 = get_time(sort_bubble, arr2, rep)
    tb3 = get_time(sort_bubble, arr3, rep)
    
    x, c1 = sort_bubble_test(arr1, len(arr1))
    x, c2 = sort_bubble_test(arr2, len(arr2))
    x, c3 = sort_bubble_test(arr3, len(arr3))
    
    print("Сорт.", tb1, c1, "Реверс.", tb2, c2, "Случ.", tb3, c3)

def sort_choise(a, n):
    for i in range(0, n-1):
        max_el = a[0]
        max_i = 0
        for j in range(n-i):
            if (a[j] > max_el):
                max_el = a[j]
                max_i = j
        a[n - 1 - i], a[max_i] = a[max_i], a[n - 1 - i]  
    return a

def sort_insert(a, n):
    for i in range(1, n):
        for j in range(i, 0, -1):
            if (a[j] < a[j-1]):
                a[j], a[j-1] = a[j-1], a[j]
            else:
                break
    return a

def get_time(alg, arr, rep):
    res = 0
    n = len(arr)
    arr1 = copy_arr(arr)
    for i in range(rep):
        start = time.process_time_ns()
        alg(arr1, n)
        end = time.process_time_ns()
        res += end-start
    return res/rep

def line(l):
    print("_"*l)

def title(l):
    line(l)
    print(" "*int(l/2-2) + "MENU" + " "*int(l/2-2))
    line(l)
    print("\t1. Manual testing")
    print("\t2. Auto testing")
    print("\tAnother choise is exit")
    print()

def gen_sortarr(l):
    arr = [] 
    for i in range(l):
        arr.append(i+1)
    return arr
        
def copy_arr(arr1):
    arr2 = [] 
    for i in arr1:
        arr2.append(i)
    return arr2
        
def gen_reversearr(l):
    arr = [] 
    for i in range(l):
        arr.append(l-i)
    return arr
        
def gen_randomarr(l):
    arr = [] 
    for i in range(l):
        arr.append(random.randint(0, l))
    return arr

def input_arr():
    arr = []
    if (input("Сгенерировать массив автоматически? (1 - да, иначе нет)") == "1"):
        l = int(input("Введите длину массива:"))
        t = int(input("Какого типа массив: 1 - сортированный, 2 - сортированный наоборот, иначе - случайный\n"))
        print(t, t == "1")
        try:
            if (int(t) == 1):
                arr = gen_sortarr(l)
            elif (int(t) == 2):
                arr = gen_reversearr(l)
            else:
                arr = gen_randomarr(l)
        except:
            arr = gen_randomarr(l)
    else:
        arr = list(map(int, input("Введите массив в строку:").split()))
    return arr

def manual_testing():
    arr = input_arr()
    arr1 = copy_arr(arr)
    for i in range(100):
        sort_bubble(arr1,len(arr1))
    arr1 = copy_arr(arr)
    print("Получен массив:", arr)
    print(arr1)
    start = time.process_time_ns()
    ans1 = sort_bubble(arr1,len(arr1))
    end = time.process_time_ns()
    print("Результат сортировки пузырьком", ans1, end-start)
    
    arr1 = copy_arr(arr)
    print(arr1)
    start = time.process_time_ns()
    ans1 = sort_insert(arr1,len(arr1))
    end = time.process_time_ns()
    print("Результат сортировки вставками", ans1, end-start)
    
    
    arr1 = copy_arr(arr)
    print(arr1)
    start = time.process_time_ns()
    ans1 = sort_choise(arr1,len(arr1))
    end = time.process_time_ns()
    print("Результат сортировки выбором", ans1, end-start)
    
def auto_testing():
    print("Таблица замеров времени")
    print("\tВремя в таблице указано в нсек.")
    rep = 1000
    l = 111
    exp_count = 21
    line(l)
    for i in range(1, 5):
        arr1 = gen_sortarr(i*5)
        get_time(sort_bubble, arr1, rep)
        get_time(sort_bubble, arr1, rep)
        
        
    print("{:5}|{:35}|{:35}|{:35}".format("Длина", "Сортировка пузырьком","Сортировка вставками","Сортировка выбором"))
    print("{:5}|{:11}|{:11}|{:11}|{:11}|{:11}|{:11}|{:11}|{:11}|{:11}".format("", "Сорт.м.","Реверс.м.","Случ.м", "Сорт.м.","Реверс.м.","Случ.м", "Сорт.м.","Реверс.м.","Случ.м"))
    for i in range(1, exp_count):
        size = i*50
        arr1 = gen_sortarr(size)
        arr2 = gen_reversearr(size)
        arr3 = gen_randomarr(size)
        tb1 = get_time(sort_bubble, arr1, rep)
        tb2 = get_time(sort_bubble, arr2, rep)
        tb3 = get_time(sort_bubble, arr3, rep)
        
        ti1 = get_time(sort_insert, arr1, rep)
        ti2 = get_time(sort_insert, arr2, rep)
        ti3 = get_time(sort_insert, arr3, rep)
        
        tc1 = get_time(sort_choise, arr1, rep)
        tc2 = get_time(sort_choise, arr2, rep)
        tc3 = get_time(sort_choise, arr3, rep)
        
        print("{:5}|{:11.4f}|{:11.4f}|{:11.4f}|{:11.4f}|{:11.4f}|{:11.4f}|{:11.4f}|{:11.4f}|{:11.4f}".format(size, tb1, tb2, tb3, ti1, ti2, ti3, tc1, tc2, tc3))
    line(l)
    

def menu():
    title_l = 111
    while True:
        title(title_l)
        choise = input("Your choise: ")
        if (choise == "1"):
            manual_testing()
        elif (choise == "2"):
            auto_testing()
        else:
            break
        line(title_l)
        
        
#menu()
test()       
        
        
