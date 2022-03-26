import tkinter as tk
import common as common

CANVAS = 0
MUTeX_CANVAS = common.Lock()

SCREEN_WIDTH = 1000 

BRICK_WIDTH = SCREEN_WIDTH / (common.W*2-1)*2
BRICK_HEIGHT = BRICK_WIDTH//2#50#SCREEN_HEIGHT / common.H
SCREEN_HEIGHT = max(BRICK_HEIGHT*common.H, 1000)

BUILDERS_COLORS = [ "dark orange", "red", "coral3", "maroon1","cyan"]#"deep sky blue","cyan","aquamarine", "lawn green", "gold", "dark orange", "red", "coral3", "maroon1"]

def construct_window(window):
    print("start Window!")
    
    greeting = tk.Label(text="Привет, Tkinter!")
    greeting.pack()
    canvas = tk.Canvas(window, width=SCREEN_WIDTH, height=SCREEN_HEIGHT, bg="white")
    #canvas.create_line(15, 25, 200, 25)
    #canvas.create_line(300, 35, 300, 200, dash=(4, 2))
    #canvas.create_line(55, 85, 155, 85, 105, 180, 55, 85)
    canvas.pack(fill=tk.BOTH, expand=1)
    return canvas

def continue_app():
    global CANVAS
    window = tk.Tk()
    CANVAS = construct_window(window)
    window.mainloop()

def convert_coord(x_y):
    x_y[1] = SCREEN_HEIGHT - x_y[1]

def draw_brick(x, y, name_of_builder):
    #if (name_of_builder == 0):
    #print("try draw",x, y, name_of_builder)
    global CANVAS, MUTeX_CANVAS
    
    # Перевод координат из абсолютных в экранные
    x_y_start = [(x)*BRICK_WIDTH, (y)*BRICK_HEIGHT]
    x_y_end = [(x+1)*BRICK_WIDTH, (y+1)*BRICK_HEIGHT]

    if (y %2 == 1):
        x_y_start[0] = (x-1)*BRICK_WIDTH + BRICK_WIDTH/2
        x_y_end[0] =  x*BRICK_WIDTH + BRICK_WIDTH/2
        if (x_y_start[0] < 0):
            x_y_start[0] = 0
        if (x_y_end[0] > SCREEN_WIDTH):
            x_y_end[0] = SCREEN_WIDTH
    if (y %2 == 0 and x == common.W-1):
        x_y_end[0] =  (x+1)*BRICK_WIDTH - BRICK_WIDTH/2
        if (x_y_end[0] > SCREEN_WIDTH):
            x_y_end[0] = SCREEN_WIDTH


    convert_coord(x_y_start)
    convert_coord(x_y_end)

    # Отображение на экране
    MUTeX_CANVAS.acquire()
    CANVAS.create_rectangle(x_y_start[0], x_y_start[1], x_y_end[0], x_y_end[1],
                   fill=BUILDERS_COLORS[name_of_builder],
                   outline='black',
                   width=3,
                   activedash=(5, 4))
    MUTeX_CANVAS.release()

    #print("draw brick!")


def make_app_thread():
    app_thread = common.Thread(target=continue_app)
    return app_thread

if __name__ == "__main__":
    at = make_app_thread()
    at.start() 
    common.sleep(1)
    for j in range(10):
        for i in range(common.W):
            draw_brick(i, j, 3)
        if (j%2 == 1):
            draw_brick(10, j, 3)

    at.join()