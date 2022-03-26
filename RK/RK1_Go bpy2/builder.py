from common import END_OF_WORK, NOT_END_OF_WORK, MUTeX_BUILDERS_BORDERS, MUTeX_BRICKS_MATRIX, MUTeX_QUENE, W, H, BUILDERS_WORK_STATUS, FLAG_END_OF_WORK
from common import BUILDERS_BORDERS, BRICKS_MATRIX, QUENE, STATUS_BUILDER_ACTIVE_WAIT
from common import sleep
from app import draw_brick, make_app_thread


from common import BUILDERS_COUNT, Thread
from main import print_bricks_matrix


def get_cur_request(thread_name):
    MUTeX_QUENE.acquire()
    if (QUENE.empty()):
        #print("Thread ", thread_name, ": ", QUENE.empty())
        MUTeX_QUENE.release()
        return "",False
    cur_req = QUENE.get()
    #print("Thread ", thread_name, ": ", cur_req)
    #print("I'm builder")
    MUTeX_QUENE.release()
    return cur_req, True

def get_current_place(thread_name):

    for j in range(H):

        MUTeX_BUILDERS_BORDERS.acquire()
        border= []
        border.append(BUILDERS_BORDERS[thread_name][0])
        border.append(BUILDERS_BORDERS[thread_name][1])
        MUTeX_BUILDERS_BORDERS.release()
        
        if (thread_name +1 == BUILDERS_COUNT and j %2 == 1):
            border[1] += 1

        for i in range(border[0],border[1]):
            con = True 
            if (j %2 == 0 and j != 0):
                con = (BRICKS_MATRIX[j-1][i] == 1 and BRICKS_MATRIX[j-1][i+1] == 1)
            if (j %2 == 1):
                if (i == 0):
                    con = BRICKS_MATRIX[j-1][i] == 1
                elif (i==W):
                    con = BRICKS_MATRIX[j-1][i-1] == 1
                else:
                    con = (BRICKS_MATRIX[j-1][i-1] == 1 and BRICKS_MATRIX[j-1][i] == 1)
            if (BRICKS_MATRIX[j][i]) == 0 and con:
                MUTeX_BRICKS_MATRIX.acquire()
                BRICKS_MATRIX[j][i] = 2
                MUTeX_BRICKS_MATRIX.release()
                return i, j, True
    #print("It's very bad error")
    return -1, -1, True

