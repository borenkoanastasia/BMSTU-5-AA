package main 

import "github.com/gotk3/gotk3/cairo" 
import "github.com/gotk3/gotk3/gtk" 

import "log"
import "fmt"

const SCREEN_WIDTH = 900
const SCREEN_HEIGHT = 700
const SCREEN_BORDERS = 0

var SCREEN_WIDGET cairo.Context 
var WIN (*gtk.Window)

var RED = []float64{1, 0, 0}
var YELLOW = []float64{1, 1, 0}
var GREEN = []float64{0, 1, 0}
var BLUE = []float64{0, 0, 1}

var COLORS = [][]float64{RED, YELLOW, GREEN, BLUE}

type screen_parameters_t struct{
	screen_width int 
	screen_height int
	brick_width int
	brick_height int
}

type screen_brick struct{
	x float64 
	y float64 
}

var BRICKS_MATRIX *bricks_matrix_t
var SCREEN_PARAMETERS screen_parameters_t

/*func convert_real_to_screen(app screen_parameters_t, r brick_t)float_coordinates_t{
	var s float_coordinates_t
	s.y = app.brick_height+int(r.real.y)
	if (r.real.y % 2 == 1){
    	s.x = app.brick_width/2 + app.brick_width*(r.real.x)
	} else {
		s.x = app.brick_width*(r.real.x)
	}
	return s
}*/

func max_2_int(a int, b int)int{
	if (a > b){
		return a
	}
	return b
}

func init_screen(width int, height int, builders_count int)screen_parameters_t{
	var app screen_parameters_t
	
	app.screen_width = SCREEN_WIDTH 
	app.brick_width = (app.screen_width)/width
	app.brick_height = app.brick_width/2
	app.screen_height = (max_2_int(SCREEN_HEIGHT, height * int(app.brick_height)))
	return app
}

func draw_brick(){
	//print("draw")
	WIN.QueueDraw()
}

func draw_brick_matrix(da *gtk.DrawingArea, cr *cairo.Context){
	//fmt.Println(SCREEN_PARAMETERS)
	cr.SetSourceRGB(1, 0, 0)
	for i:=0;i<len((*BRICKS_MATRIX));i++{
		for j:=0;j<len((*BRICKS_MATRIX)[0]);j++{
			if ((*BRICKS_MATRIX)[i][j].fixed == true){
				//print("find i = ", i, ", j = ", j, "\n")
				//print(i * SCREEN_PARAMETERS.brick_width, "  ",
				// j*SCREEN_PARAMETERS.brick_height, " ", (SCREEN_PARAMETERS.brick_width), 
				// " ", (SCREEN_PARAMETERS.brick_height), "\n")

				var x,y,width,height float64 

				x = float64(j * SCREEN_PARAMETERS.brick_width)
				y = float64((i+1)*SCREEN_PARAMETERS.brick_height)
				width = float64(SCREEN_PARAMETERS.brick_width)
				height = float64(SCREEN_PARAMETERS.brick_height)

				var color []float64 = COLORS[(*BRICKS_MATRIX)[i][j].builder_name]

				if (i % 2 == 0){
					if (j == 0){
						x = 0
						width = float64(SCREEN_PARAMETERS.brick_width) /2 
					}else{
						x = float64(j * SCREEN_PARAMETERS.brick_width) - float64(SCREEN_PARAMETERS.brick_width) /2 
					}
				} else if (j == len((*BRICKS_MATRIX)[i]) - 1){
					width = float64(SCREEN_PARAMETERS.brick_width) /2 
				}
				cr.SetSourceRGB(1,1,1)
				cr.Rectangle(x, float64(SCREEN_PARAMETERS.screen_height)-y, width, height)
				cr.Fill()

				cr.SetSourceRGB(color[0], color[1], color[2])
				cr.Rectangle(x+1,float64(SCREEN_PARAMETERS.screen_height) -(y)+1, width-2, height-2)
				cr.Fill()
			}
		}
	}
}



func app_start(system_parameters *system_parameters_t)chan bool{
	BRICKS_MATRIX = &system_parameters.bricks_matrix
	var app_end_chan chan bool = make(chan bool, 1)

	SCREEN_PARAMETERS = init_screen(system_parameters.width, system_parameters.height, system_parameters.builders_count)

	// Инициализируем GTK.
	gtk.Init(nil)

	// Создаём билдер
	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	// Загружаем в билдер окно из файла Glade
	err = b.AddFromFile("stena_fox.glade")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	// Получаем объект главного окна по ID
	obj, err := b.GetObject("window_main")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	obj2, err := b.GetObject("width_entry")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	width_entry := obj2.(*gtk.Entry)
	width_entry.GetText()

	// Преобразуем из объекта именно окно типа gtk.Window
	// и соединяем с сигналом "destroy" чтобы можно было закрыть
	// приложение при закрытии окна
	win := obj.(*gtk.Window)

	WIN = win
	
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Получаем поле ввода
	obj, _ = b.GetObject("draw_area")
	draw_area := obj.(*gtk.DrawingArea)
    (draw_area).SetSizeRequest(int(SCREEN_PARAMETERS.screen_width) - SCREEN_BORDERS, int(SCREEN_PARAMETERS.screen_height)+ SCREEN_BORDERS)

	// Получаем кнопку
	obj, _ = b.GetObject("start_button")
	button1 := obj.(*gtk.Button)

	// Получаем метку
	// Сигнал по нажатию на кнопку
	button1.Connect("clicked", func() {
        fmt.Println("start")
	})
	
	// Event handlers
	draw_area.Connect("draw", draw_brick_matrix)

	// Отображаем все виджеты в окне
	win.ShowAll()

	// Выполняем главный цикл GTK (для отрисовки). Он остановится когда
	// выполнится gtk.MainQuit()
	f := (func(){
		gtk.Main()

		app_end_chan <- true
	})
	go f()
	return app_end_chan
}