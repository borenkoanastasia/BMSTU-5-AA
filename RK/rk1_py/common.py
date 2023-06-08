from threading import Thread, Lock
from time import sleep
from random import randint
import queue

import tkinter as tk

from turtle import color

# Коснстанты
STATUS_BUILDER_BRICKS_COUNT = 0
STATUS_BUILDER_WAIT = 1

STATUS_BUILDER_ACTIVE_WAIT = True
STATUS_BUILDER_NOT_WAIT = False

END_OF_WORK = True
NOT_END_OF_WORK = False

# Вспомогательные ресурсы
LATTINICA = "abcdefghijklmnopqrstuvwxyz"
KIRILLICA = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"



# Константы - параметры системы
BUILDERS_COUNT = 4
W = 10
H = 20
MIN_LEN_STRING = 8#8
MAX_LEN_STRING = 20 #20
#QUENE_LEN = 0#221

# Разделяемые ресурсы (нужны более чем одному потоку)

QUENE = 0 # make_start_quene(10, 8, 20)
BRICKS_MATRIX = []
BUILDERS_BORDERS = []
# [[Кол-во блоков, флаг ожидания], [...], ...]
BUILDERS_WORK_STATUS = []
#
FLAG_END_OF_WORK = NOT_END_OF_WORK


# Мьютексы для разделяемых ресурсов
MUTeX_QUENE = Lock()
MUTeX_BRICKS_MATRIX = Lock()
MUTeX_BUILDERS_BORDERS = Lock()

# Функции для инициализации системы

def make_rand_string(size):
    string = ""
    for i in range(size):
        string = string + LATTINICA[randint(-1, len(LATTINICA)-1)]
    return string

def get_quene_len(width, height):
    quene_len = (width + 1 + width)*(height // 2)
    if (height % 2 == 1):
        quene_len+= width
    return quene_len

def make_start_quene(size, min_len, max_len):
    q = queue.Queue()

    for i in range(size):
        new_el = make_rand_string(randint(min_len, max_len))
        q.put(new_el)
    return q

def get_bricks_height(len_start_quene, weight):
    #print(len_start_quene%(weight*2+1), len_start_quene,(weight*2+1))
    if (len_start_quene%(weight*2+1) >= weight):
        return len_start_quene//(weight*2+1)*2+1
    else:
        return len_start_quene//(weight*2+1)*2

def make_bricks_matrix(height, weight):
    matrix = []
    #print(height, weight)
    for i in range(height):
        arr = []
        border = weight
        #if (i%2 == 1):
        #    border += 1
        for j in range(border): 
            arr.append(0)
        matrix.append(arr)
    return matrix

def get_builder_weight(builders_count, weight):
    b_w = weight//builders_count
    res = []
    if (b_w*builders_count == weight):
        for i in range(builders_count):
            res.append(i*b_w)
        res.append(weight)
        return res 
    b_w += 1
    flag = True
    cur_weight = 0
    res.append(0)
    print(b_w)
    for i in range(builders_count):
        cur_weight += b_w
        res.append(cur_weight)
        if (flag and (b_w-1)*(builders_count-i-1) == weight-cur_weight):
            b_w -=1
            flag = False
    return res

def make_builders_work_status():
    bws = []
    for i in range(BUILDERS_COUNT):
        bws.append([0, False])
    return bws

def make_builders_borders(builders_count, weight):
    builders_borders = []
    builder_weight = get_builder_weight(builders_count, weight)
    #print(builder_weight, weight, builders_count)
    for i in range(1, len(builder_weight)):
    #    new_borders = [i*builder_weight, (i+1)*builder_weight]
    #    print(new_borders)
    #    if (i == builders_count - 1 and (i+1)*builder_weight < weight):
    #        new_borders[1] = (i+1)*builder_weight
        builders_borders.append([builder_weight[i-1],builder_weight[i]])
    return builders_borders


# Вывод сложных структур системы

def print_bricks_matrix(bricks_matrix):
    print("-"*10, "bricks_matrix", "-"*10)
    for i in range(len(bricks_matrix)):
        print(bricks_matrix[i])
        #for j in range(len(bricks_matrix[i])):
        
    print("-"*10, "-"*len("bricks_matrix"), "-"*10)

def print_bulders_borders(bulders_borders):
    print("-"*10, "bulders_borders", "-"*10)
    for i in range(len(bulders_borders)):
        print(i, bulders_borders[i])
        #for j in range(len(bricks_matrix[i])):
        
    print("-"*10, "-"*len("bulders_borders"), "-"*10)


# Полная инициализация системы

def init_sistem(builders_count, weight, height, min_len_string, max_len_string):
    global QUENE, BRICKS_MATRIX, BUILDERS_BORDERS,BUILDERS_WORK_STATUS
    quene_len = get_quene_len(weight, height)
    QUENE = make_start_quene(quene_len, min_len_string, max_len_string)
    
    BRICKS_MATRIX = make_bricks_matrix(height, weight)
    
    BUILDERS_BORDERS = make_builders_borders(builders_count, weight)

    BUILDERS_WORK_STATUS = make_builders_work_status()



init_sistem(BUILDERS_COUNT, W, H, MIN_LEN_STRING, MAX_LEN_STRING)
# проверка инициализации системы
if __name__ == "__main__":
    print_bricks_matrix(BRICKS_MATRIX)
    print_bulders_borders(BUILDERS_BORDERS)
    print(get_quene_len(W, H))
    print(BUILDERS_WORK_STATUS)