def builder(i):
    global BRICKS_MATRIX, QUENE
    

    while FLAG_END_OF_WORK == NOT_END_OF_WORK:


        # расчет, есть ли возможность выполнить запрос
        x_w, y_h, flag = get_current_place(i)
        #if (i == 0):
        #    print(x_w, y_h, flag)
        if (not flag):
            break
        if (y_h == -1):
            BUILDERS_WORK_STATUS[i][1] = True
            continue
        else: 
            BUILDERS_WORK_STATUS[i][1] = False

        #получить запрос
        cur_req, flag = get_cur_request(i)
        if (not flag):
            break
        #выполнить запрос
        sleep(len(cur_req)//10)
        MUTeX_BRICKS_MATRIX.acquire()
        BRICKS_MATRIX[y_h][x_w] = 1
        MUTeX_BRICKS_MATRIX.release()

        #визуализировать результат
        #print(x_w, y_h, i)
        draw_brick(x_w, y_h, i)
        BUILDERS_WORK_STATUS[i][0]+=1
    print("Thread ", i, "ended")




def generate_new_borders(work_res, weight_borders):
    borders = []
    max_v = work_res[0][0]
    for i in range(len(work_res)):
        if (max_v < work_res[i][0]):
            max_v = work_res[i][0]
    summ = 0
    for i in range(len(work_res)):
        if (work_res[i][1] == STATUS_BUILDER_ACTIVE_WAIT):
            summ += max_v
        else:
            summ += work_res[i][0]
    cur = weight_borders[0] 
    borders.append(0)
    for i in range(len(work_res)):
        if (work_res[i][1] == STATUS_BUILDER_ACTIVE_WAIT):
            cur += (weight_borders[1]-weight_borders[0]) * max_v/summ
        else:
            cur += (weight_borders[1]-weight_borders[0]) * work_res[i][0]/summ
        borders.append(cur)
    res = []
    print(borders)
    for i in range(len(borders)-1):
        res.append([int(borders[i]),int(borders[i+1])])
    return res

def change_dispetcher_borders(bricks_matrix, w):
    res = [0,w]
    last_line_id = len(bricks_matrix) -1
    for i in range(len(bricks_matrix[last_line_id])):
        res[0] = i 
        if (bricks_matrix[last_line_id][i]) == 0:
            break
    for i in range(len(bricks_matrix[last_line_id]), 0):
        res[1] = i
        if (bricks_matrix[last_line_id][i]) == 0:
            break
    print(res)
    return res 

def get_end_cond(bricks_matrix):
    last_line_id = len(bricks_matrix) -1
    con = True
    for i in range(len(bricks_matrix[last_line_id])):
        if((bricks_matrix[last_line_id][i]) == 0):
            con = False 
            return con
    return con
        
def dispetcher():
    global FLAG_END_OF_WORK,BUILDERS_BORDERS
    con = True
    width_borders = [0, W]
    while con:
        sleep(5)
        '''max_id = 0
        max_v = BUILDERS_WORK_STATUS[0][0]
        for i in range(BUILDERS_COUNT):
            cur = BUILDERS_WORK_STATUS[i][0]
            if (cur > max_v):
                max_v = cur 
                max_id = i
        if (BUILDERS_BORDERS[min_id][1] - BUILDERS_BORDERS[min_id][0] == 0):
            continue
        delta = 0
        print(BUILDERS_BORDERS, BUILDERS_WORK_STATUS)
        MUTeX_BUILDERS_BORDERS.acquire()
        for i in range(BUILDERS_COUNT):
            BUILDERS_WORK_STATUS[i][0] = 0
            BUILDERS_BORDERS[i][0]+= delta
            if (i == max_id):
                delta -=1
            if (i == min_id):
                delta +=1
            BUILDERS_BORDERS[i][1]+=delta
        MUTeX_BUILDERS_BORDERS.release()
        print(BUILDERS_BORDERS, BUILDERS_WORK_STATUS)
        print_bricks_matrix(BRICKS_MATRIX)
        print()'''
        width_borders = change_dispetcher_borders(BRICKS_MATRIX, W)
        print(BUILDERS_BORDERS, BUILDERS_WORK_STATUS)
        print_bricks_matrix(BRICKS_MATRIX)
        MUTeX_BUILDERS_BORDERS.acquire()
        BUILDERS_BORDERS = generate_new_borders(BUILDERS_WORK_STATUS,width_borders)
        for i in range(BUILDERS_COUNT):
            BUILDERS_WORK_STATUS[i][0] = 0
        MUTeX_BUILDERS_BORDERS.release()
        print(BUILDERS_BORDERS, BUILDERS_WORK_STATUS)
        print()

        MUTeX_QUENE.acquire()
        if (get_end_cond(BRICKS_MATRIX)):
            FLAG_END_OF_WORK = END_OF_WORK
            con = False
        MUTeX_QUENE.release()
        

def director(height, weight):
    global BRICKS_MATRIX, QUENE
    
    if (QUENE.empty()):
        print("Main Thread : ", QUENE.empty())

    app_thread=make_app_thread()
    app_thread.start()
    sleep(1)

    th=[]

    disp_thread = Thread(target = dispetcher)

    for i in range(BUILDERS_COUNT):
        print("Thread:", i,"created")
        th.append(Thread(target=builder, args = (i,)))
    disp_thread.start()
    for i in range(BUILDERS_COUNT):
        th[i].start()
          
    for i in range(BUILDERS_COUNT):
        th[i].join()
    disp_thread.join()
    print("I'm director")
    print(BUILDERS_WORK_STATUS)
    app_thread.join()


if (__name__ == "__main__"):
    director(W, H)

