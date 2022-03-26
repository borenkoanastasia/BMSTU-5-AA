package main

import (
	"log"
    "fmt"
	"time"


	"github.com/gotk3/gotk3/gtk" // as gtk
)

//import "github.com/gotk3/gotk3/gdk" 
import "github.com/gotk3/gotk3/cairo" 

/*
const (
	KEY_LEFT  uint = 65361
	KEY_UP    uint = 65362
	KEY_RIGHT uint = 65363
	KEY_DOWN  uint = 65364
)*/
func draw_block(da *gtk.DrawingArea, cr *cairo.Context) {
    // Data
	unitSize := 20.0
	x := 0.0
	y := 0.0
    cr.SetSourceRGB(0, 0, 0)
    cr.Rectangle(x*unitSize, y*unitSize, unitSize, unitSize)
    cr.Fill()
}

func main() {
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

	// Преобразуем из объекта именно окно типа gtk.Window
	// и соединяем с сигналом "destroy" чтобы можно было закрыть
	// приложение при закрытии окна
	win := obj.(*gtk.Window)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Получаем поле ввода
	obj, _ = b.GetObject("draw_area")
	draw_area := obj.(*gtk.DrawingArea)

    (draw_area).SetSizeRequest(500, 500)

	// Получаем кнопку
	obj, _ = b.GetObject("start_button")
	button1 := obj.(*gtk.Button)

	// Получаем метку
	/*obj, _ = b.GetObject("label_1")
	label1 := obj.(*gtk.Label)*/

	// Сигнал по нажатию на кнопку
	button1.Connect("clicked", func() {
        fmt.Println("start")
		/*text, err := entry1.GetText()
		if err == nil {
			// Устанавливаем текст из поля ввода метке
			label1.SetText(text)
		}*/
	})
	/*keyMap := map[uint]func(){
		KEY_LEFT:  func() { x-- },
		KEY_UP:    func() { y-- },
		KEY_RIGHT: func() { x++ },
		KEY_DOWN:  func() { y++ },
	}*/
	x := 0.0
	y := 0.0
	unitSize := 20.0

	// Event handlers
	var car *cairo.Context
	draw_area.Connect("draw", func (da *gtk.DrawingArea, cr *cairo.Context) {
		// Data
		cr.SetSourceRGB(0, 0, 0)
		cr.Rectangle(x*unitSize, y*unitSize, unitSize, unitSize)
		cr.Fill()
		car = cr
	})
	/*win.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
		keyEvent := &gdk.EventKey{ev}
		if move, found := keyMap[keyEvent.KeyVal()]; found {
			move()
			win.QueueDraw()
		}
	})*/

	// Отображаем все виджеты в окне
	win.ShowAll()

	// Выполняем главный цикл GTK (для отрисовки). Он остановится когда
	// выполнится gtk.MainQuit()
	
	go gtk.Main()

	time.Sleep(1000000000)


	//unitSize = 20.0
	//x = 10.0
	//y = 10.0
	//car.SetSourceRGB(0, 0, 0)
	//car.Rectangle(x*unitSize, y*unitSize, unitSize, unitSize)
    //car.Fill()

	win.QueueDraw()
	for ; (true);{}
}