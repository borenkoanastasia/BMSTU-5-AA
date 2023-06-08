from threading import Thread
from time import sleep
from random import randint
import queue

import tkinter as tk

import threading


MUTeX_QUENE = threading.Lock()
MUTeX_BRICKS_MATRIX = threading.Lock()


BUILDERS_COUNT = 4
BRICKS_MATRIX = []
LATTINICA = "abcdefghijklmnopqrstuvwxyz"
KIRILLICA = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"

BUILDERS_BORDERS = []

CANVAS =0

def generate_builder_borders(builders_count, weight):
    builders_borders = []
    builder_weight = weight/builders_count
    for i in range(builders_count):
        new_borders = [i*builder_weight, (i+1)*builder_weight]
        if (i == builders_count - 1 and (i+1)*builder_weight < weight):
            new_borders[1] = (i+1)*builder_weight
        builders_borders.append(new_borders)
    return builders_borders

def make_bricks_matrix(height, weight):
    matrix = []
    print(height, weight)
    for i in range(height):
        arr = []
        border = weight
        if (i%2 == 1):
            border += 1
        for j in range(border): 
            arr.append(0)
        matrix.append(arr)
    return matrix

def print_bricks_matrix(bricks_matrix):
    print("-"*10, "bricks_matrix", "-"*10)
    for i in range(len(bricks_matrix)):
        print(bricks_matrix[i])
        #for j in range(len(bricks_matrix[i])):
        
    print("-"*10, "-"*len("bricks_matrix"), "-"*10)

def get_rand_string(size):
    string = ""
    for i in range(size):
        string = string + LATTINICA[randint(-1, len(LATTINICA)-1)]
    return string

def make_start_quene(size, min_len, max_len):
    q = queue.Queue()

    for i in range(size):
        new_el = get_rand_string(randint(min_len, max_len))
        q.put(new_el)
    return q

def builder(granica, i):
    global BRICKS_MATRIX, QUENE
    cur_line = 0
    
    while True:
        MUTeX_QUENE.acquire()
        if (QUENE.empty()):
           print("Thread ", i, ": ", QUENE.empty())
           MUTeX_QUENE.release()
           return
        cur_req = QUENE.get()
        print("Thread ", i, ": ", cur_req)
        #print("I'm builder")
        MUTeX_QUENE.release()
        sleep(len(cur_req)//10)
        MUTeX_BRICKS_MATRIX.acquire()
        g_start = granica[0]
        g_end = granica[1]
        if (cur_line %2 == 1):
            g_end +=1
        flag_new_line = True
        #while flag_new_line:
        #    for i in range(g_start, g_end):
        #        if (BRICKS_MATRIX[cur_line][i]) == 0:
        #            w = cur_line 
        #            h = i
        #            print("I find empty place!!!", cur_line, i)
        #            flag_new_line = False 
        #    if flag_new_line:
        #        cur_line+=1
        MUTeX_BRICKS_MATRIX.release()
        draw_brick(0, 0)
    
def get_bricks_height(len_start_quene, weight):
    print(len_start_quene%(weight*2+1), len_start_quene,(weight*2+1))
    if (len_start_quene%(weight*2+1) >= weight):
        return len_start_quene//(weight*2+1)+1
    else:
        return len_start_quene//(weight*2+1)

def dispetcher(grs):
    print()#while no


def director(height, weight):
    global BRICKS_MATRIX, QUENE
    BRICKS_MATRIX = make_bricks_matrix(height, weight)
    th = []
    app_thread = Thread(target=start_window)
    print(QUENE)
    print_bricks_matrix(BRICKS_MATRIX)
    
    if (QUENE.empty()):
        print("Main Thread : ", QUENE.empty())

    
    app_thread.start()

    for i in range(BUILDERS_COUNT):
         g = weight//BUILDERS_COUNT
         granica = [i*g, (i+1)*g]
         th.append(Thread(target=builder, args = (granica,i,)))
    for i in range(BUILDERS_COUNT):
         th[i].start()
          
    for i in range(BUILDERS_COUNT):
        th[i].join()
    app_thread.join()
    print("I'm director")

def draw_brick(w, h):
    global CANVAS
    CANVAS.create_rectangle(60, 80, 140, 190,
                   fill='yellow',
                   outline='green',
                   width=3,
                   activedash=(5, 4))
    print("draw brick!")
    
def get_quene_len(q):
    t = 0
    new_q = queue.Queue()
    while not q.empty():
        x = q.get()
        print(x)
        new_q.put(x)
        t+=1
    q = new_q
    return t, q

def start_window():
    global CANVAS
    print("start Window!")
    window = tk.Tk()
    
    greeting = tk.Label(text="Привет, Tkinter!")
    greeting.pack()
    CANVAS = tk.Canvas(window, width=1000, height=1000, bg="white")
    #canvas.create_line(15, 25, 200, 25)
    #canvas.create_line(300, 35, 300, 200, dash=(4, 2))
    #canvas.create_line(55, 85, 155, 85, 105, 180, 55, 85)
    CANVAS.pack(fill=tk.BOTH, expand=1)
    window.mainloop()

if ("__main__" == __name__):
    QUENE = make_start_quene(10, 8, 20)
    W = 10    
    x = len(QUENE)
    H = get_bricks_height(x, W)
    director(H, W)   

